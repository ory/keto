package ory.condition

eval_condition("CIDRCondition", request, options, key) {
    net.cidr_overlap(options.cidr, request.context[key], output)
    output == true
}

test_condition_boolean {
    eval_condition("CIDRCondition", { "context": {"foobar": "192.168.178.0" } }, { "cidr": "192.168.178.0/16" }, "foobar")
    eval_condition("CIDRCondition", { "context": {"foobar": "192.168.178.1" } }, { "cidr": "192.168.178.0/16" }, "foobar")

    not eval_condition("CIDRCondition", { "context": {"foobar": "92.168.178.1" } }, { "cidr": "192.168.178.0/16" }, "foobar")
    not eval_condition("CIDRCondition", { "context": {"foobar": "192.168.178.1" } }, { "cidr": "192.168.178.0/16" }, "foobar2")
    not eval_condition("CIDRCondition", { "context": {"foobar2": "192.168.178.1" } }, { "cidr": "192.168.178.0/16" }, "foobar")
}
