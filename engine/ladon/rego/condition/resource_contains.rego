package ory.condition

eval_condition("ResourceContainsCondition", request, options, key) {
    value := cast_string_empty(options, "value")
    delimiter := cast_string_empty(options, "delimiter")

	needle := concat("", [delimiter, value, delimiter])
	haystack := concat("", [delimiter, request.resource, delimiter])

	contains(haystack, needle) == true
}

test_condition_resource_contains {
    not eval_condition("ResourceContainsCondition", { "resource": "foo:bar" }, { "delimiter": ":", "value": "foo:ba" }, "")

    eval_condition("ResourceContainsCondition", { "resource": "foo:bar" }, { "delimiter": ":", "value": "foo:bar" }, "")
    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "delimiter": ":", "value": "foo:bar" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "delimiter": ":", "value": "foo:baz" }, "")

    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "delimiter": ":", "value": "bar:baz" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "delimiter": ":", "value": "foo:baz" }, "")
    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "delimiter": ":", "value": "bar" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "baz:foo:baz" }, { "delimiter": ":", "value": "bar" }, "")

    eval_condition("ResourceContainsCondition", { "resource": "foo:bar" }, { "value": "foo:ba" }, "")
    eval_condition("ResourceContainsCondition", { "resource": "foo:bar" }, { "value": "foo:bar" }, "")
    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "value": "foo:bar" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "foo:baz" }, { "value": "foo:bar" }, "")

    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "value": "bar:baz" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "value": "foo:baz" }, "")
    eval_condition("ResourceContainsCondition", { "resource": "foo:bar:baz" }, { "value": "bar" }, "")
    not eval_condition("ResourceContainsCondition", { "resource": "baz:foo:baz" }, { "value": "bar" }, "")

    not eval_condition("ResourceContainsCondition", { "resource": "abc" }, { "value": "", "delimiter": ":" }, "")
}
