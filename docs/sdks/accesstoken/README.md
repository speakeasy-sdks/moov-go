# access_token

## Overview

When making requests to Moov from a browser, you can use OAuth with JSON Web Tokens (JWT).

Our authentication flow follows the OAuth 2.0 standard. With this endpoint, you'll [create an access token](https://docs.moov.io/guides/quick-start/#create-an-access-token) that you will pass along with API requests or when initializing Moov.js. To generate an authentication token, you’ll need to specify scopes that enable the token to use a specific set of API endpoints. To learn more, read our [scopes guide](https://docs.moov.io/guides/developer-tools/scopes/).


### Available Operations

* [create](#create) - Create access token
* [revoke](#revoke) - Revoke access token

## create

Use the client_id and client_secret to generate an access token.

### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.ClientCredentialsGrantToAccessTokenRequest(
    client_id='5clTR_MdVrrkgxw2',
    client_secret='dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-',
    grant_type=shared.ClientCredentialsGrantToAccessTokenRequestGrantType.REFRESH_TOKEN,
    refresh_token='i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...',
    scope='/accounts.write',
)

res = s.access_token.create(req)

if res.access_token_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                              | Type                                                                                                                   | Required                                                                                                               | Description                                                                                                            |
| ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                              | [shared.ClientCredentialsGrantToAccessTokenRequest](../../models/shared/clientcredentialsgranttoaccesstokenrequest.md) | :heavy_check_mark:                                                                                                     | The request object to use for the request.                                                                             |


### Response

**[operations.PostOAuth2TokenResponse](../../models/operations/postoauth2tokenresponse.md)**


## revoke

Allows clients to notify the authorization server that a previously obtained refresh or access token is no longer needed

### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.RevokeTokenRequest1(
    client_id='5clTR_MdVrrkgxw2',
    client_secret='dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-',
    token='i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...',
    token_type_hint=shared.RevokeTokenRequestTokenTypeHint.REFRESH_TOKEN,
)

res = s.access_token.revoke(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `request`                                                                | [shared.RevokeTokenRequest1](../../models/shared/revoketokenrequest1.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[operations.RevokeOAuth2TokenResponse](../../models/operations/revokeoauth2tokenresponse.md)**

