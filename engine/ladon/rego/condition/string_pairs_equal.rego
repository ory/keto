package ory.condition

eval_condition("StringPairsEqualCondition", request, options, key) {
    cast_array(request.context[key], context)
    count(context, c)
    c > 0

    not any_not_string_pair(context)
}

any_not_string_pair(v) {
    cast_array(v[_], vv)
    not is_string_pair(vv)
}

is_string_pair(v) {
    count(v, c)
    c == 2
    v[0] == v[1]
}

test_condition_string_pairs_eqal {
    not eval_condition("StringPairsEqualCondition", { "context": { } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [] } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [[]] } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1"]] } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1", "2"]] } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1", "1", "2"]] } }, {}, "foobar")
    not eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1", "1"], ["2", "3"]] } }, {}, "foobar")
    eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1", "1"], ["2", "2"]] } }, {}, "foobar")
    eval_condition("StringPairsEqualCondition", { "context": { "foobar": [["1", "1"]] } }, {}, "foobar")
}
