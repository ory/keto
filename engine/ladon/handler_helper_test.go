package ladon

import "github.com/ory/keto/sdk/go/keto/swagger"

var (
	roles = map[string]Roles{
		"regex": {{
			ID:      "group1",
			Members: []string{"ken"},
		}, {
			ID:      "group2",
			Members: []string{"ken"},
		}, {
			ID:      "group3",
			Members: []string{"ken"},
		}},
		"exact": {{
			ID:      "group1",
			Members: []string{"ken"},
		}, {
			ID:      "group2",
			Members: []string{"ken"},
		}, {
			ID:      "group3",
			Members: []string{"ken"},
		}},
	}
	requests = map[string][]struct {
		req     swagger.OryAccessControlPolicyAllowedInput
		allowed bool
	}{
		"regex": {
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_subject",
					Action:   "create",
					Context: map[string]interface{}{
						"subject": map[string]string{
							"type": "EqualsSubjectCondition",
						},
					},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_subject",
					Action:   "create",
					Context: map[string]interface{}{
						"subject": map[string]string{
							"type": "EqualsSubjectCondition",
						},
					},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_condition",
					Action:   "create",
					Context: map[string]interface{}{
						"group": map[string]interface{}{
							"type": "StringEqualCondition",
							"options": map[string]string{
								"equals": "the-value-should-be-this",
							},
						},
					},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_condition",
					Action:   "create",
					Context: map[string]interface{}{
						"group": map[string]interface{}{
							"type": "StringEqualCondition",
							"options": map[string]string{
								"equals": "the-value-should-be-this",
							},
						},
					},
				},
				allowed: true,
			},
		},
		"exact": {
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  map[string]interface{}{},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_subject",
					Action:   "create",
					Context: map[string]interface{}{
						"subject": map[string]string{
							"type": "EqualsSubjectCondition",
						},
					},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_subject",
					Action:   "create",
					Context: map[string]interface{}{
						"subject": map[string]string{
							"type": "EqualsSubjectCondition",
						},
					},
				},
				allowed: true,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "forbidden_condition",
					Action:   "create",
					Context: map[string]interface{}{
						"group": map[string]interface{}{
							"type": "StringEqualCondition",
							"options": map[string]string{
								"equals": "the-value-should-be-this",
							},
						},
					},
				},
				allowed: false,
			},
			{
				req: swagger.OryAccessControlPolicyAllowedInput{
					Subject:  "ken",
					Resource: "allowed_condition",
					Action:   "create",
					Context: map[string]interface{}{
						"group": map[string]interface{}{
							"type": "StringEqualCondition",
							"options": map[string]string{
								"equals": "the-value-should-be-this",
							},
						},
					},
				},
				allowed: true,
			},
		},
	}
	policies = map[string]Policies{
		"regex": {
			Policy{
				ID:        "1",
				Subjects:  []string{"alice", "group1", "client"},
				Resources: []string{"matrix", "forbidden_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "2",
				Subjects:  []string{"siri"},
				Resources: []string{"<.*>"},
				Actions:   []string{"decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "3",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Deny,
			},
			Policy{
				ID:        "4",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "5",
				Subjects:  []string{"ken"},
				Resources: []string{"forbidden_subject", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "6",
				Subjects:  []string{"ken"},
				Resources: []string{"allowed_subject", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Allow,
			},
			Policy{
				ID:        "7",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_subject", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "8",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_subject", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Allow,
			},
			Policy{
				ID:        "9",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_condition", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"group": "the-value-should-be-this",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "10",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_condition", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"group": "the-value-should-be-this",
					},
				},
				Effect: Allow,
			},
		},
		"exact": {
			Policy{
				ID:        "1",
				Subjects:  []string{"alice", "group1", "client"},
				Resources: []string{"matrix", "forbidden_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "2",
				Subjects:  []string{"siri"},
				Resources: []string{""},
				Actions:   []string{"decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "3",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Deny,
			},
			Policy{
				ID:        "4",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_matrix", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Effect:    Allow,
			},
			Policy{
				ID:        "5",
				Subjects:  []string{"ken"},
				Resources: []string{"forbidden_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "6",
				Subjects:  []string{"ken"},
				Resources: []string{"allowed_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Allow,
			},
			Policy{
				ID:        "7",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "8",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"subject": "ken",
					},
				},
				Effect: Allow,
			},
			Policy{
				ID:        "9",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"group": "the-value-should-be-this",
					},
				},
				Effect: Deny,
			},
			Policy{
				ID:        "10",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_subject", "rn:hydra:token"},
				Actions:   []string{"create", "decide"},
				Conditions: []map[string]interface{}{
					{
						"group": "the-value-should-be-this",
					},
				},
				Effect: Allow,
			},
		},
	}
)
