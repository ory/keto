package ory.condition

eval_condition("BooleanCondition", request, options, key) {
    is_boolean(request.context[key], output)
    output == true

    request.context[key] == options.value
}

test_condition_boolean {
    eval_condition("BooleanCondition", { "context": {"foobar": false } }, { "value": false }, "foobar")
    eval_condition("BooleanCondition", { "context": {"foobar": true } }, { "value": true }, "foobar")

    not eval_condition("BooleanCondition", { "context": {"foobar": false } }, { "value": true }, "foobar")
    not eval_condition("BooleanCondition", { "context": {"foobar": true } }, { "value": false }, "foobar")
    not eval_condition("BooleanCondition", { "context": {"not-foobar": true } }, { "value": false }, "foobar")
    not eval_condition("BooleanCondition", { "context": {"foobar": true } }, { "not-value": false }, "foobar")
    not eval_condition("BooleanCondition", { "context": {"not-foobar": true } }, { "not-value": false }, "foobar")
}
