// Package profile provide functions to manage standard Outscale profiles.
//
// The profile can be loaded from environment variables or from a JSON configuration file.
// The configuration file is located at ~/.osc/config.json by
package profile

import (
	"errors"
	"fmt"
	"os"
	"reflect"
	"strings"
)

type Fields struct {
	AccessKey         string                `json:"access_key,omitempty" mergegroup:"aksk"`
	SecretKey         string                `json:"secret_key,omitempty" log:"sensitive" mergegroup:"aksk"`
	AccessKeyV2       string                `json:"access_key_v2,omitempty" mergegroup:"aksk_iamv2"`
	SecretKeyV2       string                `json:"secret_key_v2,omitempty" log:"sensitive" mergegroup:"aksk_iamv2"`
	IAMV2Services     []string              `json:"iam_v2_services,omitempty"`
	X509ClientCert    string                `json:"x509_client_cert,omitempty" mergegroup:"x509keypair"`
	X509ClientKey     string                `json:"x509_client_key,omitempty" log:"sensitive" mergegroup:"x509keypair"`
	X509ClientCertB64 string                `json:"x509_client_cert_b64,omitempty" mergegroup:"x509keypairb64"`
	X509ClientKeyB64  string                `json:"x509_client_key_b64,omitempty" log:"sensitive" mergegroup:"x509keypairb64"`
	TlsSkipVerify     *bool                 `json:"tls_skip_verify,omitempty"`
	Login             string                `json:"login,omitempty" mergegroup:"basicauth"`
	Password          string                `json:"password,omitempty" log:"sensitive" mergegroup:"basicauth"`
	Protocol          string                `json:"protocol,omitempty"`
	Region            string                `json:"region,omitempty"`
	Endpoints         map[OscService]string `json:"endpoints,omitzero"`
}

func (r *Profile) merge(source layer) error {
	targetValue := reflect.ValueOf(&r.Fields)

	if targetValue.Kind() == reflect.Pointer && targetValue.IsNil() {
		return errors.New("target must be a non-nil pointer")
	}

	targetValue = targetValue.Elem()
	sourceValue := reflect.ValueOf(source.fields)
	if sourceValue.Kind() == reflect.Pointer {
		if sourceValue.IsNil() {
			return errors.New("source is nil")
		}
		sourceValue = sourceValue.Elem()
	}

	if targetValue.Kind() != reflect.Struct || sourceValue.Kind() != reflect.Struct {
		return errors.New("target and source must be struct")
	}

	if targetValue.Type() != sourceValue.Type() {
		return errors.New("target and source must be same type")
	}

	targetType := targetValue.Type()
	groups := make(map[string][]int)
	for i := 0; i < targetType.NumField(); i++ {
		field := targetType.Field(i)

		// Maps merge key-by-key across layers; scalars skip once set.
		if targetValue.Field(i).Kind() != reflect.Map && isSet(targetValue.Field(i)) {
			continue
		}

		group := field.Tag.Get("mergegroup")
		if group == "" {
			// if not in explicit group, it's a mono field group
			group = field.Name
		}

		groups[group] = append(groups[group], i)
	}

	for group, fields := range groups {
		setCount := 0
		for _, index := range fields {
			sourceField := sourceValue.Field(index)

			if isSet(sourceField) {
				setCount += 1
			}
		}

		if setCount > 0 && setCount != len(fields) {
			return fmt.Errorf("merge-group %q is not complete", group)
		} else if setCount == 0 {
			// group is not set by other
			continue
		}

		for _, index := range fields {
			targetField := targetValue.Field(index)
			sourceField := sourceValue.Field(index)
			targetFieldName := targetType.Field(index).Name

			if targetField.Kind() == reflect.Map {
				if sourceField.IsNil() {
					continue
				}
				if targetField.IsNil() {
					targetField.Set(reflect.MakeMap(targetField.Type()))
				}
				iter := sourceField.MapRange()
				for iter.Next() {
					key := iter.Key()
					val := iter.Value()
					if !isSet(val) {
						continue
					}
					existing := targetField.MapIndex(key)
					if existing.IsValid() && isSet(existing) {
						continue
					}
					targetField.SetMapIndex(key, val)
					r.Sources[targetFieldName+"."+key.String()] = source.name
				}
				continue
			}

			targetField.Set(sourceField)
			r.Sources[targetFieldName] = source.name
		}
	}

	return nil
}

func isSet(value reflect.Value) bool {
	switch value.Kind() {
	case reflect.Pointer, reflect.Interface, reflect.Slice:
		return !value.IsNil()
	case reflect.Map:
		return !value.IsNil() && value.Len() > 0

	default:
		return !value.IsZero()
	}
}

type Options struct {
	FilePath    *string
	ProfileName *string
	Overrides   Fields

	// **HAZMAT** do not use in production unless you understand the implications.
	Default Fields
	// Nil means snapshot the process environment.
	Env map[string]string
	// Nil mean os.ReadFile. for mocking
	ReadFile func(string) ([]byte, error)
}

func NewOptions() Options {
	return Options{
		Default: Fields{
			Region:        "eu-west-2",
			Protocol:      "https",
			TlsSkipVerify: new(false),
		},
	}
}

