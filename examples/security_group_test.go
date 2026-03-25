package examples_test

import (
	"slices"
	"testing"

	"github.com/outscale/osc-sdk-go/v3/pkg/osc"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Steps done in this test:
// 1. Create a security group.
// 2. Add an inbound SSH rule.
// 3. Read the security group and validate the rule.
// 4. Delete the security group rule.
// 5. Delete the security group.
func TestSecurityGroup(t *testing.T) {
	ctx := t.Context()

	client := newOSCClient(t)

	securityGroupName := "osc-sdk-go-test-" + RandomString(10)

	// Create a security group
	createReq := osc.CreateSecurityGroupJSONRequestBody{
		SecurityGroupName: securityGroupName,
		Description:       "Test security group lifecycle",
	}

	createResp, err := client.CreateSecurityGroup(ctx, createReq)
	require.NoError(t, err)

	deleted := false
	defer func() {
		if deleted {
			return
		}

		if createResp.SecurityGroup == nil {
			return
		}

		securityGroupID := createResp.SecurityGroup.SecurityGroupId
		_, _ = client.DeleteSecurityGroup(ctx, osc.DeleteSecurityGroupJSONRequestBody{
			SecurityGroupId: &securityGroupID,
		})
	}()

	require.NotNil(t, createResp.SecurityGroup)

	securityGroupID := createResp.SecurityGroup.SecurityGroupId
	assert.NotEmpty(t, securityGroupID)

	t.Logf("Created security group: %s", securityGroupID)

	// Create a security group rule
	tcp := "tcp"
	ipRange := "0.0.0.0/0"
	fromPort := 22
	toPort := 22
	ruleReq := osc.CreateSecurityGroupRuleJSONRequestBody{
		SecurityGroupId: securityGroupID,
		Flow:            "Inbound",
		IpProtocol:      &tcp,
		FromPortRange:   &fromPort,
		ToPortRange:     &toPort,
		IpRange:         &ipRange,
	}

	ruleResp, err := client.CreateSecurityGroupRule(ctx, ruleReq)
	require.NoError(t, err)
	require.NotNil(t, ruleResp.SecurityGroup)

	t.Logf("Created security group rule for: %s", securityGroupID)

	// Read the security group
	readReq := osc.ReadSecurityGroupsJSONRequestBody{
		Filters: &osc.FiltersSecurityGroup{
			SecurityGroupIds: &[]string{securityGroupID},
		},
	}

	readResp, err := client.ReadSecurityGroups(ctx, readReq)
	require.NoError(t, err)
	require.NotNil(t, readResp.SecurityGroups)
	require.Len(t, *readResp.SecurityGroups, 1)

	sg := (*readResp.SecurityGroups)[0]
	foundSSHRule := slices.ContainsFunc(sg.InboundRules, func(rule osc.SecurityGroupRule) bool {
		return rule.FromPortRange == fromPort &&
			rule.ToPortRange == toPort &&
			rule.IpProtocol == tcp &&
			slices.Contains(rule.IpRanges, ipRange)
	})
	assert.True(t, foundSSHRule, "expected SSH rule %s %d-%d for %s", tcp, fromPort, toPort, ipRange)

	t.Logf("Successfully read security group: %s", securityGroupID)

	// Delete the security group rule
	deleteRuleReq := osc.DeleteSecurityGroupRuleJSONRequestBody{
		SecurityGroupId: securityGroupID,
		Flow:            "Inbound",
		IpProtocol:      &tcp,
		FromPortRange:   &fromPort,
		ToPortRange:     &toPort,
		IpRange:         &ipRange,
	}

	_, err = client.DeleteSecurityGroupRule(ctx, deleteRuleReq)
	require.NoError(t, err)

	t.Logf("Successfully deleted security group rule for: %s", securityGroupID)

	// Delete the security group
	deleteReq := osc.DeleteSecurityGroupJSONRequestBody{
		SecurityGroupId: &securityGroupID,
	}

	deleteResp, err := client.DeleteSecurityGroup(ctx, deleteReq)
	require.NoError(t, err)
	deleted = true
	require.NotNil(t, deleteResp.ResponseContext)
	assert.NotNil(t, deleteResp.ResponseContext.RequestId)

	t.Logf("Successfully deleted security group: %s", securityGroupID)
}
