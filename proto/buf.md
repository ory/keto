# Protocol Documentation

<a name="top"></a>

## Table of Contents

- [ory/keto/opl/v1alpha1/syntax_service.proto](#ory_keto_opl_v1alpha1_syntax_service-proto)

  - [CheckRequest](#ory-keto-opl-v1alpha1-CheckRequest)
  - [CheckResponse](#ory-keto-opl-v1alpha1-CheckResponse)
  - [ParseError](#ory-keto-opl-v1alpha1-ParseError)
  - [SourcePosition](#ory-keto-opl-v1alpha1-SourcePosition)

  - [SyntaxService](#ory-keto-opl-v1alpha1-SyntaxService)

- [ory/keto/relation_tuples/v1alpha2/relation_tuples.proto](#ory_keto_relation_tuples_v1alpha2_relation_tuples-proto)
  - [RelationQuery](#ory-keto-relation_tuples-v1alpha2-RelationQuery)
  - [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple)
  - [Subject](#ory-keto-relation_tuples-v1alpha2-Subject)
  - [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet)
- [ory/keto/relation_tuples/v1alpha2/check_service.proto](#ory_keto_relation_tuples_v1alpha2_check_service-proto)

  - [CheckRequest](#ory-keto-relation_tuples-v1alpha2-CheckRequest)
  - [CheckResponse](#ory-keto-relation_tuples-v1alpha2-CheckResponse)

  - [CheckService](#ory-keto-relation_tuples-v1alpha2-CheckService)

- [ory/keto/relation_tuples/v1alpha2/expand_service.proto](#ory_keto_relation_tuples_v1alpha2_expand_service-proto)

  - [ExpandRequest](#ory-keto-relation_tuples-v1alpha2-ExpandRequest)
  - [ExpandResponse](#ory-keto-relation_tuples-v1alpha2-ExpandResponse)
  - [SubjectTree](#ory-keto-relation_tuples-v1alpha2-SubjectTree)

  - [NodeType](#ory-keto-relation_tuples-v1alpha2-NodeType)

  - [ExpandService](#ory-keto-relation_tuples-v1alpha2-ExpandService)

- [ory/keto/relation_tuples/v1alpha2/namespaces_service.proto](#ory_keto_relation_tuples_v1alpha2_namespaces_service-proto)

  - [ListNamespacesRequest](#ory-keto-relation_tuples-v1alpha2-ListNamespacesRequest)
  - [ListNamespacesResponse](#ory-keto-relation_tuples-v1alpha2-ListNamespacesResponse)
  - [Namespace](#ory-keto-relation_tuples-v1alpha2-Namespace)

  - [NamespacesService](#ory-keto-relation_tuples-v1alpha2-NamespacesService)

- [ory/keto/relation_tuples/v1alpha2/read_service.proto](#ory_keto_relation_tuples_v1alpha2_read_service-proto)

  - [ListRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest)
  - [ListRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest-Query)
  - [ListRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesResponse)

  - [ReadService](#ory-keto-relation_tuples-v1alpha2-ReadService)

- [ory/keto/relation_tuples/v1alpha2/version.proto](#ory_keto_relation_tuples_v1alpha2_version-proto)

  - [GetVersionRequest](#ory-keto-relation_tuples-v1alpha2-GetVersionRequest)
  - [GetVersionResponse](#ory-keto-relation_tuples-v1alpha2-GetVersionResponse)

  - [VersionService](#ory-keto-relation_tuples-v1alpha2-VersionService)

- [ory/keto/relation_tuples/v1alpha2/write_service.proto](#ory_keto_relation_tuples_v1alpha2_write_service-proto)

  - [DeleteRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest)
  - [DeleteRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest-Query)
  - [DeleteRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesResponse)
  - [RelationTupleDelta](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta)
  - [TransactRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesRequest)
  - [TransactRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesResponse)

  - [RelationTupleDelta.Action](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta-Action)

  - [WriteService](#ory-keto-relation_tuples-v1alpha2-WriteService)

- [ory/keto/v1beta/relation_tuples.proto](#ory_keto_v1beta_relation_tuples-proto)
  - [RelationQuery](#ory-keto-v1beta-RelationQuery)
  - [RelationTuple](#ory-keto-v1beta-RelationTuple)
  - [Subject](#ory-keto-v1beta-Subject)
  - [SubjectSet](#ory-keto-v1beta-SubjectSet)
- [ory/keto/v1beta/check_service.proto](#ory_keto_v1beta_check_service-proto)

  - [CheckRequest](#ory-keto-v1beta-CheckRequest)
  - [CheckResponse](#ory-keto-v1beta-CheckResponse)

  - [CheckService](#ory-keto-v1beta-CheckService)

- [Scalar Value Types](#scalar-value-types)

<a name="ory_keto_opl_v1alpha1_syntax_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/opl/v1alpha1/syntax_service.proto

<a name="ory-keto-opl-v1alpha1-CheckRequest"></a>

### CheckRequest

| Field   | Type            | Label | Description |
| ------- | --------------- | ----- | ----------- |
| content | [bytes](#bytes) |       |             |

<a name="ory-keto-opl-v1alpha1-CheckResponse"></a>

### CheckResponse

| Field        | Type                                            | Label    | Description |
| ------------ | ----------------------------------------------- | -------- | ----------- |
| parse_errors | [ParseError](#ory-keto-opl-v1alpha1-ParseError) | repeated |             |

<a name="ory-keto-opl-v1alpha1-ParseError"></a>

### ParseError

| Field   | Type                                                    | Label | Description |
| ------- | ------------------------------------------------------- | ----- | ----------- |
| message | [string](#string)                                       |       |             |
| start   | [SourcePosition](#ory-keto-opl-v1alpha1-SourcePosition) |       |             |
| end     | [SourcePosition](#ory-keto-opl-v1alpha1-SourcePosition) |       |             |

<a name="ory-keto-opl-v1alpha1-SourcePosition"></a>

### SourcePosition

| Field  | Type              | Label | Description |
| ------ | ----------------- | ----- | ----------- |
| line   | [uint32](#uint32) |       |             |
| column | [uint32](#uint32) |       |             |

<a name="ory-keto-opl-v1alpha1-SyntaxService"></a>

### SyntaxService

The service that checks the syntax of an OPL file.

| Method Name | Request Type                                        | Response Type                                         | Description                      |
| ----------- | --------------------------------------------------- | ----------------------------------------------------- | -------------------------------- |
| Check       | [CheckRequest](#ory-keto-opl-v1alpha1-CheckRequest) | [CheckResponse](#ory-keto-opl-v1alpha1-CheckResponse) | Performs a syntax check request. |

<a name="ory_keto_relation_tuples_v1alpha2_relation_tuples-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/relation_tuples.proto

<a name="ory-keto-relation_tuples-v1alpha2-RelationQuery"></a>

### RelationQuery

The query for listing relationships. Clients can specify any optional field to
partially filter for specific relationships.

Example use cases (namespace is always required):

- object only: display a list of all permissions referring to a specific object
- relation only: get all groups that have members; get all directories that have
  content
- object &amp; relation: display all subjects that have a specific permission
  relation
- subject &amp; relation: display all groups a subject belongs to; display all
  objects a subject has access to
- object &amp; relation &amp; subject: check whether the relation tuple already
  exists

| Field     | Type                                                  | Label    | Description                                                                                                                           |
| --------- | ----------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| namespace | [string](#string)                                     | optional | The namespace this relation tuple lives in.                                                                                           |
| object    | [string](#string)                                     | optional | The object related by this tuple. It is an object in the namespace of the tuple.                                                      |
| relation  | [string](#string)                                     | optional | The relation between an Object and a Subject.                                                                                         |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject) | optional | The subject related by this tuple. A Subject either represents a concrete subject id or a `SubjectSet` that expands to more Subjects. |

<a name="ory-keto-relation_tuples-v1alpha2-RelationTuple"></a>

### RelationTuple

RelationTuple defines a relation between an Object and a Subject.

| Field     | Type                                                  | Label | Description                                                                                                                           |
| --------- | ----------------------------------------------------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------- |
| namespace | [string](#string)                                     |       | The namespace this relation tuple lives in.                                                                                           |
| object    | [string](#string)                                     |       | The object related by this tuple. It is an object in the namespace of the tuple.                                                      |
| relation  | [string](#string)                                     |       | The relation between an Object and a Subject.                                                                                         |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject) |       | The subject related by this tuple. A Subject either represents a concrete subject id or a `SubjectSet` that expands to more Subjects. |

<a name="ory-keto-relation_tuples-v1alpha2-Subject"></a>

### Subject

Subject is either a concrete subject id or a `SubjectSet` expanding to more
Subjects.

| Field | Type                                                        | Label | Description                                                                                                             |
| ----- | ----------------------------------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- |
| id    | [string](#string)                                           |       | A concrete id of the subject.                                                                                           |
| set   | [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet) |       | A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-relation_tuples-v1alpha2-SubjectSet"></a>

### SubjectSet

SubjectSet refers to all subjects who have the same `relation` on an `object`.

| Field     | Type              | Label | Description                                                              |
| --------- | ----------------- | ----- | ------------------------------------------------------------------------ |
| namespace | [string](#string) |       | The namespace of the object and relation referenced in this subject set. |
| object    | [string](#string) |       | The object related by this subject set.                                  |
| relation  | [string](#string) |       | The relation between the object and the subjects.                        |

<a name="ory_keto_relation_tuples_v1alpha2_check_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/check_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-CheckRequest"></a>

### CheckRequest

The request for a CheckService.Check RPC. Checks whether a specific subject is
related to an object.

| Field     | Type              | Label | Description                                          |
| --------- | ----------------- | ----- | ---------------------------------------------------- |
| namespace | [string](#string) |       | **Deprecated.** The namespace to evaluate the check. |

Note: If you use the expand-API and the check evaluates a RelationTuple
specifying a SubjectSet as subject or due to a rewrite rule in a namespace
config this check request may involve other namespaces automatically. | | object
| [string](#string) | | **Deprecated.** The related object in this check. | |
relation | [string](#string) | | **Deprecated.** The relation between the Object
and the Subject. | | subject |
[Subject](#ory-keto-relation_tuples-v1alpha2-Subject) | | **Deprecated.** The
related subject in this check. | | tuple |
[RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple) | | | | latest
| [bool](#bool) | | This field is not implemented yet and has no effect. &lt;!--
Set this field to `true` in case your application needs to authorize depending
on up to date ACLs, also called a &#34;content-change check&#34;.

If set to `true` the `snaptoken` field is ignored, the check is evaluated at the
latest snapshot (globally consistent) and the response includes a snaptoken for
clients to store along with object contents that can be used for subsequent
checks of the same content version.

Example use case: - You need to authorize a user to modify/delete some resource
and it is unacceptable that if the permission to do that had just been revoked
some seconds ago so that the change had not yet been fully replicated to all
availability zones. --&gt; | | snaptoken | [string](#string) | | This field is
not implemented yet and has no effect. &lt;!-- Optional. Like reads, a check is
always evaluated at a consistent snapshot no earlier than the given snaptoken.

Leave this field blank if you want to evaluate the check based on eventually
consistent ACLs, benefiting from very low latency, but possibly slightly stale
results.

If the specified token is too old and no longer known, the server falls back as
if no snaptoken had been specified.

If not specified the server tries to evaluate the check on the best snapshot
version where it is very likely that ACLs had already been replicated to all
availability zones. --&gt; | | max_depth | [int32](#int32) | | The maximum depth
to search for a relation.

If the value is less than 1 or greater than the global max-depth then the global
max-depth will be used instead. |

<a name="ory-keto-relation_tuples-v1alpha2-CheckResponse"></a>

### CheckResponse

The response for a CheckService.Check rpc.

| Field   | Type          | Label | Description                                                            |
| ------- | ------------- | ----- | ---------------------------------------------------------------------- |
| allowed | [bool](#bool) |       | Whether the specified subject (id) is related to the requested object. |

It is false by default if no ACL matches. | | snaptoken | [string](#string) | |
This field is not implemented yet and has no effect. &lt;!-- The last known
snapshot token ONLY specified if the request had not specified a snaptoken,
since this performed a &#34;content-change request&#34; and consistently fetched
the last known snapshot token.

This field is not set if the request had specified a snaptoken!

If set, clients should cache and use this token for subsequent requests to have
minimal latency, but allow slightly stale responses (only some milliseconds or
seconds). --&gt; |

<a name="ory-keto-relation_tuples-v1alpha2-CheckService"></a>

### CheckService

The service that performs authorization checks based on the stored Access
Control Lists.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).

| Method Name | Request Type                                                    | Response Type                                                     | Description                      |
| ----------- | --------------------------------------------------------------- | ----------------------------------------------------------------- | -------------------------------- |
| Check       | [CheckRequest](#ory-keto-relation_tuples-v1alpha2-CheckRequest) | [CheckResponse](#ory-keto-relation_tuples-v1alpha2-CheckResponse) | Performs an authorization check. |

<a name="ory_keto_relation_tuples_v1alpha2_expand_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/expand_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-ExpandRequest"></a>

### ExpandRequest

The request for an ExpandService.Expand RPC. Expands the given subject set.

| Field     | Type                                                  | Label | Description                         |
| --------- | ----------------------------------------------------- | ----- | ----------------------------------- |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject) |       | The subject to expand.              |
| max_depth | [int32](#int32)                                       |       | The maximum depth of tree to build. |

If the value is less than 1 or greater than the global max-depth then the global
max-depth will be used instead.

It is important to set this parameter to a meaningful value. Ponder how deep you
really want to display this. | | snaptoken | [string](#string) | | This field is
not implemented yet and has no effect. &lt;!-- Optional. Like reads, a expand is
always evaluated at a consistent snapshot no earlier than the given snaptoken.

Leave this field blank if you want to expand based on eventually consistent
ACLs, benefiting from very low latency, but possibly slightly stale results.

If the specified token is too old and no longer known, the server falls back as
if no snaptoken had been specified.

If not specified the server tries to build the tree on the best snapshot version
where it is very likely that ACLs had already been replicated to all
availability zones. --&gt; |

<a name="ory-keto-relation_tuples-v1alpha2-ExpandResponse"></a>

### ExpandResponse

The response for a ExpandService.Expand RPC.

| Field | Type                                                          | Label | Description                                                                                          |
| ----- | ------------------------------------------------------------- | ----- | ---------------------------------------------------------------------------------------------------- |
| tree  | [SubjectTree](#ory-keto-relation_tuples-v1alpha2-SubjectTree) |       | The tree the requested subject set expands to. The requested subject set is the subject of the root. |

This field can be nil in some circumstances. |

<a name="ory-keto-relation_tuples-v1alpha2-SubjectTree"></a>

### SubjectTree

| Field     | Type                                                              | Label    | Description                                                                                                         |
| --------- | ----------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------- |
| node_type | [NodeType](#ory-keto-relation_tuples-v1alpha2-NodeType)           |          | The type of the node.                                                                                               |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject)             |          | **Deprecated.** The subject this node represents. Deprecated: More information is now available in the tuple field. |
| tuple     | [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple) |          | The relation tuple this node represents.                                                                            |
| children  | [SubjectTree](#ory-keto-relation_tuples-v1alpha2-SubjectTree)     | repeated | The children of this node.                                                                                          |

This is never set if `node_type` == `NODE_TYPE_LEAF`. |

<a name="ory-keto-relation_tuples-v1alpha2-NodeType"></a>

### NodeType

| Name                   | Number | Description                                                                                                |
| ---------------------- | ------ | ---------------------------------------------------------------------------------------------------------- |
| NODE_TYPE_UNSPECIFIED  | 0      |                                                                                                            |
| NODE_TYPE_UNION        | 1      | This node expands to a union of all children.                                                              |
| NODE_TYPE_EXCLUSION    | 2      | Not implemented yet.                                                                                       |
| NODE_TYPE_INTERSECTION | 3      | Not implemented yet.                                                                                       |
| NODE_TYPE_LEAF         | 4      | This node is a leaf and contains no children. Its subject is a `SubjectID` unless `max_depth` was reached. |

<a name="ory-keto-relation_tuples-v1alpha2-ExpandService"></a>

### ExpandService

The service that performs subject set expansion based on the stored Access
Control Lists.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).

| Method Name | Request Type                                                      | Response Type                                                       | Description                                      |
| ----------- | ----------------------------------------------------------------- | ------------------------------------------------------------------- | ------------------------------------------------ |
| Expand      | [ExpandRequest](#ory-keto-relation_tuples-v1alpha2-ExpandRequest) | [ExpandResponse](#ory-keto-relation_tuples-v1alpha2-ExpandResponse) | Expands the subject set into a tree of subjects. |

<a name="ory_keto_relation_tuples_v1alpha2_namespaces_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/namespaces_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-ListNamespacesRequest"></a>

### ListNamespacesRequest

Request for ReadService.ListNamespaces RPC.

<a name="ory-keto-relation_tuples-v1alpha2-ListNamespacesResponse"></a>

### ListNamespacesResponse

| Field      | Type                                                      | Label    | Description |
| ---------- | --------------------------------------------------------- | -------- | ----------- |
| namespaces | [Namespace](#ory-keto-relation_tuples-v1alpha2-Namespace) | repeated |             |

<a name="ory-keto-relation_tuples-v1alpha2-Namespace"></a>

### Namespace

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| name  | [string](#string) |       |             |

<a name="ory-keto-relation_tuples-v1alpha2-NamespacesService"></a>

### NamespacesService

The service to query namespaces.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).

| Method Name    | Request Type                                                                      | Response Type                                                                       | Description      |
| -------------- | --------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ---------------- |
| ListNamespaces | [ListNamespacesRequest](#ory-keto-relation_tuples-v1alpha2-ListNamespacesRequest) | [ListNamespacesResponse](#ory-keto-relation_tuples-v1alpha2-ListNamespacesResponse) | Lists Namespaces |

<a name="ory_keto_relation_tuples_v1alpha2_read_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/read_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest"></a>

### ListRelationTuplesRequest

Request for ReadService.ListRelationTuples RPC. See
`ListRelationTuplesRequest_Query` for how to filter the query.

| Field | Type                                                                                                  | Label | Description                                                                         |
| ----- | ----------------------------------------------------------------------------------------------------- | ----- | ----------------------------------------------------------------------------------- |
| query | [ListRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest-Query) |       | **Deprecated.** All query constraints are concatenated with a logical AND operator. |

The RelationTuple list from ListRelationTuplesResponse is ordered from the
newest RelationTuple to the oldest. | | relation_query |
[RelationQuery](#ory-keto-relation_tuples-v1alpha2-RelationQuery) | | | |
expand_mask | [google.protobuf.FieldMask](#google-protobuf-FieldMask) | | This
field is not implemented yet and has no effect. &lt;!-- Optional. The list of
fields to be expanded in the RelationTuple list returned in
`ListRelationTuplesResponse`. Leaving this field unspecified means all fields
are expanded.

Available fields: &#34;object&#34;, &#34;relation&#34;, &#34;subject&#34;,
&#34;namespace&#34;, &#34;subject.id&#34;, &#34;subject.namespace&#34;,
&#34;subject.object&#34;, &#34;subject.relation&#34; --&gt; | | snaptoken |
[string](#string) | | This field is not implemented yet and has no effect.
&lt;!-- Optional. The snapshot token for this read. --&gt; | | page_size |
[int32](#int32) | | Optional. The maximum number of RelationTuples to return in
the response.

Default: 100 | | page_token | [string](#string) | | Optional. An opaque
pagination token returned from a previous call to `ListRelationTuples` that
indicates where the page should start at.

An empty token denotes the first page. All successive pages require the token
from the previous page. |

<a name="ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest-Query"></a>

### ListRelationTuplesRequest.Query

The query for listing relationships. Clients can specify any optional field to
partially filter for specific relationships.

Example use cases (namespace is always required):

- object only: display a list of all permissions referring to a specific object
- relation only: get all groups that have members; get all directories that have
  content
- object &amp; relation: display all subjects that have a specific permission
  relation
- subject &amp; relation: display all groups a subject belongs to; display all
  objects a subject has access to
- object &amp; relation &amp; subject: check whether the relation tuple already
  exists

| Field     | Type                                                  | Label | Description                          |
| --------- | ----------------------------------------------------- | ----- | ------------------------------------ |
| namespace | [string](#string)                                     |       | Required. The namespace to query.    |
| object    | [string](#string)                                     |       | Optional. The object to query for.   |
| relation  | [string](#string)                                     |       | Optional. The relation to query for. |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject) |       | Optional. The subject to query for.  |

<a name="ory-keto-relation_tuples-v1alpha2-ListRelationTuplesResponse"></a>

### ListRelationTuplesResponse

The response of a ReadService.ListRelationTuples RPC.

| Field           | Type                                                              | Label    | Description                                                                                            |
| --------------- | ----------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------ |
| relation_tuples | [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple) | repeated | The relationships matching the list request.                                                           |
| next_page_token | [string](#string)                                                 |          | The token required to get the next page. If this is the last page, the token will be the empty string. |

<a name="ory-keto-relation_tuples-v1alpha2-ReadService"></a>

### ReadService

The service to query relationships.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).

| Method Name        | Request Type                                                                              | Response Type                                                                               | Description              |
| ------------------ | ----------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------ |
| ListRelationTuples | [ListRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest) | [ListRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesResponse) | Lists ACL relationships. |

<a name="ory_keto_relation_tuples_v1alpha2_version-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/version.proto

<a name="ory-keto-relation_tuples-v1alpha2-GetVersionRequest"></a>

### GetVersionRequest

Request for the VersionService.GetVersion RPC.

<a name="ory-keto-relation_tuples-v1alpha2-GetVersionResponse"></a>

### GetVersionResponse

Response of the VersionService.GetVersion RPC.

| Field   | Type              | Label | Description                                  |
| ------- | ----------------- | ----- | -------------------------------------------- |
| version | [string](#string) |       | The version string of the Ory Keto instance. |

<a name="ory-keto-relation_tuples-v1alpha2-VersionService"></a>

### VersionService

The service returning the specific Ory Keto instance version.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis)
and [write-APIs](../concepts/api-overview.mdx#write-apis).

| Method Name | Request Type                                                              | Response Type                                                               | Description                                   |
| ----------- | ------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------- |
| GetVersion  | [GetVersionRequest](#ory-keto-relation_tuples-v1alpha2-GetVersionRequest) | [GetVersionResponse](#ory-keto-relation_tuples-v1alpha2-GetVersionResponse) | Returns the version of the Ory Keto instance. |

<a name="ory_keto_relation_tuples_v1alpha2_write_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/write_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest"></a>

### DeleteRelationTuplesRequest

| Field          | Type                                                                                                      | Label | Description     |
| -------------- | --------------------------------------------------------------------------------------------------------- | ----- | --------------- |
| query          | [DeleteRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest-Query) |       | **Deprecated.** |
| relation_query | [RelationQuery](#ory-keto-relation_tuples-v1alpha2-RelationQuery)                                         |       |                 |

<a name="ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest-Query"></a>

### DeleteRelationTuplesRequest.Query

The query for deleting relationships

| Field     | Type                                                  | Label | Description                          |
| --------- | ----------------------------------------------------- | ----- | ------------------------------------ |
| namespace | [string](#string)                                     |       | Optional. The namespace to query.    |
| object    | [string](#string)                                     |       | Optional. The object to query for.   |
| relation  | [string](#string)                                     |       | Optional. The relation to query for. |
| subject   | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject) |       | Optional. The subject to query for.  |

<a name="ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesResponse"></a>

### DeleteRelationTuplesResponse

<a name="ory-keto-relation_tuples-v1alpha2-RelationTupleDelta"></a>

### RelationTupleDelta

Write-delta for a TransactRelationTuplesRequest.

| Field          | Type                                                                                      | Label | Description                            |
| -------------- | ----------------------------------------------------------------------------------------- | ----- | -------------------------------------- |
| action         | [RelationTupleDelta.Action](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta-Action) |       | The action to do on the RelationTuple. |
| relation_tuple | [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple)                         |       | The target RelationTuple.              |

<a name="ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesRequest"></a>

### TransactRelationTuplesRequest

The request of a WriteService.TransactRelationTuples RPC.

| Field                 | Type                                                                        | Label    | Description                                                                                                                              |
| --------------------- | --------------------------------------------------------------------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------- |
| relation_tuple_deltas | [RelationTupleDelta](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta) | repeated | The write delta for the relationships operated in one single transaction. Either all actions succeed or no change takes effect on error. |

<a name="ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesResponse"></a>

### TransactRelationTuplesResponse

The response of a WriteService.TransactRelationTuples rpc.

| Field      | Type              | Label    | Description                                                                                                                                                                                                                                                |
| ---------- | ----------------- | -------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| snaptokens | [string](#string) | repeated | This field is not implemented yet and has no effect. &lt;!-- The list of the new latest snapshot tokens of the affected RelationTuple, with the same index as specified in the `relation_tuple_deltas` field of the TransactRelationTuplesRequest request. |

If the RelationTupleDelta_Action was DELETE the snaptoken is empty at the same
index. --&gt; |

<a name="ory-keto-relation_tuples-v1alpha2-RelationTupleDelta-Action"></a>

### RelationTupleDelta.Action

| Name               | Number | Description                                                                                                 |
| ------------------ | ------ | ----------------------------------------------------------------------------------------------------------- |
| ACTION_UNSPECIFIED | 0      | Unspecified. The `TransactRelationTuples` RPC ignores this RelationTupleDelta if an action was unspecified. |
| ACTION_INSERT      | 1      | Insertion of a new RelationTuple. It is ignored if already existing.                                        |
| ACTION_DELETE      | 2      | Deletion of the RelationTuple. It is ignored if it does not exist.                                          |

<a name="ory-keto-relation_tuples-v1alpha2-WriteService"></a>

### WriteService

The write service to create and delete Access Control Lists.

This service is part of the
[write-APIs](../concepts/api-overview.mdx#write-apis).

| Method Name            | Request Type                                                                                      | Response Type                                                                                       | Description                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| TransactRelationTuples | [TransactRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesRequest) | [TransactRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesResponse) | Writes one or more relationships in a single transaction. |
| DeleteRelationTuples   | [DeleteRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest)     | [DeleteRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesResponse)     | Deletes relationships based on relation query             |

<a name="ory_keto_v1beta_relation_tuples-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/v1beta/relation_tuples.proto

<a name="ory-keto-v1beta-RelationQuery"></a>

### RelationQuery

The query for listing relation tuples. Clients can specify any optional field to
partially filter for specific relation tuples.

Example use cases (namespace is always required):

- object only: display a list of all permissions referring to a specific object
- relation only: get all groups that have members; get all directories that have
  content
- object &amp; relation: display all subjects that have a specific permission
  relation
- subject &amp; relation: display all groups a subject belongs to; display all
  objects a subject has access to
- object &amp; relation &amp; subject: check whether the relation tuple already
  exists

| Field     | Type                                | Label    | Description                                                                                                                           |
| --------- | ----------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------- |
| namespace | [string](#string)                   | optional | The namespace this relation tuple lives in.                                                                                           |
| object    | [string](#string)                   | optional | The object related by this tuple. It is an object in the namespace of the tuple.                                                      |
| relation  | [string](#string)                   | optional | The relation between an Object and a Subject.                                                                                         |
| subject   | [Subject](#ory-keto-v1beta-Subject) | optional | The subject related by this tuple. A Subject either represents a concrete subject id or a `SubjectSet` that expands to more Subjects. |

<a name="ory-keto-v1beta-RelationTuple"></a>

### RelationTuple

RelationTuple defines a relation between an Object and a Subject.

| Field     | Type                                | Label | Description                                                                                                                           |
| --------- | ----------------------------------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------- |
| namespace | [string](#string)                   |       | The namespace this relation tuple lives in.                                                                                           |
| object    | [string](#string)                   |       | The object related by this tuple. It is an object in the namespace of the tuple.                                                      |
| relation  | [string](#string)                   |       | The relation between an Object and a Subject.                                                                                         |
| subject   | [Subject](#ory-keto-v1beta-Subject) |       | The subject related by this tuple. A Subject either represents a concrete subject id or a `SubjectSet` that expands to more Subjects. |

<a name="ory-keto-v1beta-Subject"></a>

### Subject

Subject is either a concrete subject id or a `SubjectSet` expanding to more
Subjects.

| Field | Type                                      | Label | Description                                                                                                             |
| ----- | ----------------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- |
| id    | [string](#string)                         |       | A concrete id of the subject.                                                                                           |
| set   | [SubjectSet](#ory-keto-v1beta-SubjectSet) |       | A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-v1beta-SubjectSet"></a>

### SubjectSet

SubjectSet refers to all subjects who have the same `relation` on an `object`.

| Field     | Type              | Label | Description                                                              |
| --------- | ----------------- | ----- | ------------------------------------------------------------------------ |
| namespace | [string](#string) |       | The namespace of the object and relation referenced in this subject set. |
| object    | [string](#string) |       | The object related by this subject set.                                  |
| relation  | [string](#string) |       | The relation between the object and the subjects.                        |

<a name="ory_keto_v1beta_check_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/v1beta/check_service.proto

<a name="ory-keto-v1beta-CheckRequest"></a>

### CheckRequest

The request for a CheckService.Check RPC. Checks whether a specific subject is
related to an object.

| Field     | Type              | Label | Description                                          |
| --------- | ----------------- | ----- | ---------------------------------------------------- |
| namespace | [string](#string) |       | **Deprecated.** The namespace to evaluate the check. |

Note: If you use the expand-API and the check evaluates a RelationTuple
specifying a SubjectSet as subject or due to a rewrite rule in a namespace
config this check request may involve other namespaces automatically. | | object
| [string](#string) | | **Deprecated.** The related object in this check. | |
relation | [string](#string) | | **Deprecated.** The relation between the Object
and the Subject. | | subject | [Subject](#ory-keto-v1beta-Subject) | |
**Deprecated.** The related subject in this check. | | tuple |
[RelationTuple](#ory-keto-v1beta-RelationTuple) | | | | latest | [bool](#bool) |
| This field is not implemented yet and has no effect. &lt;!-- Set this field to
`true` in case your application needs to authorize depending on up to date ACLs,
also called a &#34;content-change check&#34;.

If set to `true` the `snaptoken` field is ignored, the check is evaluated at the
latest snapshot (globally consistent) and the response includes a snaptoken for
clients to store along with object contents that can be used for subsequent
checks of the same content version.

Example use case: - You need to authorize a user to modify/delete some resource
and it is unacceptable that if the permission to do that had just been revoked
some seconds ago so that the change had not yet been fully replicated to all
availability zones. --&gt; | | snaptoken | [string](#string) | | This field is
not implemented yet and has no effect. &lt;!-- Optional. Like reads, a check is
always evaluated at a consistent snapshot no earlier than the given snaptoken.

Leave this field blank if you want to evaluate the check based on eventually
consistent ACLs, benefiting from very low latency, but possibly slightly stale
results.

If the specified token is too old and no longer known, the server falls back as
if no snaptoken had been specified.

If not specified the server tries to evaluate the check on the best snapshot
version where it is very likely that ACLs had already been replicated to all
availability zones. --&gt; | | max_depth | [int32](#int32) | | The maximum depth
to search for a relation.

If the value is less than 1 or greater than the global max-depth then the global
max-depth will be used instead. |

<a name="ory-keto-v1beta-CheckResponse"></a>

### CheckResponse

The response for a CheckService.Check rpc.

| Field   | Type          | Label | Description                                                            |
| ------- | ------------- | ----- | ---------------------------------------------------------------------- |
| allowed | [bool](#bool) |       | Whether the specified subject (id) is related to the requested object. |

It is false by default if no ACL matches. | | snaptoken | [string](#string) | |
This field is not implemented yet and has no effect. &lt;!-- The last known
snapshot token ONLY specified if the request had not specified a snaptoken,
since this performed a &#34;content-change request&#34; and consistently fetched
the last known snapshot token.

This field is not set if the request had specified a snaptoken!

If set, clients should cache and use this token for subsequent requests to have
minimal latency, but allow slightly stale responses (only some milliseconds or
seconds). --&gt; |

<a name="ory-keto-v1beta-CheckService"></a>

### CheckService

The service that performs authorization checks based on the stored Access
Control Lists.

This service is part of the [read-APIs](../concepts/api-overview.mdx#read-apis).

| Method Name | Request Type                                  | Response Type                                   | Description                      |
| ----------- | --------------------------------------------- | ----------------------------------------------- | -------------------------------- |
| Check       | [CheckRequest](#ory-keto-v1beta-CheckRequest) | [CheckResponse](#ory-keto-v1beta-CheckResponse) | Performs an authorization check. |

## Scalar Value Types

| .proto Type                    | Notes                                                                                                                                           | C++    | Java       | Python      | Go      | C#         | PHP            | Ruby                           |
| ------------------------------ | ----------------------------------------------------------------------------------------------------------------------------------------------- | ------ | ---------- | ----------- | ------- | ---------- | -------------- | ------------------------------ |
| <a name="double" /> double     |                                                                                                                                                 | double | double     | float       | float64 | double     | float          | Float                          |
| <a name="float" /> float       |                                                                                                                                                 | float  | float      | float       | float32 | float      | float          | Float                          |
| <a name="int32" /> int32       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="int64" /> int64       | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="uint32" /> uint32     | Uses variable-length encoding.                                                                                                                  | uint32 | int        | int/long    | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64     | Uses variable-length encoding.                                                                                                                  | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s.                            | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64     | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s.                            | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="fixed32" /> fixed32   | Always four bytes. More efficient than uint32 if values are often greater than 2^28.                                                            | uint32 | int        | int         | uint32  | uint       | integer        | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64   | Always eight bytes. More efficient than uint64 if values are often greater than 2^56.                                                           | uint64 | long       | int/long    | uint64  | ulong      | integer/string | Bignum                         |
| <a name="sfixed32" /> sfixed32 | Always four bytes.                                                                                                                              | int32  | int        | int         | int32   | int        | integer        | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes.                                                                                                                             | int64  | long       | int/long    | int64   | long       | integer/string | Bignum                         |
| <a name="bool" /> bool         |                                                                                                                                                 | bool   | boolean    | boolean     | bool    | bool       | boolean        | TrueClass/FalseClass           |
| <a name="string" /> string     | A string must always contain UTF-8 encoded or 7-bit ASCII text.                                                                                 | string | String     | str/unicode | string  | string     | string         | String (UTF-8)                 |
| <a name="bytes" /> bytes       | May contain any arbitrary sequence of bytes.                                                                                                    | string | ByteString | str         | []byte  | ByteString | string         | String (ASCII-8BIT)            |
