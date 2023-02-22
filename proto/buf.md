# Protocol Documentation

<a name="top"></a>

## Table of Contents

- [google/api/http.proto](#google_api_http-proto)
  - [CustomHttpPattern](#google-api-CustomHttpPattern)
  - [Http](#google-api-Http)
  - [HttpRule](#google-api-HttpRule)
- [google/api/annotations.proto](#google_api_annotations-proto)
  - [File-level Extensions](#google_api_annotations-proto-extensions)
- [protoc-gen-openapiv2/options/openapiv2.proto](#protoc-gen-openapiv2_options_openapiv2-proto)

  - [Contact](#grpc-gateway-protoc_gen_openapiv2-options-Contact)
  - [ExternalDocumentation](#grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation)
  - [Header](#grpc-gateway-protoc_gen_openapiv2-options-Header)
  - [HeaderParameter](#grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter)
  - [Info](#grpc-gateway-protoc_gen_openapiv2-options-Info)
  - [Info.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Info-ExtensionsEntry)
  - [JSONSchema](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema)
  - [JSONSchema.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-ExtensionsEntry)
  - [JSONSchema.FieldConfiguration](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-FieldConfiguration)
  - [License](#grpc-gateway-protoc_gen_openapiv2-options-License)
  - [Operation](#grpc-gateway-protoc_gen_openapiv2-options-Operation)
  - [Operation.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Operation-ExtensionsEntry)
  - [Operation.ResponsesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Operation-ResponsesEntry)
  - [Parameters](#grpc-gateway-protoc_gen_openapiv2-options-Parameters)
  - [Response](#grpc-gateway-protoc_gen_openapiv2-options-Response)
  - [Response.ExamplesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-ExamplesEntry)
  - [Response.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-ExtensionsEntry)
  - [Response.HeadersEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-HeadersEntry)
  - [Schema](#grpc-gateway-protoc_gen_openapiv2-options-Schema)
  - [Scopes](#grpc-gateway-protoc_gen_openapiv2-options-Scopes)
  - [Scopes.ScopeEntry](#grpc-gateway-protoc_gen_openapiv2-options-Scopes-ScopeEntry)
  - [SecurityDefinitions](#grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions)
  - [SecurityDefinitions.SecurityEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions-SecurityEntry)
  - [SecurityRequirement](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement)
  - [SecurityRequirement.SecurityRequirementEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementEntry)
  - [SecurityRequirement.SecurityRequirementValue](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementValue)
  - [SecurityScheme](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme)
  - [SecurityScheme.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-ExtensionsEntry)
  - [Swagger](#grpc-gateway-protoc_gen_openapiv2-options-Swagger)
  - [Swagger.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Swagger-ExtensionsEntry)
  - [Swagger.ResponsesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Swagger-ResponsesEntry)
  - [Tag](#grpc-gateway-protoc_gen_openapiv2-options-Tag)
  - [Tag.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Tag-ExtensionsEntry)

  - [HeaderParameter.Type](#grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter-Type)
  - [JSONSchema.JSONSchemaSimpleTypes](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-JSONSchemaSimpleTypes)
  - [Scheme](#grpc-gateway-protoc_gen_openapiv2-options-Scheme)
  - [SecurityScheme.Flow](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Flow)
  - [SecurityScheme.In](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-In)
  - [SecurityScheme.Type](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Type)

- [protoc-gen-openapiv2/options/annotations.proto](#protoc-gen-openapiv2_options_annotations-proto)
  - [File-level Extensions](#protoc-gen-openapiv2_options_annotations-proto-extensions)
  - [File-level Extensions](#protoc-gen-openapiv2_options_annotations-proto-extensions)
  - [File-level Extensions](#protoc-gen-openapiv2_options_annotations-proto-extensions)
  - [File-level Extensions](#protoc-gen-openapiv2_options_annotations-proto-extensions)
  - [File-level Extensions](#protoc-gen-openapiv2_options_annotations-proto-extensions)
- [ory/keto/opl/v1alpha1/syntax_service.proto](#ory_keto_opl_v1alpha1_syntax_service-proto)

  - [CheckRequest](#ory-keto-opl-v1alpha1-CheckRequest)
  - [CheckResponse](#ory-keto-opl-v1alpha1-CheckResponse)
  - [ParseError](#ory-keto-opl-v1alpha1-ParseError)
  - [SourcePosition](#ory-keto-opl-v1alpha1-SourcePosition)

  - [SyntaxService](#ory-keto-opl-v1alpha1-SyntaxService)

- [google/api/visibility.proto](#google_api_visibility-proto)

  - [Visibility](#google-api-Visibility)
  - [VisibilityRule](#google-api-VisibilityRule)

  - [File-level Extensions](#google_api_visibility-proto-extensions)
  - [File-level Extensions](#google_api_visibility-proto-extensions)
  - [File-level Extensions](#google_api_visibility-proto-extensions)
  - [File-level Extensions](#google_api_visibility-proto-extensions)
  - [File-level Extensions](#google_api_visibility-proto-extensions)
  - [File-level Extensions](#google_api_visibility-proto-extensions)

- [google/api/field_behavior.proto](#google_api_field_behavior-proto)

  - [FieldBehavior](#google-api-FieldBehavior)

  - [File-level Extensions](#google_api_field_behavior-proto-extensions)

- [ory/keto/relation_tuples/v1alpha2/relation_tuples.proto](#ory_keto_relation_tuples_v1alpha2_relation_tuples-proto)
  - [RelationQuery](#ory-keto-relation_tuples-v1alpha2-RelationQuery)
  - [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple)
  - [Subject](#ory-keto-relation_tuples-v1alpha2-Subject)
  - [SubjectQuery](#ory-keto-relation_tuples-v1alpha2-SubjectQuery)
  - [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet)
  - [SubjectSetQuery](#ory-keto-relation_tuples-v1alpha2-SubjectSetQuery)
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

- [ory/keto/relation_tuples/v1alpha2/openapi.proto](#ory_keto_relation_tuples_v1alpha2_openapi-proto)
  - [ErrorResponse](#ory-keto-relation_tuples-v1alpha2-ErrorResponse)
  - [ErrorResponse.Error](#ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error)
  - [ErrorResponse.Error.DetailsEntry](#ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error-DetailsEntry)
- [ory/keto/relation_tuples/v1alpha2/read_service.proto](#ory_keto_relation_tuples_v1alpha2_read_service-proto)

  - [ListRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest)
  - [ListRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesRequest-Query)
  - [ListRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-ListRelationTuplesResponse)

  - [ReadService](#ory-keto-relation_tuples-v1alpha2-ReadService)

- [ory/keto/relation_tuples/v1alpha2/version.proto](#ory_keto_relation_tuples_v1alpha2_version-proto)

  - [GetVersionRequest](#ory-keto-relation_tuples-v1alpha2-GetVersionRequest)
  - [GetVersionResponse](#ory-keto-relation_tuples-v1alpha2-GetVersionResponse)

  - [VersionService](#ory-keto-relation_tuples-v1alpha2-VersionService)

- [validate/validate.proto](#validate_validate-proto)

  - [AnyRules](#validate-AnyRules)
  - [BoolRules](#validate-BoolRules)
  - [BytesRules](#validate-BytesRules)
  - [DoubleRules](#validate-DoubleRules)
  - [DurationRules](#validate-DurationRules)
  - [EnumRules](#validate-EnumRules)
  - [FieldRules](#validate-FieldRules)
  - [Fixed32Rules](#validate-Fixed32Rules)
  - [Fixed64Rules](#validate-Fixed64Rules)
  - [FloatRules](#validate-FloatRules)
  - [Int32Rules](#validate-Int32Rules)
  - [Int64Rules](#validate-Int64Rules)
  - [MapRules](#validate-MapRules)
  - [MessageRules](#validate-MessageRules)
  - [RepeatedRules](#validate-RepeatedRules)
  - [SFixed32Rules](#validate-SFixed32Rules)
  - [SFixed64Rules](#validate-SFixed64Rules)
  - [SInt32Rules](#validate-SInt32Rules)
  - [SInt64Rules](#validate-SInt64Rules)
  - [StringRules](#validate-StringRules)
  - [TimestampRules](#validate-TimestampRules)
  - [UInt32Rules](#validate-UInt32Rules)
  - [UInt64Rules](#validate-UInt64Rules)

  - [KnownRegex](#validate-KnownRegex)

  - [File-level Extensions](#validate_validate-proto-extensions)
  - [File-level Extensions](#validate_validate-proto-extensions)
  - [File-level Extensions](#validate_validate-proto-extensions)
  - [File-level Extensions](#validate_validate-proto-extensions)

- [ory/keto/relation_tuples/v1alpha2/write_service.proto](#ory_keto_relation_tuples_v1alpha2_write_service-proto)

  - [CreateRelationTupleRequest](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest)
  - [CreateRelationTupleRequest.Relationship](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest-Relationship)
  - [CreateRelationTupleResponse](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleResponse)
  - [DeleteRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest)
  - [DeleteRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest-Query)
  - [DeleteRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesResponse)
  - [RelationTupleDelta](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta)
  - [TransactRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesRequest)
  - [TransactRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesResponse)

  - [RelationTupleDelta.Action](#ory-keto-relation_tuples-v1alpha2-RelationTupleDelta-Action)

  - [WriteService](#ory-keto-relation_tuples-v1alpha2-WriteService)

- [Scalar Value Types](#scalar-value-types)

<a name="google_api_http-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## google/api/http.proto

<a name="google-api-CustomHttpPattern"></a>

### CustomHttpPattern

A custom pattern is used for defining custom HTTP verb.

| Field | Type              | Label | Description                           |
| ----- | ----------------- | ----- | ------------------------------------- |
| kind  | [string](#string) |       | The name of this custom HTTP verb.    |
| path  | [string](#string) |       | The path matched by this custom verb. |

<a name="google-api-Http"></a>

### Http

Defines the HTTP configuration for an API service. It contains a list of
[HttpRule][google.api.HttpRule], each specifying the mapping of an RPC method to
one or more HTTP REST API methods.

| Field | Type                             | Label    | Description                                                              |
| ----- | -------------------------------- | -------- | ------------------------------------------------------------------------ |
| rules | [HttpRule](#google-api-HttpRule) | repeated | A list of HTTP configuration rules that apply to individual API methods. |

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order.
| | fully_decode_reserved_expansion | [bool](#bool) | | When set to true, URL
path parameters will be fully URI-decoded except in cases of single segment
matches in reserved expansion, where &#34;%2F&#34; will be left encoded.

The default behavior is to not decode RFC 6570 reserved characters in multi
segment matches. |

<a name="google-api-HttpRule"></a>

### HttpRule

# gRPC Transcoding

gRPC Transcoding is a feature for mapping between a gRPC method and one or more
HTTP REST endpoints. It allows developers to build a single API service that
supports both gRPC APIs and REST APIs. Many systems, including
[Google APIs](https://github.com/googleapis/googleapis),
[Cloud Endpoints](https://cloud.google.com/endpoints),
[gRPC Gateway](https://github.com/grpc-ecosystem/grpc-gateway), and
[Envoy](https://github.com/envoyproxy/envoy) proxy support this feature and use
it for large scale production services.

`HttpRule` defines the schema of the gRPC/REST mapping. The mapping specifies
how different portions of the gRPC request message are mapped to the URL path,
URL query parameters, and HTTP request body. It also controls how the gRPC
response message is mapped to the HTTP response body. `HttpRule` is typically
specified as an `google.api.http` annotation on the gRPC method.

Each mapping specifies a URL path template and an HTTP method. The path template
may refer to one or more fields in the gRPC request message, as long as each
field is a non-repeated field with a primitive (non-message) type. The path
template controls how fields of the request message are mapped to the URL path.

Example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get: &#34;/v1/{name=messages/*}&#34;
        };
      }
    }
    message GetMessageRequest {
      string name = 1; // Mapped to URL path.
    }
    message Message {
      string text = 1; // The resource content.
    }

This enables an HTTP REST to gRPC mapping as below:

| HTTP                      | gRPC                                          |
| ------------------------- | --------------------------------------------- |
| `GET /v1/messages/123456` | `GetMessage(name: &#34;messages/123456&#34;)` |

Any fields in the request message which are not bound by the path template
automatically become HTTP query parameters if there is no HTTP request body. For
example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
            get:&#34;/v1/messages/{message_id}&#34;
        };
      }
    }
    message GetMessageRequest {
      message SubMessage {
        string subfield = 1;
      }
      string message_id = 1; // Mapped to URL path.
      int64 revision = 2;    // Mapped to URL query parameter `revision`.
      SubMessage sub = 3;    // Mapped to URL query parameter `sub.subfield`.
    }

This enables a HTTP JSON to RPC mapping as below:

| HTTP                                                      | gRPC |
| --------------------------------------------------------- | ---- |
| `GET /v1/messages/123456?revision=2&amp;sub.subfield=foo` |

`GetMessage(message_id: &#34;123456&#34; revision: 2 sub: SubMessage(subfield: &#34;foo&#34;))`

Note that fields which are mapped to URL query parameters must have a primitive
type or a repeated primitive type or a non-repeated message type. In the case of
a repeated type, the parameter can be repeated in the URL as
`...?param=A&amp;param=B`. In the case of a message type, each field of the
message is mapped to a separate parameter, such as
`...?foo.a=A&amp;foo.b=B&amp;foo.c=C`.

For HTTP methods that allow a request body, the `body` field specifies the
mapping. Consider a REST update method on the message resource collection:

    service Messaging {
      rpc UpdateMessage(UpdateMessageRequest) returns (Message) {
        option (google.api.http) = {
          patch: &#34;/v1/messages/{message_id}&#34;
          body: &#34;message&#34;
        };
      }
    }
    message UpdateMessageRequest {
      string message_id = 1; // mapped to the URL
      Message message = 2;   // mapped to the body
    }

The following HTTP JSON to RPC mapping is enabled, where the representation of
the JSON in the request body is determined by protos JSON encoding:

| HTTP                                                          | gRPC                       |
| ------------------------------------------------------------- | -------------------------- |
| `PATCH /v1/messages/123456 { &#34;text&#34;: &#34;Hi!&#34; }` | `UpdateMessage(message_id: |

&#34;123456&#34; message { text: &#34;Hi!&#34; })`

The special name `*` can be used in the body mapping to define that every field
not bound by the path template should be mapped to the request body. This
enables the following alternative definition of the update method:

    service Messaging {
      rpc UpdateMessage(Message) returns (Message) {
        option (google.api.http) = {
          patch: &#34;/v1/messages/{message_id}&#34;
          body: &#34;*&#34;
        };
      }
    }
    message Message {
      string message_id = 1;
      string text = 2;
    }

The following HTTP JSON to RPC mapping is enabled:

| HTTP                                                          | gRPC                       |
| ------------------------------------------------------------- | -------------------------- |
| `PATCH /v1/messages/123456 { &#34;text&#34;: &#34;Hi!&#34; }` | `UpdateMessage(message_id: |

&#34;123456&#34; text: &#34;Hi!&#34;)`

Note that when using `*` in the body mapping, it is not possible to have HTTP
parameters, as all fields not bound by the path end in the body. This makes this
option more rarely used in practice when defining REST APIs. The common usage of
`*` is in custom methods which don&#39;t use the URL at all for transferring
data.

It is possible to define multiple HTTP methods for one RPC by using the
`additional_bindings` option. Example:

    service Messaging {
      rpc GetMessage(GetMessageRequest) returns (Message) {
        option (google.api.http) = {
          get: &#34;/v1/messages/{message_id}&#34;
          additional_bindings {
            get: &#34;/v1/users/{user_id}/messages/{message_id}&#34;
          }
        };
      }
    }
    message GetMessageRequest {
      string message_id = 1;
      string user_id = 2;
    }

This enables the following two alternative HTTP JSON to RPC mappings:

| HTTP                               | gRPC                                          |
| ---------------------------------- | --------------------------------------------- |
| `GET /v1/messages/123456`          | `GetMessage(message_id: &#34;123456&#34;)`    |
| `GET /v1/users/me/messages/123456` | `GetMessage(user_id: &#34;me&#34; message_id: |

&#34;123456&#34;)`

## Rules for HTTP mapping

1. Leaf request fields (recursive expansion nested messages in the request
   message) are classified into three categories:
   - Fields referred by the path template. They are passed via the URL path.
   - Fields referred by the [HttpRule.body][google.api.HttpRule.body]. They are
     passed via the HTTP request body.
   - All other fields are passed via the URL query parameters, and the parameter
     name is the field path in the request message. A repeated field can be
     represented as multiple query parameters under the same name.
2. If [HttpRule.body][google.api.HttpRule.body] is &#34;\*&#34;, there is no URL
   query parameter, all fields are passed via URL path and HTTP request body.
3. If [HttpRule.body][google.api.HttpRule.body] is omitted, there is no HTTP
   request body, all fields are passed via URL path and URL query parameters.

### Path template syntax

    Template = &#34;/&#34; Segments [ Verb ] ;
    Segments = Segment { &#34;/&#34; Segment } ;
    Segment  = &#34;*&#34; | &#34;**&#34; | LITERAL | Variable ;
    Variable = &#34;{&#34; FieldPath [ &#34;=&#34; Segments ] &#34;}&#34; ;
    FieldPath = IDENT { &#34;.&#34; IDENT } ;
    Verb     = &#34;:&#34; LITERAL ;

The syntax `*` matches a single URL path segment. The syntax `**` matches zero
or more URL path segments, which must be the last part of the URL path except
the `Verb`.

The syntax `Variable` matches part of the URL path as specified by its template.
A variable template must not contain other variables. If a variable matches a
single path segment, its template may be omitted, e.g. `{var}` is equivalent to
`{var=*}`.

The syntax `LITERAL` matches literal text in the URL path. If the `LITERAL`
contains any reserved character, such characters should be percent-encoded
before the matching.

If a variable contains exactly one path segment, such as `&#34;{var}&#34;` or
`&#34;{var=*}&#34;`, when such a variable is expanded into a URL path on the
client side, all characters except `[-_.~0-9a-zA-Z]` are percent-encoded. The
server side does the reverse decoding. Such variables show up in the
[Discovery Document](https://developers.google.com/discovery/v1/reference/apis)
as `{var}`.

If a variable contains multiple path segments, such as `&#34;{var=foo/*}&#34;`
or `&#34;{var=**}&#34;`, when such a variable is expanded into a URL path on the
client side, all characters except `[-_.~/0-9a-zA-Z]` are percent-encoded. The
server side does the reverse decoding, except &#34;%2F&#34; and &#34;%2f&#34;
are left unchanged. Such variables show up in the
[Discovery Document](https://developers.google.com/discovery/v1/reference/apis)
as `{&#43;var}`.

## Using gRPC API Service Configuration

gRPC API Service Configuration (service config) is a configuration language for
configuring a gRPC service to become a user-facing product. The service config
is simply the YAML representation of the `google.api.Service` proto message.

As an alternative to annotating your proto file, you can configure gRPC
transcoding in your service config YAML files. You do this by specifying a
`HttpRule` that maps the gRPC method to a REST endpoint, achieving the same
effect as the proto annotation. This can be particularly useful if you have a
proto that is reused in multiple services. Note that any transcoding specified
in the service config will override any matching transcoding configuration in
the proto.

Example:

    http:
      rules:
        # Selects a gRPC method and applies HttpRule to it.
        - selector: example.v1.Messaging.GetMessage
          get: /v1/messages/{message_id}/{sub.subfield}

## Special notes

When gRPC Transcoding is used to map a gRPC to JSON REST endpoints, the proto to
JSON conversion must follow the
[proto3 specification](https://developers.google.com/protocol-buffers/docs/proto3#json).

While the single segment variable follows the semantics of
[RFC 6570](https://tools.ietf.org/html/rfc6570) Section 3.2.2 Simple String
Expansion, the multi segment variable **does not** follow RFC 6570 Section 3.2.3
Reserved Expansion. The reason is that the Reserved Expansion does not expand
special characters like `?` and `#`, which would lead to invalid URLs. As the
result, gRPC Transcoding uses a custom encoding for multi segment variables.

The path variables **must not** refer to any repeated or mapped field, because
client libraries are not capable of handling such variable expansion.

The path variables **must not** capture the leading &#34;/&#34; character. The
reason is that the most common use case &#34;{var}&#34; does not capture the
leading &#34;/&#34; character. For consistency, all path variables must share
the same behavior.

Repeated message fields must not be mapped to URL query parameters, because no
client library can support such complicated mapping.

If an API needs to use a JSON array for request or response body, it can map the
request or response body to a repeated field. However, some gRPC Transcoding
implementations may not support this feature.

| Field    | Type              | Label | Description                                  |
| -------- | ----------------- | ----- | -------------------------------------------- |
| selector | [string](#string) |       | Selects a method to which this rule applies. |

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| get | [string](#string) | | Maps to HTTP GET. Used for listing and getting
information about resources. | | put | [string](#string) | | Maps to HTTP PUT.
Used for replacing a resource. | | post | [string](#string) | | Maps to HTTP
POST. Used for creating a resource or performing an action. | | delete |
[string](#string) | | Maps to HTTP DELETE. Used for deleting a resource. | |
patch | [string](#string) | | Maps to HTTP PATCH. Used for updating a resource.
| | custom | [CustomHttpPattern](#google-api-CustomHttpPattern) | | The custom
pattern is used for specifying an HTTP method that is not included in the
`pattern` field, such as HEAD, or &#34;_&#34; to leave the HTTP method
unspecified for this rule. The wild-card rule is useful for services that
provide content to Web (HTML) clients. | | body | [string](#string) | | The name
of the request field whose value is mapped to the HTTP request body, or `_` for
mapping all request fields not captured by the path pattern to the HTTP body, or
omitted for not having any HTTP request body.

NOTE: the referred field must be present at the top-level of the request message
type. | | response_body | [string](#string) | | Optional. The name of the
response field whose value is mapped to the HTTP response body. When omitted,
the entire response message will be used as the HTTP response body.

NOTE: The referred field must be present at the top-level of the response
message type. | | additional_bindings | [HttpRule](#google-api-HttpRule) |
repeated | Additional HTTP bindings for the selector. Nested bindings must not
contain an `additional_bindings` field themselves (that is, the nesting may only
be one level deep). |

<a name="google_api_annotations-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## google/api/annotations.proto

<a name="google_api_annotations-proto-extensions"></a>

### File-level Extensions

| Extension | Type     | Base                           | Number   | Description     |
| --------- | -------- | ------------------------------ | -------- | --------------- |
| http      | HttpRule | .google.protobuf.MethodOptions | 72295728 | See `HttpRule`. |

<a name="protoc-gen-openapiv2_options_openapiv2-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## protoc-gen-openapiv2/options/openapiv2.proto

<a name="grpc-gateway-protoc_gen_openapiv2-options-Contact"></a>

### Contact

`Contact` is a representation of OpenAPI v2 specification&#39;s Contact object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#contactObject

Example:

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = { info: {
... contact: { name: &#34;gRPC-Gateway project&#34;; url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway&#34;; email:
&#34;none@example.com&#34;; }; ... }; ... };

| Field | Type              | Label | Description                                                                                      |
| ----- | ----------------- | ----- | ------------------------------------------------------------------------------------------------ |
| name  | [string](#string) |       | The identifying name of the contact person/organization.                                         |
| url   | [string](#string) |       | The URL pointing to the contact information. MUST be in the format of a URL.                     |
| email | [string](#string) |       | The email address of the contact person/organization. MUST be in the format of an email address. |

<a name="grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation"></a>

### ExternalDocumentation

`ExternalDocumentation` is a representation of OpenAPI v2 specification&#39;s
ExternalDocumentation object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#externalDocumentationObject

Example:

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = { ...
external_docs: { description: &#34;More about gRPC-Gateway&#34;; url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway&#34;; } ... };

| Field       | Type              | Label | Description                                                                                           |
| ----------- | ----------------- | ----- | ----------------------------------------------------------------------------------------------------- |
| description | [string](#string) |       | A short description of the target documentation. GFM syntax can be used for rich text representation. |
| url         | [string](#string) |       | The URL for the target documentation. Value MUST be in the format of a URL.                           |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Header"></a>

### Header

`Header` is a representation of OpenAPI v2 specification&#39;s Header object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#headerObject

| Field       | Type              | Label | Description                                                                                                                                                                                                                                               |
| ----------- | ----------------- | ----- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| description | [string](#string) |       | `Description` is a short description of the header.                                                                                                                                                                                                       |
| type        | [string](#string) |       | The type of the object. The value MUST be one of &#34;string&#34;, &#34;number&#34;, &#34;integer&#34;, or &#34;boolean&#34;. The &#34;array&#34; type is not supported.                                                                                  |
| format      | [string](#string) |       | `Format` The extending format for the previously mentioned type.                                                                                                                                                                                          |
| default     | [string](#string) |       | `Default` Declares the value of the header that the server will use if none is provided. See: https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-6.2. Unlike JSON Schema this value MUST conform to the defined type for the header. |
| pattern     | [string](#string) |       | &#39;Pattern&#39; See https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.2.3.                                                                                                                                                      |

<a name="grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter"></a>

### HeaderParameter

`HeaderParameter` a HTTP header parameter. See:
https://swagger.io/specification/v2/#parameter-object

| Field       | Type                                                                                    | Label | Description                                                                                                                                                                                                                                 |
| ----------- | --------------------------------------------------------------------------------------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name        | [string](#string)                                                                       |       | `Name` is the header name.                                                                                                                                                                                                                  |
| description | [string](#string)                                                                       |       | `Description` is a short description of the header.                                                                                                                                                                                         |
| type        | [HeaderParameter.Type](#grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter-Type) |       | `Type` is the type of the object. The value MUST be one of &#34;string&#34;, &#34;number&#34;, &#34;integer&#34;, or &#34;boolean&#34;. The &#34;array&#34; type is not supported. See: https://swagger.io/specification/v2/#parameterType. |
| format      | [string](#string)                                                                       |       | `Format` The extending format for the previously mentioned type.                                                                                                                                                                            |
| required    | [bool](#bool)                                                                           |       | `Required` indicates if the header is optional                                                                                                                                                                                              |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Info"></a>

### Info

`Info` is a representation of OpenAPI v2 specification&#39;s Info object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#infoObject

Example:

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = { info: {
title: &#34;Echo API&#34;; version: &#34;1.0&#34;; description: &#34;&#34;;
contact: { name: &#34;gRPC-Gateway project&#34;; url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway&#34;; email:
&#34;none@example.com&#34;; }; license: { name: &#34;BSD 3-Clause License&#34;;
url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE.txt&#34;;
}; }; ... };

| Field            | Type                                                                                    | Label    | Description                                                                                                                                                                                                                               |
| ---------------- | --------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| title            | [string](#string)                                                                       |          | The title of the application.                                                                                                                                                                                                             |
| description      | [string](#string)                                                                       |          | A short description of the application. GFM syntax can be used for rich text representation.                                                                                                                                              |
| terms_of_service | [string](#string)                                                                       |          | The Terms of Service for the API.                                                                                                                                                                                                         |
| contact          | [Contact](#grpc-gateway-protoc_gen_openapiv2-options-Contact)                           |          | The contact information for the exposed API.                                                                                                                                                                                              |
| license          | [License](#grpc-gateway-protoc_gen_openapiv2-options-License)                           |          | The license information for the exposed API.                                                                                                                                                                                              |
| version          | [string](#string)                                                                       |          | Provides the version of the application API (not to be confused with the specification version).                                                                                                                                          |
| extensions       | [Info.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Info-ExtensionsEntry) | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/ |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Info-ExtensionsEntry"></a>

### Info.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-JSONSchema"></a>

### JSONSchema

`JSONSchema` represents properties from JSON Schema taken, and as used, in the
OpenAPI v2 spec.

This includes changes made by OpenAPI v2.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject

See also: https://cswr.github.io/JsonSchema/spec/basic_types/,
https://github.com/json-schema-org/json-schema-spec/blob/master/schema.json

Example:

message SimpleMessage { option
(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_schema) = { json_schema: {
title: &#34;SimpleMessage&#34; description: &#34;A simple message.&#34;
required: [&#34;id&#34;] } };

// Id represents the message identifier. string id = 1; [
(grpc.gateway.protoc_gen_openapiv2.options.openapiv2_field) = { description:
&#34;The unique identifier of the simple message.&#34; }]; }

| Field               | Type                                                                                                            | Label    | Description                                                                                                                                                                                                                                                                                                                                                       |
| ------------------- | --------------------------------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| ref                 | [string](#string)                                                                                               |          | Ref is used to define an external reference to include in the message. This could be a fully qualified proto message reference, and that type must be imported into the protofile. If no message is identified, the Ref will be used verbatim in the output. For example: `ref: &#34;.google.protobuf.Timestamp&#34;`.                                            |
| title               | [string](#string)                                                                                               |          | The title of the schema.                                                                                                                                                                                                                                                                                                                                          |
| description         | [string](#string)                                                                                               |          | A short description of the schema.                                                                                                                                                                                                                                                                                                                                |
| default             | [string](#string)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| read_only           | [bool](#bool)                                                                                                   |          |                                                                                                                                                                                                                                                                                                                                                                   |
| example             | [string](#string)                                                                                               |          | A free-form property to include a JSON example of this field. This is copied verbatim to the output swagger.json. Quotes must be escaped. This property is the same for 2.0 and 3.0.0 https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/3.0.0.md#schemaObject https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject |
| multiple_of         | [double](#double)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| maximum             | [double](#double)                                                                                               |          | Maximum represents an inclusive upper limit for a numeric instance. The value of MUST be a number,                                                                                                                                                                                                                                                                |
| exclusive_maximum   | [bool](#bool)                                                                                                   |          |                                                                                                                                                                                                                                                                                                                                                                   |
| minimum             | [double](#double)                                                                                               |          | minimum represents an inclusive lower limit for a numeric instance. The value of MUST be a number,                                                                                                                                                                                                                                                                |
| exclusive_minimum   | [bool](#bool)                                                                                                   |          |                                                                                                                                                                                                                                                                                                                                                                   |
| max_length          | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| min_length          | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| pattern             | [string](#string)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| max_items           | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| min_items           | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| unique_items        | [bool](#bool)                                                                                                   |          |                                                                                                                                                                                                                                                                                                                                                                   |
| max_properties      | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| min_properties      | [uint64](#uint64)                                                                                               |          |                                                                                                                                                                                                                                                                                                                                                                   |
| required            | [string](#string)                                                                                               | repeated |                                                                                                                                                                                                                                                                                                                                                                   |
| array               | [string](#string)                                                                                               | repeated | Items in &#39;array&#39; must be unique.                                                                                                                                                                                                                                                                                                                          |
| type                | [JSONSchema.JSONSchemaSimpleTypes](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-JSONSchemaSimpleTypes) | repeated |                                                                                                                                                                                                                                                                                                                                                                   |
| format              | [string](#string)                                                                                               |          | `Format`                                                                                                                                                                                                                                                                                                                                                          |
| enum                | [string](#string)                                                                                               | repeated | Items in `enum` must be unique https://tools.ietf.org/html/draft-fge-json-schema-validation-00#section-5.5.1                                                                                                                                                                                                                                                      |
| field_configuration | [JSONSchema.FieldConfiguration](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-FieldConfiguration)       |          | Additional field level properties used when generating the OpenAPI v2 file.                                                                                                                                                                                                                                                                                       |
| extensions          | [JSONSchema.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-ExtensionsEntry)             | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/                                                                                                                         |

<a name="grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-ExtensionsEntry"></a>

### JSONSchema.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-FieldConfiguration"></a>

### JSONSchema.FieldConfiguration

&#39;FieldConfiguration&#39; provides additional field level properties used
when generating the OpenAPI v2 file. These properties are not defined by
OpenAPIv2, but they are used to control the generation.

| Field           | Type              | Label | Description                                                                                                                                                                                                                                       |
| --------------- | ----------------- | ----- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| path_param_name | [string](#string) |       | Alternative parameter name when used as path parameter. If set, this will be used as the complete parameter name when this field is used as a path parameter. Use this to avoid having auto generated path parameter names for overlapping paths. |

<a name="grpc-gateway-protoc_gen_openapiv2-options-License"></a>

### License

`License` is a representation of OpenAPI v2 specification&#39;s License object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#licenseObject

Example:

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = { info: {
... license: { name: &#34;BSD 3-Clause License&#34;; url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE.txt&#34;;
}; ... }; ... };

| Field | Type              | Label | Description                                                            |
| ----- | ----------------- | ----- | ---------------------------------------------------------------------- |
| name  | [string](#string) |       | The license name used for the API.                                     |
| url   | [string](#string) |       | A URL to the license used for the API. MUST be in the format of a URL. |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Operation"></a>

### Operation

`Operation` is a representation of OpenAPI v2 specification&#39;s Operation
object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#operationObject

Example:

service EchoService { rpc Echo(SimpleMessage) returns (SimpleMessage) { option
(google.api.http) = { get: &#34;/v1/example/echo/{id}&#34; };

     option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_operation) = {
       summary: &#34;Get a message.&#34;;
       operation_id: &#34;getMessage&#34;;
       tags: &#34;echo&#34;;
       responses: {
         key: &#34;200&#34;
           value: {
           description: &#34;OK&#34;;
         }
       }
     };

} }

| Field         | Type                                                                                              | Label    | Description                                                                                                                                                                                                                                                                                                                                               |
| ------------- | ------------------------------------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| tags          | [string](#string)                                                                                 | repeated | A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.                                                                                                                                                                                                                    |
| summary       | [string](#string)                                                                                 |          | A short summary of what the operation does. For maximum readability in the swagger-ui, this field SHOULD be less than 120 characters.                                                                                                                                                                                                                     |
| description   | [string](#string)                                                                                 |          | A verbose explanation of the operation behavior. GFM syntax can be used for rich text representation.                                                                                                                                                                                                                                                     |
| external_docs | [ExternalDocumentation](#grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation)         |          | Additional external documentation for this operation.                                                                                                                                                                                                                                                                                                     |
| operation_id  | [string](#string)                                                                                 |          | Unique string used to identify the operation. The id MUST be unique among all operations described in the API. Tools and libraries MAY use the operationId to uniquely identify an operation, therefore, it is recommended to follow common programming naming conventions.                                                                               |
| consumes      | [string](#string)                                                                                 | repeated | A list of MIME types the operation can consume. This overrides the consumes definition at the OpenAPI Object. An empty value MAY be used to clear the global definition. Value MUST be as described under Mime Types.                                                                                                                                     |
| produces      | [string](#string)                                                                                 | repeated | A list of MIME types the operation can produce. This overrides the produces definition at the OpenAPI Object. An empty value MAY be used to clear the global definition. Value MUST be as described under Mime Types.                                                                                                                                     |
| responses     | [Operation.ResponsesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Operation-ResponsesEntry)   | repeated | The list of possible responses as they are returned from executing this operation.                                                                                                                                                                                                                                                                        |
| schemes       | [Scheme](#grpc-gateway-protoc_gen_openapiv2-options-Scheme)                                       | repeated | The transfer protocol for the operation. Values MUST be from the list: &#34;http&#34;, &#34;https&#34;, &#34;ws&#34;, &#34;wss&#34;. The value overrides the OpenAPI Object schemes definition.                                                                                                                                                           |
| deprecated    | [bool](#bool)                                                                                     |          | Declares this operation to be deprecated. Usage of the declared operation should be refrained. Default value is false.                                                                                                                                                                                                                                    |
| security      | [SecurityRequirement](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement)             | repeated | A declaration of which security schemes are applied for this operation. The list of values describes alternative security schemes that can be used (that is, there is a logical OR between the security requirements). This definition overrides any declared top-level security. To remove a top-level security declaration, an empty array can be used. |
| extensions    | [Operation.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Operation-ExtensionsEntry) | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/                                                                                                                 |
| parameters    | [Parameters](#grpc-gateway-protoc_gen_openapiv2-options-Parameters)                               |          | Custom parameters such as HTTP request headers. See: https://swagger.io/docs/specification/2-0/describing-parameters/ and https://swagger.io/specification/v2/#parameter-object.                                                                                                                                                                          |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Operation-ExtensionsEntry"></a>

### Operation.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Operation-ResponsesEntry"></a>

### Operation.ResponsesEntry

| Field | Type                                                            | Label | Description |
| ----- | --------------------------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                                               |       |             |
| value | [Response](#grpc-gateway-protoc_gen_openapiv2-options-Response) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Parameters"></a>

### Parameters

`Parameters` is a representation of OpenAPI v2 specification&#39;s parameters
object. Note: This technically breaks compatibility with the OpenAPI 2
definition structure as we only allow header parameters to be set here since we
do not want users specifying custom non-header parameters beyond those inferred
from the Protobuf schema. See:
https://swagger.io/specification/v2/#parameter-object

| Field   | Type                                                                          | Label    | Description                                                                                                                             |
| ------- | ----------------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------------------------------------------------------- |
| headers | [HeaderParameter](#grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter) | repeated | `Headers` is one or more HTTP header parameter. See: https://swagger.io/docs/specification/2-0/describing-parameters/#header-parameters |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Response"></a>

### Response

`Response` is a representation of OpenAPI v2 specification&#39;s Response
object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#responseObject

| Field       | Type                                                                                            | Label    | Description                                                                                                                                                                                                                               |
| ----------- | ----------------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| description | [string](#string)                                                                               |          | `Description` is a short description of the response. GFM syntax can be used for rich text representation.                                                                                                                                |
| schema      | [Schema](#grpc-gateway-protoc_gen_openapiv2-options-Schema)                                     |          | `Schema` optionally defines the structure of the response. If `Schema` is not provided, it means there is no content to the response.                                                                                                     |
| headers     | [Response.HeadersEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-HeadersEntry)       | repeated | `Headers` A list of headers that are sent with the response. `Header` name is expected to be a string in the canonical format of the MIME header key See: https://golang.org/pkg/net/textproto/#CanonicalMIMEHeaderKey                    |
| examples    | [Response.ExamplesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-ExamplesEntry)     | repeated | `Examples` gives per-mimetype response examples. See: https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#example-object                                                                                              |
| extensions  | [Response.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Response-ExtensionsEntry) | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/ |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Response-ExamplesEntry"></a>

### Response.ExamplesEntry

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| key   | [string](#string) |       |             |
| value | [string](#string) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Response-ExtensionsEntry"></a>

### Response.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Response-HeadersEntry"></a>

### Response.HeadersEntry

| Field | Type                                                        | Label | Description |
| ----- | ----------------------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                                           |       |             |
| value | [Header](#grpc-gateway-protoc_gen_openapiv2-options-Header) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Schema"></a>

### Schema

`Schema` is a representation of OpenAPI v2 specification&#39;s Schema object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#schemaObject

| Field         | Type                                                                                      | Label | Description                                                                                                                                                                                                                                                                                                                                        |
| ------------- | ----------------------------------------------------------------------------------------- | ----- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| json_schema   | [JSONSchema](#grpc-gateway-protoc_gen_openapiv2-options-JSONSchema)                       |       |                                                                                                                                                                                                                                                                                                                                                    |
| discriminator | [string](#string)                                                                         |       | Adds support for polymorphism. The discriminator is the schema property name that is used to differentiate between other schema that inherit this schema. The property name used MUST be defined at this schema and it MUST be in the required property list. When used, the value MUST be the name of this schema or any schema that inherits it. |
| read_only     | [bool](#bool)                                                                             |       | Relevant only for Schema &#34;properties&#34; definitions. Declares the property as &#34;read only&#34;. This means that it MAY be sent as part of a response but MUST NOT be sent as part of the request. Properties marked as readOnly being true SHOULD NOT be in the required list of the defined schema. Default value is false.              |
| external_docs | [ExternalDocumentation](#grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation) |       | Additional external documentation for this schema.                                                                                                                                                                                                                                                                                                 |
| example       | [string](#string)                                                                         |       | A free-form property to include an example of an instance for this schema in JSON. This is copied verbatim to the output.                                                                                                                                                                                                                          |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Scopes"></a>

### Scopes

`Scopes` is a representation of OpenAPI v2 specification&#39;s Scopes object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#scopesObject

Lists the available scopes for an OAuth2 security scheme.

| Field | Type                                                                              | Label    | Description                                                                                 |
| ----- | --------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------- |
| scope | [Scopes.ScopeEntry](#grpc-gateway-protoc_gen_openapiv2-options-Scopes-ScopeEntry) | repeated | Maps between a name of a scope to a short description of it (as the value of the property). |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Scopes-ScopeEntry"></a>

### Scopes.ScopeEntry

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| key   | [string](#string) |       |             |
| value | [string](#string) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions"></a>

### SecurityDefinitions

`SecurityDefinitions` is a representation of OpenAPI v2 specification&#39;s
Security Definitions object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityDefinitionsObject

A declaration of the security schemes available to be used in the specification.
This does not enforce the security schemes on the operations and only serves to
provide the relevant details for each scheme.

| Field    | Type                                                                                                              | Label    | Description                                                                             |
| -------- | ----------------------------------------------------------------------------------------------------------------- | -------- | --------------------------------------------------------------------------------------- |
| security | [SecurityDefinitions.SecurityEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions-SecurityEntry) | repeated | A single security scheme definition, mapping a &#34;name&#34; to the scheme it defines. |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions-SecurityEntry"></a>

### SecurityDefinitions.SecurityEntry

| Field | Type                                                                        | Label | Description |
| ----- | --------------------------------------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                                                           |       |             |
| value | [SecurityScheme](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement"></a>

### SecurityRequirement

`SecurityRequirement` is a representation of OpenAPI v2 specification&#39;s
Security Requirement object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securityRequirementObject

Lists the required security schemes to execute this operation. The object can
have multiple security schemes declared in it which are all required (that is,
there is a logical AND between the schemes).

The name used for each property MUST correspond to a security scheme declared in
the Security Definitions.

| Field                | Type                                                                                                                                    | Label    | Description                                                                                                                                                                                                                                                                     |
| -------------------- | --------------------------------------------------------------------------------------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| security_requirement | [SecurityRequirement.SecurityRequirementEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementEntry) | repeated | Each name must correspond to a security scheme which is declared in the Security Definitions. If the security scheme is of type &#34;oauth2&#34;, then the value is a list of scope names required for the execution. For other security scheme types, the array MUST be empty. |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementEntry"></a>

### SecurityRequirement.SecurityRequirementEntry

| Field | Type                                                                                                                                    | Label | Description |
| ----- | --------------------------------------------------------------------------------------------------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                                                                                                                       |       |             |
| value | [SecurityRequirement.SecurityRequirementValue](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementValue) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement-SecurityRequirementValue"></a>

### SecurityRequirement.SecurityRequirementValue

If the security scheme is of type &#34;oauth2&#34;, then the value is a list of
scope names required for the execution. For other security scheme types, the
array MUST be empty.

| Field | Type              | Label    | Description |
| ----- | ----------------- | -------- | ----------- |
| scope | [string](#string) | repeated |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme"></a>

### SecurityScheme

`SecurityScheme` is a representation of OpenAPI v2 specification&#39;s Security
Scheme object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#securitySchemeObject

Allows the definition of a security scheme that can be used by the operations.
Supported schemes are basic authentication, an API key (either as a header or as
a query parameter) and OAuth2&#39;s common flows (implicit, password,
application and access code).

| Field             | Type                                                                                                        | Label    | Description                                                                                                                                                                                                                               |
| ----------------- | ----------------------------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| type              | [SecurityScheme.Type](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Type)                       |          | The type of the security scheme. Valid values are &#34;basic&#34;, &#34;apiKey&#34; or &#34;oauth2&#34;.                                                                                                                                  |
| description       | [string](#string)                                                                                           |          | A short description for security scheme.                                                                                                                                                                                                  |
| name              | [string](#string)                                                                                           |          | The name of the header or query parameter to be used. Valid for apiKey.                                                                                                                                                                   |
| in                | [SecurityScheme.In](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-In)                           |          | The location of the API key. Valid values are &#34;query&#34; or &#34;header&#34;. Valid for apiKey.                                                                                                                                      |
| flow              | [SecurityScheme.Flow](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Flow)                       |          | The flow used by the OAuth2 security scheme. Valid values are &#34;implicit&#34;, &#34;password&#34;, &#34;application&#34; or &#34;accessCode&#34;. Valid for oauth2.                                                                    |
| authorization_url | [string](#string)                                                                                           |          | The authorization URL to be used for this flow. This SHOULD be in the form of a URL. Valid for oauth2/implicit and oauth2/accessCode.                                                                                                     |
| token_url         | [string](#string)                                                                                           |          | The token URL to be used for this flow. This SHOULD be in the form of a URL. Valid for oauth2/password, oauth2/application and oauth2/accessCode.                                                                                         |
| scopes            | [Scopes](#grpc-gateway-protoc_gen_openapiv2-options-Scopes)                                                 |          | The available scopes for the OAuth2 security scheme. Valid for oauth2.                                                                                                                                                                    |
| extensions        | [SecurityScheme.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-ExtensionsEntry) | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/ |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-ExtensionsEntry"></a>

### SecurityScheme.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Swagger"></a>

### Swagger

`Swagger` is a representation of OpenAPI v2 specification&#39;s Swagger object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#swaggerObject

Example:

option (grpc.gateway.protoc_gen_openapiv2.options.openapiv2_swagger) = { info: {
title: &#34;Echo API&#34;; version: &#34;1.0&#34;; description: &#34;&#34;;
contact: { name: &#34;gRPC-Gateway project&#34;; url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway&#34;; email:
&#34;none@example.com&#34;; }; license: { name: &#34;BSD 3-Clause License&#34;;
url:
&#34;https://github.com/grpc-ecosystem/grpc-gateway/blob/main/LICENSE.txt&#34;;
}; }; schemes: HTTPS; consumes: &#34;application/json&#34;; produces:
&#34;application/json&#34;; };

| Field                | Type                                                                                          | Label    | Description                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                      |
| -------------------- | --------------------------------------------------------------------------------------------- | -------- | -------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| swagger              | [string](#string)                                                                             |          | Specifies the OpenAPI Specification version being used. It can be used by the OpenAPI UI and other clients to interpret the API listing. The value MUST be &#34;2.0&#34;.                                                                                                                                                                                                                                                                                                                                                                                                        |
| info                 | [Info](#grpc-gateway-protoc_gen_openapiv2-options-Info)                                       |          | Provides metadata about the API. The metadata can be used by the clients if needed.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                              |
| host                 | [string](#string)                                                                             |          | The host (name or ip) serving the API. This MUST be the host only and does not include the scheme nor sub-paths. It MAY include a port. If the host is not included, the host serving the documentation is to be used (including the port). The host does not support path templating.                                                                                                                                                                                                                                                                                           |
| base_path            | [string](#string)                                                                             |          | The base path on which the API is served, which is relative to the host. If it is not included, the API is served directly under the host. The value MUST start with a leading slash (/). The basePath does not support path templating. Note that using `base_path` does not change the endpoint paths that are generated in the resulting OpenAPI file. If you wish to use `base_path` with relatively generated OpenAPI paths, the `base_path` prefix must be manually removed from your `google.api.http` paths and your code changed to serve the API from the `base_path`. |
| schemes              | [Scheme](#grpc-gateway-protoc_gen_openapiv2-options-Scheme)                                   | repeated | The transfer protocol of the API. Values MUST be from the list: &#34;http&#34;, &#34;https&#34;, &#34;ws&#34;, &#34;wss&#34;. If the schemes is not included, the default scheme to be used is the one used to access the OpenAPI definition itself.                                                                                                                                                                                                                                                                                                                             |
| consumes             | [string](#string)                                                                             | repeated | A list of MIME types the APIs can consume. This is global to all APIs but can be overridden on specific API calls. Value MUST be as described under Mime Types.                                                                                                                                                                                                                                                                                                                                                                                                                  |
| produces             | [string](#string)                                                                             | repeated | A list of MIME types the APIs can produce. This is global to all APIs but can be overridden on specific API calls. Value MUST be as described under Mime Types.                                                                                                                                                                                                                                                                                                                                                                                                                  |
| responses            | [Swagger.ResponsesEntry](#grpc-gateway-protoc_gen_openapiv2-options-Swagger-ResponsesEntry)   | repeated | An object to hold responses that can be used across operations. This property does not define global responses for all operations.                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| security_definitions | [SecurityDefinitions](#grpc-gateway-protoc_gen_openapiv2-options-SecurityDefinitions)         |          | Security scheme definitions that can be used across the specification.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| security             | [SecurityRequirement](#grpc-gateway-protoc_gen_openapiv2-options-SecurityRequirement)         | repeated | A declaration of which security schemes are applied for the API as a whole. The list of values describes alternative security schemes that can be used (that is, there is a logical OR between the security requirements). Individual operations can override this definition.                                                                                                                                                                                                                                                                                                   |
| tags                 | [Tag](#grpc-gateway-protoc_gen_openapiv2-options-Tag)                                         | repeated | A list of tags for API documentation control. Tags can be used for logical grouping of operations by resources or any other qualifier.                                                                                                                                                                                                                                                                                                                                                                                                                                           |
| external_docs        | [ExternalDocumentation](#grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation)     |          | Additional external documentation.                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                                               |
| extensions           | [Swagger.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Swagger-ExtensionsEntry) | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/                                                                                                                                                                                                                                                                                                                                        |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Swagger-ExtensionsEntry"></a>

### Swagger.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Swagger-ResponsesEntry"></a>

### Swagger.ResponsesEntry

| Field | Type                                                            | Label | Description |
| ----- | --------------------------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                                               |       |             |
| value | [Response](#grpc-gateway-protoc_gen_openapiv2-options-Response) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Tag"></a>

### Tag

`Tag` is a representation of OpenAPI v2 specification&#39;s Tag object.

See:
https://github.com/OAI/OpenAPI-Specification/blob/3.0.0/versions/2.0.md#tagObject

| Field         | Type                                                                                      | Label    | Description                                                                                                                                                                                                                               |
| ------------- | ----------------------------------------------------------------------------------------- | -------- | ----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| name          | [string](#string)                                                                         |          | The name of the tag. Use it to allow override of the name of a global Tag object, then use that name to reference the tag throughout the OpenAPI file.                                                                                    |
| description   | [string](#string)                                                                         |          | A short description for the tag. GFM syntax can be used for rich text representation.                                                                                                                                                     |
| external_docs | [ExternalDocumentation](#grpc-gateway-protoc_gen_openapiv2-options-ExternalDocumentation) |          | Additional external documentation for this tag.                                                                                                                                                                                           |
| extensions    | [Tag.ExtensionsEntry](#grpc-gateway-protoc_gen_openapiv2-options-Tag-ExtensionsEntry)     | repeated | Custom properties that start with &#34;x-&#34; such as &#34;x-foo&#34; used to describe extra functionality that is not covered by the standard OpenAPI Specification. See: https://swagger.io/docs/specification/2-0/swagger-extensions/ |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Tag-ExtensionsEntry"></a>

### Tag.ExtensionsEntry

| Field | Type                                            | Label | Description |
| ----- | ----------------------------------------------- | ----- | ----------- |
| key   | [string](#string)                               |       |             |
| value | [google.protobuf.Value](#google-protobuf-Value) |       |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-HeaderParameter-Type"></a>

### HeaderParameter.Type

`Type` is a a supported HTTP header type. See
https://swagger.io/specification/v2/#parameterType.

| Name    | Number | Description |
| ------- | ------ | ----------- |
| UNKNOWN | 0      |             |
| STRING  | 1      |             |
| NUMBER  | 2      |             |
| INTEGER | 3      |             |
| BOOLEAN | 4      |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-JSONSchema-JSONSchemaSimpleTypes"></a>

### JSONSchema.JSONSchemaSimpleTypes

| Name    | Number | Description |
| ------- | ------ | ----------- |
| UNKNOWN | 0      |             |
| ARRAY   | 1      |             |
| BOOLEAN | 2      |             |
| INTEGER | 3      |             |
| NULL    | 4      |             |
| NUMBER  | 5      |             |
| OBJECT  | 6      |             |
| STRING  | 7      |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-Scheme"></a>

### Scheme

Scheme describes the schemes supported by the OpenAPI Swagger and Operation
objects.

| Name    | Number | Description |
| ------- | ------ | ----------- |
| UNKNOWN | 0      |             |
| HTTP    | 1      |             |
| HTTPS   | 2      |             |
| WS      | 3      |             |
| WSS     | 4      |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Flow"></a>

### SecurityScheme.Flow

The flow used by the OAuth2 security scheme. Valid values are
&#34;implicit&#34;, &#34;password&#34;, &#34;application&#34; or
&#34;accessCode&#34;.

| Name             | Number | Description |
| ---------------- | ------ | ----------- |
| FLOW_INVALID     | 0      |             |
| FLOW_IMPLICIT    | 1      |             |
| FLOW_PASSWORD    | 2      |             |
| FLOW_APPLICATION | 3      |             |
| FLOW_ACCESS_CODE | 4      |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-In"></a>

### SecurityScheme.In

The location of the API key. Valid values are &#34;query&#34; or
&#34;header&#34;.

| Name       | Number | Description |
| ---------- | ------ | ----------- |
| IN_INVALID | 0      |             |
| IN_QUERY   | 1      |             |
| IN_HEADER  | 2      |             |

<a name="grpc-gateway-protoc_gen_openapiv2-options-SecurityScheme-Type"></a>

### SecurityScheme.Type

The type of the security scheme. Valid values are &#34;basic&#34;,
&#34;apiKey&#34; or &#34;oauth2&#34;.

| Name         | Number | Description |
| ------------ | ------ | ----------- |
| TYPE_INVALID | 0      |             |
| TYPE_BASIC   | 1      |             |
| TYPE_API_KEY | 2      |             |
| TYPE_OAUTH2  | 3      |             |

<a name="protoc-gen-openapiv2_options_annotations-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## protoc-gen-openapiv2/options/annotations.proto

<a name="protoc-gen-openapiv2_options_annotations-proto-extensions"></a>

### File-level Extensions

| Extension       | Type       | Base                          | Number | Description                                                                            |
| --------------- | ---------- | ----------------------------- | ------ | -------------------------------------------------------------------------------------- |
| openapiv2_field | JSONSchema | .google.protobuf.FieldOptions | 1042   | ID assigned by protobuf-global-extension-registry@google.com for gRPC-Gateway project. |

All IDs are the same, as assigned. It is okay that they are the same, as they
extend different descriptor messages. | | openapiv2_swagger | Swagger |
.google.protobuf.FileOptions | 1042 | ID assigned by
protobuf-global-extension-registry@google.com for gRPC-Gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they
extend different descriptor messages. | | openapiv2_schema | Schema |
.google.protobuf.MessageOptions | 1042 | ID assigned by
protobuf-global-extension-registry@google.com for gRPC-Gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they
extend different descriptor messages. | | openapiv2_operation | Operation |
.google.protobuf.MethodOptions | 1042 | ID assigned by
protobuf-global-extension-registry@google.com for gRPC-Gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they
extend different descriptor messages. | | openapiv2_tag | Tag |
.google.protobuf.ServiceOptions | 1042 | ID assigned by
protobuf-global-extension-registry@google.com for gRPC-Gateway project.

All IDs are the same, as assigned. It is okay that they are the same, as they
extend different descriptor messages. |

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

| Field  | Type                                            | Label    | Description |
| ------ | ----------------------------------------------- | -------- | ----------- |
| errors | [ParseError](#ory-keto-opl-v1alpha1-ParseError) | repeated |             |

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

<a name="google_api_visibility-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## google/api/visibility.proto

<a name="google-api-Visibility"></a>

### Visibility

`Visibility` restricts service consumer&#39;s access to service elements, such
as whether an application can call a visibility-restricted method. The
restriction is expressed by applying visibility labels on service elements. The
visibility labels are elsewhere linked to service consumers.

A service can define multiple visibility labels, but a service consumer should
be granted at most one visibility label. Multiple visibility labels for a single
service consumer are not supported.

If an element and all its parents have no visibility label, its visibility is
unconditionally granted.

Example:

    visibility:
      rules:
      - selector: google.calendar.Calendar.EnhancedSearch
        restriction: PREVIEW
      - selector: google.calendar.Calendar.Delegate
        restriction: INTERNAL

Here, all methods are publicly visible except for the restricted methods
EnhancedSearch and Delegate.

| Field | Type                                         | Label    | Description                                                       |
| ----- | -------------------------------------------- | -------- | ----------------------------------------------------------------- |
| rules | [VisibilityRule](#google-api-VisibilityRule) | repeated | A list of visibility rules that apply to individual API elements. |

**NOTE:** All service configuration rules follow &#34;last one wins&#34; order.
|

<a name="google-api-VisibilityRule"></a>

### VisibilityRule

A visibility rule provides visibility configuration for an individual API
element.

| Field    | Type              | Label | Description                                                                |
| -------- | ----------------- | ----- | -------------------------------------------------------------------------- |
| selector | [string](#string) |       | Selects methods, messages, fields, enums, etc. to which this rule applies. |

Refer to [selector][google.api.DocumentationRule.selector] for syntax details. |
| restriction | [string](#string) | | A comma-separated list of visibility
labels that apply to the `selector`. Any of the listed labels can be used to
grant the visibility.

If a rule has multiple labels, removing one of the labels but not all of them
can break clients.

Example:

visibility: rules: - selector: google.calendar.Calendar.EnhancedSearch
restriction: INTERNAL, PREVIEW

Removing INTERNAL from this restriction will break clients that rely on this
method and only had access to it through INTERNAL. |

<a name="google_api_visibility-proto-extensions"></a>

### File-level Extensions

| Extension          | Type           | Base                              | Number   | Description           |
| ------------------ | -------------- | --------------------------------- | -------- | --------------------- |
| enum_visibility    | VisibilityRule | .google.protobuf.EnumOptions      | 72295727 | See `VisibilityRule`. |
| value_visibility   | VisibilityRule | .google.protobuf.EnumValueOptions | 72295727 | See `VisibilityRule`. |
| field_visibility   | VisibilityRule | .google.protobuf.FieldOptions     | 72295727 | See `VisibilityRule`. |
| message_visibility | VisibilityRule | .google.protobuf.MessageOptions   | 72295727 | See `VisibilityRule`. |
| method_visibility  | VisibilityRule | .google.protobuf.MethodOptions    | 72295727 | See `VisibilityRule`. |
| api_visibility     | VisibilityRule | .google.protobuf.ServiceOptions   | 72295727 | See `VisibilityRule`. |

<a name="google_api_field_behavior-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## google/api/field_behavior.proto

<a name="google-api-FieldBehavior"></a>

### FieldBehavior

An indicator of the behavior of a given field (for example, that a field is
required in requests, or given as output but ignored as input). This **does
not** change the behavior in protocol buffers itself; it only denotes the
behavior and may affect how API tooling handles the field.

Note: This enum **may** receive new values in the future.

| Name                       | Number | Description                                                                                                                                                                                                                                                         |
| -------------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| FIELD_BEHAVIOR_UNSPECIFIED | 0      | Conventional default for enums. Do not use this.                                                                                                                                                                                                                    |
| OPTIONAL                   | 1      | Specifically denotes a field as optional. While all fields in protocol buffers are optional, this may be specified for emphasis if appropriate.                                                                                                                     |
| REQUIRED                   | 2      | Denotes a field as required. This indicates that the field **must** be provided as part of the request, and failure to do so will cause an error (usually `INVALID_ARGUMENT`).                                                                                      |
| OUTPUT_ONLY                | 3      | Denotes a field as output only. This indicates that the field is provided in responses, but including the field in a request does nothing (the server _must_ ignore it and _must not_ throw an error as a result of the field&#39;s presence).                      |
| INPUT_ONLY                 | 4      | Denotes a field as input only. This indicates that the field is provided in requests, and the corresponding field is not included in output.                                                                                                                        |
| IMMUTABLE                  | 5      | Denotes a field as immutable. This indicates that the field may be set once in a request to create a resource, but may not be changed thereafter.                                                                                                                   |
| UNORDERED_LIST             | 6      | Denotes that a (repeated) field is an unordered list. This indicates that the service may provide the elements of the list in any arbitrary order, rather than the order the user originally provided. Additionally, the list&#39;s order may or may not be stable. |
| NON_EMPTY_DEFAULT          | 7      | Denotes that this field returns a non-empty default value if not set. This indicates that if the user provides the empty value in a request, a non-empty value will be returned. The user will not be aware of what non-empty value to expect.                      |

<a name="google_api_field_behavior-proto-extensions"></a>

### File-level Extensions

| Extension      | Type          | Base                          | Number | Description                                                                                    |
| -------------- | ------------- | ----------------------------- | ------ | ---------------------------------------------------------------------------------------------- |
| field_behavior | FieldBehavior | .google.protobuf.FieldOptions | 1052   | A designation of a specific field behavior (required, output only, etc.) in protobuf messages. |

Examples:

string name = 1 [(google.api.field_behavior) = REQUIRED]; State state = 1
[(google.api.field_behavior) = OUTPUT_ONLY]; google.protobuf.Duration ttl = 1
[(google.api.field_behavior) = INPUT_ONLY]; google.protobuf.Timestamp
expire_time = 1 [(google.api.field_behavior) = OUTPUT_ONLY,
(google.api.field_behavior) = IMMUTABLE]; |

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

| Field       | Type                                                        | Label | Description                                                                                                                             |
| ----------- | ----------------------------------------------------------- | ----- | --------------------------------------------------------------------------------------------------------------------------------------- |
| namespace   | [string](#string)                                           |       | The namespace this relation tuple lives in.                                                                                             |
| object      | [string](#string)                                           |       | The object related by this tuple. It is an object in the namespace of the tuple.                                                        |
| relation    | [string](#string)                                           |       | The relation between an Object and a Subject.                                                                                           |
| subject     | [Subject](#ory-keto-relation_tuples-v1alpha2-Subject)       |       | The subject related by this tuple. A Subject either represents a concrete subject id or a `SubjectSet` that expands to more Subjects.   |
| subject_id  | [string](#string)                                           |       | **Deprecated.** A concrete id of the subject.                                                                                           |
| subject_set | [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet) |       | **Deprecated.** A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-relation_tuples-v1alpha2-Subject"></a>

### Subject

Subject is either a concrete subject id or a `SubjectSet` expanding to more
Subjects.

| Field | Type                                                        | Label | Description                                                                                                             |
| ----- | ----------------------------------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- |
| id    | [string](#string)                                           |       | A concrete id of the subject.                                                                                           |
| set   | [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet) |       | A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-relation_tuples-v1alpha2-SubjectQuery"></a>

### SubjectQuery

SubjectQuery is either a concrete subject id or a `SubjectSet` expanding to more
Subjects.

| Field | Type                                                                  | Label | Description                                                                                                             |
| ----- | --------------------------------------------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- |
| id    | [string](#string)                                                     |       | A concrete id of the subject.                                                                                           |
| set   | [SubjectSetQuery](#ory-keto-relation_tuples-v1alpha2-SubjectSetQuery) |       | A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-relation_tuples-v1alpha2-SubjectSet"></a>

### SubjectSet

SubjectSet refers to all subjects who have the same `relation` on an `object`.

| Field     | Type              | Label | Description                                                              |
| --------- | ----------------- | ----- | ------------------------------------------------------------------------ |
| namespace | [string](#string) |       | The namespace of the object and relation referenced in this subject set. |
| object    | [string](#string) |       | The object related by this subject set.                                  |
| relation  | [string](#string) |       | The relation between the object and the subjects.                        |

<a name="ory-keto-relation_tuples-v1alpha2-SubjectSetQuery"></a>

### SubjectSetQuery

SubjectSetQuery refers to all subjects who have the same `relation` on an
`object`.

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
related subject in this check. | | subject_id | [string](#string) | |
**Deprecated.** A concrete id of the subject. | | subject_set |
[SubjectSetQuery](#ory-keto-relation_tuples-v1alpha2-SubjectSetQuery) | |
**Deprecated.** A subject set that expands to more Subjects. More information
are available under [concepts](../concepts/subjects.mdx). | | tuple |
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
availability zones. --&gt; | | namespace | [string](#string) | | **Deprecated.**
The namespace of the object and relation referenced in this subject set. | |
object | [string](#string) | | **Deprecated.** The object related by this
subject set. | | relation | [string](#string) | | **Deprecated.** The relation
between the object and the subjects. |

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

| Name                           | Number | Description                                                                                                |
| ------------------------------ | ------ | ---------------------------------------------------------------------------------------------------------- |
| unspecified                    | 0      |                                                                                                            |
| NODE_TYPE_UNSPECIFIED          | 0      |                                                                                                            |
| union                          | 1      | This node expands to a union of all children.                                                              |
| NODE_TYPE_UNION                | 1      |                                                                                                            |
| exclusion                      | 2      | Not implemented yet.                                                                                       |
| NODE_TYPE_EXCLUSION            | 2      |                                                                                                            |
| intersection                   | 3      | Not implemented yet.                                                                                       |
| NODE_TYPE_INTERSECTION         | 3      |                                                                                                            |
| leaf                           | 4      | This node is a leaf and contains no children. Its subject is a `SubjectID` unless `max_depth` was reached. |
| NODE_TYPE_LEAF                 | 4      |                                                                                                            |
| tuple_to_subject_set           | 5      | This node is a leaf and contains no children. Its subject is a `SubjectID` unless `max_depth` was reached. |
| NODE_TYPE_TUPLE_TO_SUBJECT_SET | 5      |                                                                                                            |
| computed_subject_set           | 6      | This node is a leaf and contains no children. Its subject is a `SubjectID` unless `max_depth` was reached. |
| NODE_TYPE_COMPUTED_SUBJECT_SET | 6      |                                                                                                            |
| not                            | 7      | This node is a leaf and contains no children. Its subject is a `SubjectID` unless `max_depth` was reached. |
| NODE_TYPE_NOT                  | 7      |                                                                                                            |

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

Get all namespaces. |

<a name="ory_keto_relation_tuples_v1alpha2_openapi-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/openapi.proto

<a name="ory-keto-relation_tuples-v1alpha2-ErrorResponse"></a>

### ErrorResponse

JSON API Error Response

The standard Ory JSON API error format.

| Field | Type                                                                          | Label | Description |
| ----- | ----------------------------------------------------------------------------- | ----- | ----------- |
| error | [ErrorResponse.Error](#ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error) |       |             |

<a name="ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error"></a>

### ErrorResponse.Error

| Field | Type              | Label | Description       |
| ----- | ----------------- | ----- | ----------------- |
| code  | [int64](#int64)   |       | The status code   |
| debug | [string](#string) |       | Debug information |

Debug information is often not exposed to protect against leaking sensitive
information. | | details |
[ErrorResponse.Error.DetailsEntry](#ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error-DetailsEntry)
| repeated | Further error details

Further details about the error. | | id | [string](#string) | | The error ID

The error ID is useful when trying to identify various errors in application
logic. | | message | [string](#string) | | The error message

The error&#39;s message (required). | | reason | [string](#string) | | The error
reason

Reason holds a human-readable reason for the error. | | request |
[string](#string) | | The request ID

The request ID is often exposed internally in order to trace errors across
service architectures. This is often a UUID. | | status | [string](#string) | |
The status description

Status holds the human-readable HTTP status code. |

<a name="ory-keto-relation_tuples-v1alpha2-ErrorResponse-Error-DetailsEntry"></a>

### ErrorResponse.Error.DetailsEntry

| Field | Type              | Label | Description |
| ----- | ----------------- | ----- | ----------- |
| key   | [string](#string) |       |             |
| value | [string](#string) |       |             |

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
from the previous page. | | namespace | [string](#string) | | **Deprecated.**
The namespace | | object | [string](#string) | | **Deprecated.** The related
object in this check. | | relation | [string](#string) | | **Deprecated.** The
relation between the Object and the Subject. | | subject_id | [string](#string)
| | A concrete id of the subject. | | subject_set |
[SubjectSetQuery](#ory-keto-relation_tuples-v1alpha2-SubjectSetQuery) | | A
subject set that expands to more Subjects. More information are available under
[concepts](../concepts/subjects.mdx). |

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

This endpoint returns the service version typically notated using semantic
versioning.

If the service supports TLS Edge Termination, this endpoint does not require the
X-Forwarded-Proto header to be set. |

<a name="validate_validate-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## validate/validate.proto

<a name="validate-AnyRules"></a>

### AnyRules

AnyRules describe constraints applied exclusively to the `google.protobuf.Any`
well-known type

| Field    | Type              | Label    | Description                                                                                        |
| -------- | ----------------- | -------- | -------------------------------------------------------------------------------------------------- |
| required | [bool](#bool)     | optional | Required specifies that this field must be set                                                     |
| in       | [string](#string) | repeated | In specifies that this field&#39;s `type_url` must be equal to one of the specified values.        |
| not_in   | [string](#string) | repeated | NotIn specifies that this field&#39;s `type_url` must not be equal to any of the specified values. |

<a name="validate-BoolRules"></a>

### BoolRules

BoolRules describes the constraints applied to `bool` values

| Field | Type          | Label    | Description                                                         |
| ----- | ------------- | -------- | ------------------------------------------------------------------- |
| const | [bool](#bool) | optional | Const specifies that this field must be exactly the specified value |

<a name="validate-BytesRules"></a>

### BytesRules

BytesRules describe the constraints applied to `bytes` values

| Field        | Type              | Label    | Description                                                                                                                                             |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [bytes](#bytes)   | optional | Const specifies that this field must be exactly the specified value                                                                                     |
| len          | [uint64](#uint64) | optional | Len specifies that this field must be the specified number of bytes                                                                                     |
| min_len      | [uint64](#uint64) | optional | MinLen specifies that this field must be the specified number of bytes at a minimum                                                                     |
| max_len      | [uint64](#uint64) | optional | MaxLen specifies that this field must be the specified number of bytes at a maximum                                                                     |
| pattern      | [string](#string) | optional | Pattern specifes that this field must match against the specified regular expression (RE2 syntax). The included expression should elide any delimiters. |
| prefix       | [bytes](#bytes)   | optional | Prefix specifies that this field must have the specified bytes at the beginning of the string.                                                          |
| suffix       | [bytes](#bytes)   | optional | Suffix specifies that this field must have the specified bytes at the end of the string.                                                                |
| contains     | [bytes](#bytes)   | optional | Contains specifies that this field must have the specified bytes anywhere in the string.                                                                |
| in           | [bytes](#bytes)   | repeated | In specifies that this field must be equal to one of the specified values                                                                               |
| not_in       | [bytes](#bytes)   | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                          |
| ip           | [bool](#bool)     | optional | Ip specifies that the field must be a valid IP (v4 or v6) address in byte format                                                                        |
| ipv4         | [bool](#bool)     | optional | Ipv4 specifies that the field must be a valid IPv4 address in byte format                                                                               |
| ipv6         | [bool](#bool)     | optional | Ipv6 specifies that the field must be a valid IPv6 address in byte format                                                                               |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                        |

<a name="validate-DoubleRules"></a>

### DoubleRules

DoubleRules describes the constraints applied to `double` values

| Field        | Type              | Label    | Description                                                                                                                                                                     |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [double](#double) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [double](#double) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [double](#double) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [double](#double) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [double](#double) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [double](#double) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [double](#double) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-DurationRules"></a>

### DurationRules

DurationRules describe the constraints applied exclusively to the
`google.protobuf.Duration` well-known type

| Field    | Type                                                  | Label    | Description                                                                       |
| -------- | ----------------------------------------------------- | -------- | --------------------------------------------------------------------------------- |
| required | [bool](#bool)                                         | optional | Required specifies that this field must be set                                    |
| const    | [google.protobuf.Duration](#google-protobuf-Duration) | optional | Const specifies that this field must be exactly the specified value               |
| lt       | [google.protobuf.Duration](#google-protobuf-Duration) | optional | Lt specifies that this field must be less than the specified value, exclusive     |
| lte      | [google.protobuf.Duration](#google-protobuf-Duration) | optional | Lt specifies that this field must be less than the specified value, inclusive     |
| gt       | [google.protobuf.Duration](#google-protobuf-Duration) | optional | Gt specifies that this field must be greater than the specified value, exclusive  |
| gte      | [google.protobuf.Duration](#google-protobuf-Duration) | optional | Gte specifies that this field must be greater than the specified value, inclusive |
| in       | [google.protobuf.Duration](#google-protobuf-Duration) | repeated | In specifies that this field must be equal to one of the specified values         |
| not_in   | [google.protobuf.Duration](#google-protobuf-Duration) | repeated | NotIn specifies that this field cannot be equal to one of the specified values    |

<a name="validate-EnumRules"></a>

### EnumRules

EnumRules describe the constraints applied to enum values

| Field        | Type            | Label    | Description                                                                                                                 |
| ------------ | --------------- | -------- | --------------------------------------------------------------------------------------------------------------------------- |
| const        | [int32](#int32) | optional | Const specifies that this field must be exactly the specified value                                                         |
| defined_only | [bool](#bool)   | optional | DefinedOnly specifies that this field must be only one of the defined values for this enum, failing on any undefined value. |
| in           | [int32](#int32) | repeated | In specifies that this field must be equal to one of the specified values                                                   |
| not_in       | [int32](#int32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                              |

<a name="validate-FieldRules"></a>

### FieldRules

FieldRules encapsulates the rules for each type of field. Depending on the
field, the correct set should be used to ensure proper validations.

| Field     | Type                                       | Label    | Description            |
| --------- | ------------------------------------------ | -------- | ---------------------- |
| message   | [MessageRules](#validate-MessageRules)     | optional |                        |
| float     | [FloatRules](#validate-FloatRules)         | optional | Scalar Field Types     |
| double    | [DoubleRules](#validate-DoubleRules)       | optional |                        |
| int32     | [Int32Rules](#validate-Int32Rules)         | optional |                        |
| int64     | [Int64Rules](#validate-Int64Rules)         | optional |                        |
| uint32    | [UInt32Rules](#validate-UInt32Rules)       | optional |                        |
| uint64    | [UInt64Rules](#validate-UInt64Rules)       | optional |                        |
| sint32    | [SInt32Rules](#validate-SInt32Rules)       | optional |                        |
| sint64    | [SInt64Rules](#validate-SInt64Rules)       | optional |                        |
| fixed32   | [Fixed32Rules](#validate-Fixed32Rules)     | optional |                        |
| fixed64   | [Fixed64Rules](#validate-Fixed64Rules)     | optional |                        |
| sfixed32  | [SFixed32Rules](#validate-SFixed32Rules)   | optional |                        |
| sfixed64  | [SFixed64Rules](#validate-SFixed64Rules)   | optional |                        |
| bool      | [BoolRules](#validate-BoolRules)           | optional |                        |
| string    | [StringRules](#validate-StringRules)       | optional |                        |
| bytes     | [BytesRules](#validate-BytesRules)         | optional |                        |
| enum      | [EnumRules](#validate-EnumRules)           | optional | Complex Field Types    |
| repeated  | [RepeatedRules](#validate-RepeatedRules)   | optional |                        |
| map       | [MapRules](#validate-MapRules)             | optional |                        |
| any       | [AnyRules](#validate-AnyRules)             | optional | Well-Known Field Types |
| duration  | [DurationRules](#validate-DurationRules)   | optional |                        |
| timestamp | [TimestampRules](#validate-TimestampRules) | optional |                        |

<a name="validate-Fixed32Rules"></a>

### Fixed32Rules

Fixed32Rules describes the constraints applied to `fixed32` values

| Field        | Type                | Label    | Description                                                                                                                                                                     |
| ------------ | ------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [fixed32](#fixed32) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [fixed32](#fixed32) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [fixed32](#fixed32) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [fixed32](#fixed32) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [fixed32](#fixed32) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [fixed32](#fixed32) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [fixed32](#fixed32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)       | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-Fixed64Rules"></a>

### Fixed64Rules

Fixed64Rules describes the constraints applied to `fixed64` values

| Field        | Type                | Label    | Description                                                                                                                                                                     |
| ------------ | ------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [fixed64](#fixed64) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [fixed64](#fixed64) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [fixed64](#fixed64) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [fixed64](#fixed64) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [fixed64](#fixed64) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [fixed64](#fixed64) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [fixed64](#fixed64) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)       | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-FloatRules"></a>

### FloatRules

FloatRules describes the constraints applied to `float` values

| Field        | Type            | Label    | Description                                                                                                                                                                     |
| ------------ | --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [float](#float) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [float](#float) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [float](#float) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [float](#float) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [float](#float) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [float](#float) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [float](#float) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)   | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-Int32Rules"></a>

### Int32Rules

Int32Rules describes the constraints applied to `int32` values

| Field        | Type            | Label    | Description                                                                                                                                                                     |
| ------------ | --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [int32](#int32) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [int32](#int32) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [int32](#int32) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [int32](#int32) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [int32](#int32) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [int32](#int32) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [int32](#int32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)   | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-Int64Rules"></a>

### Int64Rules

Int64Rules describes the constraints applied to `int64` values

| Field        | Type            | Label    | Description                                                                                                                                                                     |
| ------------ | --------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [int64](#int64) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [int64](#int64) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [int64](#int64) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [int64](#int64) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [int64](#int64) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [int64](#int64) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [int64](#int64) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)   | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-MapRules"></a>

### MapRules

MapRules describe the constraints applied to `map` values

| Field        | Type                               | Label    | Description                                                                                                                                                                     |
| ------------ | ---------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| min_pairs    | [uint64](#uint64)                  | optional | MinPairs specifies that this field must have the specified number of KVs at a minimum                                                                                           |
| max_pairs    | [uint64](#uint64)                  | optional | MaxPairs specifies that this field must have the specified number of KVs at a maximum                                                                                           |
| no_sparse    | [bool](#bool)                      | optional | NoSparse specifies values in this field cannot be unset. This only applies to map&#39;s with message value types.                                                               |
| keys         | [FieldRules](#validate-FieldRules) | optional | Keys specifies the constraints to be applied to each key in the field.                                                                                                          |
| values       | [FieldRules](#validate-FieldRules) | optional | Values specifies the constraints to be applied to the value of each key in the field. Message values will still have their validations evaluated unless skip is specified here. |
| ignore_empty | [bool](#bool)                      | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-MessageRules"></a>

### MessageRules

MessageRules describe the constraints applied to embedded message values. For
message-type fields, validation is performed recursively.

| Field    | Type          | Label    | Description                                                                    |
| -------- | ------------- | -------- | ------------------------------------------------------------------------------ |
| skip     | [bool](#bool) | optional | Skip specifies that the validation rules of this field should not be evaluated |
| required | [bool](#bool) | optional | Required specifies that this field must be set                                 |

<a name="validate-RepeatedRules"></a>

### RepeatedRules

RepeatedRules describe the constraints applied to `repeated` values

| Field        | Type                               | Label    | Description                                                                                                                                                                    |
| ------------ | ---------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| min_items    | [uint64](#uint64)                  | optional | MinItems specifies that this field must have the specified number of items at a minimum                                                                                        |
| max_items    | [uint64](#uint64)                  | optional | MaxItems specifies that this field must have the specified number of items at a maximum                                                                                        |
| unique       | [bool](#bool)                      | optional | Unique specifies that all elements in this field must be unique. This contraint is only applicable to scalar and enum types (messages are not supported).                      |
| items        | [FieldRules](#validate-FieldRules) | optional | Items specifies the contraints to be applied to each item in the field. Repeated message fields will still execute validation against each item unless skip is specified here. |
| ignore_empty | [bool](#bool)                      | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                               |

<a name="validate-SFixed32Rules"></a>

### SFixed32Rules

SFixed32Rules describes the constraints applied to `sfixed32` values

| Field        | Type                  | Label    | Description                                                                                                                                                                     |
| ------------ | --------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [sfixed32](#sfixed32) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [sfixed32](#sfixed32) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [sfixed32](#sfixed32) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [sfixed32](#sfixed32) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [sfixed32](#sfixed32) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [sfixed32](#sfixed32) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [sfixed32](#sfixed32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)         | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-SFixed64Rules"></a>

### SFixed64Rules

SFixed64Rules describes the constraints applied to `sfixed64` values

| Field        | Type                  | Label    | Description                                                                                                                                                                     |
| ------------ | --------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [sfixed64](#sfixed64) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [sfixed64](#sfixed64) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [sfixed64](#sfixed64) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [sfixed64](#sfixed64) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [sfixed64](#sfixed64) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [sfixed64](#sfixed64) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [sfixed64](#sfixed64) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)         | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-SInt32Rules"></a>

### SInt32Rules

SInt32Rules describes the constraints applied to `sint32` values

| Field        | Type              | Label    | Description                                                                                                                                                                     |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [sint32](#sint32) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [sint32](#sint32) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [sint32](#sint32) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [sint32](#sint32) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [sint32](#sint32) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [sint32](#sint32) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [sint32](#sint32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-SInt64Rules"></a>

### SInt64Rules

SInt64Rules describes the constraints applied to `sint64` values

| Field        | Type              | Label    | Description                                                                                                                                                                     |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [sint64](#sint64) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [sint64](#sint64) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [sint64](#sint64) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [sint64](#sint64) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [sint64](#sint64) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [sint64](#sint64) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [sint64](#sint64) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-StringRules"></a>

### StringRules

StringRules describe the constraints applied to `string` values

| Field            | Type                               | Label    | Description                                                                                                                                                                                                                                                                                                                           |
| ---------------- | ---------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const            | [string](#string)                  | optional | Const specifies that this field must be exactly the specified value                                                                                                                                                                                                                                                                   |
| len              | [uint64](#uint64)                  | optional | Len specifies that this field must be the specified number of characters (Unicode code points). Note that the number of characters may differ from the number of bytes in the string.                                                                                                                                                 |
| min_len          | [uint64](#uint64)                  | optional | MinLen specifies that this field must be the specified number of characters (Unicode code points) at a minimum. Note that the number of characters may differ from the number of bytes in the string.                                                                                                                                 |
| max_len          | [uint64](#uint64)                  | optional | MaxLen specifies that this field must be the specified number of characters (Unicode code points) at a maximum. Note that the number of characters may differ from the number of bytes in the string.                                                                                                                                 |
| len_bytes        | [uint64](#uint64)                  | optional | LenBytes specifies that this field must be the specified number of bytes                                                                                                                                                                                                                                                              |
| min_bytes        | [uint64](#uint64)                  | optional | MinBytes specifies that this field must be the specified number of bytes at a minimum                                                                                                                                                                                                                                                 |
| max_bytes        | [uint64](#uint64)                  | optional | MaxBytes specifies that this field must be the specified number of bytes at a maximum                                                                                                                                                                                                                                                 |
| pattern          | [string](#string)                  | optional | Pattern specifes that this field must match against the specified regular expression (RE2 syntax). The included expression should elide any delimiters.                                                                                                                                                                               |
| prefix           | [string](#string)                  | optional | Prefix specifies that this field must have the specified substring at the beginning of the string.                                                                                                                                                                                                                                    |
| suffix           | [string](#string)                  | optional | Suffix specifies that this field must have the specified substring at the end of the string.                                                                                                                                                                                                                                          |
| contains         | [string](#string)                  | optional | Contains specifies that this field must have the specified substring anywhere in the string.                                                                                                                                                                                                                                          |
| not_contains     | [string](#string)                  | optional | NotContains specifies that this field cannot have the specified substring anywhere in the string.                                                                                                                                                                                                                                     |
| in               | [string](#string)                  | repeated | In specifies that this field must be equal to one of the specified values                                                                                                                                                                                                                                                             |
| not_in           | [string](#string)                  | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                                                                                                                                                                        |
| email            | [bool](#bool)                      | optional | Email specifies that the field must be a valid email address as defined by RFC 5322                                                                                                                                                                                                                                                   |
| hostname         | [bool](#bool)                      | optional | Hostname specifies that the field must be a valid hostname as defined by RFC 1034. This constraint does not support internationalized domain names (IDNs).                                                                                                                                                                            |
| ip               | [bool](#bool)                      | optional | Ip specifies that the field must be a valid IP (v4 or v6) address. Valid IPv6 addresses should not include surrounding square brackets.                                                                                                                                                                                               |
| ipv4             | [bool](#bool)                      | optional | Ipv4 specifies that the field must be a valid IPv4 address.                                                                                                                                                                                                                                                                           |
| ipv6             | [bool](#bool)                      | optional | Ipv6 specifies that the field must be a valid IPv6 address. Valid IPv6 addresses should not include surrounding square brackets.                                                                                                                                                                                                      |
| uri              | [bool](#bool)                      | optional | Uri specifies that the field must be a valid, absolute URI as defined by RFC 3986                                                                                                                                                                                                                                                     |
| uri_ref          | [bool](#bool)                      | optional | UriRef specifies that the field must be a valid URI as defined by RFC 3986 and may be relative or absolute.                                                                                                                                                                                                                           |
| address          | [bool](#bool)                      | optional | Address specifies that the field must be either a valid hostname as defined by RFC 1034 (which does not support internationalized domain names or IDNs), or it can be a valid IP (v4 or v6).                                                                                                                                          |
| uuid             | [bool](#bool)                      | optional | Uuid specifies that the field must be a valid UUID as defined by RFC 4122                                                                                                                                                                                                                                                             |
| well_known_regex | [KnownRegex](#validate-KnownRegex) | optional | WellKnownRegex specifies a common well known pattern defined as a regex.                                                                                                                                                                                                                                                              |
| strict           | [bool](#bool)                      | optional | This applies to regexes HTTP_HEADER_NAME and HTTP_HEADER_VALUE to enable strict header validation. By default, this is true, and HTTP header validations are RFC-compliant. Setting to false will enable a looser validations that only disallows \r\n\0 characters, which can be used to bypass header matching rules. Default: true |
| ignore_empty     | [bool](#bool)                      | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                                                                                                                                                                      |

<a name="validate-TimestampRules"></a>

### TimestampRules

TimestampRules describe the constraints applied exclusively to the
`google.protobuf.Timestamp` well-known type

| Field    | Type                                                    | Label    | Description                                                                                                                                             |
| -------- | ------------------------------------------------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------- |
| required | [bool](#bool)                                           | optional | Required specifies that this field must be set                                                                                                          |
| const    | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | Const specifies that this field must be exactly the specified value                                                                                     |
| lt       | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                           |
| lte      | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | Lte specifies that this field must be less than the specified value, inclusive                                                                          |
| gt       | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | Gt specifies that this field must be greater than the specified value, exclusive                                                                        |
| gte      | [google.protobuf.Timestamp](#google-protobuf-Timestamp) | optional | Gte specifies that this field must be greater than the specified value, inclusive                                                                       |
| lt_now   | [bool](#bool)                                           | optional | LtNow specifies that this must be less than the current time. LtNow can only be used with the Within rule.                                              |
| gt_now   | [bool](#bool)                                           | optional | GtNow specifies that this must be greater than the current time. GtNow can only be used with the Within rule.                                           |
| within   | [google.protobuf.Duration](#google-protobuf-Duration)   | optional | Within specifies that this field must be within this duration of the current time. This constraint can be used alone or with the LtNow and GtNow rules. |

<a name="validate-UInt32Rules"></a>

### UInt32Rules

UInt32Rules describes the constraints applied to `uint32` values

| Field        | Type              | Label    | Description                                                                                                                                                                     |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [uint32](#uint32) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [uint32](#uint32) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [uint32](#uint32) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [uint32](#uint32) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [uint32](#uint32) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [uint32](#uint32) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [uint32](#uint32) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-UInt64Rules"></a>

### UInt64Rules

UInt64Rules describes the constraints applied to `uint64` values

| Field        | Type              | Label    | Description                                                                                                                                                                     |
| ------------ | ----------------- | -------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| const        | [uint64](#uint64) | optional | Const specifies that this field must be exactly the specified value                                                                                                             |
| lt           | [uint64](#uint64) | optional | Lt specifies that this field must be less than the specified value, exclusive                                                                                                   |
| lte          | [uint64](#uint64) | optional | Lte specifies that this field must be less than or equal to the specified value, inclusive                                                                                      |
| gt           | [uint64](#uint64) | optional | Gt specifies that this field must be greater than the specified value, exclusive. If the value of Gt is larger than a specified Lt or Lte, the range is reversed.               |
| gte          | [uint64](#uint64) | optional | Gte specifies that this field must be greater than or equal to the specified value, inclusive. If the value of Gte is larger than a specified Lt or Lte, the range is reversed. |
| in           | [uint64](#uint64) | repeated | In specifies that this field must be equal to one of the specified values                                                                                                       |
| not_in       | [uint64](#uint64) | repeated | NotIn specifies that this field cannot be equal to one of the specified values                                                                                                  |
| ignore_empty | [bool](#bool)     | optional | IgnoreEmpty specifies that the validation rules of this field should be evaluated only if the field is not empty                                                                |

<a name="validate-KnownRegex"></a>

### KnownRegex

WellKnownRegex contain some well-known patterns.

| Name              | Number | Description                               |
| ----------------- | ------ | ----------------------------------------- |
| UNKNOWN           | 0      |                                           |
| HTTP_HEADER_NAME  | 1      | HTTP header name as defined by RFC 7230.  |
| HTTP_HEADER_VALUE | 2      | HTTP header value as defined by RFC 7230. |

<a name="validate_validate-proto-extensions"></a>

### File-level Extensions

| Extension | Type       | Base                            | Number | Description                                                                                                                           |
| --------- | ---------- | ------------------------------- | ------ | ------------------------------------------------------------------------------------------------------------------------------------- |
| rules     | FieldRules | .google.protobuf.FieldOptions   | 1071   | Rules specify the validations to be performed on this field. By default, no validation is performed against a field.                  |
| disabled  | bool       | .google.protobuf.MessageOptions | 1071   | Disabled nullifies any validation rules for this message, including any message fields associated with it that do support validation. |
| ignored   | bool       | .google.protobuf.MessageOptions | 1072   | Ignore skips generation of validation methods for this message.                                                                       |
| required  | bool       | .google.protobuf.OneofOptions   | 1071   | Required ensures that exactly one the field options in a oneof is set; validation fails if no fields in the oneof are set.            |

<a name="ory_keto_relation_tuples_v1alpha2_write_service-proto"></a>

<p align="right"><a href="#top">Top</a></p>

## ory/keto/relation_tuples/v1alpha2/write_service.proto

<a name="ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest"></a>

### CreateRelationTupleRequest

The request to create a new relationship.

| Field          | Type                                                                                                                  | Label | Description                 |
| -------------- | --------------------------------------------------------------------------------------------------------------------- | ----- | --------------------------- |
| relation_tuple | [CreateRelationTupleRequest.Relationship](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest-Relationship) |       | The relationship to create. |

<a name="ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest-Relationship"></a>

### CreateRelationTupleRequest.Relationship

| Field       | Type                                                        | Label | Description                                                                                                             |
| ----------- | ----------------------------------------------------------- | ----- | ----------------------------------------------------------------------------------------------------------------------- |
| namespace   | [string](#string)                                           |       | The namespace this relation tuple lives in.                                                                             |
| object      | [string](#string)                                           |       | The object related by this tuple. It is an object in the namespace of the tuple.                                        |
| relation    | [string](#string)                                           |       | The relation between an Object and a Subject.                                                                           |
| subject_id  | [string](#string)                                           |       | A concrete id of the subject.                                                                                           |
| subject_set | [SubjectSet](#ory-keto-relation_tuples-v1alpha2-SubjectSet) |       | A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

<a name="ory-keto-relation_tuples-v1alpha2-CreateRelationTupleResponse"></a>

### CreateRelationTupleResponse

The response from creating a new relationship.

| Field          | Type                                                              | Label | Description               |
| -------------- | ----------------------------------------------------------------- | ----- | ------------------------- |
| relation_tuple | [RelationTuple](#ory-keto-relation_tuples-v1alpha2-RelationTuple) |       | The created relationship. |

<a name="ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest"></a>

### DeleteRelationTuplesRequest

| Field          | Type                                                                                                      | Label | Description                                                                                                                             |
| -------------- | --------------------------------------------------------------------------------------------------------- | ----- | --------------------------------------------------------------------------------------------------------------------------------------- |
| query          | [DeleteRelationTuplesRequest.Query](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest-Query) |       | **Deprecated.**                                                                                                                         |
| relation_query | [RelationQuery](#ory-keto-relation_tuples-v1alpha2-RelationQuery)                                         |       |                                                                                                                                         |
| namespace      | [string](#string)                                                                                         |       | **Deprecated.** The namespace this relation tuple lives in.                                                                             |
| object         | [string](#string)                                                                                         |       | **Deprecated.** The object related by this tuple. It is an object in the namespace of the tuple.                                        |
| relation       | [string](#string)                                                                                         |       | **Deprecated.** The relation between an Object and a Subject.                                                                           |
| subject_id     | [string](#string)                                                                                         |       | **Deprecated.** A concrete id of the subject.                                                                                           |
| subject_set    | [SubjectSetQuery](#ory-keto-relation_tuples-v1alpha2-SubjectSetQuery)                                     |       | **Deprecated.** A subject set that expands to more Subjects. More information are available under [concepts](../concepts/subjects.mdx). |

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
| insert             | 1      | Insertion of a new RelationTuple. It is ignored if already existing.                                        |
| ACTION_DELETE      | 2      | Deletion of the RelationTuple. It is ignored if it does not exist.                                          |
| delete             | 2      | Deletion of the RelationTuple. It is ignored if it does not exist.                                          |

<a name="ory-keto-relation_tuples-v1alpha2-WriteService"></a>

### WriteService

The write service to create and delete Access Control Lists.

This service is part of the
[write-APIs](../concepts/api-overview.mdx#write-apis).

| Method Name            | Request Type                                                                                      | Response Type                                                                                       | Description                                               |
| ---------------------- | ------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------- | --------------------------------------------------------- |
| TransactRelationTuples | [TransactRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesRequest) | [TransactRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-TransactRelationTuplesResponse) | Writes one or more relationships in a single transaction. |
| CreateRelationTuple    | [CreateRelationTupleRequest](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleRequest)       | [CreateRelationTupleResponse](#ory-keto-relation_tuples-v1alpha2-CreateRelationTupleResponse)       | Creates a relationship                                    |
| DeleteRelationTuples   | [DeleteRelationTuplesRequest](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesRequest)     | [DeleteRelationTuplesResponse](#ory-keto-relation_tuples-v1alpha2-DeleteRelationTuplesResponse)     | Deletes relationships based on relation query             |

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
