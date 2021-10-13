package credentials

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

var STD_PATH = filepath.Join(UserHomeDir(), ".osc", "config.json")

type Credentials struct {
	AccessKey  string `json:"access_key"`
	SecretKey  string `json:"secret_key"`
	RegionName string `json:"region_name"`
	Profile    string
}

func NewCredentials(ak, sk, region, profile string) *Credentials {
	return &Credentials{AccessKey: ak, SecretKey: sk, RegionName: region, Profile: profile}
}

func CredentialsFromEnv() (Credentials, error) {
	var err error
	e := Credentials{}
	if os.Getenv("OSC_ACCESS_KEY") != "" {
		e.AccessKey = os.Getenv("OSC_ACCESS_KEY")
	} else {
		err = fmt.Errorf("access key  not set")
		return Credentials{}, err
	}

	if os.Getenv("OSC_SECRET_KEY") != "" {
		e.SecretKey = os.Getenv("OSC_SECRET_KEY")
	} else {
		err := fmt.Errorf("secret key  not set")
		return Credentials{}, err
	}

	if os.Getenv("OSC_REGION") != "" {
		e.RegionName = os.Getenv("OSC_REGION")
	} else {
		err := fmt.Errorf("region  not set")
		return Credentials{}, err
	}

	return Credentials{
		AccessKey:  e.AccessKey,
		SecretKey:  e.SecretKey,
		RegionName: e.RegionName,
	}, err
}

func CredentialsFromFile(profile string) (Credentials, error) {
	var Result map[string]Credentials
	jsonFile, err := os.Open(STD_PATH)
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		fmt.Println(err)
	}
	json.Unmarshal([]byte(byteValue), &Result)
	if err != nil {
		fmt.Println("error osscured while extract json file data:", err)
	}
	creds, ok := Result[profile]
	if !ok {
		fmt.Println("profile name not set and it is not default", err)
		return Credentials{}, err
	}

	ak := creds.AccessKey
	if len(ak) == 0 {
		fmt.Errorf("credentials file don't contain access key: %s", err)
		return Credentials{}, err
	}

	sk := creds.SecretKey
	if len(sk) == 0 {
		fmt.Errorf("credentials file don't contain security key: %s", err)
		return Credentials{}, err
	}

	reg := creds.RegionName
	if len(reg) == 0 {
		fmt.Println("credentials file don't contain region")
	}

	return Credentials{
		AccessKey:  ak,
		SecretKey:  sk,
		RegionName: reg,
		Profile:    profile,
	}, err
}

func GetProfile() string {
	profile := os.Getenv("OSC_PROFILE")
	if profile == "" {
		profile = "default"
	}
	return profile
}

func UserHomeDir() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir
}
