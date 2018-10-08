package ory.exact

import data.store.ory.exact as store
import data.ory.core as core
import data.ory.condition as condition
import input as request

default allow = false

allow {
    decide_allow(store.policies, store.roles)
}

decide_allow(policies, roles) {
	effects := [effect | effect := policies[i].effect
			policies[i].resources[_] == request.resource
			match_subjects(policies[i].subjects, roles, request.subject)
			policies[i].actions[_] == request.action
			condition.all_conditions_true(policies[i])
		]

    count(effects, c)
    c > 0

	core.effect_allow(effects)
}

match_subjects(matches, roles, subject) {
    matches[_] == subject
} {
    r := core.role_ids(roles, subject)
    matches[_] == r[_]
}
