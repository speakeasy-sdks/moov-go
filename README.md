<div align="center">
   <img src="https://github.com/speakeasy-sdks/moov-go/assets/68016351/fd9ab239-6d52-474d-800b-385ee20082d8" width="300">
   <h1>Go SDK</h1>
   <p><strong>Accept, store, send, & spend – all in one place.</strong></p>
   <a href="https://docs.moov.io/api/"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://slack.moov.io/"><img src="https://img.shields.io/static/v1?label=Slack&message=Join&color=7289da&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/moov-go
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Accounts](docs/sdks/accounts/README.md)

* [AssignCountry](docs/sdks/accounts/README.md#assigncountry) - Assign Account Countries
* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetTosToken](docs/sdks/accounts/README.md#gettostoken) - Get terms of service token
* [List](docs/sdks/accounts/README.md#list) - List accounts
* [ListCountries](docs/sdks/accounts/README.md#listcountries) - Get Account Countries
* [Update](docs/sdks/accounts/README.md#update) - Patch account

### [Cards](docs/sdks/cards/README.md)

* [LinkApplePayToken](docs/sdks/cards/README.md#linkapplepaytoken) - Link Apple Pay token
* [LinkCard](docs/sdks/cards/README.md#linkcard) - Link card
* [ListCards](docs/sdks/cards/README.md#listcards) - List cards
* [CreateApplePaySession](docs/sdks/cards/README.md#createapplepaysession) - Create Apple Pay session
* [Delete](docs/sdks/cards/README.md#delete) - Disable card
* [Get](docs/sdks/cards/README.md#get) - Get card
* [ListApplePayDomains](docs/sdks/cards/README.md#listapplepaydomains) - Get Apple Pay domains
* [RegisterApplePayDomain](docs/sdks/cards/README.md#registerapplepaydomain) - Register Apple Pay domains
* [Update](docs/sdks/cards/README.md#update) - Update card
* [UpdateApplePayDomains](docs/sdks/cards/README.md#updateapplepaydomains) - Update Apple Pay domains

### [BankAccounts](docs/sdks/bankaccounts/README.md)

* [InitiateMicroDeposits](docs/sdks/bankaccounts/README.md#initiatemicrodeposits) - Initiate micro-deposits
* [CompleteMicroDeposits](docs/sdks/bankaccounts/README.md#completemicrodeposits) - Complete micro-deposits
* [Delete](docs/sdks/bankaccounts/README.md#delete) - Delete bank account
* [Get](docs/sdks/bankaccounts/README.md#get) - Get bank account
* [Link](docs/sdks/bankaccounts/README.md#link) - Bank account
* [List](docs/sdks/bankaccounts/README.md#list) - List bank accounts

### [Capabilities](docs/sdks/capabilities/README.md)

* [Delete](docs/sdks/capabilities/README.md#delete) - Disable a capability for an account
* [Get](docs/sdks/capabilities/README.md#get) - Get capability for account
* [List](docs/sdks/capabilities/README.md#list) - List capabilities for account
* [Request](docs/sdks/capabilities/README.md#request) - Request capabilities

### [Files](docs/sdks/files/README.md)

* [Get](docs/sdks/files/README.md#get) - Get File Details
* [List](docs/sdks/files/README.md#list) - List files
* [Upload](docs/sdks/files/README.md#upload) - Upload File

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [Representatives](docs/sdks/representatives/README.md)

* [Create](docs/sdks/representatives/README.md#create) - Create representative
* [Delete](docs/sdks/representatives/README.md#delete) - Delete a representative
* [Get](docs/sdks/representatives/README.md#get) - Get representative
* [List](docs/sdks/representatives/README.md#list) - List representatives
* [Update](docs/sdks/representatives/README.md#update) - Patch representative

### [Underwriting](docs/sdks/underwriting/README.md)

* [Get](docs/sdks/underwriting/README.md#get) - Retrieve underwriting details
* [Update](docs/sdks/underwriting/README.md#update) - Update underwriting details

### [Wallets](docs/sdks/wallets/README.md)

* [Get](docs/sdks/wallets/README.md#get) - Get wallet
* [GetTransaction](docs/sdks/wallets/README.md#gettransaction) - Get wallet transaction
* [List](docs/sdks/wallets/README.md#list) - List wallets
* [ListTransactions](docs/sdks/wallets/README.md#listtransactions) - List wallet transactions

### [Analytics](docs/sdks/analytics/README.md)

* [CountAccountsCreated](docs/sdks/analytics/README.md#countaccountscreated) - Count the number of profiles created by an individual or business
* [CountTransferStatuses](docs/sdks/analytics/README.md#counttransferstatuses) - Count the transfer statuses
* [LargestTransfer](docs/sdks/analytics/README.md#largesttransfer) - Return the largest number of transfers
* [SmallestTransfer](docs/sdks/analytics/README.md#smallesttransfer) - Return the smallest number of transfers
* [SumTransfers](docs/sdks/analytics/README.md#sumtransfers) - Sum all transfers across intervals

### [Enrichment](docs/sdks/enrichment/README.md)

* [GetAddress](docs/sdks/enrichment/README.md#getaddress) - Get address suggestions
* [GetAvatar](docs/sdks/enrichment/README.md#getavatar) - Get avatar
* [GetIndustries](docs/sdks/enrichment/README.md#getindustries) - List all industries
* [GetProfile](docs/sdks/enrichment/README.md#getprofile) - Get enriched profile

### [Disputes](docs/sdks/disputes/README.md)

* [Get](docs/sdks/disputes/README.md#get) - Get Dispute by ID
* [List](docs/sdks/disputes/README.md#list) - List of all disputes

### [Institutions](docs/sdks/institutions/README.md)

* [Search](docs/sdks/institutions/README.md#search) - Search institutions

### [CardIssuing](docs/sdks/cardissuing/README.md)

* [RequestCard](docs/sdks/cardissuing/README.md#requestcard) - Request card
* [GetCard](docs/sdks/cardissuing/README.md#getcard) - Get issued card
* [GetCardFullDetails](docs/sdks/cardissuing/README.md#getcardfulldetails) - Get full card details
* [ListCards](docs/sdks/cardissuing/README.md#listcards) - List issued cards
* [UpdateCard](docs/sdks/cardissuing/README.md#updatecard) - Update issued card

### [Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - Get account transactions

### [AccessToken](docs/sdks/accesstoken/README.md)

* [Create](docs/sdks/accesstoken/README.md#create) - Create access token
* [Revoke](docs/sdks/accesstoken/README.md#revoke) - Revoke access token

### [Transfers](docs/sdks/transfers/README.md)

* [Cancel](docs/sdks/transfers/README.md#cancel) - Cancel or refund a card transfer
* [Create](docs/sdks/transfers/README.md#create) - Create a transfer
* [GenerateOptions](docs/sdks/transfers/README.md#generateoptions) - Generate transfer options
* [Get](docs/sdks/transfers/README.md#get) - Get a transfer
* [GetRefund](docs/sdks/transfers/README.md#getrefund) - Get refund details
* [ListRefunds](docs/sdks/transfers/README.md#listrefunds) - Get a list of refunds for a card transfer
* [Refund](docs/sdks/transfers/README.md#refund) - Refund a transfer
* [Update](docs/sdks/transfers/README.md#update) - Patch transfer metadata
<!-- End Available Resources and Operations [operations] -->







<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->



<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object                     | Status Code                      | Content Type                     |
| -------------------------------- | -------------------------------- | -------------------------------- |
| sdkerrors.CapabilityRequestError | 409                              | application/json                 |
| sdkerrors.SDKError               | 400-600                          | */*                              |

### Example

```go
package main

import (
	"context"
	"errors"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"log"
)

func main() {
	s := moovgo.New(
		moovgo.WithSecurity(shared.Security{
			AccessToken: moovgo.String(""),
		}),
	)

	addCapabilityRequest := shared.AddCapabilityRequest{
		Capabilities: []shared.CapabilityID{
			shared.CapabilityIDTransfers,
		},
	}

	var accountID string = "12e6e103-56d1-4f09-9ae6-2352496ce763"

	ctx := context.Background()
	res, err := s.Capabilities.Request(ctx, addCapabilityRequest, accountID)
	if err != nil {

		var e *sdkerrors.CapabilityRequestError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->



<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.moov.io` | None |

#### Example

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
		moovgo.WithServerIndex(0),
		moovgo.WithSecurity(shared.Security{
			AccessToken: moovgo.String(""),
		}),
	)

	countries := shared.Countries{
		Countries: []string{
			"U",
			"n",
			"i",
			"t",
			"e",
			"d",
			" ",
			"S",
			"t",
			"a",
			"t",
			"e",
			"s",
		},
	}

	var accountID string = "f51a6841-150c-4bc7-8c84-e981f74cfa3f"

	ctx := context.Background()
	res, err := s.Accounts.AssignCountry(ctx, countries, accountID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Countries != nil {
		// handle response
	}
}

```


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		moovgo.WithServerURL("https://api.moov.io"),
		moovgo.WithSecurity(shared.Security{
			AccessToken: moovgo.String(""),
		}),
	)

	countries := shared.Countries{
		Countries: []string{
			"U",
			"n",
			"i",
			"t",
			"e",
			"d",
			" ",
			"S",
			"t",
			"a",
			"t",
			"e",
			"s",
		},
	}

	var accountID string = "f51a6841-150c-4bc7-8c84-e981f74cfa3f"

	ctx := context.Background()
	res, err := s.Accounts.AssignCountry(ctx, countries, accountID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Countries != nil {
		// handle response
	}
}

```

### Override Server URL Per-Operation

The server URL can also be overridden on a per-operation basis, provided a server list was specified for the operation. For example:
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
	res, err := s.Cards.LinkCard(ctx, operations.WithServerURL("https://cards.moov.io/api"), cardRequest, accountID, xWaitFor)
	if err != nil {
		log.Fatal(err)
	}

	if res.Card != nil {
		// handle response
	}
}

```
<!-- End Server Selection [server] -->



<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->



<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security schemes globally:

| Name                | Type                | Scheme              |
| ------------------- | ------------------- | ------------------- |
| `AccessToken`       | oauth2              | OAuth2 token        |
| `GatewayAuth`       | http                | HTTP Bearer         |
| `OAuth2Credentials` | http                | HTTP Basic          |

You can set the security parameters through the `WithSecurity` option when initializing the SDK client instance. The selected scheme will be used by default to authenticate with the API for all operations that support it. For example:
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

	countries := shared.Countries{
		Countries: []string{
			"U",
			"n",
			"i",
			"t",
			"e",
			"d",
			" ",
			"S",
			"t",
			"a",
			"t",
			"e",
			"s",
		},
	}

	var accountID string = "f51a6841-150c-4bc7-8c84-e981f74cfa3f"

	ctx := context.Background()
	res, err := s.Accounts.AssignCountry(ctx, countries, accountID)
	if err != nil {
		log.Fatal(err)
	}

	if res.Countries != nil {
		// handle response
	}
}

```
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->



### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
