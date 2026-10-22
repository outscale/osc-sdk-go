package osc_test

import (
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
)

func TestUpdateAccessKeyRequestStateUsesAccessKeyState(t *testing.T) {
	state := osc.AccessKeyStateACTIVE
	request := osc.UpdateAccessKeyRequest{State: &state}
	if request.State == nil || *request.State != osc.AccessKeyStateACTIVE {
		t.Fatalf("State = %v, want %s", request.State, osc.AccessKeyStateACTIVE)
	}
}
