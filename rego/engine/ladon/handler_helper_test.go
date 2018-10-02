package ladon

import (
	"github.com/ory/ladon"
)

var (
	roles = map[string]Roles{
		"regex": {{
			ID:      "group1",
			Members: []string{"ken"},
		}},
	}
	requests = map[string][]struct {
		req     input
		allowed bool
	}{
		"regex": {
			{
				req: input{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  ladon.Context{},
				},
				allowed: false,
			},
			{
				req: input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  ladon.Context{},
				},
				allowed: false,
			},
			{
				req: input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  ladon.Context{},
				},
				allowed: true,
			},
			{
				req: input{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  ladon.Context{},
				},
				allowed: false,
			},
			{
				req: input{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  ladon.Context{},
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
				Effect:    ladon.AllowAccess,
			},
			Policy{
				ID:        "2",
				Subjects:  []string{"siri"},
				Resources: []string{"<.*>"},
				Actions:   []string{"decide"},
				Effect:    ladon.AllowAccess,
			},
			Policy{
				ID:        "3",
				Subjects:  []string{"group1"},
				Resources: []string{"forbidden_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    ladon.DenyAccess,
			},
			Policy{
				ID:        "4",
				Subjects:  []string{"group1"},
				Resources: []string{"allowed_matrix", "rn:hydra:token<.*>"},
				Actions:   []string{"create", "decide"},
				Effect:    ladon.AllowAccess,
			},
		},
	}
)
