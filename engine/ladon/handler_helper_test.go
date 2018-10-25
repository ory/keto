package ladon

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
		req     Input
		allowed bool
	}{
		"regex": {
			{
				req: Input{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: true,
			},
			{
				req: Input{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: true,
			},
		},
		"exact": {
			{
				req: Input{
					Subject:  "alice",
					Resource: "other-thing",
					Action:   "create",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "delete",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "alice",
					Resource: "matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: true,
			},
			{
				req: Input{
					Subject:  "ken",
					Resource: "forbidden_matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: false,
			},
			{
				req: Input{
					Subject:  "ken",
					Resource: "allowed_matrix",
					Action:   "create",
					Context:  Context{},
				},
				allowed: true,
			},
		},
	}
	policies = map[string][]Policies{
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
		},
	}
)
