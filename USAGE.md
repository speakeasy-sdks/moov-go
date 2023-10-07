<!-- Start SDK Example Usage -->


```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    cardRequest := shared.CardRequest{
        AdditionalProperties: map[string]interface{}{
            "transmit": "sultan",
        },
        BillingAddress: &shared.Address{
            AdditionalProperties: map[string]interface{}{
                "female": "McAllen",
            },
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        CardCvv: moovgo.String("0123"),
        Expiration: &shared.CardExpiration{
            Month: moovgo.String("01"),
            Year: moovgo.String("21"),
        },
        HolderName: moovgo.String("Jules Jackson"),
    }
    var accountID string = "cf1dc7b4-8ba0-4e01-bb33-b1b7a8af1bc7"
    var xWaitFor *shared.SchemasWaitFor = shared.SchemasWaitForPaymentMethod

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