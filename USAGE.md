<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    cardRequest := shared.CardRequest{
        BillingAddress: &shared.Address{
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        CardCvv: moovgo.String("0123"),
        CardNumber: moovgo.String("lavender parallel"),
        CardOnFile: moovgo.Bool(false),
        Expiration: &shared.CardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
        HolderName: moovgo.String("Jules Jackson"),
        MerchantAccountID: moovgo.String("cf08cf1d-c7b4-48ba-8e01-3b33b1b7a8af"),
    }
    accountID := "1bc7a2ba-2e98-4e21-8b15-24f359b041e3"
    xWaitFor := "circuit"

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