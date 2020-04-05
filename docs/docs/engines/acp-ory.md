---
id: acp-ory
title: ORY Access Control Policies
---

At ORY, we use an Access Control Policy DSL modeled after AWS IAM Policies.
These policies define `effects` for `subjects` who perform `actions` on
`resources`. For example, `Alice` (subject aka identity aka user) is `allowed`
(effect) to `delete` (action) blog article with ID `my-first-blog-post`
(`resource`). This is comparable to how ACLs work:

```json
{
  "subjects": ["alice"],
  "resources": ["blog_posts:my-first-blog-post"],
  "actions": ["delete"],
  "effect": "allow"
}
```

The policy above allows `Alice` to `delete` `blog_posts:my-first-blog-post`. We
could apply this policy to more subjects and also more actions or resources, if
we want to:

```json
{
  "subjects": ["alice", "bob"],
  "resources": [
    "blog_posts:my-first-blog-post",
    "blog_posts:2",
    "blog_posts:3"
  ],
  "actions": ["delete", "create", "read", "modify"],
  "effect": "allow"
}
```

Well, this looks like ACL in disguise so far. So what's different?

## Precedence

The first difference is that we can explicitly deny access:

```json
{
  "subjects": ["peter"],
  "resources": [
    "blog_posts:my-first-blog-post",
    "blog_posts:2",
    "blog_posts:3"
  ],
  "actions": ["delete", "create", "read", "modify"],
  "effect": "deny"
}
```

The policy decision point (the one checking if something is allowed or not)
applies the following rule set when deciding if something is allowed or not:

1. If a policy for a given subject, action, and resource matches, and the effect
   is `deny`, the request is always denied.
2. If no policy with effect `deny` matches, and at least one policy with effect
   `allow`, the request is allowed.
3. If no policies match at all, the request is denied.

## Pattern Matching Strategies

ORY Keto has implements several ORY ACP pattern matching strategies.

### Case Sensitive Equality

The easiest pattern matching strategy is the case sensitive equality check. This
strategy simply checks if two strings are exactly the same. Assuming a policy
defines `{"subjects": ["alice", "boB"] }`, then it will match exactly subjects
`alice` and `boB`.

### Glob Pattern Matching

ORY Keto supports matching URNs with glob pattern matching. Policy

```json
{
  "subjects": ["users:*"],
  "actions": ["get", "create"],
  "resources": ["resources:articles:*", "resources:{accounts,profiles}:*"],
  "effect": "allow"
}
```

for example will match the following inputs:

```json
{
  "subject": "users:maria",
  "action": "get",
  "resource": "resources:profiles:foo"
}
```

The `:` is a delimiter in ORY Access Control Policies. Other supported syntax
is:

- **single symbol wildcard:** `?at` matches `cat` and `bat` but not `at`
- **wildcard:** `foo:*:bar` matches `foo:baz:bar` and `foo:zab:bar` but not
  `foo:bar` nor `foo:baz:baz:bar`
- **super wildcard:** `foo:**:bar` matches `foo:baz:baz:bar`, `foo:baz:bar`, and
  `foo:bar`, but not `foobar` or `foo:baz`
- **character list:** `[cb]at` matches `cat` and `bat` but not `mat` nor `at`.
- **negated character list:** `[!cb]at` matches `tat` and `mat` but not `cat`
  nor `bat`.
- **ranged character list:** `[a-c]at` `cat` and `bat` but not `mat` nor `at`.
- **negated ranged character list:** `[!a-c]at` matches `mat` and `tat` but not
  `cat` nor `bat`.
- **alternatives list:** `{cat,bat,[mt]at}` matches `cat`, `bat`, `mat`, `tat`
  and nothing else.
- **backslash:** `foo\\bar` matches `foo\bar` and nothing else. `foo\bar`
  matches `foobar` and nothing else. `foo\*bar` matches `foo*bar` and nothing
  else. Please note that when using JSON you need to double escape backslashes:
  `foo\\bar` becomes `{"...": "foo\\\\bar"}`.

The pattern syntax is:

