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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    countries := shared.Countries{
        Countries: []string{
            "United States",
        },
    }
    accountID := "92059293-96fe-4a75-96eb-10faaa2352c5"

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

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `countries`                                           | [shared.Countries](../../models/shared/countries.md)  | :heavy_check_mark:                                    | N/A                                                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, shared.CreateAccountRequest{
        AccountType: shared.AccountTypeBusiness,
        Capabilities: []shared.CapabilityID{
            shared.CapabilityIDSendFunds,
            shared.CapabilityIDSendFunds,
            shared.CapabilityIDCollectFunds,
        },
        CustomerSupport: &shared.CreateAccountRequestCustomerSupport{
            Address: &shared.CreateAccountRequestCustomerSupportAddress{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            Email: moov.String("amanda@classbooker.dev"),
            Phone: &shared.CreateAccountRequestCustomerSupportPhone{
                CountryCode: moov.String("1"),
                Number: moov.String("8185551212"),
            },
            Website: moov.String("www.wholebodyfitnessgym.com"),
        },
        ForeignID: moov.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Metadata: map[string]string{
            "iure": "culpa",
        },
        Mode: shared.ModeProduction.ToPointer(),
        Profile: shared.CreateProfile{
            Business: &shared.CreateProfileBusiness{
                Address: &shared.CreateProfileBusinessAddress{
                    AddressLine1: moov.String("123 Main Street"),
                    AddressLine2: moov.String("Apt 302"),
                    City: moov.String("Boulder"),
                    Country: moov.String("US"),
                    PostalCode: moov.String("80301"),
                    StateOrProvince: moov.String("CO"),
                },
                BusinessType: shared.BusinessTypeLlc,
                Description: moov.String("Local fitness center paying out instructors"),
                DoingBusinessAs: moov.String("Whole Body Fitness"),
                Email: moov.String("amanda@classbooker.dev"),
                IndustryCodes: &shared.CreateProfileBusinessIndustryCodes{
                    Mcc: moov.String("7997"),
                    Naics: moov.String("713940"),
                    Sic: moov.String("7991"),
                },
                LegalBusinessName: "Whole Body Fitness LLC",
                Phone: &shared.CreateProfileBusinessPhone{
                    CountryCode: moov.String("1"),
                    Number: moov.String("8185551212"),
                },
                TaxID: &shared.CreateProfileBusinessTaxID{
                    Ein: &shared.Ein{
                        Number: moov.String("123-45-6789"),
                    },
                },
                Website: moov.String("www.wholebodyfitnessgym.com"),
            },
            Individual: &shared.CreateProfileIndividual{
                Address: &shared.CreateProfileIndividualAddress{
                    AddressLine1: moov.String("123 Main Street"),
                    AddressLine2: moov.String("Apt 302"),
                    City: moov.String("Boulder"),
                    Country: moov.String("US"),
                    PostalCode: moov.String("80301"),
                    StateOrProvince: moov.String("CO"),
                },
                BirthDate: &shared.CreateProfileIndividualBirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                Email: moov.String("amanda@classbooker.dev"),
                GovernmentID: &shared.CreateProfileIndividualGovernmentID{
                    Itin: &shared.CreateProfileIndividualGovernmentIDItin{
                        Full: moov.String("123-45-6789"),
                        LastFour: moov.String("6789"),
                    },
                    Ssn: &shared.CreateProfileIndividualGovernmentIDSsn{
                        Full: moov.String("123-45-6789"),
                        LastFour: moov.String("6789"),
                    },
                },
                Name: shared.Name{
                    FirstName: moov.String("Amanda"),
                    LastName: moov.String("Yang"),
                    MiddleName: moov.String("Amanda"),
                    Suffix: moov.String("Jr"),
                },
                Phone: &shared.CreateProfileIndividualPhone{
                    CountryCode: moov.String("1"),
                    Number: moov.String("8185551212"),
                },
            },
        },
        Settings: &shared.CreateAccountRequestSettings{
            AchPayment: &shared.CreateAccountRequestSettingsAchPayment{
                CompanyName: moov.String("Whole Body Fitness"),
            },
            CardPayment: &shared.CreateAccountRequestSettingsCardPayment{
                StatementDescriptor: moov.String("Whole Body Fitness"),
            },
        },
        TermsOfService: &shared.CreateAccountRequestTermsOfService{
            Token: moov.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "ff1a3a2f-a946-4773-9251-aa52c3f5ad01"

    ctx := context.Background()
    res, err := s.Accounts.Get(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.List(ctx, operations.ListAccountsRequest{
        Count: moov.Int64(622846),
        Email: moov.String("Margie_Boyer87@hotmail.com"),
        ForeignID: moov.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        IncludeDisconnected: moov.Bool(false),
        Name: moov.String("Miss Irma Wolff"),
        Skip: moov.Int64(739264),
        Type: shared.AccountTypeBusiness.ToPointer(),
        VerificationStatus: shared.AccountVerificationStatusUnverified.ToPointer(),
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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "074f1547-1b5e-46e1-bb99-d488e1e91e45"

    ctx := context.Background()
    res, err := s.Accounts.ListCountries(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Countries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    patchAccountRequest := shared.PatchAccountRequest{
        CustomerSupport: &shared.PatchAccountRequestCustomerSupport{
            Address: &shared.PatchAccountRequestCustomerSupportAddress{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            Email: moov.String("amanda@classbooker.dev"),
            Phone: &shared.PatchAccountRequestCustomerSupportPhone{
                CountryCode: moov.String("1"),
                Number: moov.String("8185551212"),
            },
            Website: moov.String("www.wholebodyfitnessgym.com"),
        },
        ForeignID: moov.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Metadata: map[string]string{
            "est": "quibusdam",
        },
        Profile: &shared.PatchAccountRequestProfile{
            Business: &shared.PatchAccountRequestProfileBusiness{
                Address: &shared.PatchAccountRequestProfileBusinessAddress{
                    AddressLine1: moov.String("123 Main Street"),
                    AddressLine2: moov.String("Apt 302"),
                    City: moov.String("Boulder"),
                    Country: moov.String("US"),
                    PostalCode: moov.String("80301"),
                    StateOrProvince: moov.String("CO"),
                },
                BusinessType: shared.PatchAccountRequestProfileBusinessBusinessTypeLlc.ToPointer(),
                Description: moov.String("Local fitness center paying out instructors"),
                DoingBusinessAs: moov.String("Whole Body Fitness"),
                Email: moov.String("amanda@classbooker.dev"),
                IndustryCodes: &shared.PatchAccountRequestProfileBusinessIndustryCodes{
                    Mcc: moov.String("7997"),
                    Naics: moov.String("713940"),
                    Sic: moov.String("7991"),
                },
                LegalBusinessName: moov.String("Whole Body Fitness LLC"),
                OwnersProvided: moov.Bool(false),
                Phone: &shared.PatchAccountRequestProfileBusinessPhone{
                    CountryCode: moov.String("1"),
                    Number: moov.String("8185551212"),
                },
                TaxID: &shared.PatchAccountRequestProfileBusinessTaxID{
                    Ein: &shared.Ein{
                        Number: moov.String("123-45-6789"),
                    },
                },
                Website: moov.String("www.wholebodyfitnessgym.com"),
            },
            Individual: &shared.PatchAccountRequestProfileIndividual{
                Address: &shared.PatchAccountRequestProfileIndividualAddress{
                    AddressLine1: moov.String("123 Main Street"),
                    AddressLine2: moov.String("Apt 302"),
                    City: moov.String("Boulder"),
                    Country: moov.String("US"),
                    PostalCode: moov.String("80301"),
                    StateOrProvince: moov.String("CO"),
                },
                BirthDate: &shared.PatchAccountRequestProfileIndividualBirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                Email: moov.String("amanda@classbooker.dev"),
                GovernmentID: &shared.PatchAccountRequestProfileIndividualGovernmentID{
                    Itin: &shared.PatchAccountRequestProfileIndividualGovernmentIDItin{
                        Full: moov.String("123-45-6789"),
                        LastFour: moov.String("6789"),
                    },
                    Ssn: &shared.PatchAccountRequestProfileIndividualGovernmentIDSsn{
                        Full: moov.String("123-45-6789"),
                        LastFour: moov.String("6789"),
                    },
                },
                Name: &shared.PatchAccountRequestProfileIndividualName{
                    FirstName: moov.String("Amanda"),
                    LastName: moov.String("Yang"),
                    MiddleName: moov.String("Amanda"),
                    Suffix: moov.String("Jr"),
                },
                Phone: &shared.PatchAccountRequestProfileIndividualPhone{
                    CountryCode: moov.String("1"),
                    Number: moov.String("8185551212"),
                },
            },
        },
        Settings: &shared.PatchAccountRequestSettings{
            AchPayment: &shared.PatchAccountRequestSettingsAchPayment{
                CompanyName: moov.String("Whole Body Fitness"),
            },
            CardPayment: &shared.PatchAccountRequestSettingsCardPayment{
                StatementDescriptor: moov.String("Whole Body Fitness"),
            },
        },
        TermsOfService: &shared.PatchAccountRequestTermsOfService{
            Token: moov.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
        },
    }
    accountID := "2abd4426-9802-4d50-aa94-bb4f63c969e9"

    ctx := context.Background()
    res, err := s.Accounts.Update(ctx, patchAccountRequest, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Account != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `patchAccountRequest`                                                    | [shared.PatchAccountRequest](../../models/shared/patchaccountrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `accountID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | ID of the account                                                        |


### Response

**[*operations.PatchAccountResponse](../../models/operations/patchaccountresponse.md), error**

