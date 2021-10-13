package credentials

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"os"
	"testing"
)

func TestCredentialsFromEnv(t *testing.T) {
	defer os.Clearenv()
	os.Setenv("OSC_ACCESS_KEY", "access")
	os.Setenv("OSC_SECRET_KEY", "secret")
	os.Setenv("OSC_REGION", "region")

	creds, err := CredentialsFromEnv()
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}
	if e, a := "access", creds.AccessKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "secret", creds.SecretKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestEnvNoAccessKey(t *testing.T) {
	defer os.Clearenv()

	creds, err := CredentialsFromEnv()
	if err == nil {
		t.Errorf("expect error, got %v", err)
	}
	if e, a := "", creds.AccessKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}

func TestDefaultProfileCredentials(t *testing.T) {
	creds, err := CredentialsFromFile("default")
	if err != nil {
		fmt.Println("problem to extract creds from file", err)
	}
	spew.Dump(creds)
}

func TestLoadCredentialsWithOSC_PROFILE(t *testing.T) {
	os.Clearenv()
	os.Setenv("OSC_PROFILE", "dev")

	profile := GetProfile()
	fmt.Println("the profile", profile)
	creds, err := CredentialsFromFile(profile)
	if err != nil {
		t.Errorf("expect nil, got %v", err)
	}
	if e, a := "", creds.AccessKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
	if e, a := "", creds.SecretKey; e != a {
		t.Errorf("expect %v, got %v", e, a)
	}
}
