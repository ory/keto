package ory.regex

policies = [
    {
    	"id": "1",
        "resources": [`articles:1`],
        "subjects": [`subjects:1`],
        "actions": [`actions:1`],
        "effect": "allow",
        "conditions": {
            "foobar": {
                "type": "StringEqualCondition",
                "options": {
                    "equals": "the-value-should-be-this"
                }
            }
        }
    },
    {
    	"id": "2",
        "resources": [`articles:2`],
        "subjects": [`subjects:2`],
        "actions": [`actions:2`],
        "effect": "deny",
    },
    {
    	"id": "3-1",
        "resources": [`articles:3`],
        "subjects": [`subjects:3`],
        "actions": [`actions:3`],
        "effect": "allow",
    },
    {
    	"id": "3-2",
        "resources": [`articles:3`],
        "subjects": [`subjects:3`],
        "actions": [`actions:3`],
        "effect": "deny",
    },
    {
    	"id": "3-3",
        "resources": [`articles:3`],
        "subjects": [`subjects:3`],
        "actions": [`actions:3`],
        "effect": "allow",
    },
    {
    	"id": "4",
        "resources": [`articles:<4|44>`],
        "subjects": [`subjects:<4|44>`],
        "actions": [`actions:<4|44>`],
        "effect": "allow",
    },
    {
    	"id": "5",
        "resources": [`articles:5`],
        "subjects": [`subjects:5`],
        "actions": [`actions:5`],
        "effect": "allow",
        "conditions": {
            "foobar": {
                "type": "InvalidCondition"
            }
        }
    },
    {
    	"id": "6",
        "resources": [`articles:6`],
        "subjects": [`<subjects|groups>:6`],
        "actions": [`actions:6`],
        "effect": "allow"
    },
]

test_allow_policy {
    decide_allow(policies, []) with input as {"resource": "articles:4", "subject": "subjects:4", "action": "actions:4"}
    decide_allow(policies, []) with input as {"resource": "articles:44", "subject": "subjects:44", "action": "actions:44"}
    not decide_allow(policies, []) with input as {"resource": "articles:454", "subject": "subjects:454", "action": "actions:454"}
}

test_allow_group_policy {
    decide_allow(policies, []) with input as {"resource": "articles:6", "subject": "subjects:6", "action": "actions:6"}
    decide_allow(policies, []) with input as {"resource": "articles:6", "subject": "groups:6", "action": "actions:6"}

    decide_allow(policies, [{"id": "groups:6", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "group-subject", "action": "actions:6"}

    not decide_allow(policies, [{"id": "groups:6", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "not-group-subject", "action": "actions:6"}
    not decide_allow(policies, [{"id": "not-groups", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "group-subject", "action": "actions:6"}
}

test_deny_policy {
    not decide_allow(policies, []) with input as {"resource": "articles:2", "subject": "subjects:2", "action": "actions:2"}
}

test_deny_overrides {
    not decide_allow(policies, []) with input as {"resource": "articles:3", "subject": "subjects:3", "action": "actions:3"}
}

test_deny_without_match {
	not decide_allow(policies, []) with input as {"resource": "unknown", "subject": "unknown", "action": "unknown", "context": {"unknown": "unknown"}}
}

test_with_condition {
    decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": "the-value-should-be-this"}}
    not decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": "not-the-value-should-be-this"}}
    not decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"not-foobar": "the-value-should-be-this"}}
    not decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": 1234}}
    not decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {}}
    not decide_allow(policies, []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1"}
}

test_with_unknown_condition {
    not decide_allow(policies, []) with input as {"resource": "articles:5", "subject": "subjects:5", "action": "actions:5", "context": {"foobar": {}}}
}
