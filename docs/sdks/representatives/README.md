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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Create(ctx, operations.CreateRepresentativeRequest{
        CreateRepresentative: shared.CreateRepresentative{
            Address: &shared.CreateRepresentativeAddress{
                AddressLine1: petstore.String("123 Main Street"),
                AddressLine2: petstore.String("Apt 302"),
                City: petstore.String("Boulder"),
                Country: petstore.String("US"),
                PostalCode: petstore.String("80301"),
                StateOrProvince: petstore.String("CO"),
            },
            BirthDate: &shared.CreateRepresentativeBirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            Email: petstore.String("amanda@classbooker.dev"),
            GovernmentID: &shared.CreateRepresentativeGovernmentID{
                Itin: &shared.CreateRepresentativeGovernmentIDItin{
                    Full: petstore.String("123-45-6789"),
                    LastFour: petstore.String("6789"),
                },
                Ssn: &shared.CreateRepresentativeGovernmentIDSsn{
                    Full: petstore.String("123-45-6789"),
                    LastFour: petstore.String("6789"),
                },
            },
            Name: &shared.CreateRepresentativeName{
                FirstName: "Amanda",
                LastName: "Yang",
                MiddleName: petstore.String("Amanda"),
                Suffix: petstore.String("Jr"),
            },
            Phone: &shared.CreateRepresentativePhone{
                CountryCode: petstore.String("1"),
                Number: petstore.String("8185551212"),
            },
            Responsibilities: &shared.CreateRepresentativeResponsibilities{
                IsController: false,
                IsOwner: true,
                JobTitle: "CEO",
                OwnershipPercentage: 38,
            },
        },
        AccountID: "512c1032-648d-4c2f-a151-99ebfd0e9fe6",
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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Delete(ctx, operations.DeleteRepresentativeRequest{
        AccountID: "c632ca3a-ed01-4179-9631-2fde04771778",
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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Get(ctx, operations.GetRepresentativeRequest{
        AccountID: "ff61d017-4763-460a-95db-6a660659a1ad",
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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.List(ctx, operations.ListRepresentativesRequest{
        AccountID: "eaab5851-d6c6-445b-88b6-1891baa0fe1a",
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
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Representatives.Update(ctx, operations.PatchRepresentativeRequest{
        PatchRepresentativeRequest: shared.PatchRepresentativeRequest{
            Address: &shared.PatchRepresentativeRequestAddress{
                AddressLine1: petstore.String("123 Main Street"),
                AddressLine2: petstore.String("Apt 302"),
                City: petstore.String("Boulder"),
                Country: petstore.String("US"),
                PostalCode: petstore.String("80301"),
                StateOrProvince: petstore.String("CO"),
            },
            BirthDate: &shared.PatchRepresentativeRequestBirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            Email: petstore.String("amanda@classbooker.dev"),
            GovernmentID: &shared.PatchRepresentativeRequestGovernmentID{
                Itin: &shared.PatchRepresentativeRequestGovernmentIDItin{
                    Full: petstore.String("123-45-6789"),
                    LastFour: petstore.String("6789"),
                },
                Ssn: &shared.PatchRepresentativeRequestGovernmentIDSsn{
                    Full: petstore.String("123-45-6789"),
                    LastFour: petstore.String("6789"),
                },
            },
            Name: &shared.PatchRepresentativeRequestName{
                FirstName: petstore.String("Amanda"),
                LastName: petstore.String("Yang"),
                MiddleName: petstore.String("Amanda"),
                Suffix: petstore.String("Jr"),
            },
            Phone: &shared.PatchRepresentativeRequestPhone{
                CountryCode: petstore.String("1"),
                Number: petstore.String("8185551212"),
            },
            Responsibilities: &shared.PatchRepresentativeRequestResponsibilities{
                IsController: petstore.Bool(false),
                IsOwner: petstore.Bool(true),
                JobTitle: petstore.String("CEO"),
                OwnershipPercentage: petstore.Int64(38),
            },
        },
        AccountID: "de008e6f-8c5f-4350-98cd-b5a341814301",
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

