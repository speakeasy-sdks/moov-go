# AccessToken
(*AccessToken*)

## Overview

When making requests to Moov from a browser, you can use OAuth with JSON Web Tokens (JWT).

Our authentication flow follows the OAuth 2.0 standard. With this endpoint, you'll [create an access token](https://docs.moov.io/guides/quick-start/#create-an-access-token) that you will pass along with API requests or when initializing Moov.js. To generate an authentication token, you’ll need to specify scopes that enable the token to use a specific set of API endpoints. To learn more, read our [scopes guide](https://docs.moov.io/guides/developer-tools/scopes/).


### Available Operations

* [Create](#create) - Create access token
* [Revoke](#revoke) - Revoke access token

## Create

Use the client_id and client_secret to generate an access token.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.AccessToken.Create(ctx, shared.ClientCredentialsGrantToAccessTokenRequest{
        ClientID: moovgo.String("5clTR_MdVrrkgxw2"),
        ClientSecret: moovgo.String("dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-"),
        GrantType: shared.GrantTypeClientCredentials,
        RefreshToken: moovgo.String("i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6..."),
        Scope: moovgo.String("/accounts.write"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccessTokenResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [shared.ClientCredentialsGrantToAccessTokenRequest](../../pkg/models/shared/clientcredentialsgranttoaccesstokenrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.PostOAuth2TokenResponse](../../pkg/models/operations/postoauth2tokenresponse.md), error**
| Error Object                       | Status Code                        | Content Type                       |
| ---------------------------------- | ---------------------------------- | ---------------------------------- |
| sdkerrors.AccessTokenErrorResponse | 400                                | application/json                   |
| sdkerrors.SDKError                 | 400-600                            | */*                                |

## Revoke

Allows clients to notify the authorization server that a previously obtained refresh or access token is no longer needed

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
	"net/http"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.AccessToken.Revoke(ctx, shared.RevokeTokenRequest{
        ClientID: moovgo.String("5clTR_MdVrrkgxw2"),
        ClientSecret: moovgo.String("dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-"),
        Token: "i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.RevokeTokenRequest](../../pkg/models/shared/revoketokenrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.RevokeOAuth2TokenResponse](../../pkg/models/operations/revokeoauth2tokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
