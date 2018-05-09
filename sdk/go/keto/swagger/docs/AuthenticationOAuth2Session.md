# AuthenticationOAuth2Session

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Allowed** | **bool** | Allowed is true if the request is allowed and false otherwise. | [optional] [default to null]
**Aud** | **[]string** |  | [optional] [default to null]
**ClientId** | **string** | ClientID is the id of the OAuth2 client that requested the token. | [optional] [default to null]
**Exp** | [**time.Time**](time.Time.md) | ExpiresAt is the expiry timestamp. | [optional] [default to null]
**Iat** | [**time.Time**](time.Time.md) | IssuedAt is the token creation time stamp. | [optional] [default to null]
**Iss** | **string** | Issuer is the id of the issuer, typically an hydra instance. | [optional] [default to null]
**Nbf** | [**time.Time**](time.Time.md) |  | [optional] [default to null]
**Scope** | **string** | GrantedScopes is a list of scopes that the subject authorized when asked for consent. | [optional] [default to null]
**Session** | [**map[string]interface{}**](interface{}.md) | Session represents arbitrary session data. | [optional] [default to null]
**Sub** | **string** | Subject is the identity that authorized issuing the token, for example a user or an OAuth2 app. This is usually a uuid but you can choose a urn or some other id too. | [optional] [default to null]
**Username** | **string** |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


