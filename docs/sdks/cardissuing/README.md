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
            FundingWalletID: moov.String("doloribus"),
            Memo: moov.String("iusto"),
            Type: shared.IssuedCardTypeSingleUse.ToPointer(),
        },
        AccountID: "c70a4562-6d43-4681-bf16-d9f5fce6c556",
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
        AccountID: "146c3e25-0fb0-408c-82e1-41aac366c8dd",
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
        AccountID: "6b144290-7474-4778-a7bd-466d28c10ab3",
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
        AccountID: "cdca4251-904e-4523-87e0-bc7178e4796f",
        Count: moov.Int64(174112),
        Skip: moov.Int64(645570),
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
            Memo: moov.String("porro"),
            State: shared.IssuedCardStateInactive.ToPointer(),
        },
        AccountID: "88282aa4-8256-42f2-a2e9-817ee17cbe61",
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

