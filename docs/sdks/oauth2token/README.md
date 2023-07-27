# o_auth2_token

### Available Operations

* [post](#post) - Create access token
* [revoke](#revoke) - Revoke access token

## post

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
    grant_type=shared.ClientCredentialsGrantToAccessTokenRequestGrantType.CLIENT_CREDENTIALS,
    refresh_token='i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...',
    scope='/accounts.write',
)

res = s.o_auth2_token.post(req)

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
    token_type_hint=shared.RevokeTokenRequestTokenTypeHint.ACCESS_TOKEN,
)

res = s.o_auth2_token.revoke(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `request`                                                                | [shared.RevokeTokenRequest1](../../models/shared/revoketokenrequest1.md) | :heavy_check_mark:                                                       | The request object to use for the request.                               |


### Response

**[operations.RevokeOAuth2TokenResponse](../../models/operations/revokeoauth2tokenresponse.md)**

