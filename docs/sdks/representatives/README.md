# Representatives

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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Create(ctx, operations.CreateRepresentativeRequest{
        CreateRepresentative: shared.CreateRepresentative{
            Address: &shared.CreateRepresentativeAddress{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            BirthDate: &shared.CreateRepresentativeBirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            Email: moov.String("amanda@classbooker.dev"),
            GovernmentID: &shared.CreateRepresentativeGovernmentID{
                Itin: &shared.CreateRepresentativeGovernmentIDItin{
                    Full: moov.String("123-45-6789"),
                    LastFour: moov.String("6789"),
                },
                Ssn: &shared.CreateRepresentativeGovernmentIDSsn{
                    Full: moov.String("123-45-6789"),
                    LastFour: moov.String("6789"),
                },
            },
            Name: &shared.CreateRepresentativeName{
                FirstName: "Amanda",
                LastName: "Yang",
                MiddleName: moov.String("Amanda"),
                Suffix: moov.String("Jr"),
            },
            Phone: &shared.CreateRepresentativePhone{
                CountryCode: moov.String("1"),
                Number: moov.String("8185551212"),
            },
            Responsibilities: &shared.CreateRepresentativeResponsibilities{
                IsController: false,
                IsOwner: true,
                JobTitle: "CEO",
                OwnershipPercentage: 38,
            },
        },
        AccountID: "de047717-78ff-461d-8174-76360a15db6a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.CreateRepresentativeRequest](../../models/operations/createrepresentativerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Delete(ctx, operations.DeleteRepresentativeRequest{
        AccountID: "660659a1-adea-4ab5-851d-6c645b08b618",
        RepresentativeID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.DeleteRepresentativeRequest](../../models/operations/deleterepresentativerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Get(ctx, operations.GetRepresentativeRequest{
        AccountID: "91baa0fe-1ade-4008-a6f8-c5f350d8cdb5",
        RepresentativeID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetRepresentativeRequest](../../models/operations/getrepresentativerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.List(ctx, operations.ListRepresentativesRequest{
        AccountID: "a3418143-0104-4218-93d5-208ece7e253b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Representatives != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.ListRepresentativesRequest](../../models/operations/listrepresentativesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Update(ctx, operations.PatchRepresentativeRequest{
        PatchRepresentativeRequest: shared.PatchRepresentativeRequest{
            Address: &shared.PatchRepresentativeRequestAddress{
                AddressLine1: moov.String("123 Main Street"),
                AddressLine2: moov.String("Apt 302"),
                City: moov.String("Boulder"),
                Country: moov.String("US"),
                PostalCode: moov.String("80301"),
                StateOrProvince: moov.String("CO"),
            },
            BirthDate: &shared.PatchRepresentativeRequestBirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            Email: moov.String("amanda@classbooker.dev"),
            GovernmentID: &shared.PatchRepresentativeRequestGovernmentID{
                Itin: &shared.PatchRepresentativeRequestGovernmentIDItin{
                    Full: moov.String("123-45-6789"),
                    LastFour: moov.String("6789"),
                },
                Ssn: &shared.PatchRepresentativeRequestGovernmentIDSsn{
                    Full: moov.String("123-45-6789"),
                    LastFour: moov.String("6789"),
                },
            },
            Name: &shared.PatchRepresentativeRequestName{
                FirstName: moov.String("Amanda"),
                LastName: moov.String("Yang"),
                MiddleName: moov.String("Amanda"),
                Suffix: moov.String("Jr"),
            },
            Phone: &shared.PatchRepresentativeRequestPhone{
                CountryCode: moov.String("1"),
                Number: moov.String("8185551212"),
            },
            Responsibilities: &shared.PatchRepresentativeRequestResponsibilities{
                IsController: moov.Bool(false),
                IsOwner: moov.Bool(true),
                JobTitle: moov.String("CEO"),
                OwnershipPercentage: moov.Int64(38),
            },
        },
        AccountID: "668451c6-c6e2-405e-96de-ab3fec9578a6",
        RepresentativeID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Representative != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.PatchRepresentativeRequest](../../models/operations/patchrepresentativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.PatchRepresentativeResponse](../../models/operations/patchrepresentativeresponse.md), error**

