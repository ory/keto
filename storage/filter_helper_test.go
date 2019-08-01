package storage

var (
	rolReq = Roles{
		{
			ID:      "role1",
			Members: []string{"mem1"},
		},
		{
			ID:      "role2",
			Members: []string{"mem1", "mem2"},
		},
	}
	polReq = Policies{
		{
			ID:        "policy1",
			Actions:   []string{"create"},
			Subjects:  []string{"mem1", "mem2"},
			Resources: []string{"res1", "res2"},
		},
		{
			ID:        "policy2",
			Actions:   []string{"create"},
			Subjects:  []string{"mem3", "mem4"},
			Resources: []string{"res1", "res2"},
		},
	}
	paramsReq = []map[string][]string{
		{"members": {"mem1"}},
		{"members": {"mem2"}},
		{"members": {"mem3"}},
		{"actions": {"create"}},
		{"subjects": {"mem3", "mem4"}},
		{"actions": {"create"}, "subjects": {"mem3"}, "resources": {"res2"}},
		{"actions": {"create"}, "subjects": {"mem3"}, "resources": {"res3"}},
		{"actions": {"delete"}},
	}
	rolRes = []Roles{
		Roles{
			{
				ID:      "role1",
				Members: []string{"mem1"},
			},
			{
				ID:      "role2",
				Members: []string{"mem1", "mem2"},
			},
		},
		Roles{
			{
				ID:      "role2",
				Members: []string{"mem1", "mem2"},
			},
		},
		nil,
		rolReq,
		rolReq,
		rolReq,
		rolReq,
		rolReq,
	}
	polRes = []Policies{
		polReq,
		polReq,
		polReq,
		Policies{
			{
				ID:        "policy1",
				Actions:   []string{"create"},
				Subjects:  []string{"mem1", "mem2"},
				Resources: []string{"res1", "res2"},
			},
			{
				ID:        "policy2",
				Actions:   []string{"create"},
				Subjects:  []string{"mem3", "mem4"},
				Resources: []string{"res1", "res2"},
			},
		},
		Policies{
			{
				ID:        "policy2",
				Actions:   []string{"create"},
				Subjects:  []string{"mem3", "mem4"},
				Resources: []string{"res1", "res2"},
			},
		},
		Policies{
			{
				ID:        "policy2",
				Actions:   []string{"create"},
				Subjects:  []string{"mem3", "mem4"},
				Resources: []string{"res1", "res2"},
			},
		},
		nil,
		nil,
	}
)