```
  pattern:
      { term }

  term:
      `*`         matches any sequence of non-separator characters
      `**`        matches any sequence of characters
      `?`         matches any single non-separator character
      `[` [ `!` ] { character-range } `]`
                  character class (must be non-empty)
      `{` pattern-list `}`
                  pattern alternatives
      c           matches character c (c != `*`, `**`, `?`, `\`, `[`, `{`, `}`)
      `\` c       matches character c

  character-range:
      c           matches character c (c != `\\`, `-`, `]`)
      `\` c       matches character c
      lo `-` hi   matches character c for lo <= c <= hi

  pattern-list:
      pattern { `,` pattern }
                  comma-separated (without spaces) patterns
```

### Regular Expressions

ORY Keto also allows pattern matching with regular expressions. This depend on
how you name your subjects, resources, and actions. More on that topic in the
[Best Practices](#best-practices) section.

```json
{
  "subjects": ["users:<.*>"]
}
```

In this example, the (incomplete) policy would match every subject that is
prefixed with `users:`, so for example `users:alice`, `users:bob`. ORY Ladon and
ORY Keto delimit regular expressions with `<` and `>`. For example, `"users:.*"`
is not a valid regular expression, just a simple string.

The next example will allow all subjects with prefix `user:` to read
(`actions:read`) all resources that match `resources:blog_posts:<[0-9]+>` (e.g.
`resources:blog_posts:1234` but not `resources:blog_posts:abcde`):

```json
{
  "subjects": ["users:<.*>"],
  "resources": ["resources:blog_posts:<[0-9]+>"],
  "actions": ["actions:read"],
  "effect": "allow"
}
```

### Computational Overhead

Some pattern matching strategies can introduce computational complexity.
Consider the performance implications when choosing an approach:

- Case sensitive equality: no computational overhead
- Glob pattern matching: little computational overhead
- Regex: considerable computational overhead

## Conditions

So far, we covered that an ORY ACP applies to a list of `subjects`, `resources`,
and `actions`. Conditions narrow down the use cases in which a certain ACP
applies. A condition may, for example, mandate that the IP Address of the client
making the request has to match `192.168.0.0/16` or that the subject is also the
owner of the resource. Here is an example for the former condition:

```json
{
  "description": "One policy to rule them all",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "remoteIPAddress": {
      "type": "CIDRCondition",
      "options": {
        "cidr": "192.168.0.0/16"
      }
    }
  }
}
```

Conditions are a part of policies. They determine if a policy can decide the
current access request in the current context. Context is the larger environment
in which the access request happens. A condition has this JSON format:

```json
{
  "subjects": ["..."],
  "actions": ["..."],
  "effect": "allow",
  "resources": ["..."],
  "conditions": {
    "this-key-will-be-matched-with-the-context": {
      "type": "SomeConditionType",
      "options": {
        "some": "configuration options set by the condition type"
      }
    }
  }
}
```

Conditions are functions returning true or false given a context. Because
conditions implement logic, they must be programmed. ORY Keto provides the
following commonly used conditions out of the box. You can improve or extend
them.

### CIDR Condition

The CIDR condition matches CIDR IP Ranges. A possible policy definition could
look like this:

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "remoteIPAddress": {
      "type": "CIDRCondition",
      "options": {
        "cidr": "192.168.0.0/16"
      }
    }
  }
}
```

The following access request would be allowed.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "remoteIPAddress": "192.168.0.5"
  }
}
```

The next access request would be denied as the condition is not fulfilled and
thus no policy matches.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "remoteIPAddress": "255.255.0.0"
  }
}
```

The next access request would also be denied as the context is not using the key
`remoteIPAddress` but instead `someOtherKey`.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "someOtherKey": "192.168.0.5"
  }
}
```

##### String Equal Condition

This condition matches if the value passed in the access request's context is
identical with the string defined in the condition.

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "myKey": {
      "type": "StringEqualCondition",
      "options": {
        "equals": "expected-value"
      }
    }
  }
}
```

The following access request would be allowed.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "myKey": "expected-value"
  }
}
```

The following access request would be denied.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "meKey": "another-value"
  }
}
```

### String Match Condition

This condition applies when the value passed in the access request's context
matches the regular expression in the condition.

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "someKeyName": {
      "type": "StringMatchCondition",
      "options": {
        "matches": "foo.+"
      }
    }
  }
}
```

