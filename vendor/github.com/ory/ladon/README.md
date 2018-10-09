<h1 align="center"><img src="./docs/images/banner_ladon.png" alt="ORY Ladon - Policy-based Access Control"></h1>

[![Join the chat at https://discord.gg/PAMQWkr](https://img.shields.io/badge/join-chat-00cc99.svg)](https://discord.gg/PAMQWkr)
[![Join newsletter](https://img.shields.io/badge/join-newsletter-00cc99.svg)](http://eepurl.com/bKT3N9)
[![Follow twitter](https://img.shields.io/badge/follow-twitter-00cc99.svg)](https://twitter.com/_aeneasr)
[![Follow GitHub](https://img.shields.io/badge/follow-github-00cc99.svg)](https://github.com/arekkas)
[![Become a patron!](https://img.shields.io/badge/support%20us-on%20patreon-green.svg)](https://patreon.com/user?u=4298803)

[![Build Status](https://travis-ci.org/ory/ladon.svg?branch=master)](https://travis-ci.org/ory/ladon)
[![Coverage Status](https://coveralls.io/repos/ory/ladon/badge.svg?branch=master&service=github)](https://coveralls.io/github/ory/ladon?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/ory/ladon)](https://goreportcard.com/report/github.com/ory/ladon)
[![GoDoc](https://godoc.org/github.com/ory/ladon?status.png)](https://godoc.org/github.com/ory/ladon)

[Ladon](https://en.wikipedia.org/wiki/Ladon_%28mythology%29) is the serpent dragon protecting your resources.

Ladon is a library written in [Go](https://golang.org) for access control policies, similar to [Role Based Access Control](https://en.wikipedia.org/wiki/Role-based_access_control)
or [Access Control Lists](https://en.wikipedia.org/wiki/Access_control_list).
In contrast to [ACL](https://en.wikipedia.org/wiki/Access_control_list) and [RBAC](https://en.wikipedia.org/wiki/Role-based_access_control)
you get fine-grained access control with the ability to answer questions in complex environments such as multi-tenant or distributed applications
and large organizations. Ladon is inspired by [AWS IAM Policies](http://docs.aws.amazon.com/IAM/latest/UserGuide/access_policies.html).

Ladon officially ships with storage adapters for SQL (officially supported: MySQL 5.5+, PostgreSQL 9.2+) and in-memory. Community adapters are available for [CockroachDB](https://github.com/wehco/ladon-crdb).

---

ORY builds solutions for better internet security and accessibility. We have a couple more projects you might enjoy:

* **[Hydra](https://github.com/ory/hydra)**, a security-first open source OAuth2 and OpenID Connect server for new and existing infrastructures that uses Ladon for access control.
* **[ORY Editor](https://github.com/ory/editor)**, an extensible, modern WYSI editor for the web written in React.
* **[Fosite](https://github.com/ory/fosite)**, an extensible security first OAuth 2.0 and OpenID Connect SDK for Go.
* **[Dockertest](https://github.com/ory/dockertest)**: Write better integration tests with dockertest!

---

<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Installation](#installation)
- [Concepts](#concepts)
- [Usage](#usage)
  - [Policies](#policies)
    - [Conditions](#conditions)
      - [CIDR Condition](#cidr-condition)
      - [String Equal Condition](#string-equal-condition)
      - [Boolean Condition](#boolean-condition)
      - [String Match Condition](#string-match-condition)
      - [Subject Condition](#subject-condition)
      - [String Pairs Equal Condition](#string-pairs-equal-condition)
      - [Resource Contains Condition](#resource-contains-condition)
      - [Adding Custom Conditions](#adding-custom-conditions)
    - [Persistence](#persistence)
  - [Access Control (Warden)](#access-control-warden)
  - [Audit Log (Warden)](#audit-log-warden)
- [Limitations](#limitations)
  - [Regular expressions](#regular-expressions)
- [Examples](#examples)
- [Good to know](#good-to-know)
- [Useful commands](#useful-commands)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

Ladon utilizes ory-am/dockertest for tests.
Please refer to [ory-am/dockertest](https://github.com/ory-am/dockertest) for more information of how to setup testing environment.

## Installation

```
go get github.com/ory/ladon
```

We recommend to use [Dep](https://github.com/golang/dep) for dependency management. Ladon uses [semantic
versioning](http://semver.org/) and versions beginning with zero (`0.1.2`) might introduce backwards compatibility
breaks with [each minor version](http://semver.org/#how-should-i-deal-with-revisions-in-the-0yz-initial-development-phase).

## Concepts

Ladon is an access control library that answers the question:

> **Who** is **able** to do **what** on **something** given some **context**

* **Who**: An arbitrary unique subject name, for example "ken" or "printer-service.mydomain.com".
* **Able**: The effect which can be either "allow" or "deny".
* **What**: An arbitrary action name, for example "delete", "create" or "scoped:action:something".
* **Something**: An arbitrary unique resource name, for example "something", "resources.articles.1234" or some uniform
    resource name like "urn:isbn:3827370191".
* **Context**: The current context containing information about the environment such as the IP Address,
    request date, the resource owner name, the department ken is working in or any other information you want to pass along.
    (optional)

To decide what the answer is, Ladon uses policy documents which can be represented as JSON

```json
{
  "description": "One policy to rule them all.",
  "subjects": ["users:<peter|ken>", "users:maria", "groups:admins"],
  "actions" : ["delete", "<create|update>"],
  "effect": "allow",
  "resources": [
    "resources:articles:<.*>",
    "resources:printer"
  ],
  "conditions": {
    "remoteIP": {
        "type": "CIDRCondition",
        "options": {
            "cidr": "192.168.0.1/16"
        }
    }
  }
}
```

and can answer access requests that look like:

```json
{
  "subject": "users:peter",
  "action" : "delete",
  "resource": "resources:articles:ladon-introduction",
  "context": {
    "remoteIP": "192.168.0.5"
  }
}
```

However, Ladon does not come with a HTTP or server implementation. It does not restrict JSON either. We believe that it is your job to decide
if you want to use Protobuf, RESTful, HTTP, AMPQ, or some other protocol. It's up to you to write server!

The following example should give you an idea what a RESTful flow *could* look like. Initially we create a policy by
POSTing it to an artificial HTTP endpoint:

```
> curl \
      -X POST \
      -H "Content-Type: application/json" \
      -d@- \
      "https://my-ladon-implementation.localhost/policies" <<EOF
        {
          "description": "One policy to rule them all.",
          "subjects": ["users:<peter|ken>", "users:maria", "groups:admins"],
          "actions" : ["delete", "<create|update>"],
          "effect": "allow",
          "resources": [
            "resources:articles:<.*>",
            "resources:printer"
          ],
          "conditions": {
            "remoteIP": {
                "type": "CIDRCondition",
                "options": {
                    "cidr": "192.168.0.1/16"
                }
            }
          }
        }
  EOF
```

Then we test if "peter" (ip: "192.168.0.5") is allowed to "delete" the "ladon-introduction" article:

```
> curl \
      -X POST \
      -H "Content-Type: application/json" \
      -d@- \
      "https://my-ladon-implementation.localhost/warden" <<EOF
        {
          "subject": "users:peter",
          "action" : "delete",
          "resource": "resources:articles:ladon-introduction",
          "context": {
            "remoteIP": "192.168.0.5"
          }
        }
  EOF

{
    "allowed": true
}
```

## Usage

We already discussed two essential parts of Ladon: policies and access control requests. Let's take a closer look at those two.

### Policies

Policies are the basis for access control decisions. Think of them as a set of rules. In this library, policies
are abstracted as the `ladon.Policy` interface, and Ladon comes with a standard implementation of this interface
which is `ladon.DefaultPolicy`. Creating such a policy could look like:

```go
import "github.com/ory/ladon"

var pol = &ladon.DefaultPolicy{
	// A required unique identifier. Used primarily for database retrieval.
	ID: "68819e5a-738b-41ec-b03c-b58a1b19d043",

	// A optional human readable description.
	Description: "something humanly readable",

	// A subject can be an user or a service. It is the "who" in "who is allowed to do what on something".
	// As you can see here, you can use regular expressions inside < >.
	Subjects: []string{"max", "peter", "<zac|ken>"},

	// Which resources this policy affects.
	// Again, you can put regular expressions in inside < >.
	Resources: []string{"myrn:some.domain.com:resource:123", "myrn:some.domain.com:resource:345", "myrn:something:foo:<.+>"},

	// Which actions this policy affects. Supports RegExp
	// Again, you can put regular expressions in inside < >.
	Actions: []string{"<create|delete>", "get"},

	// Should access be allowed or denied?
	// Note: If multiple policies match an access request, ladon.DenyAccess will always override ladon.AllowAccess
	// and thus deny access.
	Effect: ladon.AllowAccess,

	// Under which conditions this policy is "active".
	Conditions: ladon.Conditions{
		// In this example, the policy is only "active" when the requested subject is the owner of the resource as well.
		"resourceOwner": &ladon.EqualsSubjectCondition{},

		// Additionally, the policy will only match if the requests remote ip address matches address range 127.0.0.1/32
		"remoteIPAddress": &ladon.CIDRCondition{
			CIDR: "127.0.0.1/32",
		},
	},
}
```

#### Conditions

Conditions are functions returning true or false given a context. Because conditions implement logic, they must
be programmed. Adding conditions to a policy consist of two parts, a key name and an implementation of `ladon.Condition`:

```go
// StringEqualCondition is an exemplary condition.
type StringEqualCondition struct {
	Equals string `json:"equals"`
}

// Fulfills returns true if the given value is a string and is the
// same as in StringEqualCondition.Equals
func (c *StringEqualCondition) Fulfills(value interface{}, _ *ladon.Request) bool {
	s, ok := value.(string)

	return ok && s == c.Equals
}

// GetName returns the condition's name.
func (c *StringEqualCondition) GetName() string {
	return "StringEqualCondition"
}

var pol = &ladon.DefaultPolicy{
    // ...
    Conditions: ladon.Conditions{
        "some-arbitrary-key": &StringEqualCondition{
            Equals: "the-value-should-be-this"
        }
    },
}
```

The default implementation of `Policy` supports JSON un-/marshalling. In JSON, this policy would look like:

```json
{
  "conditions": {
    "some-arbitrary-key": {
        "type": "StringEqualCondition",
        "options": {
            "equals": "the-value-should-be-this"
        }
    }
  }
}
```

As you can see, `type` is the value that `StringEqualCondition.GetName()` is returning and `options` is used to
set the value of `StringEqualCondition.Equals`.

This condition is fulfilled by (we will cover the warden in the next section)

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
        "some-arbitrary-key": "the-value-should-be-this",
    },
}
```

but not by

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
        "some-arbitrary-key": "some other value",
    },
}
```

and neither by:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
        "same value but other key": "the-value-should-be-this",
    },
}
```

Ladon ships with a couple of default conditions:

##### [CIDR Condition](condition_cidr.go)

The CIDR condition matches CIDR IP Ranges. Using this condition would look like this in JSON:

```json
{
    "conditions": {
        "remoteIPAddress": {
            "type": "CIDRCondition",
            "options": {
                "cidr": "192.168.0.1/16"
            }
        }
    }
}
```

and in Go:

```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
        "remoteIPAddress": &ladon.CIDRCondition{
            CIDR: "192.168.0.1/16",
        },
    },
}
```

In this case, we expect that the context of an access request contains a field `"remoteIpAddress"` matching
the CIDR `"192.168.0.1/16"`, for example `"192.168.0.5"`.


##### [String Equal Condition](condition_string_equal.go)

Checks if the value passed in the access request's context is identical with the string that was given initially

```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
        "some-arbitrary-key": &ladon.StringEqualCondition{
            Equals: "the-value-should-be-this"
        }
    },
}
```

and would match in the following case:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
         "some-arbitrary-key": "the-value-should-be-this",
    },
}
```

##### [Boolean Condition](condition_boolean.go)

Checks if the boolean value passed in the access request's context is identical with the expected boolean value in the policy
```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
        "some-arbitrary-key": &ladon.BooleanCondition{
            BooleanValue: true,
        }
    },
}
```

and would match in the following case:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
        "some-arbitrary-key": true,
    },
})
```

This condition type is particularly useful if you need to assert a policy dynamically on resources for multiple subjects. For example, consider
if you wanted to enforce policy that only allows individuals that own a resource to view that resource. You'd have to be able to create a Ladon
policy that permits access to every resource for every subject that enters your system.

With the Boolean Condition type, you can use conditional logic at runtime to create a match for a policy's condition.

##### [String Match Condition](condition_string_match.go)

Checks if the value passed in the access request's context matches the regular expression that was given initially

```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
      "some-arbitrary-key": &ladon.StringMatchCondition{
          Matches: "regex-pattern-here.+"
      }
    }
}
```

and would match in the following case:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
          "some-arbitrary-key": "regex-pattern-here111"
    }
  }
})
```

##### [Subject Condition](condition_subject_equal.go)

Checks if the access request's subject is identical with the string that was given initially

```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
        "some-arbitrary-key": &ladon.EqualsSubjectCondition{}
    },
}
```

and would match

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Subject: "peter",
    Context: &ladon.Context{
         "some-arbitrary-key": "peter",
    },
}
```

but not:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Subject: "peter",
    Context: &ladon.Context{
         "some-arbitrary-key": "max",
    },
}
```

##### [String Pairs Equal Condition](condition_string_pairs_equal.go)

Checks if the value passed in the access request's context contains two-element arrays
and that both elements in each pair are equal.

```go
var pol = &ladon.DefaultPolicy{
    Conditions: ladon.Conditions{
        "some-arbitrary-key": &ladon.StringPairsEqualCondition{}
    },
}
```

and would match

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
         "some-arbitrary-key": [
             ["some-arbitrary-pair-value", "some-arbitrary-pair-value"],
             ["some-other-arbitrary-pair-value", "some-other-arbitrary-pair-value"],
         ]
    },
}
```

but not:

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Context: &ladon.Context{
         "some-arbitrary-key": [
             ["some-arbitrary-pair-value", "some-other-arbitrary-pair-value"],
         ]
    },
}
```


##### [Resource Contains Condition](condition_resource_contains.go)

Checks if the string value passed in the access request's context is present in the resource string.

The Condition requires a value string and an optional delimiter (needs to match the resource string) to be passed.

A resource could for instance be: `myrn:some.domain.com:resource:123` and `myrn:some.otherdomain.com:resource:123` (the `:` is then considered a delimiter, and used by the condition to be able to separate the resource components from each other) to allow an action to the resources on `myrn:some.otherdomain.com` you could for instance create a resource condition  with

{value: `myrn:some.otherdomain.com`, Delimiter: ":"}

alternatively:

{value: `myrn:some.otherdomain.com`}

> The delimiter is optional *but needed for* the condition to be able to separate resource string components:
> i.e. to make sure the value `foo:bar` matches `foo:bar` but not `foo:bara` nor `foo:bara:baz`.
>
> That is, a delimiter is necessary to separate:
>
> `{value: "myrn:fo", delimiter: ":"}` from `{value: "myrn:foo", delimiter: ":"}` or
> `{value: "myid:12"}` from `{value: "myid:123"}`.



This condition is fulfilled by this (allow for all resources containing `part:north`):

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Resource: "rn:city:laholm:part:north"
    Context: &ladon.Context{
      delimiter: ":",
      value: "part:north"
    },
}
```

or ( allow all resources with `city:laholm`)

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Resource: "rn:city:laholm:part:north"
    Context: &ladon.Context{
      delimiter: ":",
      value: "city:laholm"
    },
}
```

but not (allow for all resources containing `part:west`, the resource does not contain `part:west`):

```go
var err = warden.IsAllowed(&ladon.Request{
    // ...
    Resource: "rn:city:laholm:part:north"
    Context: &ladon.Context{
      delimiter: ":",
      value: "part:west"
    },
}
```


##### Adding Custom Conditions

You can add custom conditions by appending it to `ladon.ConditionFactories`:

```go
import "github.com/ory/ladon"

func main() {
    // ...

    ladon.ConditionFactories[new(CustomCondition).GetName()] = func() Condition {
        return new(CustomCondition)
    }

    // ...
}
```

#### Persistence

Obviously, creating such a policy is not enough. You want to persist it too. Ladon ships an interface `ladon.Manager` for
this purpose with default implementations for In-Memory and SQL (PostgreSQL, MySQL) via [sqlx](github.com/jmoiron/sqlx). 
There are also adapters available written by the community [for Redis and RethinkDB](https://github.com/ory/ladon-community)

Let's take a look how to instantiate those:

**In-Memory** (officially supported)

```go
import (
	"github.com/ory/ladon"
	manager "github.com/ory/ladon/manager/memory"
)


func main() {
	warden := &ladon.Ladon{
		Manager: manager.NewMemoryManager(),
	}
	err := warden.Manager.Create(pol)

    // ...
}
```

**SQL** (officially supported)

```go
import "github.com/ory/ladon"
import manager "github.com/ory/ladon/manager/sql"
import "github.com/jmoiron/sqlx"
import _ "github.com/go-sql-driver/mysql"

func main() {
    // The database manager expects a sqlx.DB object
    //
    // For MySQL, be sure to include parseTime=true in the connection string
    // You can find all of the supported MySQL connection string options for the
    // driver at: https://github.com/go-sql-driver/mysql
    //
    db, err = sqlx.Open("mysql", "user:pass@tcp(127.0.0.1:3306)/?parseTime=true")
    // Or, if using postgres:
    //  import _ "github.com/lib/pq"
    //
    //  db, err = sqlx.Open("postgres", "postgres://foo:bar@localhost/ladon")
    if err != nil {
      log.Fatalf("Could not connect to database: %s", err)
    }

    warden := &ladon.Ladon{
        Manager: manager.NewSQLManager(db, nil),
    }

    // You must call SQLManager.CreateSchemas(schema, table) before use
    // to apply the necessary SQL migrations
    //
    // You can provide your own schema and table name or pass
    // empty strings to use the default
    n, err := warden.Manager.CreateSchemas("", "")
    if err != nil {
      log.Fatalf("Failed to create schemas: %s", err)
    }
    log.Printf("applied %d migrations", n)

    // ...
}
```

### Access Control (Warden)

Now that we have defined our policies, we can use the warden to check if a request is valid.
`ladon.Ladon`, which is the default implementation for the `ladon.Warden` interface defines `ladon.Ladon.IsAllowed()` which
will return `nil` if the access request can be granted and an error otherwise.

```go
import "github.com/ory/ladon"

func main() {
    // ...

    if err := warden.IsAllowed(&ladon.Request{
        Subject: "peter",
        Action: "delete",
        Resource: "myrn:some.domain.com:resource:123",
        Context: ladon.Context{
            "ip": "127.0.0.1",
        },
    }); err != nil {
        log.Fatal("Access denied")
    }

    // ...
}
```

### Audit Log (Warden)

In order to keep track of authorization grants and denials, it is possible to attach a `ladon.AuditLogger`.
The provided `ladon.AuditLoggerInfo` outputs information about the policies involved when responding to authorization requests.

```go
import "github.com/ory/ladon"
import manager "github.com/ory/ladon/manager/memory"

func main() {

    warden := ladon.Ladon{
        Manager: manager.NewMemoryManager(),
        AuditLogger: ladon.AuditLoggerInfo{}
    }

    // ...

```

It will output to `stderr` by default.

## Limitations

Ladon's limitations are listed here.

### Regular expressions

Matching regular expressions has a complexity of `O(n)` and databases such as MySQL or Postgres can not
leverage indexes when parsing regular expressions. Thus, there is considerable overhead when using regular
expressions.

We have implemented various strategies for reducing policy matching time:

1. An LRU cache is used for caching frequently compiled regular expressions. This reduces cpu complexity
significantly for memory manager implementations.
2. The SQL schema is 3NF normalized.
3. Policies, subjects and actions are stored uniquely, reducing the total number of rows.
4. Only one query per look up is executed.
5. If no regular expression is used, a simple equal match is done in SQL back-ends.

You will get the best performance with the in-memory manager. The SQL adapters perform about
1000:1 compared to the in-memory solution. Please note that these
tests where in laboratory environments with Docker, without an SSD, and single-threaded. You might get better
results on your system. We are thinking about introducing It would be possible a simple cache strategy such as
LRU with a maximum age to further reduce runtime complexity.

We are also considering to offer different matching strategies (e.g. wildcard match) in the future, which will perform better
with SQL databases. If you have ideas or suggestions, leave us an issue.

## Examples

Check out [ladon_test.go](ladon_test.go) which includes a couple of policies and tests cases. You can run the code with `go test -run=TestLadon -v .`

## Good to know

* All checks are *case sensitive* because subject values could be case sensitive IDs.
* If `ladon.Ladon` is not able to match a policy with the request, it will default to denying the request and return an error.

Ladon does not use reflection for matching conditions to their appropriate structs due to security considerations.

## Useful commands

**Create mocks**
```sh
mockgen -package ladon_test -destination manager_mock_test.go github.com/ory/ladon Manager
```
