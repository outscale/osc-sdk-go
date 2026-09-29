package profile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path"
	"path/filepath"
	"sort"
)

var ErrProfileNotFound = errors.New("profile not found")

const (
	DefaultProfile     = "default"
	ConfigFileEnvName  = "OSC_CONFIG_PATH"
	ProfileNameEnvName = "OSC_PROFILE"
)

type profile struct {
	Fields
	Default bool `json:"default,omitempty"`
}

type ConfigFile struct {
	FilePath            *string
	SelectedProfileName string
	DefaultProfileName  string

	// raw profile content
	rawProfiles map[string]json.RawMessage

	// Flags for explicit biaviour
	explicitFile    bool
	explicitProfile bool
}

func (cf *ConfigFile) SelectedProfile() (string, Fields, error) {
	f, e := cf.Profile(cf.SelectedProfileName)
	if e != nil && errors.Is(e, ErrProfileNotFound) && !cf.explicitFile && !cf.explicitProfile {
		return cf.DefaultProfileName, Fields{}, nil
	}

	return cf.SelectedProfileName, f, e
}

func (cf *ConfigFile) DefaultProfile() (string, Fields, error) {
	f, e := cf.Profile(cf.DefaultProfileName)
	if e != nil && errors.Is(e, ErrProfileNotFound) && !cf.explicitFile {
		return cf.DefaultProfileName, Fields{}, nil
	}

	return cf.DefaultProfileName, f, e
}

func (cf *ConfigFile) Profile(p string) (Fields, error) {
	var selectedProfile profile
	if raw, ok := cf.rawProfiles[p]; ok {
		if err := decodeOne(raw, &selectedProfile); err != nil {
			return Fields{}, fmt.Errorf("decode profile %q: %w", p, err)
		}
		return selectedProfile.Fields, nil
	}
	return Fields{}, ErrProfileNotFound
}

func (cf *ConfigFile) ProfileList() []string {
	profiles := make([]string, 0, len(cf.rawProfiles))
	for name := range cf.rawProfiles {
		profiles = append(profiles, name)
	}
	sort.Strings(profiles)
	return profiles
}

func (cf *ConfigFile) ProfileAdd(name string, fields Fields) error {
	data, err := json.Marshal(fields) //nolint:gosec
	if err != nil {
		return fmt.Errorf("encode new profile %q: %w", name, err)
	}

	cf.rawProfiles[name] = json.RawMessage(data)

	return nil
}

func (cf *ConfigFile) ProfileRemove(name string) error {
	if _, ok := cf.rawProfiles[name]; !ok {
		return fmt.Errorf("profile %q: %w", name, ErrProfileNotFound)
	}
	delete(cf.rawProfiles, name)

	// TODO: ensure default source exist

	return nil
}

func LoadConfigFile(o Options) (*ConfigFile, error) {
	if o.Env == nil {
		o.Env = snapshotEnv()
	}
	if o.ReadFile == nil {
		o.ReadFile = os.ReadFile
	}

	return loadConfigFile(o)
}

func loadConfigFile(o Options) (*ConfigFile, error) {
	var cf ConfigFile

	// 1. First load config file
	filePath, explicitFile := choose(o.FilePath, o.Env, ConfigFileEnvName)
	if filePath == "" {
		if explicitFile {
			return nil, errors.New("config path is explicitly empty")
		}
		filePath, _ = DefaultConfigPath()
	}
	cf.explicitFile = explicitFile

	cf.rawProfiles = map[string]json.RawMessage{}
	data, err := o.ReadFile(filePath)
	switch {
	case err == nil:
		if err := decodeOne(data, &cf.rawProfiles); err != nil {
			return nil, fmt.Errorf("parse config file: %w", err)
		}
		cf.FilePath = &filePath

	case errors.Is(err, os.ErrNotExist) && !explicitFile:
		// A missing default file is ok

	default:
		return nil, fmt.Errorf("reading config file: %w", err)
	}

	// 2. Find default profile
	for name, raw := range cf.rawProfiles {
		var marker struct {
			Default *bool `json:"default"`
		}
		if err := json.Unmarshal(raw, &marker); err != nil {
			return nil, fmt.Errorf("profile %q: invalid default marker: %w", name, err)
		}
		if marker.Default != nil && *marker.Default {
			if cf.DefaultProfileName != "" {
				return nil, fmt.Errorf("more than one profile is marked default. first: %q, second: %q", cf.DefaultProfileName, name)
			}
			cf.DefaultProfileName = name
		}
	}
	if cf.DefaultProfileName == "" {
		cf.DefaultProfileName = DefaultProfile
	}

	// 3. Find selected profile
	cf.SelectedProfileName, cf.explicitProfile = choose(o.ProfileName, o.Env, ProfileNameEnvName)
	if !cf.explicitProfile {
		// Looking for marked default profile
		cf.SelectedProfileName = cf.DefaultProfileName
	}

	return &cf, nil
}

