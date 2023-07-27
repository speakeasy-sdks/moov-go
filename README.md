# Petstore

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

req = shared.CreateAccountRequest(
    account_type=shared.AccountType.BUSINESS,
    capabilities=[
        shared.CapabilityID.COLLECT_FUNDS,
        shared.CapabilityID.WALLET,
        shared.CapabilityID.CARD_ISSUING,
    ],
    customer_support=shared.CreateAccountRequestCustomerSupport(
        address=shared.CreateAccountRequestCustomerSupportAddress(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        email='amanda@classbooker.dev',
        phone=shared.CreateAccountRequestCustomerSupportPhone(
            country_code='1',
            number='8185551212',
        ),
        website='www.wholebodyfitnessgym.com',
    ),
    foreign_id='4528aba-b9a1-11eb-8529-0242ac13003',
    metadata={
        "nulla": 'corrupti',
        "illum": 'vel',
        "error": 'deserunt',
    },
    mode=shared.Mode.PRODUCTION,
    profile=shared.CreateProfile(
        business=shared.CreateProfileBusiness(
            address=shared.CreateProfileBusinessAddress(
                address_line1='123 Main Street',
                address_line2='Apt 302',
                city='Boulder',
                country='US',
                postal_code='80301',
                state_or_province='CO',
            ),
            business_type=shared.BusinessType.LLC,
            description='Local fitness center paying out instructors',
            doing_business_as='Whole Body Fitness',
            email='amanda@classbooker.dev',
            industry_codes=shared.CreateProfileBusinessIndustryCodes(
                mcc='7997',
                naics='713940',
                sic='7991',
            ),
            legal_business_name='Whole Body Fitness LLC',
            phone=shared.CreateProfileBusinessPhone(
                country_code='1',
                number='8185551212',
            ),
            tax_id=shared.CreateProfileBusinessTaxID(
                ein=shared.Ein(
                    number='123-45-6789',
                ),
            ),
            website='www.wholebodyfitnessgym.com',
        ),
        individual=shared.CreateProfileIndividual(
            address=shared.CreateProfileIndividualAddress(
                address_line1='123 Main Street',
                address_line2='Apt 302',
                city='Boulder',
                country='US',
                postal_code='80301',
                state_or_province='CO',
            ),
            birth_date=shared.CreateProfileIndividualBirthDate(
                day=9,
                month=11,
                year=1989,
            ),
            email='amanda@classbooker.dev',
            government_id=shared.CreateProfileIndividualGovernmentID(
                itin=shared.CreateProfileIndividualGovernmentIDItin(
                    full='123-45-6789',
                    last_four='6789',
                ),
                ssn=shared.CreateProfileIndividualGovernmentIDSsn(
                    full='123-45-6789',
                    last_four='6789',
                ),
            ),
            name=shared.Name(
                first_name='Amanda',
                last_name='Yang',
                middle_name='Amanda',
                suffix='Jr',
            ),
            phone=shared.CreateProfileIndividualPhone(
                country_code='1',
                number='8185551212',
            ),
        ),
    ),
    settings=shared.CreateAccountRequestSettings(
        ach_payment=shared.CreateAccountRequestSettingsAchPayment(
            company_name='Whole Body Fitness',
        ),
        card_payment=shared.CreateAccountRequestSettingsCardPayment(
            statement_descriptor='Whole Body Fitness',
        ),
    ),
    terms_of_service=shared.CreateAccountRequestTermsOfService(
        token='kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4',
    ),
)

res = s.account.create(req)

if res.account is not None:
    # handle response
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [account](docs/sdks/account/README.md)

* [create](docs/sdks/account/README.md#create) - Create account
* [get](docs/sdks/account/README.md#get) - Get account
* [patch](docs/sdks/account/README.md#patch) - Patch account

### [account_countries](docs/sdks/accountcountries/README.md)

* [get](docs/sdks/accountcountries/README.md#get) - Get Account Countries
* [put](docs/sdks/accountcountries/README.md#put) - Assign Account Countries

### [account_issued_card_transactions](docs/sdks/accountissuedcardtransactions/README.md)

* [list](docs/sdks/accountissuedcardtransactions/README.md#list) - Get account transactions

### [analytics_accounts_created](docs/sdks/analyticsaccountscreated/README.md)

* [post](docs/sdks/analyticsaccountscreated/README.md#post) - Count the number of profiles created by an individual or business

### [analytics_transfer_largest](docs/sdks/analyticstransferlargest/README.md)

* [post](docs/sdks/analyticstransferlargest/README.md#post) - Return the largest number of transfers

### [analytics_transfer_smallest](docs/sdks/analyticstransfersmallest/README.md)

* [post](docs/sdks/analyticstransfersmallest/README.md#post) - Return the smallest number of transfers

### [analytics_transfer_statuses](docs/sdks/analyticstransferstatuses/README.md)

* [post](docs/sdks/analyticstransferstatuses/README.md#post) - Count the transfer statuses

### [analytics_transfer_sum](docs/sdks/analyticstransfersum/README.md)

* [post](docs/sdks/analyticstransfersum/README.md#post) - Sum all transfers across intervals

### [apple_pay_merchant_domains](docs/sdks/applepaymerchantdomains/README.md)

* [get](docs/sdks/applepaymerchantdomains/README.md#get) - Get Apple Pay domains
* [post](docs/sdks/applepaymerchantdomains/README.md#post) - Register Apple Pay domains
* [update](docs/sdks/applepaymerchantdomains/README.md#update) - Update Apple Pay domains

### [apple_pay_session](docs/sdks/applepaysession/README.md)

* [post](docs/sdks/applepaysession/README.md#post) - Create Apple Pay session

### [avatar](docs/sdks/avatar/README.md)

* [get](docs/sdks/avatar/README.md#get) - Get avatar

### [bank_account](docs/sdks/bankaccount/README.md)

* [post](docs/sdks/bankaccount/README.md#post) - Bank account

### [bank_id](docs/sdks/bankid/README.md)

* [delete](docs/sdks/bankid/README.md#delete) - Delete bank account
* [get](docs/sdks/bankid/README.md#get) - Get bank account

### [capability](docs/sdks/capability/README.md)

* [delete](docs/sdks/capability/README.md#delete) - Disable a capability for an account
* [get](docs/sdks/capability/README.md#get) - Get capability for account
* [post](docs/sdks/capability/README.md#post) - Request capabilities

### [card](docs/sdks/card/README.md)

* [delete](docs/sdks/card/README.md#delete) - Disable card
* [get](docs/sdks/card/README.md#get) - Get card
* [update](docs/sdks/card/README.md#update) - Update card

### [complete_micro_deposits](docs/sdks/completemicrodeposits/README.md)

* [put](docs/sdks/completemicrodeposits/README.md#put) - Complete micro-deposits

### [dispute](docs/sdks/dispute/README.md)

* [get](docs/sdks/dispute/README.md#get) - Get Dispute by ID

### [disputes](docs/sdks/disputes/README.md)

* [list](docs/sdks/disputes/README.md#list) - List of all disputes

### [enrichment_address](docs/sdks/enrichmentaddress/README.md)

* [get](docs/sdks/enrichmentaddress/README.md#get) - Get address suggestions

### [enrichment_profile](docs/sdks/enrichmentprofile/README.md)

* [get](docs/sdks/enrichmentprofile/README.md#get) - Get enriched profile

### [file](docs/sdks/file/README.md)

* [upload](docs/sdks/file/README.md#upload) - Upload File

### [file_details](docs/sdks/filedetails/README.md)

* [get](docs/sdks/filedetails/README.md#get) - Get File Details

### [files](docs/sdks/files/README.md)

* [list](docs/sdks/files/README.md#list) - List files

### [full_issued_card](docs/sdks/fullissuedcard/README.md)

* [get](docs/sdks/fullissuedcard/README.md#get) - Get full card details

### [industries](docs/sdks/industries/README.md)

* [list](docs/sdks/industries/README.md#list) - List all industries

### [initiate_micro_deposits](docs/sdks/initiatemicrodeposits/README.md)

* [post](docs/sdks/initiatemicrodeposits/README.md#post) - Initiate micro-deposits

### [institutions](docs/sdks/institutions/README.md)

* [search](docs/sdks/institutions/README.md#search) - Search institutions

### [issued_card](docs/sdks/issuedcard/README.md)

* [get](docs/sdks/issuedcard/README.md#get) - Get issued card
* [update](docs/sdks/issuedcard/README.md#update) - Update issued card

### [link_apple_pay_token](docs/sdks/linkapplepaytoken/README.md)

* [post](docs/sdks/linkapplepaytoken/README.md#post) - Link Apple Pay token

### [link_card](docs/sdks/linkcard/README.md)

* [post](docs/sdks/linkcard/README.md#post) - Link card

### [list_accounts](docs/sdks/listaccounts/README.md)

* [get](docs/sdks/listaccounts/README.md#get) - List accounts

### [list_bank_accounts](docs/sdks/listbankaccounts/README.md)

* [get](docs/sdks/listbankaccounts/README.md#get) - List bank accounts

### [list_capability](docs/sdks/listcapability/README.md)

* [get](docs/sdks/listcapability/README.md#get) - List capabilities for account

### [list_cards](docs/sdks/listcards/README.md)

* [get](docs/sdks/listcards/README.md#get) - List cards

### [list_issued_cards](docs/sdks/listissuedcards/README.md)

* [get](docs/sdks/listissuedcards/README.md#get) - List issued cards

### [network_i_ds](docs/sdks/networkids/README.md)

* [get](docs/sdks/networkids/README.md#get) - Get network IDs of an account

### [o_auth2_token](docs/sdks/oauth2token/README.md)

* [post](docs/sdks/oauth2token/README.md#post) - Create access token
* [revoke](docs/sdks/oauth2token/README.md#revoke) - Revoke access token

### [payment_method](docs/sdks/paymentmethod/README.md)

* [get](docs/sdks/paymentmethod/README.md#get) - Get payment method

### [payment_methods](docs/sdks/paymentmethods/README.md)

* [get](docs/sdks/paymentmethods/README.md#get) - List payment methods

### [refund](docs/sdks/refund/README.md)

* [get](docs/sdks/refund/README.md#get) - Get refund details

### [representative](docs/sdks/representative/README.md)

* [create](docs/sdks/representative/README.md#create) - Create representative
* [delete](docs/sdks/representative/README.md#delete) - Delete a representative
* [get](docs/sdks/representative/README.md#get) - Get representative
* [patch](docs/sdks/representative/README.md#patch) - Patch representative

### [representatives](docs/sdks/representatives/README.md)

* [list](docs/sdks/representatives/README.md#list) - List representatives

### [request_card](docs/sdks/requestcard/README.md)

* [post](docs/sdks/requestcard/README.md#post) - Request card

### [terms_of_service_token](docs/sdks/termsofservicetoken/README.md)

* [get](docs/sdks/termsofservicetoken/README.md#get) - Get terms of service token

### [transfer](docs/sdks/transfer/README.md)

* [cancel](docs/sdks/transfer/README.md#cancel) - Cancel or refund a card transfer
* [create](docs/sdks/transfer/README.md#create) - Create a transfer
* [get](docs/sdks/transfer/README.md#get) - Get a transfer
* [patch](docs/sdks/transfer/README.md#patch) - Patch transfer metadata
* [refund](docs/sdks/transfer/README.md#refund) - Refund a transfer

### [transfer_options](docs/sdks/transferoptions/README.md)

* [create](docs/sdks/transferoptions/README.md#create) - Generate transfer options

### [transfers](docs/sdks/transfers/README.md)

* [get](docs/sdks/transfers/README.md#get) - Get a list of refunds for a card transfer

### [underwriting](docs/sdks/underwriting/README.md)

* [get](docs/sdks/underwriting/README.md#get) - Retrieve underwriting details
* [update](docs/sdks/underwriting/README.md#update) - Update underwriting details

### [wallet_transaction](docs/sdks/wallettransaction/README.md)

* [get](docs/sdks/wallettransaction/README.md#get) - Get wallet transaction

### [wallet_transactions](docs/sdks/wallettransactions/README.md)

* [list](docs/sdks/wallettransactions/README.md#list) - List wallet transactions

### [wallets_for_account](docs/sdks/walletsforaccount/README.md)

* [get](docs/sdks/walletsforaccount/README.md#get) - Get wallet
* [list](docs/sdks/walletsforaccount/README.md#list) - List wallets
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
