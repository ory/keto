# The "Ory Permission Language" Specification

Enforcing fine-grained permissions is a critical building block of mature
technology solutions that protect privacy and identity in the information age.
Several proprietary languages to represent permission already exist, namely the
[Authzed Schema Language](https://docs.authzed.com/reference/schema-lang) and
[Auth0 FGA](https://docs.fga.dev/). Most permissions are defined by normal
developers who typically are most familiar with Web technologies like JavaScript
or Typescript. We see the need for a developer-friendly configuration language
for permissions that has such a small learning curve that most developers can
understand and use it with close to no effort. We therefore chose to define our
permissions configuration language as a subset of the most common
general-purpose programming language: JavaScript/TypeScript.

The Ory Permission Language is a syntactical subset of TypeScript. Along with
type definitions for the syntax elements of the language (such as `Namespace` or
`Context`), users can get context help from their IDE while writing the
configuration.

## Notation

The syntax is specified using the Extended Backus-Naur Form (EBNF):

```ebnf
Production  = production_name "=" [ Expression ] "." .
Expression  = Alternative { "|" Alternative } .
Alternative = Term { Term } .
Term        = production_name | token [ "…" token ] | Group | Option | Repetition .
Group       = "(" Expression ")" .
Option      = "[" Expression "]" .
Repetition  = "{" Expression "}" .
```

Productions are expressions constructed from terms and the following operators,
in increasing precedence:

```ebnf
|   alternation
()  grouping
[]  option (0 or 1 times)
{}  repetition (0 to n times)
```

Lower-case production names are used to identify lexical tokens. Non-terminals
are in CamelCase. Lexical tokens are enclosed in double quotes `""` or single
quotes `''`.

The form `a … b` represents the set of characters from a through b as
alternatives. The horizontal ellipsis `…` is also used elsewhere in the spec to
informally denote various enumerations or code snippets that are not further
specified.

## Configuraton text representation

The configuration is encoded in UTF-8.

## Lexical elements

### Comments

1. Line comments start with the character sequence `//` and stop at the end of
   the line.
2. General comments start with the character sequence `/*` and stop with the
   first subsequent character sequence `*/`.
3. Documentation comments start with the character sequence `/**` and stop with
   the first subsequent character sequence `*/`.

### Identifiers

Identifiers name program entities such as variables and types. An identifier is
a sequence of one or more letters and digits. The first character in an
identifier must be a letter.

```ebnf
identifier = letter { letter | digit } .
digit      = "0" … "9" .
letter     = "A" … "Z" | "a" … "z" | "_" .
```

### String literals

String literals represent string constants as sequences of characters.

```ebnf
string_lit    = single_quoted | double_quoted .
single_quoted = "'" identifier "'" .
double_quoted = '"' identifier '"' .
```

### Keywords

The configuration language has the following keywords:

- `class`
- `implements`
- `related`
- `permits`
- `this`
- `ctx`
- `id`
- `imports`
- `exports`
- `as`

### Builtin Types

The following types are built in:

- `Context`
- `Namespace`
- `Namespace[]`
- `boolean`
- `string`
- `SubjectSet<T, R>`

In TypeScript, they would be defined as follows:

```typescript
type Context = { subject: never }

interface Namespace {
  related?: { [relation: string]: Namespace[] }
  permits?: { [method: string]: (ctx: Context) => boolean }
}

interface Array<Namespace> {
  includes(element: Namespace): boolean
  traverse(iteratorfn: (element: Namespace) => boolean): boolean
}

type SubjectSet<
  A extends Namespace,
  R extends keyof A["related"],
> = A["related"][R] extends Array<infer T> ? T : never
```

### Operators

The following character sequences represent boolean operators:

| Operator | Signature      | Semantic                             |
| -------- | -------------- | ------------------------------------ |
| `&&`     | _x_ `&&` _y_   | true iff. both _x_ and _y_ are true  |
| `\|\|`   | _x_ `\|\|` _y_ | true iff. either _x_ or _y_ are true |

The following character sequences represent miscellaneous operators:

| Operator | Example                                  | Semantic                                                          |
| -------- | ---------------------------------------- | ----------------------------------------------------------------- |
| `(`, `)` | `x()`                                    | call function `x`                                                 |
| `[]`     | `T[]`                                    | `T` is an array type                                              |
| `<`, `>` | `SubjectSet<Group, "members">`           | Type reference to the "members" relation of the "Group" namespace |
| `{`, `}` | `class x {}`                             | Scope delimiters                                                  |
| `=>`     | `list.transitive(el => f(el))`           | Lambda definition token.                                          |
| `.`      | `this.x`                                 | Property traversal token                                          |
| `:`      | `relation: type`                         | Relation/type separator token                                     |
| `=`      | `permits = {...}`                        | Assignment token                                                  |
| `,`      | `{x: 1, y: 2}`                           | Property separator                                                |
| `'`      | `'string'`                               | Single quoted string literal                                      |
| `"`      | `"string"`                               | Double quoted string literal                                      |
| `*`      | `import * from @ory/permission-language` | Import glob                                                       |
| `\|`     | `(User \| Group)[]`                      | Type union                                                        |

## Statements

### Type declaration

The top level of the configuration consists of a list of `class` declarations
for each namespace. Each `class` consists of relation declarations and
permission declarations.

```ebnf
Config          = [ ClassDecl ] .
ClassDecl       = "class" identifier "implements" "Namespace" "{" ClassSpec "}" .
ClassSpec       = [ RelationDecls ] | [ PermissionDefns] .
```

The following example declares the type _User_.

```ts
class User implements Namespace {}
```

### Relation declaration

The `related` section of type declarations defines relations. Unlike regular
TypeScript, `RelationName` must be a unique identifier used as the relation
strings in the tuples. The `TypeName` must be the name of another `type` that is
defined above or below (in TypeScript: a class that implements `Namespace`).

Type unions (`|`) can be used to denote that a relation can have subjects of
multiple types, e.g., `viewers: (User | SubjectSet<Group, "members">)[]`,
meaning that the subject of the "viewer" relation can be either a "User", or a
subject set "Group#members".

```ebnf
RelationDecls   = "related" "=" "{" { RelationName ":" ArrayType } "}" .
RelationName    = identifier .
ArrayType       = RelationType | ( "(" RelationType { "|" RelationType } ")" ) "[]" .
RelationType    = SubjectType | SubjectSetType .
SubjectType     = TypeName .
SubjectSetType  = "SubjectSet" "<" TypeName, string_lit ">" .
TypeName        = identifier .
```

Note that all relations are defined as array types `T[]` because there are
naturally only many-to-many relations in Keto.

The following declares a type _Document_ with three relations: _owners_ and
_viewers_, both of which have _users_ as subjects. Additionally, the relation
_parent_ has type _Document_.

```ts
class User {}

class Document {
  related = {
    parents: Document[]
    owners: User[]
    viewers: User[]
  }
}
```

### Permission definition

Permissions are defined as functions within class declarations that take a
parameter `ctx` of type `Context` and evaluate to a boolean `true` or `false`.

The type annotations for `ctx` and the return value are optional.

```ebnf
PermissionDefns = "permits" "=" "{" Permission [ "," Permission ] "}" .
Permission      = PermissionSign "=>" PermissionBody .
PermissionSign  = PermissionName ":" "(" "ctx" [ ":" "Context" ] ")" [ ":" "boolean" ] .
PermissionName  = identifier .
```

The `ctx` object is a fixed parameter that contains the `subject` for which the
permission check should be conducted:

```ts
ctx = { subject: "some_user_id" }
```

The context will contain more fields in the future, e.g., the IP range or
geolocation, the time of day, or the security level of the device making the
request.

```ebnf
PermissionBody  = ( "(" PermissionBody ")" ) | ( PermissionCheck | { Operator PermissionBody } ) .
Operator        = "||" | "&&" .
PermissionCheck = TransitiveCheck | IncludesCheck .
```

The body of a permission check is either one of:

- a `IncludesCheck`, a check that something is in a set, e.g.,
  `this.related.viewers.includes(ctx.subject)`:

  ```ebnf
  IncludesCheck = Var "." "related" "." RelationName "." "includes" "(" "ctx" "." "subject" ")" .
  Var           = identifier .
  ```

- a `TranstitiveCheck`, a call to a permission on a relation, e.g.,
  `this.related.parents.transitive(p => p.permits.view(ctx))`:

  ```ebnf
  TransitiveCheck = "this" "." "related" "." RelationName "." "transitive" "(" Var "=>" ( PermissionCall | IncludesCheck ) ")" .
  PermissionCall  = Var "." "permits" "." PermissionName "(" "ctx" ")" .
  ```

## Implementation notes

`IncludeCheck` and `TransitiveCheck` translate to Zanzibar concepts as follows:

| Keto Config                                               | Zanzibar AST                                                                         |
| --------------------------------------------------------- | ------------------------------------------------------------------------------------ |
| `this.related.R.includes(ctx.subject)`                    | `computed_userset { relation: "R" } }`                                               |
| `this.related.R.transitive(x = x.permits.P(ctx.subject))` | `tuple_to_userset { tupleset { relation: "R" } computed_userset { relation: "P" } }` |

## Type checking

The following type checks are performed once the config is fully parsed:

- Given a `TypeName` as `X` (e.g., in `RelationDecls`), we check that there
  exists a class declaration for `X`.
- Given a `SubjectSetType` as `SubjectSet<T, R>`, we check that `R` is a
  relation defined for `T`.
- Given an `IncludesCheck` as `this.related.R.includes(ctx.subject)`, we check
  that
  - `R` is a relation defined for the current namespace.
- Given a `TransitiveCheck` as
  `this.related.R.transitive(x = x.permits.P(ctx.subject))`, we check that
  - `R` is a relation defined for the current namespace and that
  - `P` is a permission defined for all types referenced by `R`.
- Given a `TransitiveCheck` as
  `this.related.R.transitive(x = x.related.S.includes(ctx.subject))`, we check
  that
  - `R` is a relation defined for the current namespace and that
  - `S` is a relation defined for all types referenced by `R`.

## Examples

The config can be type-checked in `strict` mode by TypeScript with the
[noLib](https://www.typescriptlang.org/tsconfig#noLib) option (preventing the
standard globals), and the
[strictPropertyInitialization](https://www.typescriptlang.org/tsconfig#strictPropertyInitialization)
option (allowing uninitialized properties).

```ts
class User implements Namespace {
  related: {
    manager: User[]
  }
}

class Group implements Namespace {
  related: {
    members: (User | Group)[]
  }
}

class Folder implements Namespace {
  related: {
    parents: File[]
    viewers: (User | SubjectSet<Group, "members">)[]
  }

  permits = {
    view: (ctx: Context): boolean => this.related.viewers.includes(ctx.subject),
  }
}

class File implements Namespace {
  related: {
    parents: (File | Folder)[]
    viewers: (User | SubjectSet<Group, "members">)[]
    owners: (User | SubjectSet<Group, "members">)[]
    siblings: File[]
  }

  permits = {
    view: (ctx: Context): boolean =>
      this.related.parents.traverse((p) =>
        p.related.viewers.includes(ctx.subject),
      ) ||
      this.related.parents.traverse((p) => p.permits.view(ctx)) ||
      this.related.viewers.includes(ctx.subject) ||
      this.related.owners.includes(ctx.subject),

    edit: (ctx: Context) => this.related.owners.includes(ctx.subject),

    rename: (ctx: Context) =>
      this.related.siblings.traverse((s) => s.permits.edit(ctx)),
  }
}
```
