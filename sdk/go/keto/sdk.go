package keto

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ory/keto/sdk/go/keto/swagger"
)

type SDK interface {
	AddOryAccessControlPolicyRoleMembers(flavor string, id string, body swagger.AddOryAccessControlPolicyRoleMembersBody) (*swagger.OryAccessControlPolicyRole, *swagger.APIResponse, error)
	DeleteOryAccessControlPolicy(flavor string, id string) (*swagger.APIResponse, error)
	DeleteOryAccessControlPolicyRole(flavor string, id string) (*swagger.APIResponse, error)
	DoOryAccessControlPoliciesAllow(flavor string, body swagger.OryAccessControlPolicyAllowedInput) (*swagger.AuthorizationResult, *swagger.APIResponse, error)
	GetOryAccessControlPolicy(flavor string, id string) (*swagger.OryAccessControlPolicy, *swagger.APIResponse, error)
	GetOryAccessControlPolicyRole(flavor string, id string) (*swagger.OryAccessControlPolicyRole, *swagger.APIResponse, error)
	ListOryAccessControlPolicies(flavor string, limit int64, offset int64) ([]swagger.OryAccessControlPolicy, *swagger.APIResponse, error)
	ListOryAccessControlPolicyRoles(flavor string, limit int64, offset int64) ([]swagger.OryAccessControlPolicyRole, *swagger.APIResponse, error)
	RemoveOryAccessControlPolicyRoleMembers(flavor string, id string, member string) (*swagger.APIResponse, error)
	UpsertOryAccessControlPolicy(flavor string, body swagger.OryAccessControlPolicy) (*swagger.OryAccessControlPolicy, *swagger.APIResponse, error)
	UpsertOryAccessControlPolicyRole(flavor string, body swagger.OryAccessControlPolicyRole) (*swagger.OryAccessControlPolicyRole, *swagger.APIResponse, error)
}

var _ SDK = new(CodeGenSDK)

type CodeGenSDK struct {
	*swagger.EnginesApi

	Configuration *Configuration
}

// Configuration configures the CodeGenSDK.
type Configuration struct {
	// EndpointURL is the URL of the ORY Keto API
	EndpointURL string
}

// NewCodeGenSDK instantiates a new CodeGenSDK instance or returns an error.
func NewCodeGenSDK(c *Configuration) (*CodeGenSDK, error) {
	if c.EndpointURL == "" {
		return nil, errors.New("Please specify the ORY Keto endpoint URL")
	}

	c.EndpointURL = strings.TrimRight(c.EndpointURL, "/")
	sdk := &CodeGenSDK{
		EnginesApi: swagger.NewEnginesApiWithBasePath(c.EndpointURL),
	}

	return sdk, nil
}
