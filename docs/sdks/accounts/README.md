# Accounts

## Overview

[Accounts](https://docs.moov.io/guides/accounts/) represent a legal entity (either a business or an individual) in Moov. You can create an account for yourself or set up accounts for others, requesting different [capabilities](https://docs.moov.io/guides/accounts/capabilities/) depending on what you need to be able to do with that account. You can retrieve an account to get details on the business or individual account holder, such as an email address or employer identification number (EIN). 

Based on the type of account and its requested capabilities, we have certain [verification requirements](https://docs.moov.io/guides/accounts/identity-verification/). To see what capabilities that account has, you can use the [GET capability endpoint](https://docs.moov.io/api/#operation/getCapability). 

When you sign up for the Moov Dashboard, you will have a **facilitator account** which can be used to facilitate money movement between other accounts. A facilitator account will not show up in your list of accounts and cannot be created via API. To update your facilitator account information, use the settings page of the Moov Dashboard. 


### Available Operations

* [AssignCountry](#assigncountry) - Assign Account Countries
* [Create](#create) - Create account
* [Get](#get) - Get account
* [GetTosToken](#gettostoken) - Get terms of service token
* [List](#list) - List accounts
* [ListCountries](#listcountries) - Get Account Countries
* [Update](#update) - Patch account

## AssignCountry

Assign the countries of operation for an account. This endpoint will always overwrite the previously assigned values. <br><br> To update the account countries, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.AssignCountry(ctx, operations.PutAccountCountriesRequest{
        Countries: shared.Countries{
            Countries: []string{
                "United States",
                "United States",
                "United States",
                "United States",
            },
        },
        AccountID: "9d8d69a6-74e0-4f46-bcc8-796ed151a05d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Countries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PutAccountCountriesRequest](../../models/operations/putaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PutAccountCountriesResponse](../../models/operations/putaccountcountriesresponse.md), error**


## Create

You can create accounts for your users by passing the required information to Moov. <br><br> Note that `mode` field is only required when creating a facilitator account. All non-facilitator account creation requests will ignore the mode field provided and be set to the calling facilitator's mode. <br><br> If you are creating an account with the business type "llc", "partnership", or "privateCorporation", you will need to also provide [business representatives](https://docs.moov.io/api/#tag/Representatives) after creating the account for verification purposes. Once you've added your business owners as representatives, you'll then need to [patch your Moov account](https://docs.moov.io/api/#operation/patchAccount) to indicate that ownership information is complete. Read more on our [business verification requirements here](https://docs.moov.io/guides/accounts/business-verification/). <br><br> When creating an account, you will need to specify the `/accounts.write` scope.

### Example Usage

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
    res, err := s.Accounts.Create(ctx, shared.CreateAccountRequest{
        AccountType: shared.AccountTypeBusiness,
        Capabilities: []shared.CapabilityID{
            shared.CapabilityIDWallet,
            shared.CapabilityIDTransfers,
            shared.CapabilityIDCardIssuing,
            shared.CapabilityIDCardIssuing,
        },
        CustomerSupport: &shared.CreateAccountRequestCustomerSupport{
            Address: &shared.CreateAccountRequestCustomerSupportAddress{
                AddressLine1: petstore.String("123 Main Street"),
                AddressLine2: petstore.String("Apt 302"),
                City: petstore.String("Boulder"),
                Country: petstore.String("US"),
                PostalCode: petstore.String("80301"),
                StateOrProvince: petstore.String("CO"),
            },
            Email: petstore.String("amanda@classbooker.dev"),
            Phone: &shared.CreateAccountRequestCustomerSupportPhone{
                CountryCode: petstore.String("1"),
                Number: petstore.String("8185551212"),
            },
            Website: petstore.String("www.wholebodyfitnessgym.com"),
        },
        ForeignID: petstore.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Metadata: map[string]string{
            "molestiae": "quod",
            "quod": "esse",
            "totam": "porro",
            "dolorum": "dicta",
        },
        Mode: shared.ModeProduction.ToPointer(),
        Profile: shared.CreateProfile{
            Business: &shared.CreateProfileBusiness{
                Address: &shared.CreateProfileBusinessAddress{
                    AddressLine1: petstore.String("123 Main Street"),
                    AddressLine2: petstore.String("Apt 302"),
                    City: petstore.String("Boulder"),
                    Country: petstore.String("US"),
                    PostalCode: petstore.String("80301"),
                    StateOrProvince: petstore.String("CO"),
                },
                BusinessType: shared.BusinessTypeLlc,
                Description: petstore.String("Local fitness center paying out instructors"),
                DoingBusinessAs: petstore.String("Whole Body Fitness"),
                Email: petstore.String("amanda@classbooker.dev"),
                IndustryCodes: &shared.CreateProfileBusinessIndustryCodes{
                    Mcc: petstore.String("7997"),
                    Naics: petstore.String("713940"),
                    Sic: petstore.String("7991"),
                },
                LegalBusinessName: "Whole Body Fitness LLC",
                Phone: &shared.CreateProfileBusinessPhone{
                    CountryCode: petstore.String("1"),
                    Number: petstore.String("8185551212"),
                },
                TaxID: &shared.CreateProfileBusinessTaxID{
                    Ein: &shared.Ein{
                        Number: petstore.String("123-45-6789"),
                    },
                },
                Website: petstore.String("www.wholebodyfitnessgym.com"),
            },
            Individual: &shared.CreateProfileIndividual{
                Address: &shared.CreateProfileIndividualAddress{
                    AddressLine1: petstore.String("123 Main Street"),
                    AddressLine2: petstore.String("Apt 302"),
                    City: petstore.String("Boulder"),
                    Country: petstore.String("US"),
                    PostalCode: petstore.String("80301"),
                    StateOrProvince: petstore.String("CO"),
                },
                BirthDate: &shared.CreateProfileIndividualBirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                Email: petstore.String("amanda@classbooker.dev"),
                GovernmentID: &shared.CreateProfileIndividualGovernmentID{
                    Itin: &shared.CreateProfileIndividualGovernmentIDItin{
                        Full: petstore.String("123-45-6789"),
                        LastFour: petstore.String("6789"),
                    },
                    Ssn: &shared.CreateProfileIndividualGovernmentIDSsn{
                        Full: petstore.String("123-45-6789"),
                        LastFour: petstore.String("6789"),
                    },
                },
                Name: shared.Name{
                    FirstName: petstore.String("Amanda"),
                    LastName: petstore.String("Yang"),
                    MiddleName: petstore.String("Amanda"),
                    Suffix: petstore.String("Jr"),
                },
                Phone: &shared.CreateProfileIndividualPhone{
                    CountryCode: petstore.String("1"),
                    Number: petstore.String("8185551212"),
                },
            },
        },
        Settings: &shared.CreateAccountRequestSettings{
            AchPayment: &shared.CreateAccountRequestSettingsAchPayment{
                CompanyName: petstore.String("Whole Body Fitness"),
            },
            CardPayment: &shared.CreateAccountRequestSettingsCardPayment{
                StatementDescriptor: petstore.String("Whole Body Fitness"),
            },
        },
        TermsOfService: &shared.CreateAccountRequestTermsOfService{
            Token: petstore.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.CreateAccountRequest](../../models/shared/createaccountrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateAccountResponse](../../models/operations/createaccountresponse.md), error**


## Get

Retrieves details for the account with the specified ID. <br><br> To get an account, you will need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Get(ctx, operations.GetAccountRequest{
        AccountID: "ba928fc8-1674-42cb-b392-05929396fea7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAccountRequest](../../models/operations/getaccountrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetAccountResponse](../../models/operations/getaccountresponse.md), error**


## GetTosToken

Generates a non-expiring token that can then be used to accept Moov's terms of service. This token can only be generated via API. Any Moov account requesting the `collect-funds`, `send-funds`, `wallet`, or `card-issuing` capabilities must accept Moov's terms of service, then have the generated terms of service token patched to the account. Read more on our [quick start guide](https://docs.moov.io/guides/quick-start/#request-capabilities).

### Example Usage

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
    res, err := s.Accounts.GetTosToken(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.TermsOfServiceToken != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetTermsOfServiceTokenResponse](../../models/operations/gettermsofservicetokenresponse.md), error**


## List

List or search accounts to which the caller is connected.<br><br>
All supported query parameters are optional. If none are provided
the response will include all connected accounts. Pagination is
supported via the `skip` and `count` query parameters.<br><br>
Searching by name and email will overlap and return results based on relevance.
<br><br> To list connected accounts, you must specify the `/accounts.read` scope when retrieving the authentication token.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.List(ctx, operations.ListAccountsRequest{
        Count: petstore.Int64(359508),
        Email: petstore.String("Humberto.Turcotte6@yahoo.com"),
        ForeignID: petstore.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        IncludeDisconnected: petstore.Bool(false),
        Name: petstore.String("Carlton O'Hara"),
        Skip: petstore.Int64(210382),
        Type: shared.AccountTypeBusiness.ToPointer(),
        VerificationStatus: shared.AccountVerificationStatusResubmit.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Accounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListAccountsRequest](../../models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListAccountsResponse](../../models/operations/listaccountsresponse.md), error**


## ListCountries

Retrieve the specified countries of operation for an account. <br><br> To get the list of countries, you'll need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.ListCountries(ctx, operations.GetAccountCountriesRequest{
        AccountID: "2c595590-7aff-41a3-a2fa-9467739251aa",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Countries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetAccountCountriesRequest](../../models/operations/getaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetAccountCountriesResponse](../../models/operations/getaccountcountriesresponse.md), error**


## Update

To patch an account, you must specify the `/accounts/{accountID}/profile.write` scope and provide the changed information.  

When **can** profile data be updated:  
  + For unverified accounts, all profile data can be edited.
  + During the verification process, missing or incomplete profile data can be edited.
  + Verified accounts can only add missing profile data.  

  When **can't** profile data be updated:  
  + Verified accounts cannot change any existing profile data.  

If you need to update information in a locked state, please contact Moov support.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Update(ctx, operations.PatchAccountRequest{
        PatchAccountRequest: shared.PatchAccountRequest{
            CustomerSupport: &shared.PatchAccountRequestCustomerSupport{
                Address: &shared.PatchAccountRequestCustomerSupportAddress{
                    AddressLine1: petstore.String("123 Main Street"),
                    AddressLine2: petstore.String("Apt 302"),
                    City: petstore.String("Boulder"),
                    Country: petstore.String("US"),
                    PostalCode: petstore.String("80301"),
                    StateOrProvince: petstore.String("CO"),
                },
                Email: petstore.String("amanda@classbooker.dev"),
                Phone: &shared.PatchAccountRequestCustomerSupportPhone{
                    CountryCode: petstore.String("1"),
                    Number: petstore.String("8185551212"),
                },
                Website: petstore.String("www.wholebodyfitnessgym.com"),
            },
            ForeignID: petstore.String("4528aba-b9a1-11eb-8529-0242ac13003"),
            Metadata: map[string]string{
                "odit": "quo",
                "sequi": "tenetur",
            },
            Profile: &shared.PatchAccountRequestProfile{
                Business: &shared.PatchAccountRequestProfileBusiness{
                    Address: &shared.PatchAccountRequestProfileBusinessAddress{
                        AddressLine1: petstore.String("123 Main Street"),
                        AddressLine2: petstore.String("Apt 302"),
                        City: petstore.String("Boulder"),
                        Country: petstore.String("US"),
                        PostalCode: petstore.String("80301"),
                        StateOrProvince: petstore.String("CO"),
                    },
                    BusinessType: shared.PatchAccountRequestProfileBusinessBusinessTypeLlc.ToPointer(),
                    Description: petstore.String("Local fitness center paying out instructors"),
                    DoingBusinessAs: petstore.String("Whole Body Fitness"),
                    Email: petstore.String("amanda@classbooker.dev"),
                    IndustryCodes: &shared.PatchAccountRequestProfileBusinessIndustryCodes{
                        Mcc: petstore.String("7997"),
                        Naics: petstore.String("713940"),
                        Sic: petstore.String("7991"),
                    },
                    LegalBusinessName: petstore.String("Whole Body Fitness LLC"),
                    OwnersProvided: petstore.Bool(false),
                    Phone: &shared.PatchAccountRequestProfileBusinessPhone{
                        CountryCode: petstore.String("1"),
                        Number: petstore.String("8185551212"),
                    },
                    TaxID: &shared.PatchAccountRequestProfileBusinessTaxID{
                        Ein: &shared.Ein{
                            Number: petstore.String("123-45-6789"),
                        },
                    },
                    Website: petstore.String("www.wholebodyfitnessgym.com"),
                },
                Individual: &shared.PatchAccountRequestProfileIndividual{
                    Address: &shared.PatchAccountRequestProfileIndividualAddress{
                        AddressLine1: petstore.String("123 Main Street"),
                        AddressLine2: petstore.String("Apt 302"),
                        City: petstore.String("Boulder"),
                        Country: petstore.String("US"),
                        PostalCode: petstore.String("80301"),
                        StateOrProvince: petstore.String("CO"),
                    },
                    BirthDate: &shared.PatchAccountRequestProfileIndividualBirthDate{
                        Day: 9,
                        Month: 11,
                        Year: 1989,
                    },
                    Email: petstore.String("amanda@classbooker.dev"),
                    GovernmentID: &shared.PatchAccountRequestProfileIndividualGovernmentID{
                        Itin: &shared.PatchAccountRequestProfileIndividualGovernmentIDItin{
                            Full: petstore.String("123-45-6789"),
                            LastFour: petstore.String("6789"),
                        },
                        Ssn: &shared.PatchAccountRequestProfileIndividualGovernmentIDSsn{
                            Full: petstore.String("123-45-6789"),
                            LastFour: petstore.String("6789"),
                        },
                    },
                    Name: &shared.PatchAccountRequestProfileIndividualName{
                        FirstName: petstore.String("Amanda"),
                        LastName: petstore.String("Yang"),
                        MiddleName: petstore.String("Amanda"),
                        Suffix: petstore.String("Jr"),
                    },
                    Phone: &shared.PatchAccountRequestProfileIndividualPhone{
                        CountryCode: petstore.String("1"),
                        Number: petstore.String("8185551212"),
                    },
                },
            },
            Settings: &shared.PatchAccountRequestSettings{
                AchPayment: &shared.PatchAccountRequestSettingsAchPayment{
                    CompanyName: petstore.String("Whole Body Fitness"),
                },
                CardPayment: &shared.PatchAccountRequestSettingsCardPayment{
                    StatementDescriptor: petstore.String("Whole Body Fitness"),
                },
            },
            TermsOfService: &shared.PatchAccountRequestTermsOfService{
                Token: petstore.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
            },
        },
        AccountID: "5ad019da-1ffe-478f-897b-0074f15471b5",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.PatchAccountRequest](../../models/operations/patchaccountrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.PatchAccountResponse](../../models/operations/patchaccountresponse.md), error**

