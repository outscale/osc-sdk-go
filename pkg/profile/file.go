package profile

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"slices"
)

var ErrProfileNotFound = errors.New("profile not found")

const (
	DefaultProfile     = "default"
	ConfigFileEnvName  = "OSC_CONFIG_FILE"
	ProfileNameEnvName = "OSC_PROFILE"
)

type ConfigFile struct {
	FilePath *string

	rawProfiles      map[string]json.RawMessage // We use RawMessage for forward compibility
	markedDefault    string
	hasMarkedDefault bool
}

func (cf *ConfigFile) defaultName() string {
	if cf.hasMarkedDefault {
		return cf.markedDefault
	}
	return DefaultProfile
}

// DefaultProfile returns the marked profile, or the entry named "default" if none is marked.
func (cf *ConfigFile) DefaultProfile() (Profile, error) {
	return cf.Profile(cf.defaultName())
}

// Profile returns a saved entry without applying environment or SDK defaults.
func (cf *ConfigFile) Profile(name string) (Profile, error) {
	raw, ok := cf.rawProfiles[name]
	if !ok {
		return Profile{}, ErrProfileNotFound
	}
	p := Profile{Name: name}
	if err := decodeOne(raw, &p); err != nil {
		return Profile{}, fmt.Errorf("decode profile %q: %w", name, err)
	}
	p.Default = name == cf.defaultName()
	return p, nil
}

// ListProfiles returns saved entries in name order, or an error if an entry cannot be decoded.
func (cf *ConfigFile) ListProfiles() ([]Profile, error) {
	names := make([]string, 0, len(cf.rawProfiles))
	for name := range cf.rawProfiles {
		names = append(names, name)
	}
	slices.Sort(names)
	profiles := make([]Profile, 0, len(names))
	for _, name := range names {
		p, err := cf.Profile(name)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, p)
	}
	return profiles, nil
}

// AddProfile writes or replaces a saved entry. If p.Default is true, it marks
// this entry as the default and clears any previous mark.
func (cf *ConfigFile) AddProfile(p Profile) error {
	data, err := json.Marshal(p)
	if err != nil {
		return fmt.Errorf("encode profile %q: %w", p.Name, err)
	}

	if p.Default && cf.hasMarkedDefault && cf.markedDefault != p.Name {
		previous, err := cf.Profile(cf.markedDefault)
		if err != nil {
			return err
		}
		previous.Default = false
		previousData, err := json.Marshal(previous)
		if err != nil {
			return fmt.Errorf("encode old default profile %q: %w", previous.Name, err)
		}
		cf.rawProfiles[previous.Name] = previousData
	}

	if cf.rawProfiles == nil {
		cf.rawProfiles = make(map[string]json.RawMessage)
	}
	cf.rawProfiles[p.Name] = data
	if p.Default {
		cf.markedDefault, cf.hasMarkedDefault = p.Name, true
	} else if p.Name == cf.markedDefault {
		cf.markedDefault, cf.hasMarkedDefault = "", false
	}

	return nil
}

func (cf *ConfigFile) RemoveProfile(name string) error {
	if _, ok := cf.rawProfiles[name]; !ok {
		return fmt.Errorf("profile %q: %w", name, ErrProfileNotFound)
	}
	delete(cf.rawProfiles, name)
	if name == cf.markedDefault {
		cf.markedDefault, cf.hasMarkedDefault = "", false
	}

	return nil
}

func LoadConfigFile(o Options) (*ConfigFile, error) {
	cf, _, err := loadConfigFile(o)
	return cf, err
}

// loadConfigFile returns whether the caller or environment explicitly chose the file.
func loadConfigFile(o Options) (*ConfigFile, bool, error) {
	if o.ReadFile == nil {
		o.ReadFile = os.ReadFile
	}
	var cf ConfigFile

	filePath, explicitFile := choose(o.FilePath, ConfigFileEnvName)
	if filePath == "" {
		if explicitFile {
			return nil, explicitFile, errors.New("config path is explicitly empty")
		}
		filePath, _ = DefaultConfigPath()
	}

	cf.rawProfiles = map[string]json.RawMessage{}
	data, err := o.ReadFile(filePath)
	switch {
	case err == nil:
		if err := decodeOne(data, &cf.rawProfiles); err != nil {
			return nil, explicitFile, fmt.Errorf("parse config file: %w", err)
		}
		cf.FilePath = &filePath

	case errors.Is(err, os.ErrNotExist) && !explicitFile:
		// A missing default file is ok

	default:
		return nil, explicitFile, fmt.Errorf("reading config file: %w", err)
	}

	for name, raw := range cf.rawProfiles {
		var marker struct {
			Default *bool `json:"default"`
		}
		if err := json.Unmarshal(raw, &marker); err != nil {
			return nil, explicitFile, fmt.Errorf("profile %q: invalid default marker: %w", name, err)
		}
		if marker.Default == nil || !*marker.Default {
			continue
		}
		if cf.hasMarkedDefault {
			return nil, explicitFile, fmt.Errorf("more than one profile is marked default. first: %q, second: %q", cf.markedDefault, name)
		}
		cf.markedDefault, cf.hasMarkedDefault = name, true
	}
	return &cf, explicitFile, nil
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

func (cf *ConfigFile) SetDefault(name string) error {
	p, err := cf.Profile(name)
	if err != nil {
		return fmt.Errorf("setting default profile %q: %w", name, err)
	}
	p.Default = true
	return cf.AddProfile(p)
}

func DefaultConfigPath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(home, ".osc", "config.json"), nil
}

func decodeOne(data []byte, dst any) error {
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(dst); err != nil {
		return err
	}
	return nil
}

func choose(explicit *string, name string) (string, bool) {
	// helper for selecting a value: explicit parameter takes precedence,
	// then environment variable, and finally fallback to default.
	if explicit != nil {
		return *explicit, true
	}
	return os.LookupEnv(name)
}
