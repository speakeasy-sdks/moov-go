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
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    createRepresentative := shared.CreateRepresentative{
        Address: &shared.CreateRepresentativeAddress{
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        BirthDate: &shared.CreateRepresentativeBirthDate{
            Day: 9,
            Month: 11,
            Year: 1989,
        },
        Email: moovgo.String("amanda@classbooker.dev"),
        GovernmentID: &shared.CreateRepresentativeGovernmentID{
            Itin: &shared.CreateRepresentativeGovernmentIDItin{
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
            Ssn: &shared.CreateRepresentativeGovernmentIDSsn{
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
        },
        Name: &shared.CreateRepresentativeName{
            FirstName: "Amanda",
            LastName: "Yang",
            MiddleName: moovgo.String("Amanda"),
            Suffix: moovgo.String("Jr"),
        },
        Phone: &shared.CreateRepresentativePhone{
            CountryCode: moovgo.String("1"),
            Number: moovgo.String("8185551212"),
        },
        Responsibilities: &shared.CreateRepresentativeResponsibilities{
            IsController: moovgo.Bool(false),
            IsOwner: moovgo.Bool(true),
            JobTitle: "CEO",
            OwnershipPercentage: 38,
        },
    }
    accountID := "b5851d6c-645b-408b-a189-1baa0fe1ade0"

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "08e6f8c5-f350-4d8c-9b5a-341814301042"
    representativeID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "1813d520-8ece-47e2-93b6-68451c6c6e20"
    representativeID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "5e16deab-3fec-4957-8a64-584273a8418d"

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    patchRepresentativeRequest := shared.PatchRepresentativeRequest{
        Address: &shared.PatchRepresentativeRequestAddress{
            AddressLine1: moovgo.String("123 Main Street"),
            AddressLine2: moovgo.String("Apt 302"),
            City: moovgo.String("Boulder"),
            Country: moovgo.String("US"),
            PostalCode: moovgo.String("80301"),
            StateOrProvince: moovgo.String("CO"),
        },
        BirthDate: &shared.PatchRepresentativeRequestBirthDate{
            Day: 9,
            Month: 11,
            Year: 1989,
        },
        Email: moovgo.String("amanda@classbooker.dev"),
        GovernmentID: &shared.PatchRepresentativeRequestGovernmentID{
            Itin: &shared.PatchRepresentativeRequestGovernmentIDItin{
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
            Ssn: &shared.PatchRepresentativeRequestGovernmentIDSsn{
                Full: moovgo.String("123-45-6789"),
                LastFour: moovgo.String("6789"),
            },
        },
        Name: &shared.PatchRepresentativeRequestName{
            FirstName: moovgo.String("Amanda"),
            LastName: moovgo.String("Yang"),
            MiddleName: moovgo.String("Amanda"),
            Suffix: moovgo.String("Jr"),
        },
        Phone: &shared.PatchRepresentativeRequestPhone{
            CountryCode: moovgo.String("1"),
            Number: moovgo.String("8185551212"),
        },
        Responsibilities: &shared.PatchRepresentativeRequestResponsibilities{
            IsController: moovgo.Bool(false),
            IsOwner: moovgo.Bool(true),
            JobTitle: moovgo.String("CEO"),
            OwnershipPercentage: moovgo.Int64(38),
        },
    }
    accountID := "162309fb-0929-4921-aefb-9f58c4d86e68"
    representativeID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

