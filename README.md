<div align="center">
   <img src="https://github.com/speakeasy-sdks/moov-go/assets/6267663/84c3c912-695e-4a03-8fa3-45076cccac26" width="300">
   <h1>Go SDK</h1>
   <p><strong>The modern way of proving identity</strong></p>
   <a href="https://docs.moov.io/api/"><img src="https://img.shields.io/static/v1?label=Docs&message=Docs&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/moov-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/moov-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://slack.moov.io/"><img src="https://img.shields.io/static/v1?label=Discord&message=Join&color=7289da&style=for-the-badge" /></a>
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/moov-go
```
<!-- End SDK Installation -->

## SDK Example Usage
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

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [AccessToken](docs/sdks/accesstoken/README.md)

* [Create](docs/sdks/accesstoken/README.md#create) - Create access token
* [Revoke](docs/sdks/accesstoken/README.md#revoke) - Revoke access token

### [Accounts](docs/sdks/accounts/README.md)

* [AssignCountry](docs/sdks/accounts/README.md#assigncountry) - Assign Account Countries
* [Create](docs/sdks/accounts/README.md#create) - Create account
* [Get](docs/sdks/accounts/README.md#get) - Get account
* [GetTosToken](docs/sdks/accounts/README.md#gettostoken) - Get terms of service token
* [List](docs/sdks/accounts/README.md#list) - List accounts
* [ListCountries](docs/sdks/accounts/README.md#listcountries) - Get Account Countries
* [Update](docs/sdks/accounts/README.md#update) - Patch account

### [Analytics](docs/sdks/analytics/README.md)

* [CountAccountsCreated](docs/sdks/analytics/README.md#countaccountscreated) - Count the number of profiles created by an individual or business
* [CountTransferStatuses](docs/sdks/analytics/README.md#counttransferstatuses) - Count the transfer statuses
* [LargestTransfer](docs/sdks/analytics/README.md#largesttransfer) - Return the largest number of transfers
* [SmallestTransfer](docs/sdks/analytics/README.md#smallesttransfer) - Return the smallest number of transfers
* [SumTransfers](docs/sdks/analytics/README.md#sumtransfers) - Sum all transfers across intervals

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

### [CardIssuing](docs/sdks/cardissuing/README.md)

* [RequestCard](docs/sdks/cardissuing/README.md#requestcard) - Request card
* [GetCard](docs/sdks/cardissuing/README.md#getcard) - Get issued card
* [GetCardFullDetails](docs/sdks/cardissuing/README.md#getcardfulldetails) - Get full card details
* [ListCards](docs/sdks/cardissuing/README.md#listcards) - List issued cards
* [UpdateCard](docs/sdks/cardissuing/README.md#updatecard) - Update issued card

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

### [Disputes](docs/sdks/disputes/README.md)

* [Get](docs/sdks/disputes/README.md#get) - Get Dispute by ID
* [List](docs/sdks/disputes/README.md#list) - List of all disputes

### [Enrichment](docs/sdks/enrichment/README.md)

* [GetAddress](docs/sdks/enrichment/README.md#getaddress) - Get address suggestions
* [GetAvatar](docs/sdks/enrichment/README.md#getavatar) - Get avatar
* [GetIndustries](docs/sdks/enrichment/README.md#getindustries) - List all industries
* [GetProfile](docs/sdks/enrichment/README.md#getprofile) - Get enriched profile

### [Files](docs/sdks/files/README.md)

* [Get](docs/sdks/files/README.md#get) - Get File Details
* [List](docs/sdks/files/README.md#list) - List files
* [Upload](docs/sdks/files/README.md#upload) - Upload File

### [Institutions](docs/sdks/institutions/README.md)

* [Search](docs/sdks/institutions/README.md#search) - Search institutions

### [PaymentMethods](docs/sdks/paymentmethods/README.md)

* [Get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [List](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [Representatives](docs/sdks/representatives/README.md)

* [Create](docs/sdks/representatives/README.md#create) - Create representative
* [Delete](docs/sdks/representatives/README.md#delete) - Delete a representative
* [Get](docs/sdks/representatives/README.md#get) - Get representative
* [List](docs/sdks/representatives/README.md#list) - List representatives
* [Update](docs/sdks/representatives/README.md#update) - Patch representative

### [Transactions](docs/sdks/transactions/README.md)

* [List](docs/sdks/transactions/README.md#list) - Get account transactions

### [Transfers](docs/sdks/transfers/README.md)

* [Cancel](docs/sdks/transfers/README.md#cancel) - Cancel or refund a card transfer
* [Create](docs/sdks/transfers/README.md#create) - Create a transfer
* [GenerateOptions](docs/sdks/transfers/README.md#generateoptions) - Generate transfer options
* [Get](docs/sdks/transfers/README.md#get) - Get a transfer
* [GetRefund](docs/sdks/transfers/README.md#getrefund) - Get refund details
* [ListRefunds](docs/sdks/transfers/README.md#listrefunds) - Get a list of refunds for a card transfer
* [Refund](docs/sdks/transfers/README.md#refund) - Refund a transfer
* [Update](docs/sdks/transfers/README.md#update) - Patch transfer metadata

### [Underwriting](docs/sdks/underwriting/README.md)

* [Get](docs/sdks/underwriting/README.md#get) - Retrieve underwriting details
* [Update](docs/sdks/underwriting/README.md#update) - Update underwriting details

### [Wallets](docs/sdks/wallets/README.md)

* [Get](docs/sdks/wallets/README.md#get) - Get wallet
* [GetTransaction](docs/sdks/wallets/README.md#gettransaction) - Get wallet transaction
* [List](docs/sdks/wallets/README.md#list) - List wallets
* [ListTransactions](docs/sdks/wallets/README.md#listtransactions) - List wallet transactions
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
