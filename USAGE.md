<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.AccessToken.Create(ctx, shared.ClientCredentialsGrantToAccessTokenRequest{
        ClientID: petstore.String("5clTR_MdVrrkgxw2"),
        ClientSecret: petstore.String("dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-"),
        GrantType: shared.ClientCredentialsGrantToAccessTokenRequestGrantTypeRefreshToken,
        RefreshToken: petstore.String("i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6..."),
        Scope: petstore.String("/accounts.write"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AccessTokenResponse != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->