# Representatives
(*Representatives*)

## Overview

We think of a representative as an individual who can take major actions on behalf of a business. A representative can be the business owner, or anyone who owns 25% or more of the business. Some examples of representatives are the CEO, CFO, COO, or president. We require all business accounts to have valid information for at least one representative. Moov accounts must have verified business representatives before a business account can send funds, collect money from other accounts, or store funds in a wallet. To learn more, read our guide on [business representatives](https://docs.moov.io/guides/accounts/business-representatives/).

### Available Operations

* [Create](#create) - Create representative
* [Delete](#delete) - Delete a representative
* [Get](#get) - Get representative
* [List](#list) - List representatives
* [Update](#update) - Patch representative

## Create

Moov accounts associated with businesses require information regarding individuals who represent the business. You can provide this information by creating a representative. Each account is allowed a maximum of 7 representatives.<br><br> To create a representative, you must specify the `/accounts/{accountID}/representatives.write` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    createRepresentative := shared.CreateRepresentative{
        AdditionalProperties: map[string]interface{}{
            "online": "Configuration",
        },
        Address: &shared.CreateRepresentativeAddress{
            AdditionalProperties: map[string]interface{}{
                "Money": "blue",
            },
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        BirthDate: &shared.CreateRepresentativeBirthDate{
            AdditionalProperties: map[string]interface{}{
                "shred": "abnormally",
            },
            Day: 9,
            Month: 11,
            Year: 1989,
        },
        Email: moovgo.String("amanda@classbooker.dev"),
        GovernmentID: &shared.CreateRepresentativeGovernmentID{
            AdditionalProperties: map[string]interface{}{
                "deposit": "evolve",
            },
            Itin: &shared.CreateRepresentativeGovernmentIDItin{
                AdditionalProperties: map[string]interface{}{
                    "male": "SUV",
                },
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
            Ssn: &shared.CreateRepresentativeGovernmentIDSsn{
                AdditionalProperties: map[string]interface{}{
                    "quantify": "Polestar",
                },
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
        },
        Name: &shared.CreateRepresentativeName{
            AdditionalProperties: map[string]interface{}{
                "mobile": "National",
            },
            FirstName: "Amanda",
            LastName: "Yang",
            MiddleName: moovgo.String("Amanda"),
            Suffix: moovgo.String("Jr"),
        },
        Phone: &shared.CreateRepresentativePhone{
            AdditionalProperties: map[string]interface{}{
                "Durham": "after",
            },
            CountryCode: moovgo.String("1"),
            Number: moovgo.String("8185551212"),
        },
        Responsibilities: &shared.CreateRepresentativeResponsibilities{
            AdditionalProperties: map[string]interface{}{
                "overriding": "Bike",
            },
            IsController: moovgo.Bool(false),
            IsOwner: moovgo.Bool(true),
            JobTitle: "CEO",
            OwnershipPercentage: 38,
        },
    }
    var accountID string = "b1d5e261-915a-425d-8d9e-a1320e8504aa"

    ctx := context.Background()
    res, err := s.Representatives.Create(ctx, createRepresentative, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `createRepresentative`                                                     | [shared.CreateRepresentative](../../models/shared/createrepresentative.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `accountID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | ID of the account                                                          |


### Response

**[*operations.CreateRepresentativeResponse](../../models/operations/createrepresentativeresponse.md), error**


## Delete

Deletes a business representative associated with a Moov account. <br><br> To use this endpoint, you'll need to specify the `/accounts/{accountID}/representatives.write` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    var accountID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"
    var representativeID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Representatives.Delete(ctx, accountID, representativeID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `representativeID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the representative                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.DeleteRepresentativeResponse](../../models/operations/deleterepresentativeresponse.md), error**


## Get

Retrieve a specific representative associated with a given Moov account. <br><br> To get a representative, you'll need to specify the `/accounts/{accountID}/representatives.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"
    var representativeID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Representatives.Get(ctx, accountID, representativeID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `representativeID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the representative                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.GetRepresentativeResponse](../../models/operations/getrepresentativeresponse.md), error**


## List

A Moov account may have multiple representatives depending on the associated business's ownership and management structure. You can use this method to list all the representatives for a given Moov account. Note that Moov accounts associated with an individual do not have representatives. <br><br> To list representatives, you need to specify the `/accounts/{accountID}/representatives.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    var accountID string = "c184a429-302e-4aca-80db-f1718b882a50"

    ctx := context.Background()
    res, err := s.Representatives.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Representatives != nil {
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

**[*operations.ListRepresentativesResponse](../../models/operations/listrepresentativesresponse.md), error**


## Update

If a representative's information has changed you can patch the information associated with a specific representative ID.  
To patch a representative, you'll need to specify the `/accounts/{accountID}/representatives.write` scope and provide the changed information.

When **can** profile data be updated:  
  + For unverified representatives, all profile data can be edited.
  + During the verification process, missing or incomplete profile data can be edited.
  + Verified representatives can only add missing profile data.  

  When **can't** profile data be updated:  
  + Verified representatives cannot change any existing profile data.  

If you need to update information in a locked state, please contact Moov support.


### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    patchRepresentativeRequest := shared.PatchRepresentativeRequest{
        AdditionalProperties: map[string]interface{}{
            "Van": "East",
        },
        Address: &shared.PatchRepresentativeRequestAddress{
            AdditionalProperties: map[string]interface{}{
                "male": "Metal",
            },
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        BirthDate: &shared.PatchRepresentativeRequestBirthDate{
            AdditionalProperties: map[string]interface{}{
                "cheater": "Islands",
            },
            Day: 9,
            Month: 11,
            Year: 1989,
        },
        Email: moovgo.String("amanda@classbooker.dev"),
        GovernmentID: &shared.PatchRepresentativeRequestGovernmentID{
            AdditionalProperties: map[string]interface{}{
                "online": "dynamic",
            },
            Itin: &shared.PatchRepresentativeRequestGovernmentIDItin{
                AdditionalProperties: map[string]interface{}{
                    "white": "bifurcated",
                },
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
            Ssn: &shared.PatchRepresentativeRequestGovernmentIDSsn{
                AdditionalProperties: map[string]interface{}{
                    "Forward": "syndicate",
                },
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
        },
        Name: &shared.PatchRepresentativeRequestName{
            AdditionalProperties: map[string]interface{}{
                "East": "Baht",
            },
            FirstName: moovgo.String("Amanda"),
            LastName: moovgo.String("Yang"),
            MiddleName: moovgo.String("Amanda"),
            Suffix: moovgo.String("Jr"),
        },
        Phone: &shared.PatchRepresentativeRequestPhone{
            AdditionalProperties: map[string]interface{}{
                "Quality": "guestbook",
            },
            CountryCode: moovgo.String("1"),
            Number: moovgo.String("8185551212"),
        },
        Responsibilities: &shared.PatchRepresentativeRequestResponsibilities{
            AdditionalProperties: map[string]interface{}{
                "driver": "users",
            },
            IsController: moovgo.Bool(false),
            IsOwner: moovgo.Bool(true),
            JobTitle: moovgo.String("CEO"),
            OwnershipPercentage: moovgo.Int64(38),
        },
    }
    var accountID string = "5ca71871-4355-4ad7-94e1-b584578f9d86"
    var representativeID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Representatives.Update(ctx, patchRepresentativeRequest, accountID, representativeID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            | Example                                                                                |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |                                                                                        |
| `patchRepresentativeRequest`                                                           | [shared.PatchRepresentativeRequest](../../models/shared/patchrepresentativerequest.md) | :heavy_check_mark:                                                                     | N/A                                                                                    |                                                                                        |
| `accountID`                                                                            | *string*                                                                               | :heavy_check_mark:                                                                     | ID of the account                                                                      |                                                                                        |
| `representativeID`                                                                     | *string*                                                                               | :heavy_check_mark:                                                                     | ID of the representative                                                               | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                   |


### Response

**[*operations.PatchRepresentativeResponse](../../models/operations/patchrepresentativeresponse.md), error**

