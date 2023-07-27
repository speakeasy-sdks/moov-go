<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    cardRequest := shared.CardRequest{
        BillingAddress: &shared.Address{
            AddressLine1: moov.String("123 Main Street"),
            AddressLine2: moov.String("Apt 302"),
            City: moov.String("Boulder"),
            Country: moov.String("US"),
            PostalCode: moov.String("80301"),
            StateOrProvince: moov.String("CO"),
        },
        CardCvv: moov.String("0123"),
        CardNumber: moov.String("corrupti"),
        CardOnFile: moov.Bool(false),
        Expiration: &shared.CardExpiration{
            Month: moov.String("01"),
            Year: moov.String("21"),
        },
        HolderName: moov.String("Jules Jackson"),
        MerchantAccountID: moov.String("9bd9d8d6-9a67-44e0-b467-cc8796ed151a"),
    }
    accountID := "05dfc2dd-f7cc-478c-a1ba-928fc816742c"
    xWaitFor := shared.SchemasWaitForPaymentMethod

    ctx := context.Background()
    res, err := s.Cards.LinkCard(ctx, cardRequest, accountID, xWaitFor)
    if err != nil {
        log.Fatal(err)
    }

    if res.Card != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->