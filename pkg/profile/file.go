package profile

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

var ErrNoDefaultProfile = errors.New("no default profile found")

type ConfigFile struct {
	Path     string
	Profiles map[string]Profile
}

func (cf *ConfigFile) DefaultProfile() (string, Profile, error) {
	for name, p := range cf.Profiles {
		if p.Default {
			return name, p, nil
		}
	}
	if _, found := cf.Profiles[DefaultProfile]; !found {
		return "", Profile{}, ErrNoDefaultProfile
	}
	return DefaultProfile, cf.Profiles[DefaultProfile], nil
}

func (cf *ConfigFile) Profile(profile string) (Profile, error) {
	if profile == "" {
		_, p, err := cf.DefaultProfile()
		return p, err
	}

	if _, found := cf.Profiles[profile]; !found {
		return Profile{}, fmt.Errorf("profile %q does not exist", profile)
	}
	return cf.Profiles[profile], nil
}

func LoadConfigFile(path string) (*ConfigFile, error) {
	if path == "" {
		var err error
		path, err = DefaultConfigPath()
		if err != nil {
			return nil, fmt.Errorf("unable to locate config file: %w", err)
		}
	}

	configJSON, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("unable to load config file: %w", err)
	}

	cf := &ConfigFile{
		Path:     path,
		Profiles: make(map[string]Profile),
	}
	if err := json.Unmarshal(configJSON, &cf.Profiles); err != nil {
		return nil, fmt.Errorf("unable to load config file: %w", err)
	}
	return cf, nil
}

func (cf *ConfigFile) Save() error {
	dir := filepath.Dir(cf.Path)
	if _, err := os.Stat(dir); errors.Is(err, fs.ErrNotExist) {
		err := os.MkdirAll(dir, 0700)
		if err != nil {
			return fmt.Errorf("unable to save config file: %w", err)
		}
	}
	tmpFile := cf.Path + ".tmp"
	fd, err := os.OpenFile(tmpFile, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return fmt.Errorf("unable to save config file: %w", err)
	}
	enc := json.NewEncoder(fd)
	enc.SetIndent("", "  ")
	err = enc.Encode(cf.Profiles)
	if err != nil {
		_ = fd.Close()
	} else {
		err = fd.Close()
	}
	if err == nil {
		saveFile := cf.Path + ".saved"
		err = os.Rename(cf.Path, saveFile)
		if errors.Is(err, fs.ErrNotExist) {
			err = nil
		}
	}
	if err == nil {
		err = os.Rename(tmpFile, cf.Path)
	}
	if err != nil {
		_ = os.Remove(tmpFile)
		return fmt.Errorf("unable to save config file: %w", err)
	}
	return nil
}

func (cf *ConfigFile) SetDefault(profile string) error {
	name, def, err := cf.DefaultProfile()
	if err == nil && def.Default {
		def.Default = false
		cf.Profiles[name] = def
	}
	if p, found := cf.Profiles[profile]; found {
		p.Default = true
		cf.Profiles[profile] = p
		return nil
	}
	return fmt.Errorf("profile %q does not exist", profile)
}
