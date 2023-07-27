# CardIssuing

## Overview

A Moov wallet can serve as a funding source for issuing virtual cards. Note that we currently only issue Discover cards. Virtual cards can then be used to spend funds from the wallet.

<em> The `card-issuing` and `wallet` capabilities are required to be enabled before any card issuing functionality is available. Moov is in a private beta with select customers for card issuing.</em>


### Available Operations

* [RequestCard](#requestcard) - Request card
* [GetCard](#getcard) - Get issued card
* [GetCardFullDetails](#getcardfulldetails) - Get full card details
* [ListCards](#listcards) - List issued cards
* [UpdateCard](#updatecard) - Update issued card

## RequestCard

Request a virtual card be created
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.write` scope.


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
    res, err := s.CardIssuing.RequestCard(ctx, operations.PostRequestCardRequest{
        RequestCard: shared.RequestCard{
            AuthorizationControls: &shared.AuthorizationControls{
                SpendLimits: []shared.AuthorizationSpendLimitControl{
                    shared.AuthorizationSpendLimitControl{
                        Amount: moov.Int64(10000),
                        Duration: shared.AuthorizationSpendDurationTransaction.ToPointer(),
                    },
                    shared.AuthorizationSpendLimitControl{
                        Amount: moov.Int64(10000),
                        Duration: shared.AuthorizationSpendDurationTransaction.ToPointer(),
                    },
                },
            },
            AuthorizedUser: &shared.CreateAuthorizedUser{
                BirthDate: &shared.BirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                FirstName: moov.String("Jane"),
                LastName: moov.String("Doe"),
            },
            FundingWalletID: moov.String("labore"),
            Memo: moov.String("adipisci"),
            Type: shared.IssuedCardTypeSingleUse.ToPointer(),
        },
        AccountID: "a1108e0a-dcf4-4b92-9879-fce953f73ef7",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.PostRequestCardRequest](../../models/operations/postrequestcardrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.PostRequestCardResponse](../../models/operations/postrequestcardresponse.md), error**


## GetCard

Retrieve a single issued card associated with a Moov account
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.


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
    res, err := s.CardIssuing.GetCard(ctx, operations.GetIssuedCardRequest{
        AccountID: "fbc7abd7-4dd3-49c0-b5d2-cff7c70a4562",
        IssuedCardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetIssuedCardRequest](../../models/operations/getissuedcardrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetIssuedCardResponse](../../models/operations/getissuedcardresponse.md), error**


## GetCardFullDetails

Get issued card with PAN, CVV, and expiration. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read-secure` scope.


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
    res, err := s.CardIssuing.GetCardFullDetails(ctx, operations.GetFullIssuedCardRequest{
        AccountID: "6d436813-f16d-49f5-bce6-c556146c3e25",
        IssuedCardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FullIssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.GetFullIssuedCardRequest](../../models/operations/getfullissuedcardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.GetFullIssuedCardResponse](../../models/operations/getfullissuedcardresponse.md), error**


## ListCards

List Moov issued cards existing for the account.
<br><br> All supported query parameters are optional.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.


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
    res, err := s.CardIssuing.ListCards(ctx, operations.ListIssuedCardsRequest{
        AccountID: "0fb008c4-2e14-41aa-8366-c8dd6b144290",
        Count: moov.Int64(476477),
        Skip: moov.Int64(301598),
        States: shared.IssuedCardStateInactive.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.ListIssuedCardsRequest](../../models/operations/listissuedcardsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.ListIssuedCardsResponse](../../models/operations/listissuedcardsresponse.md), error**


## UpdateCard

Update a Moov issued card
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.write` scope.


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
    res, err := s.CardIssuing.UpdateCard(ctx, operations.UpdateIssuedCardRequest{
        UpdateIssuedCard: shared.UpdateIssuedCard{
            AuthorizationControls: &shared.AuthorizationControls{
                SpendLimits: []shared.AuthorizationSpendLimitControl{
                    shared.AuthorizationSpendLimitControl{
                        Amount: moov.Int64(10000),
                        Duration: shared.AuthorizationSpendDurationTransaction.ToPointer(),
                    },
                    shared.AuthorizationSpendLimitControl{
                        Amount: moov.Int64(10000),
                        Duration: shared.AuthorizationSpendDurationTransaction.ToPointer(),
                    },
                },
            },
            AuthorizedUser: &shared.CreateAuthorizedUser{
                BirthDate: &shared.BirthDate{
                    Day: 9,
                    Month: 11,
                    Year: 1989,
                },
                FirstName: moov.String("Jane"),
                LastName: moov.String("Doe"),
            },
            Memo: moov.String("esse"),
            State: shared.IssuedCardStateInactive.ToPointer(),
        },
        AccountID: "8a7bd466-d28c-410a-b3cd-ca4251904e52",
        IssuedCardID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateIssuedCardRequest](../../models/operations/updateissuedcardrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateIssuedCardResponse](../../models/operations/updateissuedcardresponse.md), error**