func NewOptionsWith(opts ...func(*Options)) Options {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

type Profile struct {
	Fields
	Sources     map[string]string
	ProfileName string
	File        *string
}

func New() (Profile, error) {
	return NewWith()
}

func NewWith(opts ...func(*Options)) (Profile, error) {
	return Load(NewOptionsWith(opts...))
}

func NewFrom(path, name string) (Profile, error) {
	return NewWith(func(o *Options) {
		if path != "" {
			o.FilePath = &path
		}

		if name != "" {
			o.ProfileName = &name
		}
	})
}

type layer struct {
	name   string
	fields Fields
}

func layerFromEnv(env map[string]string) (layer, error) {
	var fields Fields
	fields.AccessKey = env["OSC_ACCESS_KEY"]
	fields.SecretKey = env["OSC_SECRET_KEY"]
	fields.AccessKeyV2 = env["OSC_ACCESS_KEY_V2"]
	fields.SecretKeyV2 = env["OSC_SECRET_KEY_V2"]
	fields.X509ClientCert = env["OSC_X509_CLIENT_CERT"]
	fields.X509ClientCertB64 = env["OSC_X509_CLIENT_CERT_B64"]
	fields.X509ClientKey = env["OSC_X509_CLIENT_KEY"]
	fields.X509ClientKeyB64 = env["OSC_X509_CLIENT_KEY_B64"]

	v, ok := env["OSC_TLS_SKIP_VERIFY"]
	if ok {
		switch v {
		case "true":
			fields.TlsSkipVerify = new(true)
		case "false":
			fields.TlsSkipVerify = new(false)
		default:
			return layer{}, fmt.Errorf("invalid OSC_TLS_SKIP_VERIFY value: %q, expected 'true' or 'false'", v)
		}
	}

	fields.Login = env["OSC_LOGIN"]
	fields.Password = env["OSC_PASSWORD"]
	fields.Protocol = env["OSC_PROTOCOL"]
	fields.Region = env["OSC_REGION"]

	fields.Endpoints = make(map[OscService]string)
	for _, svc := range []struct {
		s OscService
		e string
	}{
		{s: OscServiceApi, e: "OSC_ENDPOINT_API"},
		{s: OscServiceOKS, e: "OSC_ENDPOINT_OKS"},
		{s: OscServiceFCU, e: "OSC_ENDPOINT_FCU"},
		{s: OscServiceLBU, e: "OSC_ENDPOINT_LBU"},
		{s: OscServiceOOS, e: "OSC_ENDPOINT_OOS"},
		{s: OscServiceEIM, e: "OSC_ENDPOINT_EIM"},
		{s: OscServiceDirectLink, e: "OSC_ENDPOINT_DIRECTLINK"},
	} {
		v, ok := env[svc.e]
		if ok {
			fields.Endpoints[svc.s] = v
		}
	}

	if v, ok := env["OSC_IAM_V2_SERVICES"]; ok {
		fields.IAMV2Services = strings.Split(v, ",")
	}

	return layer{fields: fields, name: "environment"}, nil
}

func layerFromConfigFile(cf *ConfigFile) (layer, error) {
	name, fields, err := cf.SelectedProfile()
	if err != nil {
		return layer{}, err
	}

	return layer{name: "profile:" + name, fields: fields}, nil
}

func Load(o Options) (Profile, error) {
	if o.Env == nil {
		o.Env = snapshotEnv()
	}
	if o.ReadFile == nil {
		o.ReadFile = os.ReadFile
	}

	return load(o)
}

func load(o Options) (Profile, error) {
	// 1. Load Config file and profile
	cf, err := loadConfigFile(o)
	if err != nil {
		return Profile{}, fmt.Errorf("coult not load config file: %w", err)
	}

	// 4. Layer definition and ordering
	tool := layer{name: "tool", fields: o.Overrides}

	file, err := layerFromConfigFile(cf)
	if err != nil {
		return Profile{}, fmt.Errorf("could not config file layer: %w", err)
	}

	environement, err := layerFromEnv(o.Env)
	if err != nil {
		return Profile{}, fmt.Errorf("could not load environment layer: %w", err)
	}

	builtinDefault := layer{name: "default", fields: o.Default}
	layers := []layer{tool, environement, file, builtinDefault}
	if cf.explicitProfile || cf.explicitFile {
		layers = []layer{tool, file, environement, builtinDefault}
	}

	// 5. layer merging
	r := Profile{
		Sources:     map[string]string{},
		ProfileName: cf.SelectedProfileName,
		File:        cf.FilePath,
	}

	for _, layer := range layers {
		err := r.merge(layer)
		if err != nil {
			return Profile{}, fmt.Errorf("merging layer %q: %w", layer.name, err)
		}
	}

	// 6. Add computed default
	computedDefault := layer{name: "default (computed)", fields: computedDefault(&r.Fields)}
	if err := r.merge(computedDefault); err != nil {
		return Profile{}, fmt.Errorf("merging layer default (computed): %w", err)
	}

	return r, nil
}

func computedDefault(fields *Fields) Fields {
	var computedFileds Fields
	computedFileds.Endpoints = make(map[OscService]string)
	computedFileds.Endpoints[OscServiceApi] = fmt.Sprintf("%s://api.%s.outscale.com/api/v1", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceFCU] = fmt.Sprintf("%s://fcu.%s.outscale.com", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceOOS] = fmt.Sprintf("%s://oos.%s.outscale.com", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceLBU] = fmt.Sprintf("%s://lbu.%s.outscale.com", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceEIM] = fmt.Sprintf("%s://eim.%s.outscale.com", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceDirectLink] = fmt.Sprintf("%s://directlink.%s.outscale.com", fields.Protocol, fields.Region)
	computedFileds.Endpoints[OscServiceOKS] = fmt.Sprintf("%s://api.%s.oks.outscale.com/api/v2", fields.Protocol, fields.Region)
	return computedFileds
}

func snapshotEnv() map[string]string {
	// just return a simple map of environment variables filtered by "OSC_" prefix.
	out := map[string]string{}

	for _, entry := range os.Environ() {
		k, v, ok := strings.Cut(entry, "=")
		if ok && strings.HasPrefix(k, "OSC_") {
			out[k] = v
		}
	}

	return out
}
