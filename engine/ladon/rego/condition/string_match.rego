package ory.condition

eval_condition("StringMatchCondition", request, options, key) {
    re_match(options.matches, request.context[key]) == true
}

test_condition_string_match {
    eval_condition("StringMatchCondition", { "context": {"foobar": "abc"} }, { "matches": ".*" }, "foobar")
    eval_condition("StringMatchCondition", { "context": {"foobar": "abc"} }, { "matches": "abc.*" }, "foobar")

    not eval_condition("StringMatchCondition", { "context": {"not-foobar": "abc" } }, { "matches": ".+" }, "foobar")
    not eval_condition("StringMatchCondition", { "context": {"foobar": "abc" } }, { "matches": "abc.+" }, "foobar")
}
