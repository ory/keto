package ladon

import data
import input as request

default allow_exact = false

allow_exact {
    decide_allow_exact(data.policies_exact)
}

decide_allow_exact(policies) {
	effects := [effect | effect := policies[i].effect
			policies[i].resources[_] == request.resource
			policies[i].subjects[_] == request.subject
			policies[i].actions[_] == request.action
			all_conditions_true(policies[i])
		]

    count(effects, c)
    c > 0

	effect_allow(effects)
}
