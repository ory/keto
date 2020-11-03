package storage

type ParamsMap struct {
	target map[string][]string
	offset int
	limit int
}
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
	paramsReq = []ParamsMap {
		ParamsMap{
			target: map[string][]string{"member": {"mem1"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target: map[string][]string{"member": {"mem2"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target: map[string][]string{"member": {"mem3"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target: map[string][]string{"action": {"create"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target: map[string][]string{"subject": {"mem3"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target:map[string][]string{"action": {"create"}, "subject": {"mem3"}, "resource": {"res2"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target:map[string][]string{"action": {"create"}, "subject": {"mem3"}, "resource": {"res3"}},
			offset: 0,
			limit: 100,
		},
		ParamsMap{
			target:map[string][]string{"action": {"delete"}},
			offset: 0,
			limit: 100,
		},
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
		Roles{},
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
		Policies{},
		Policies{},
	}
)
