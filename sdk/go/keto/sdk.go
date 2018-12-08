package keto

import (
	"strings"

	"github.com/pkg/errors"

	"github.com/ory/keto/sdk/go/keto/swagger"
)

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
