# AuthenticationOAuth2Session

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**allowed** | **bool** | Allowed is true if the request is allowed and false otherwise. | [optional] 
**aud** | **string[]** |  | [optional] 
**client_id** | **string** | ClientID is the id of the OAuth2 client that requested the token. | [optional] 
**exp** | [**\DateTime**](\DateTime.md) | ExpiresAt is the expiry timestamp. | [optional] 
**iat** | [**\DateTime**](\DateTime.md) | IssuedAt is the token creation time stamp. | [optional] 
**iss** | **string** | Issuer is the id of the issuer, typically an hydra instance. | [optional] 
**nbf** | [**\DateTime**](\DateTime.md) |  | [optional] 
**scope** | **string** | GrantedScopes is a list of scopes that the subject authorized when asked for consent. | [optional] 
**session** | **map[string,object]** | Session represents arbitrary session data. | [optional] 
**sub** | **string** | Subject is the identity that authorized issuing the token, for example a user or an OAuth2 app. This is usually a uuid but you can choose a urn or some other id too. | [optional] 
**username** | **string** |  | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


