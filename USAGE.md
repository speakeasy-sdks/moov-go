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
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.AccessToken.Create(ctx, shared.ClientCredentialsGrantToAccessTokenRequest{
        ClientID: moov.String("5clTR_MdVrrkgxw2"),
        ClientSecret: moov.String("dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-"),
        GrantType: shared.ClientCredentialsGrantToAccessTokenRequestGrantTypeRefreshToken,
        RefreshToken: moov.String("i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6..."),
        Scope: moov.String("/accounts.write"),
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