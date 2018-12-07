package ory.condition

eval_condition("EqualsSubjectCondition", request, options, key) {
    request.context[key] == request.subject
}

test_condition_equals_subject {
    eval_condition("EqualsSubjectCondition", { "subject": "some-subject", "context": { "foobar": "some-subject" } }, {}, "foobar")
    not eval_condition("EqualsSubjectCondition", { "subject": "some-subject", "context": { "foobar": "not-some-subject" } }, {}, "foobar")
    not eval_condition("EqualsSubjectCondition", { "subject": "some-subject", "context": { "not-foobar": "some-subject" } }, {}, "foobar")
}
