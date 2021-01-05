package ory.condition

eval_condition("TimeInterval", request, options, key) {
    is_number(request.context[key], output)
    output == true

    is_number(options.after, aok)
    aok == true

    is_number(options.before, bok)
    bok == true

    options.after <= request.context[key]
    request.context[key]  < options.before
}

test_condition_time_interval {
    eval_condition("TimeInterval", { "context": {"time": 2} }, { "after": 1, "before":5 }, "time")
    eval_condition("TimeInterval", { "context": {"time": 5} }, { "after": 5, "before":20 }, "time")

    not eval_condition("TimeInterval", { "context": {"time": 4} }, { "after": 5, "before":20 }, "time")
    not eval_condition("TimeInterval", { "context": {"time": 20} }, { "after": 5, "before":20 }, "time")
}