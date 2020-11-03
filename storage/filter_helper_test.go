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
		{
			ID:      "role3",
			Members: []string{"mem5", "mem6"},
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
		{
			ID:        "policy3",
			Actions:   []string{"update"},
			Subjects:  []string{"mem5", "mem6"},
			Resources: []string{"res3", "res4"},
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
		// test that our changes to offset and limit didnt affect getting everything else
		ParamsMap{
			target:map[string][]string{},
			offset: 0,
			limit: 1,
		},
		// From this point on we set all limits to one to make it easier to get into a scenario where
		// the first page will be an empty array, but there's still a result on page 2, and try out the various scenarios
		// that will allow us to ensure that the filter doesnt behave that way anymore

		// test with single action and limit, that a result expected to be on page 2 with page 1 returning an
		// empty array is now on page 1
		ParamsMap{
			target:map[string][]string{"action":{"update"}},
			offset: 0,
			limit: 1,
		},
		// test with single subject and limit, that a result expected to be on page 2 with page 1 returning an
		// empty array is now on page 1
		// this also adds a single member so that a single loop is compatible with both testing role query and subject query.
		ParamsMap{
			target:map[string][]string{"subject":{"mem5"}, "member":{"mem5"}},
			offset: 0,
			limit: 1,
		},
		// test with single resource and limit, that a result expected to be on page 2 with page 1 returning an
		// empty array is now on page 1
		ParamsMap{
			target:map[string][]string{"resource":{"res3"}},
			offset: 0,
			limit: 1,
		},
		// combine all together, test that a result expected to be on page 2 with page 1 returning an
		// empty array is now on page 1
		ParamsMap{
			target:map[string][]string{"resource":{"res3"}, "subject":{"mem5"},"action":{"update"} },
			offset: 0,
			limit: 1,
		},
		// role testing
		ParamsMap{
			target:map[string][]string{"member": {"mem5"} },
			offset: 0,
			limit: 1,
		},
		// test advance offset
		ParamsMap{
			target:map[string][]string{"resource":{"res1"}},
			offset: 0,
			limit: 1,
		},
		ParamsMap{
			target:map[string][]string{"resource":{"res1"}},
			offset: 1,
			limit: 1,
		},
	}
	rolRes = []Roles{
		Roles{rolReq[0],rolReq[1]},
		Roles{rolReq[1]},
		Roles{},
		rolReq,
		rolReq,
		rolReq,
		rolReq,
		rolReq,
		Roles{rolReq[0]},
		Roles{rolReq[0]},
		Roles{rolReq[2]},
		// because we do not specify an explicit param for role, even though polRes returns roles for mem3,
		// role will simply return the first result for role.
		Roles{rolReq[0]},
		Roles{rolReq[0]},
		Roles{rolReq[2]},
		Roles{rolReq[0]},
		Roles{rolReq[1]},
	}
	polRes = []Policies{
		polReq,
		polReq,
		polReq,
		Policies{polReq[0], polReq[1]},
		Policies{polReq[1]},
		Policies{polReq[1]},
		Policies{},
		Policies{},
		Policies{polReq[0]},
		Policies{polReq[2]},
		Policies{polReq[2]},
		Policies{polReq[2]},
		Policies{polReq[2]},
		Policies{polReq[0]},
		Policies{polReq[0]},
		Policies{polReq[1]},
	}
)
