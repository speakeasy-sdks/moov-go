# Moov

<!-- Start SDK Installation -->
## SDK Installation

```bash
pip install git+https://github.com/speakeasy-sdks/moov-go.git
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->


```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.ClientCredentialsGrantToAccessTokenRequest(
    client_id='5clTR_MdVrrkgxw2',
    client_secret='dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-',
    grant_type=shared.ClientCredentialsGrantToAccessTokenRequestGrantType.REFRESH_TOKEN,
    refresh_token='i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...',
    scope='/accounts.write',
)

res = s.access_token.create(req)

if res.access_token_response is not None:
    # handle response
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [access_token](docs/sdks/accesstoken/README.md)

* [create](docs/sdks/accesstoken/README.md#create) - Create access token
* [revoke](docs/sdks/accesstoken/README.md#revoke) - Revoke access token

### [accounts](docs/sdks/accounts/README.md)

* [assign_country](docs/sdks/accounts/README.md#assign_country) - Assign Account Countries
* [create](docs/sdks/accounts/README.md#create) - Create account
* [get](docs/sdks/accounts/README.md#get) - Get account
* [get_tos_token](docs/sdks/accounts/README.md#get_tos_token) - Get terms of service token
* [list](docs/sdks/accounts/README.md#list) - List accounts
* [list_countries](docs/sdks/accounts/README.md#list_countries) - Get Account Countries
* [update](docs/sdks/accounts/README.md#update) - Patch account

### [analytics](docs/sdks/analytics/README.md)

* [count_accounts_created](docs/sdks/analytics/README.md#count_accounts_created) - Count the number of profiles created by an individual or business
* [count_transfer_statuses](docs/sdks/analytics/README.md#count_transfer_statuses) - Count the transfer statuses
* [largest_transfer](docs/sdks/analytics/README.md#largest_transfer) - Return the largest number of transfers
* [smallest_transfer](docs/sdks/analytics/README.md#smallest_transfer) - Return the smallest number of transfers
* [sum_transfers](docs/sdks/analytics/README.md#sum_transfers) - Sum all transfers across intervals

### [bank_accounts](docs/sdks/bankaccounts/README.md)

* [initiate_micro_deposits](docs/sdks/bankaccounts/README.md#initiate_micro_deposits) - Initiate micro-deposits
* [complete_micro_deposits](docs/sdks/bankaccounts/README.md#complete_micro_deposits) - Complete micro-deposits
* [delete](docs/sdks/bankaccounts/README.md#delete) - Delete bank account
* [get](docs/sdks/bankaccounts/README.md#get) - Get bank account
* [link](docs/sdks/bankaccounts/README.md#link) - Bank account
* [list](docs/sdks/bankaccounts/README.md#list) - List bank accounts

### [capabilities](docs/sdks/capabilities/README.md)

* [delete](docs/sdks/capabilities/README.md#delete) - Disable a capability for an account
* [get](docs/sdks/capabilities/README.md#get) - Get capability for account
* [list](docs/sdks/capabilities/README.md#list) - List capabilities for account
* [request](docs/sdks/capabilities/README.md#request) - Request capabilities

### [card_issuing](docs/sdks/cardissuing/README.md)

* [request_card](docs/sdks/cardissuing/README.md#request_card) - Request card
* [get_card](docs/sdks/cardissuing/README.md#get_card) - Get issued card
* [get_card_full_details](docs/sdks/cardissuing/README.md#get_card_full_details) - Get full card details
* [list_cards](docs/sdks/cardissuing/README.md#list_cards) - List issued cards
* [update_card](docs/sdks/cardissuing/README.md#update_card) - Update issued card

### [cards](docs/sdks/cards/README.md)

* [link_apple_pay_token](docs/sdks/cards/README.md#link_apple_pay_token) - Link Apple Pay token
* [link_card](docs/sdks/cards/README.md#link_card) - Link card
* [list_cards](docs/sdks/cards/README.md#list_cards) - List cards
* [create_apple_pay_session](docs/sdks/cards/README.md#create_apple_pay_session) - Create Apple Pay session
* [delete](docs/sdks/cards/README.md#delete) - Disable card
* [get](docs/sdks/cards/README.md#get) - Get card
* [list_apple_pay_domains](docs/sdks/cards/README.md#list_apple_pay_domains) - Get Apple Pay domains
* [register_apple_pay_domain](docs/sdks/cards/README.md#register_apple_pay_domain) - Register Apple Pay domains
* [update](docs/sdks/cards/README.md#update) - Update card
* [update_apple_pay_domains](docs/sdks/cards/README.md#update_apple_pay_domains) - Update Apple Pay domains

### [disputes](docs/sdks/disputes/README.md)

* [get](docs/sdks/disputes/README.md#get) - Get Dispute by ID
* [list](docs/sdks/disputes/README.md#list) - List of all disputes

### [enrichment](docs/sdks/enrichment/README.md)

* [get_address](docs/sdks/enrichment/README.md#get_address) - Get address suggestions
* [get_avatar](docs/sdks/enrichment/README.md#get_avatar) - Get avatar
* [get_industries](docs/sdks/enrichment/README.md#get_industries) - List all industries
* [get_profile](docs/sdks/enrichment/README.md#get_profile) - Get enriched profile

### [files](docs/sdks/files/README.md)

* [get](docs/sdks/files/README.md#get) - Get File Details
* [list](docs/sdks/files/README.md#list) - List files
* [upload](docs/sdks/files/README.md#upload) - Upload File

### [institutions](docs/sdks/institutions/README.md)

* [search](docs/sdks/institutions/README.md#search) - Search institutions

### [payment_methods](docs/sdks/paymentmethods/README.md)

* [get](docs/sdks/paymentmethods/README.md#get) - Get payment method
* [list](docs/sdks/paymentmethods/README.md#list) - List payment methods

### [representatives](docs/sdks/representatives/README.md)

* [create](docs/sdks/representatives/README.md#create) - Create representative
* [delete](docs/sdks/representatives/README.md#delete) - Delete a representative
* [get](docs/sdks/representatives/README.md#get) - Get representative
* [list](docs/sdks/representatives/README.md#list) - List representatives
* [update](docs/sdks/representatives/README.md#update) - Patch representative

### [transactions](docs/sdks/transactions/README.md)

* [list](docs/sdks/transactions/README.md#list) - Get account transactions

### [transfers](docs/sdks/transfers/README.md)

* [cancel](docs/sdks/transfers/README.md#cancel) - Cancel or refund a card transfer
* [create](docs/sdks/transfers/README.md#create) - Create a transfer
* [generate_options](docs/sdks/transfers/README.md#generate_options) - Generate transfer options
* [get](docs/sdks/transfers/README.md#get) - Get a transfer
* [get_refund](docs/sdks/transfers/README.md#get_refund) - Get refund details
* [list_refunds](docs/sdks/transfers/README.md#list_refunds) - Get a list of refunds for a card transfer
* [refund](docs/sdks/transfers/README.md#refund) - Refund a transfer
* [update](docs/sdks/transfers/README.md#update) - Patch transfer metadata

### [underwriting](docs/sdks/underwriting/README.md)

* [get](docs/sdks/underwriting/README.md#get) - Retrieve underwriting details
* [update](docs/sdks/underwriting/README.md#update) - Update underwriting details

### [wallets](docs/sdks/wallets/README.md)

* [get](docs/sdks/wallets/README.md#get) - Get wallet
* [get_transaction](docs/sdks/wallets/README.md#get_transaction) - Get wallet transaction
* [list](docs/sdks/wallets/README.md#list) - List wallets
* [list_transactions](docs/sdks/wallets/README.md#list_transactions) - List wallet transactions
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
