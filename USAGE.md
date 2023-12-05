<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"log"
)

func main() {
	s := moovgo.New(
		moovgo.WithSecurity(shared.Security{
			AccessToken: moovgo.String(""),
		}),
	)

	cardRequest := shared.CardRequest{
		BillingAddress: &shared.Address{
			AddressLine1:    moovgo.String("123 Main Street"),
			AddressLine2:    moovgo.String("Apt 302"),
			City:            moovgo.String("Boulder"),
			Country:         moovgo.String("US"),
			PostalCode:      moovgo.String("80301"),
			StateOrProvince: moovgo.String("CO"),
		},
		CardCvv: moovgo.String("0123"),
		Expiration: &shared.CardExpiration{
			Month: moovgo.String("01"),
			Year:  moovgo.String("21"),
		},
		HolderName: moovgo.String("Jules Jackson"),
	}

	var accountID string = "8cfd9cf0-8cf1-4dc7-b48b-a0e013b33b1b"

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
<!-- End SDK Example Usage [usage] -->