package ory.glob

simple_policy = {
    "id": "4",
    "resources": [`articles:4`],
    "subjects": [`subjects:4`],
    "actions": [`actions:4`],
    "effect": "allow",
}

test_allow_policy {
    decide_allow([simple_policy], []) with input as {"resource": "articles:4", "subject": "subjects:4", "action": "actions:4"}
}

test_policy_must_match_resource_subject_and_action {
    not decide_allow([simple_policy], []) with input as {"resource": "articles:5", "subject": "subjects:4", "action": "actions:4"}
    not decide_allow([simple_policy], []) with input as {"resource": "articles:4", "subject": "subjects:5", "action": "actions:4"}
    not decide_allow([simple_policy], []) with input as {"resource": "articles:4", "subject": "subjects:4", "action": "actions:5"}
}

invalid_condition_policy = {
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
}

test_invalid_condition_policy {
    not decide_allow(invalid_condition_policy, []) with input as {"resource": "articles:5", "subject": "subjects:5", "action": "actions:5"}
}

group_policy = {
    "id": "6",
    "resources": [`articles:6`],
    "subjects": [`{subjects,groups}:6`],
    "actions": [`actions:6`],
    "effect": "allow"
}

test_allow_group_policy {
    decide_allow([group_policy], []) with input as {"resource": "articles:6", "subject": "subjects:6", "action": "actions:6"}
    decide_allow([group_policy], []) with input as {"resource": "articles:6", "subject": "groups:6", "action": "actions:6"}

    decide_allow([group_policy], [{"id": "groups:6", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "group-subject", "action": "actions:6"}

    not decide_allow([group_policy], [{"id": "groups:6", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "not-group-subject", "action": "actions:6"}
    not decide_allow([group_policy], [{"id": "not-groups", "members": ["group-subject"]}]) with input as {"resource": "articles:6", "subject": "group-subject", "action": "actions:6"}
}

deny_policies = [
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
]

test_deny_policy {
    not decide_allow(deny_policies, []) with input as {"resource": "articles:2", "subject": "subjects:2", "action": "actions:2"}
}

test_deny_overrides {
    not decide_allow(deny_policies, []) with input as {"resource": "articles:3", "subject": "subjects:3", "action": "actions:3"}
}

test_deny_without_match {
    not decide_allow(deny_policies, []) with input as {"resource": "unknown", "subject": "unknown", "action": "unknown", "context": {"unknown": "unknown"}}
}

condition_policy = {
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
}

test_with_condition {
    decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": "the-value-should-be-this"}}
    not decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": "not-the-value-should-be-this"}}
    not decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"not-foobar": "the-value-should-be-this"}}
    not decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {"foobar": 1234}}
    not decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1", "context": {}}
    not decide_allow([condition_policy], []) with input as {"resource": "articles:1", "subject": "subjects:1", "action": "actions:1"}
}

test_with_unknown_condition {
    not decide_allow([condition_policy], []) with input as {"resource": "articles:5", "subject": "subjects:5", "action": "actions:5", "context": {"foobar": {}}}
}

wildcard_policy = {
    "id": "7",
    # Allows single character
    "resources": [`articles:?:7`],
    # Allows any number of characters
    "subjects": [`subjects:*:7`],
    # Allows any number of characters spanning delimiters
    "actions": [`actions:**:7`],
    "effect": "allow"
}

test_allow_wildcards_single_char {
    decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:a:7"}
}

test_question_mark_exactly_one_character {
    not decide_allow([wildcard_policy], []) with input as {"resource": "articles::7", "subject": "subjects:a:7", "action": "actions:a:7"}
    not decide_allow([wildcard_policy], []) with input as {"resource": "articles:ab:7", "subject": "subjects:a:7", "action": "actions:a:7"}
}

test_star_respects_delimiters {
    decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:ab:7", "action": "actions:a:7"}
    not decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:7", "action": "actions:a:7"}
    not decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:b:7", "action": "actions:a:7"}
}

test_allow_star_star_with_delimiters {
    decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:ab:7"}
    decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:a:b:7"}
    decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions:7"}
}

test_star_star_requires_a_delimiter {
    not decide_allow([wildcard_policy], []) with input as {"resource": "articles:a:7", "subject": "subjects:a:7", "action": "actions7"}
}

list_policy = {
    "id": "8",
    # Character list (bat or cat)
    "resources": [`articles:[cb]at:8`],
    # Ranged character list (aat, bat, cat)
    "subjects": [`subjects:[a-c]at:8`],
    # Negated character lists and ranges
    "actions": [`actions:[!cb]a[!x-z]:8`],
    "effect": "allow"
}

test_allow_lists {
    decide_allow([list_policy], []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
    decide_allow([list_policy], []) with input as {"resource": "articles:bat:8", "subject": "subjects:bat:8", "action": "actions:dad:8"}
}

test_allow_character_list {
    not decide_allow([list_policy], []) with input as {"resource": "articles:dat:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
    not decide_allow([list_policy], []) with input as {"resource": "articles:cay:8", "subject": "subjects:cat:8", "action": "actions:tat:8"}
}

test_allow_character_ranges {
    not decide_allow([list_policy], []) with input as {"resource": "articles:cat:8", "subject": "subjects:dat:8", "action": "actions:tat:8"}
}

test_allow_negated_lists {
    not decide_allow([list_policy], []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:cat:8"}
    not decide_allow([list_policy], []) with input as {"resource": "articles:cat:8", "subject": "subjects:cat:8", "action": "actions:tay:8"}
}

# Various combinations to guard against interaction effects
combination_policy = {
    "id": "9",
    "resources": [`articles:foo**{bar,baz}:9`],
    "subjects": [`subjects:{?at,d*g}:9`],
    "actions": [`actions:{[cbm]at,d[!j-n]g}:9`],
    "effect": "allow"
}

test_allow_valid_combinations {
    decide_allow([combination_policy], []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
    decide_allow([combination_policy], []) with input as {"resource": "articles:foobarbaz:9", "subject": "subjects:doooog:9", "action": "actions:dig:9"}
    decide_allow([combination_policy], []) with input as {"resource": "articles:foo:bar::quux:baz:9", "subject": "subjects:dg:9", "action": "actions:dig:9"}
}

test_allow_combinations_star_star_preserves_surrounding_context {
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foo:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
}
test_allow_combinations_no_extra_chars {
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foobard:9", "subject": "subjects:cat:9", "action": "actions:cat:9"}
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:doog:9"}
}
test_allow_combinations_no_missing_chars {
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foobar:9", "subject": "subjects:at:9", "action": "actions:cat:9"}
}

test_allow_combinations_exactly_one_alternative {
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foobar:9", "subject": "subjects:catdog:9", "action": "actions:cat:9"}
    not decide_allow([combination_policy], []) with input as {"resource": "articles:foobar:9", "subject": "subjects:cat:9", "action": "actions:catdog:9"}
}
