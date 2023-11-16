# Accounts
(*Accounts*)

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
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

### Parameters

| Parameter                                                   | Type                                                        | Required                                                    | Description                                                 |
| ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- | ----------------------------------------------------------- |
| `ctx`                                                       | [context.Context](https://pkg.go.dev/context#Context)       | :heavy_check_mark:                                          | The context to use for the request.                         |
| `countries`                                                 | [shared.Countries](../../../pkg/models/shared/countries.md) | :heavy_check_mark:                                          | N/A                                                         |
| `accountID`                                                 | *string*                                                    | :heavy_check_mark:                                          | ID of the account                                           |


### Response

**[*operations.PutAccountCountriesResponse](../../pkg/models/operations/putaccountcountriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Create

You can create accounts for your users by passing the required information to Moov. <br><br> Note that `mode` field is only required when creating a facilitator account. All non-facilitator account creation requests will ignore the mode field provided and be set to the calling facilitator's mode. <br><br> If you are creating an account with the business type "llc", "partnership", or "privateCorporation", you will need to also provide [business representatives](https://docs.moov.io/api/#tag/Representatives) after creating the account for verification purposes. Once you've added your business owners as representatives, you'll then need to [patch your Moov account](https://docs.moov.io/api/#operation/patchAccount) to indicate that ownership information is complete. Read more on our [business verification requirements here](https://docs.moov.io/guides/accounts/business-verification/). <br><br> When creating an account, you will need to specify the `/accounts.write` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.Create(ctx, shared.CreateAccountRequest{
        AccountType: shared.AccountTypeBusiness,
        Capabilities: []shared.CapabilityID{
            shared.CapabilityIDCollectFunds,
        },
        CustomerSupport: &shared.CreateAccountRequestCustomerSupport{
            Address: &shared.CreateAccountRequestAddress{
                AddressLine1: moovgo.String("123 Main Street"),
                AddressLine2: moovgo.String("Apt 302"),
                City: moovgo.String("Boulder"),
                Country: moovgo.String("US"),
                PostalCode: moovgo.String("80301"),
                StateOrProvince: moovgo.String("CO"),
            },
            Email: moovgo.String("amanda@classbooker.dev"),
            Phone: &shared.CreateAccountRequestPhone{
                CountryCode: moovgo.String("1"),
                Number: moovgo.String("8185551212"),
            },
            Website: moovgo.String("www.wholebodyfitnessgym.com"),
        },
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Metadata: map[string]string{
            "key": "string",
        },
        Mode: shared.ModeProduction.ToPointer(),
        Profile: shared.CreateProfile{
            Business: &shared.Business{
                Address: &shared.CreateProfileAddress{
                    AddressLine1: moovgo.String("123 Main Street"),
                    AddressLine2: moovgo.String("Apt 302"),
                    City: moovgo.String("Boulder"),
                    Country: moovgo.String("US"),
                    PostalCode: moovgo.String("80301"),
                    StateOrProvince: moovgo.String("CO"),
                },
                BusinessType: shared.BusinessTypeLlc,
                Description: moovgo.String("Local fitness center paying out instructors"),
                DoingBusinessAs: moovgo.String("Whole Body Fitness"),
                Email: moovgo.String("amanda@classbooker.dev"),
                IndustryCodes: &shared.IndustryCodes{
                    Mcc: moovgo.String("7997"),
                    Naics: moovgo.String("713940"),
                    Sic: moovgo.String("7991"),
                },
                LegalBusinessName: "Whole Body Fitness LLC",
                Phone: &shared.CreateProfilePhone{
                    CountryCode: moovgo.String("1"),
                    Number: moovgo.String("8185551212"),
                },
                TaxID: &shared.TaxID{
                    Ein: &shared.Ein{
                        Number: moovgo.String("123-45-6789"),
                    },
                },
                Website: moovgo.String("www.wholebodyfitnessgym.com"),
            },
            Individual: &shared.Individual{
                Address: &shared.CreateProfileSchemasAddress{
                    AddressLine1: moovgo.String("123 Main Street"),
                    AddressLine2: moovgo.String("Apt 302"),
                    City: moovgo.String("Boulder"),
                    Country: moovgo.String("US"),
                    PostalCode: moovgo.String("80301"),
                    StateOrProvince: moovgo.String("CO"),
                },
                BirthDate: &shared.CreateProfileBirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                Email: moovgo.String("amanda@classbooker.dev"),
                GovernmentID: &shared.CreateProfileGovernmentID{
                    Itin: &shared.CreateProfileItin{
                        Full: moovgo.String("123-45-6789"),
                        LastFour: moovgo.String("6789"),
                    },
                    Ssn: &shared.CreateProfileSsn{
                        Full: moovgo.String("123-45-6789"),
                        LastFour: moovgo.String("6789"),
                    },
                },
                Name: shared.Name{
                    FirstName: moovgo.String("Amanda"),
                    LastName: moovgo.String("Yang"),
                    MiddleName: moovgo.String("Amanda"),
                    Suffix: moovgo.String("Jr"),
                },
                Phone: &shared.CreateProfileSchemasPhone{
                    CountryCode: moovgo.String("1"),
                    Number: moovgo.String("8185551212"),
                },
            },
        },
        Settings: &shared.CreateAccountRequestSettings{
            AchPayment: &shared.CreateAccountRequestAchPayment{
                CompanyName: moovgo.String("Whole Body Fitness"),
            },
            CardPayment: &shared.CreateAccountRequestCardPayment{
                StatementDescriptor: moovgo.String("Whole Body Fitness"),
            },
        },
        TermsOfService: &shared.CreateAccountRequestTermsOfService{
            Token: moovgo.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
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

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [shared.CreateAccountRequest](../../pkg/models/shared/createaccountrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.CreateAccountResponse](../../pkg/models/operations/createaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## Get

Retrieves details for the account with the specified ID. <br><br> To get an account, you will need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

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

**[*operations.GetAccountResponse](../../pkg/models/operations/getaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetTosToken

Generates a non-expiring token that can then be used to accept Moov's terms of service. This token can only be generated via API. Any Moov account requesting the `collect-funds`, `send-funds`, `wallet`, or `card-issuing` capabilities must accept Moov's terms of service, then have the generated terms of service token patched to the account. Read more on our [quick start guide](https://docs.moov.io/guides/quick-start/#request-capabilities).

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
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

**[*operations.GetTermsOfServiceTokenResponse](../../pkg/models/operations/gettermsofservicetokenresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Accounts.List(ctx, operations.ListAccountsRequest{
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Type: shared.AccountTypeBusiness.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListAccountsRequest](../../pkg/models/operations/listaccountsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListAccountsResponse](../../pkg/models/operations/listaccountsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListCountries

Retrieve the specified countries of operation for an account. <br><br> To get the list of countries, you'll need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "5694ddc2-a16c-425b-bd39-0a53b9fefa9b"

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

**[*operations.GetAccountCountriesResponse](../../pkg/models/operations/getaccountcountriesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    patchAccountRequest := shared.PatchAccountRequest{
        CustomerSupport: &shared.PatchAccountRequestCustomerSupport{
            Address: &shared.PatchAccountRequestAddress{
                AddressLine1: moovgo.String("123 Main Street"),
                AddressLine2: moovgo.String("Apt 302"),
                City: moovgo.String("Boulder"),
                Country: moovgo.String("US"),
                PostalCode: moovgo.String("80301"),
                StateOrProvince: moovgo.String("CO"),
            },
            Email: moovgo.String("amanda@classbooker.dev"),
            Phone: &shared.PatchAccountRequestPhone{
                CountryCode: moovgo.String("1"),
                Number: moovgo.String("8185551212"),
            },
            Website: moovgo.String("www.wholebodyfitnessgym.com"),
        },
        ForeignID: moovgo.String("4528aba-b9a1-11eb-8529-0242ac13003"),
        Metadata: map[string]string{
            "key": "string",
        },
        Profile: &shared.PatchAccountRequestProfile{
            Business: &shared.PatchAccountRequestBusiness{
                Address: &shared.PatchAccountRequestSchemasAddress{
                    AddressLine1: moovgo.String("123 Main Street"),
                    AddressLine2: moovgo.String("Apt 302"),
                    City: moovgo.String("Boulder"),
                    Country: moovgo.String("US"),
                    PostalCode: moovgo.String("80301"),
                    StateOrProvince: moovgo.String("CO"),
                },
                BusinessType: shared.PatchAccountRequestBusinessTypeLlc.ToPointer(),
                Description: moovgo.String("Local fitness center paying out instructors"),
                DoingBusinessAs: moovgo.String("Whole Body Fitness"),
                Email: moovgo.String("amanda@classbooker.dev"),
                IndustryCodes: &shared.PatchAccountRequestIndustryCodes{
                    Mcc: moovgo.String("7997"),
                    Naics: moovgo.String("713940"),
                    Sic: moovgo.String("7991"),
                },
                LegalBusinessName: moovgo.String("Whole Body Fitness LLC"),
                Phone: &shared.PatchAccountRequestSchemasPhone{
                    CountryCode: moovgo.String("1"),
                    Number: moovgo.String("8185551212"),
                },
                TaxID: &shared.PatchAccountRequestTaxID{
                    Ein: &shared.Ein{
                        Number: moovgo.String("123-45-6789"),
                    },
                },
                Website: moovgo.String("www.wholebodyfitnessgym.com"),
            },
            Individual: &shared.PatchAccountRequestIndividual{
                Address: &shared.PatchAccountRequestSchemasProfileAddress{
                    AddressLine1: moovgo.String("123 Main Street"),
                    AddressLine2: moovgo.String("Apt 302"),
                    City: moovgo.String("Boulder"),
                    Country: moovgo.String("US"),
                    PostalCode: moovgo.String("80301"),
                    StateOrProvince: moovgo.String("CO"),
                },
                BirthDate: &shared.PatchAccountRequestBirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                Email: moovgo.String("amanda@classbooker.dev"),
                GovernmentID: &shared.PatchAccountRequestGovernmentID{
                    Itin: &shared.PatchAccountRequestItin{
                        Full: moovgo.String("123-45-6789"),
                        LastFour: moovgo.String("6789"),
                    },
                    Ssn: &shared.PatchAccountRequestSsn{
                        Full: moovgo.String("123-45-6789"),
                        LastFour: moovgo.String("6789"),
                    },
                },
                Name: &shared.PatchAccountRequestName{
                    FirstName: moovgo.String("Amanda"),
                    LastName: moovgo.String("Yang"),
                    MiddleName: moovgo.String("Amanda"),
                    Suffix: moovgo.String("Jr"),
                },
                Phone: &shared.PatchAccountRequestSchemasProfilePhone{
                    CountryCode: moovgo.String("1"),
                    Number: moovgo.String("8185551212"),
                },
            },
        },
        Settings: &shared.PatchAccountRequestSettings{
            AchPayment: &shared.PatchAccountRequestAchPayment{
                CompanyName: moovgo.String("Whole Body Fitness"),
            },
            CardPayment: &shared.PatchAccountRequestCardPayment{
                StatementDescriptor: moovgo.String("Whole Body Fitness"),
            },
        },
        TermsOfService: &shared.PatchAccountRequestTermsOfService{
            Token: moovgo.String("kgT1uxoMAk7QKuyJcmQE8nqW_HjpyuXBabiXPi6T83fUQoxsyWYPcYzuHQTqrt7YRp4gCwyDQvb6U5REM9Pgl2EloCe35t-eiMAbUWGo3Kerxme6aqNcKrP_6-v0MTXViOEJ96IBxPFTvMV7EROI2dq3u4e-x4BbGSCedAX-ViAQND6hcreCDXwrO6sHuzh5Xi2IzSqZHxaovnWEboaxuZKRJkA3dsFID6fzitMpm2qrOh4"),
        },
    }

    var accountID string = "d0905bf4-aa77-4f20-8e77-54c352acfe54"

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

| Parameter                                                                       | Type                                                                            | Required                                                                        | Description                                                                     |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `ctx`                                                                           | [context.Context](https://pkg.go.dev/context#Context)                           | :heavy_check_mark:                                                              | The context to use for the request.                                             |
| `patchAccountRequest`                                                           | [shared.PatchAccountRequest](../../../pkg/models/shared/patchaccountrequest.md) | :heavy_check_mark:                                                              | N/A                                                                             |
| `accountID`                                                                     | *string*                                                                        | :heavy_check_mark:                                                              | ID of the account                                                               |


### Response

**[*operations.PatchAccountResponse](../../pkg/models/operations/patchaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