The following access request would be allowed.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "someKeyName": "foo-bar"
  }
}
```

The following access request would be denied.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "someKeyName": "bar"
  }
}
```

### Subject Condition

This condition matches when the access request's subject is identical with the
string specified in the condition.

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "owner": {
      "type": "EqualsSubjectCondition",
      "options": {}
    }
  }
}
```

The following access request would be allowed.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "owner": "users:maria"
  }
}
```

The following access request would be denied.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "owner": "another-user"
  }
}
```

This condition makes sense together with access tokens, where the subject is
extracted from the token.

### String Pairs Equal Condition

This condition matches when the value passed in the access request's context
contains two-element arrays and both elements in each pair are equal.

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:maria"],
  "actions": ["delete", "create", "update"],
  "effect": "allow",
  "resources": ["resources:articles:<.*>"],
  "conditions": {
    "someKey": {
      "type": "StringPairsEqualCondition",
      "options": {}
    }
  }
}
```

The following access request would be allowed.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "someKey": [
      ["foo", "foo"],
      ["bar", "bar"]
    ]
  }
}
```

The following access request would be denied.

```json
{
  "subject": "users:maria",
  "action": "delete",
  "resource": "resources:articles:12345",
  "context": {
    "someKey": [["foo", "bar"]]
  }
}
```

## Roles

Similar to RBAC, ORY ACPs support the concept of roles. This feature allows
grouping a number of subjects under the same role. Whenever making a request to
the Allowed API, it will check the roles of a subject (if there are any) and use
them when looking up the `subjects` field.

Assuming the following policies:

```json
{
  "subjects": ["bob"],
  "resources": ["blog_posts:my-first-blog-post"],
  "actions": ["create"],
  "effect": "allow"
}
```

```json
{
  "subjects": ["admin"],
  "resources": ["blog_posts:my-first-blog-post"],
  "actions": ["delete"],
  "effect": "allow"
}
```

As you can see, `bob` is allowed to create resource
`blog_posts:my-first-blog-post` and `admin` is allowed to delete it. Making the
following request to the Allowed API

```
{
  "subject": "bob",
  "action" : "delete",
  "resource": "blog_posts:my-first-blog-post"
}
```

will return `{ "allowed": false }` while this request

```
{
  "subject": "admin",
  "action" : "delete",
  "resource": "blog_posts:my-first-blog-post"
}
```

will return `{ "allowed": true }`.

## Implementation Status

ORY Access Control Policies (regex, glob, equality) are first-class citizens.

## Best Practices

This sections gives an overview of best practices for access control policies we
developed over the years at ORY.

### URNs

> “There are only two hard things in Computer Science: cache invalidation,
> naming things, and off-by-one errors.” -- Phil Karlton

URN naming is as hard as naming API endpoints. Thankfully, doing the latter
typically provides a solution for the former as well.

### Scope the Organization Name

It is good security practice is to prefix resource names with a domain that
represents the organization creating the software.

- **Do not:** `my-resource`
- **Do:** `myorg.com:my-resource`

### Scope Actions, Resources and Subjects

Provide a scope for actions, resources, and subjects to prevent name collisions:

- **Do not:** `myorg.com:<subject-id>`, `myorg.com:<resource-id>`,
  `myorg.com:<action-id>`
- **Do:** `myorg.com:subjects:<subject-id>`,
  `myorg.com:resources:<resource-id>`, `myorg.com:actions:<action-id>`
- **Do:** `subjects:myorg.com:<subject-id>`,
  `resources:myorg.com:<resource-id>`, `actions:myorg.com:<action-id>`

### Multi-Tenant Systems

Multi-tenant systems typically have resources which should not be accessed by
other tenants in the system. This can be achieved by adding the tenant id to the
URN:

- **Do:** `resources:myorg.com:tenants:<tenant-id>:<resource-id>`

Some environments have organizations and projects belonging to those
organizations. The following URN semantics can be used in these situations:

- **Do:**
  `resources:myorg.com:organizations:<organization-id>:projects:<project-id>:<resource-id>`