func (cf *ConfigFile) Save() error {
	path, err := cf.configPath()
	if err != nil {
		return fmt.Errorf("unable to save config file: %w", err)
	}

	dir := filepath.Dir(path)
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		err := os.MkdirAll(dir, 0o700)
		if err != nil {
			return fmt.Errorf("unable to save config file: %w", err)
		}
	}

	tmpFile := path + ".tmp"
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0o600) //nolint:gosec
	if err != nil {
		return fmt.Errorf("unable to save config file: %w", err)
	}
	enc := json.NewEncoder(fd)
	enc.SetIndent("", "  ")
	err = enc.Encode(cf.rawProfiles)
	if err != nil {
		_ = fd.Close()
	} else {
		err = fd.Close()
	}
	if err == nil {
		saveFile := path + ".saved"
		err = os.Rename(path, saveFile)
		if errors.Is(err, fs.ErrNotExist) {
			err = nil
		}
	}
	if err == nil {
		err = os.Rename(tmpFile, path)
	}
	if err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("unable to save config file: %w", err)
	}
	return nil
}

func (cf *ConfigFile) configPath() (string, error) {
	if cf.FilePath != nil {
		return *cf.FilePath, nil
	}
	return DefaultConfigPath()
}

func (cf *ConfigFile) SetDefault(profileName string) error {
	var tmpProfile profile

	if p, ok := cf.rawProfiles[cf.DefaultProfileName]; ok {
		if err := decodeOne(p, &tmpProfile); err != nil {
			return fmt.Errorf("decode default profile %q: %w", cf.DefaultProfileName, err)
		}

		tmpProfile.Default = false

		data, err := json.Marshal(tmpProfile)
		if err != nil {
			return fmt.Errorf("encode old default profile %q: %w", cf.DefaultProfileName, err)
		}

		cf.rawProfiles[cf.DefaultProfileName] = json.RawMessage(data)
	}

	if p, ok := cf.rawProfiles[profileName]; ok {
		if err := decodeOne(p, &tmpProfile); err != nil {
			return fmt.Errorf("decode profile %q: %w", profileName, err)
		}

		tmpProfile.Default = true
		data, err := json.Marshal(tmpProfile)
		if err != nil {
			return fmt.Errorf("encode new default profile %q: %w", profileName, err)
		}

		cf.rawProfiles[profileName] = json.RawMessage(data)

		// Update default profile and selected profile
		cf.DefaultProfileName = profileName
		if !cf.explicitProfile {
			// Looking for marked default profile
			cf.SelectedProfileName = cf.DefaultProfileName
		}

		return nil
	} else {
		return fmt.Errorf("profile %q does not exist", profileName)
	}
}

func DefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return path.Join(home, ".osc", "config.json"), nil
}

func decodeOne(data []byte, dst any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(dst); err != nil {
		return err
	}
	return nil
}

func choose(explicit *string, env map[string]string, name string) (string, bool) {
	// helper for selecting a value: explicit parameter takes precedence,
	// then environment variable, and finally fallback to default.
	if explicit != nil {
		return *explicit, true
	}
	v, ok := env[name]
	return v, ok
}
