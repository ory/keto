package ladon

import (
	"github.com/ory/keto/internal/httpclient/models"

	kstorage "github.com/ory/keto/storage"
)

var (
	roles = map[string]kstorage.Roles{
		"regex": {{
			ID:          "group1",
			Description: "group1 description",
			Members:     []string{"ken"},
		}, {
			ID:          "group2",
			Description: "group12 description",
			Members:     []string{"ken"},
		}, {
			ID:          "group3",
			Description: "group3 description",
			Members:     []string{"ben"},
		}},
		"exact": {{
			ID:          "group1",
			Description: "group1 description",
			Members:     []string{"ken"},
		}, {
			ID:          "group2",
			Description: "group2 description",
			Members:     []string{"ken"},
		}, {
			ID:          "group3",
			Description: "group3 description",
			Members:     []string{"ben"},
		}},
	}
	requests = map[string][]struct {
		req     models.OryAccessControlPolicyAllowedInput
		allowed bool
	}{
		"regex": {
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
		},
		"exact": {
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: models.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
		},
	}
	policies = map[string]kstorage.Policies{
		"regex": {
			kstorage.Policy{
				ID:        "1",
				Subjects:  []string{"alice", "group1", "client"},
				Resources: []string{"matrix", "forbidden_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			kstorage.Policy{
				ID:        "2",
				Subjects:  []string{"siri"},
				Resources: []string{"<.*>"},
				Actions:   []string{"decide"},
				Effect:    Allow,
			},
			kstorage.Policy{
				ID:        "3",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Deny,
			},
			kstorage.Policy{
				ID:        "4",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
		},
		"exact": {
			kstorage.Policy{
				ID:        "1",
				Subjects:  []string{"alice", "group1", "client"},
				Resources: []string{"matrix", "forbidden_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			kstorage.Policy{
				ID:        "2",
				Subjects:  []string{"siri"},
				Resources: []string{""},
				Actions:   []string{"decide"},
				Effect:    Allow,
			},
			kstorage.Policy{
				ID:        "3",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Deny,
			},
			kstorage.Policy{
				ID:        "4",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
		},
	}
)
