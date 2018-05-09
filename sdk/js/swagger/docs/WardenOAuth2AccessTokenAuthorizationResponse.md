# SwaggerJsClient.WardenOAuth2AccessTokenAuthorizationResponse

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**allowed** | **Boolean** | Allowed is true if the request is allowed and false otherwise. | [optional] 
**aud** | **[String]** |  | [optional] 
**clientId** | **String** | ClientID is the id of the OAuth2 client that requested the token. | [optional] 
**exp** | **Date** | ExpiresAt is the expiry timestamp. | [optional] 
**iat** | **Date** | IssuedAt is the token creation time stamp. | [optional] 
**iss** | **String** | Issuer is the id of the issuer, typically an hydra instance. | [optional] 
**nbf** | **Date** |  | [optional] 
**scope** | **String** | GrantedScopes is a list of scopes that the subject authorized when asked for consent. | [optional] 
**session** | **{String: Object}** | Session represents arbitrary session data. | [optional] 
**sub** | **String** | Subject is the identity that authorized issuing the token, for example a user or an OAuth2 app. This is usually a uuid but you can choose a urn or some other id too. | [optional] 
**username** | **String** |  | [optional] 


