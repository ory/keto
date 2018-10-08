package ory.condition

all_conditions_true(policy) {
    not any_condition_false(policy)
}

any_condition_false(policy) {
    c := policy.conditions[condition_key]
    not condition_true(policy, c, condition_key)
}

condition_true(policy, c, condition_key) {
    eval_condition(c.type, input, c.options, condition_key)
} {
    false
}
