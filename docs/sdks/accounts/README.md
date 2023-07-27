# accounts

## Overview

[Accounts](https://docs.moov.io/guides/accounts/) represent a legal entity (either a business or an individual) in Moov. You can create an account for yourself or set up accounts for others, requesting different [capabilities](https://docs.moov.io/guides/accounts/capabilities/) depending on what you need to be able to do with that account. You can retrieve an account to get details on the business or individual account holder, such as an email address or employer identification number (EIN). 

Based on the type of account and its requested capabilities, we have certain [verification requirements](https://docs.moov.io/guides/accounts/identity-verification/). To see what capabilities that account has, you can use the [GET capability endpoint](https://docs.moov.io/api/#operation/getCapability). 

When you sign up for the Moov Dashboard, you will have a **facilitator account** which can be used to facilitate money movement between other accounts. A facilitator account will not show up in your list of accounts and cannot be created via API. To update your facilitator account information, use the settings page of the Moov Dashboard. 


### Available Operations

* [assign_country](#assign_country) - Assign Account Countries
* [create](#create) - Create account
* [get](#get) - Get account
* [get_tos_token](#get_tos_token) - Get terms of service token
* [list](#list) - List accounts
* [list_countries](#list_countries) - Get Account Countries
* [update](#update) - Patch account

## assign_country

Assign the countries of operation for an account. This endpoint will always overwrite the previously assigned values. <br><br> To update the account countries, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PutAccountCountriesRequest(
    countries=shared.Countries(
        countries=[
            'United States',
            'United States',
            'United States',
            'United States',
        ],
    ),
    account_id='9d8d69a6-74e0-4f46-bcc8-796ed151a05d',
)

res = s.accounts.assign_country(req)

if res.countries is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.PutAccountCountriesRequest](../../models/operations/putaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.PutAccountCountriesResponse](../../models/operations/putaccountcountriesresponse.md)**


## create

You can create accounts for your users by passing the required information to Moov. <br><br> Note that `mode` field is only required when creating a facilitator account. All non-facilitator account creation requests will ignore the mode field provided and be set to the calling facilitator's mode. <br><br> If you are creating an account with the business type "llc", "partnership", or "privateCorporation", you will need to also provide [business representatives](https://docs.moov.io/api/#tag/Representatives) after creating the account for verification purposes. Once you've added your business owners as representatives, you'll then need to [patch your Moov account](https://docs.moov.io/api/#operation/patchAccount) to indicate that ownership information is complete. Read more on our [business verification requirements here](https://docs.moov.io/guides/accounts/business-verification/). <br><br> When creating an account, you will need to specify the `/accounts.write` scope.

### Example Usage

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
        shared.CapabilityID.WALLET,
        shared.CapabilityID.TRANSFERS,
        shared.CapabilityID.CARD_ISSUING,
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
        "molestiae": 'quod',
        "quod": 'esse',
        "totam": 'porro',
        "dolorum": 'dicta',
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

res = s.accounts.create(req)

if res.account is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [shared.CreateAccountRequest](../../models/shared/createaccountrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.CreateAccountResponse](../../models/operations/createaccountresponse.md)**


## get

Retrieves details for the account with the specified ID. <br><br> To get an account, you will need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAccountRequest(
    account_id='ba928fc8-1674-42cb-b392-05929396fea7',
)

res = s.accounts.get(req)

if res.account is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetAccountRequest](../../models/operations/getaccountrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetAccountResponse](../../models/operations/getaccountresponse.md)**


## get_tos_token

Generates a non-expiring token that can then be used to accept Moov's terms of service. This token can only be generated via API. Any Moov account requesting the `collect-funds`, `send-funds`, `wallet`, or `card-issuing` capabilities must accept Moov's terms of service, then have the generated terms of service token patched to the account. Read more on our [quick start guide](https://docs.moov.io/guides/quick-start/#request-capabilities).

### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)


res = s.accounts.get_tos_token()

if res.terms_of_service_token is not None:
    # handle response
```


### Response

**[operations.GetTermsOfServiceTokenResponse](../../models/operations/gettermsofservicetokenresponse.md)**


## list

List or search accounts to which the caller is connected.<br><br>
All supported query parameters are optional. If none are provided
the response will include all connected accounts. Pagination is
supported via the `skip` and `count` query parameters.<br><br>
Searching by name and email will overlap and return results based on relevance.
<br><br> To list connected accounts, you must specify the `/accounts.read` scope when retrieving the authentication token.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListAccountsRequest(
    count=359508,
    email='Humberto.Turcotte6@yahoo.com',
    foreign_id='4528aba-b9a1-11eb-8529-0242ac13003',
    include_disconnected=False,
    name='Carlton O'Hara',
    skip=210382,
    type=shared.AccountType.BUSINESS,
    verification_status=shared.AccountVerificationStatus.RESUBMIT,
)

res = s.accounts.list(req)

if res.accounts is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.ListAccountsRequest](../../models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.ListAccountsResponse](../../models/operations/listaccountsresponse.md)**


## list_countries

Retrieve the specified countries of operation for an account. <br><br> To get the list of countries, you'll need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAccountCountriesRequest(
    account_id='2c595590-7aff-41a3-a2fa-9467739251aa',
)

res = s.accounts.list_countries(req)

if res.countries is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetAccountCountriesRequest](../../models/operations/getaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetAccountCountriesResponse](../../models/operations/getaccountcountriesresponse.md)**


## update

To patch an account, you must specify the `/accounts/{accountID}/profile.write` scope and provide the changed information.  

When **can** profile data be updated:  
  + For unverified accounts, all profile data can be edited.
  + During the verification process, missing or incomplete profile data can be edited.
  + Verified accounts can only add missing profile data.  

  When **can't** profile data be updated:  
  + Verified accounts cannot change any existing profile data.  

If you need to update information in a locked state, please contact Moov support.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PatchAccountRequest(
    patch_account_request=shared.PatchAccountRequest(
        customer_support=shared.PatchAccountRequestCustomerSupport(
            address=shared.PatchAccountRequestCustomerSupportAddress(
                address_line1='123 Main Street',
                address_line2='Apt 302',
                city='Boulder',
                country='US',
                postal_code='80301',
                state_or_province='CO',
            ),
            email='amanda@classbooker.dev',
            phone=shared.PatchAccountRequestCustomerSupportPhone(
                country_code='1',
                number='8185551212',
            ),
            website='www.wholebodyfitnessgym.com',
        ),
        foreign_id='4528aba-b9a1-11eb-8529-0242ac13003',
        metadata={
            "odit": 'quo',
            "sequi": 'tenetur',
        },
        profile=shared.PatchAccountRequestProfile(
            business=shared.PatchAccountRequestProfileBusiness(
                address=shared.PatchAccountRequestProfileBusinessAddress(
                    address_line1='123 Main Street',
                    address_line2='Apt 302',
                    city='Boulder',
                    country='US',
                    postal_code='80301',
                    state_or_province='CO',
                ),
                business_type=shared.PatchAccountRequestProfileBusinessBusinessType.LLC,
                description='Local fitness center paying out instructors',
                doing_business_as='Whole Body Fitness',
                email='amanda@classbooker.dev',
                industry_codes=shared.PatchAccountRequestProfileBusinessIndustryCodes(
                    mcc='7997',
                    naics='713940',
                    sic='7991',
                ),
                legal_business_name='Whole Body Fitness LLC',
                owners_provided=False,
                phone=shared.PatchAccountRequestProfileBusinessPhone(
                    country_code='1',
                    number='8185551212',
                ),
                tax_id=shared.PatchAccountRequestProfileBusinessTaxID(
                    ein=shared.Ein(
                        number='123-45-6789',
                    ),
                ),
                website='www.wholebodyfitnessgym.com',
            ),
            individual=shared.PatchAccountRequestProfileIndividual(
                address=shared.PatchAccountRequestProfileIndividualAddress(
                    address_line1='123 Main Street',
                    address_line2='Apt 302',
                    city='Boulder',
                    country='US',
                    postal_code='80301',
                    state_or_province='CO',
                ),
                birth_date=shared.PatchAccountRequestProfileIndividualBirthDate(
                    day=9,
                    month=11,
                    year=1989,
                ),
                email='amanda@classbooker.dev',
                government_id=shared.PatchAccountRequestProfileIndividualGovernmentID(
                    itin=shared.PatchAccountRequestProfileIndividualGovernmentIDItin(
                        full='123-45-6789',
                        last_four='6789',
                    ),
                    ssn=shared.PatchAccountRequestProfileIndividualGovernmentIDSsn(
                        full='123-45-6789',
                        last_four='6789',
                    ),
                ),
                name=shared.PatchAccountRequestProfileIndividualName(
                    first_name='Amanda',
                    last_name='Yang',
                    middle_name='Amanda',
                    suffix='Jr',
                ),
                phone=shared.PatchAccountRequestProfileIndividualPhone(
                    country_code='1',
                    number='8185551212',
                ),
            ),
        ),
        settings=shared.PatchAccountRequestSettings(
            ach_payment=shared.PatchAccountRequestSettingsAchPayment(
                company_name='Whole Body Fitness',
            ),
            card_payment=shared.PatchAccountRequestSettingsCardPayment(
                statement_descriptor='Whole Body Fitness',
            ),
        ),
        terms_of_service=shared.PatchAccountRequestTermsOfService(
            token='kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4',
        ),
    ),
    account_id='5ad019da-1ffe-478f-897b-0074f15471b5',
)

res = s.accounts.update(req)

if res.account is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.PatchAccountRequest](../../models/operations/patchaccountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.PatchAccountResponse](../../models/operations/patchaccountresponse.md)**

