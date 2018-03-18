package keto

import "github.com/ory/keto/sdk/go/keto/swagger"

type PolicySDK interface {
	CreatePolicy(body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error)
	DeletePolicy(id string) (*swagger.APIResponse, error)
	GetPolicy(id string) (*swagger.Policy, *swagger.APIResponse, error)
	ListPolicies(offset int64, limit int64) ([]swagger.Policy, *swagger.APIResponse, error)
	UpdatePolicy(id string, body swagger.Policy) (*swagger.Policy, *swagger.APIResponse, error)
}
