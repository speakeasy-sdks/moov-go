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