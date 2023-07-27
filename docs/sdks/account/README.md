# account

### Available Operations

* [create](#create) - Create account
* [get](#get) - Get account
* [patch](#patch) - Patch account

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
        shared.CapabilityID.COLLECT_FUNDS,
        shared.CapabilityID.SEND_FUNDS,
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
        "ipsa": 'delectus',
        "tempora": 'suscipit',
        "molestiae": 'minus',
        "placeat": 'voluptatum',
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
    account_id='796ed151-a05d-4fc2-9df7-cc78ca1ba928',
)

res = s.account.get(req)

if res.account is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetAccountRequest](../../models/operations/getaccountrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetAccountResponse](../../models/operations/getaccountresponse.md)**


## patch

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
            "optio": 'totam',
            "beatae": 'commodi',
            "molestiae": 'modi',
            "qui": 'impedit',
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
    account_id='b7392059-2939-46fe-a759-6eb10faaa235',
)

res = s.account.patch(req)

if res.account is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.PatchAccountRequest](../../models/operations/patchaccountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.PatchAccountResponse](../../models/operations/patchaccountresponse.md)**

