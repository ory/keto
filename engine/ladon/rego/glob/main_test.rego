package ory.glob

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
        "resources": [`articles:?4`],
        "subjects": [`subjects:?4`],
        "actions": [`actions:?4`],
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
        "subjects": [`{subjects,groups}:6`],
        "actions": [`actions:6`],
        "effect": "allow"
    },
    {
        "id": "7",
        "resources": [`articles:?:7`],
        "subjects": [`subjects:*:7`],
        "actions": [`actions:**:7`],
        "effect": "allow"
    },
    {
        "id": "8",
        "resources": [`articles:[cb]at:8`],
        "subjects": [`subjects:[a-c]at:8`],
        "actions": [`actions:[!cb]a[!x-z]:8`],
        "effect": "allow"
    },
    {
        "id": "9",
        "resources": [`articles:foo**{bar,baz}:9`],
        "subjects": [`subjects:{?at,d*g}:9`],
        "actions": [`actions:{[cbm]at,d[!j-n]g}:9`],
        "effect": "allow"
    },
]

test_allow_policy {
    decide_allow(policies, []) with input as {"resource": "articles:44", "subject": "subjects:44", "action": "actions:44"}
    decide_allow(policies, []) with input as {"resource": "articles:54", "subject": "subjects:54", "action": "actions:54"}
    not decide_allow(policies, []) with input as {"resource": "articles:45", "subject": "subjects:45", "action": "actions:45"}
    not decide_allow(policies, []) with input as {"resource": "articles:4", "subject": "subjects:4", "action": "actions:4"}
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

test_allow_wildcards {
    decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:a:7"}
    decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:ab:7", "action": "actions:a:7"}
    decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:ab:7"}
    decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:a:b:7"}
    decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:7"}

    not decide_allow(policies, []) with input as {"resource": "articles:ab:7", "subject": "subjects:a:7", "action": "actions:a:7"}
    not decide_allow(policies, []) with input as {"resource": "articles:a:7", "subject": "subjects:a:b:7", "action": "actions:a:7"}
}

test_allow_lists {
    decide_allow(policies, []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
    decide_allow(policies, []) with input as {"resource": "articles:bat:8", "subject": "subjects:bat:8", "action": "actions:dad:8"}
    not decide_allow(policies, []) with input as {"resource": "articles:dat:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
    not decide_allow(policies, []) with input as {"resource": "articles:cat:8", "subject": "subjects:dat:8", "action": "actions:tat:8"}
    not decide_allow(policies, []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:cat:8"}
    not decide_allow(policies, []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:tay:8"}
    not decide_allow(policies, []) with input as {"resource": "articles:cay:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
}

test_allow_combinations {
    decide_allow(policies, []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
    decide_allow(policies, []) with input as {"resource": "articles:foobarbaz:9", "subject": "subjects:doooog:9", "action": "actions:dig:9"}
    decide_allow(policies, []) with input as {"resource": "articles:foo:bar::quux:baz:9", "subject": "subjects:dg:9", "action": "actions:dig:9"}

    not decide_allow(policies, []) with input as {"resource": "articles:foo:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
    not decide_allow(policies, []) with input as {"resource": "articles:foobard:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
    not decide_allow(policies, []) with input as {"resource": "articles:foobar:9", "subject": "subjects:at:9", "action": "actions:cat:9"}
    not decide_allow(policies, []) with input as {"resource": "articles:foobar:9", "subject": "subjects:catdog:9", "action": "actions:cat:9"}
    not decide_allow(policies, []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:catdog:9"}
    not decide_allow(policies, []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:doog:9"}
}
