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
			AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
		}),
	)

	cardRequest := shared.CardRequest{
		CardCvv:    moovgo.String("0123"),
		HolderName: moovgo.String("Jules Jackson"),
	}

	var accountID string = "8cfd9cf0-8cf1-4dc7-b48b-a0e013b33b1b"

	var xWaitFor *shared.SchemasWaitFor = shared.SchemasWaitForPaymentMethod.ToPointer()

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