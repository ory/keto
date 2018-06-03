package keto

import (
	"strings"

	"github.com/ory/keto/sdk/go/keto/swagger"
	"github.com/pkg/errors"
)

type SDK interface {
	RoleSDK
	WardenSDK
	PolicySDK
}

var testSDK SDK = new(CodeGenSDK)

type CodeGenSDK struct {
	*swagger.RoleApi
	*swagger.WardenApi
	*swagger.PolicyApi

	Configuration *Configuration
}

// Configuration configures the CodeGenSDK.
type Configuration struct {
	// EndpointURL should point to the url of ORY Keto, for example: http://localhost:4466
	EndpointURL string
}

// CodeGenSDK instantiates a new CodeGenSDK instance or returns an error.
func NewCodeGenSDK(c *Configuration) (*CodeGenSDK, error) {
	if c.EndpointURL == "" {
		return nil, errors.New("Please specify the ORY Keto endpoint URL")
	}

	c.EndpointURL = strings.TrimRight(c.EndpointURL, "/")
	sdk := &CodeGenSDK{
		RoleApi:   swagger.NewRoleApiWithBasePath(c.EndpointURL),
		WardenApi: swagger.NewWardenApiWithBasePath(c.EndpointURL),
		PolicyApi: swagger.NewPolicyApiWithBasePath(c.EndpointURL),
	}

	return sdk, nil
}
