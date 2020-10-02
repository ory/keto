package storage

var (
	rolReq = []Role{
		{
			ID:      "role1",
			Members: []string{"mem1"},
		},
		{
			ID:      "role2",
			Members: []string{"mem1", "mem2"},
		},
	}
	polReq = []Policy{
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
		{"member": {"mem1"}},
		{"member": {"mem2"}},
		{"member": {"mem3"}},
		{"action": {"create"}},
		{"subject": {"mem3"}},
		{"action": {"create"}, "subject": {"mem3"}, "resource": {"res2"}},
		{"action": {"create"}, "subject": {"mem3"}, "resource": {"res3"}},
		{"action": {"delete"}},
	}
	rolRes = []Roles{
		[]Role{
			{
				ID:      "role1",
				Members: []string{"mem1"},
			},
			{
				ID:      "role2",
				Members: []string{"mem1", "mem2"},
			},
		},
		[]Role{
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
		[]Policy{
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
		[]Policy{
			{
				ID:        "policy2",
				Actions:   []string{"create"},
				Subjects:  []string{"mem3", "mem4"},
				Resources: []string{"res1", "res2"},
			},
		},
		[]Policy{
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
