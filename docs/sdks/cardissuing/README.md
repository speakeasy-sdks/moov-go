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
    requestCard := shared.RequestCard{
        AuthorizationControls: &shared.AuthorizationControls{
            SpendLimits: []shared.AuthorizationSpendLimitControl{
                shared.AuthorizationSpendLimitControl{
                    Amount: moovgo.Int64(10000),
                    Duration: moovgo.String("inventore"),
                },
            },
        },
        AuthorizedUser: &shared.CreateAuthorizedUser{
            BirthDate: &shared.BirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            FirstName: moovgo.String("Jane"),
            LastName: moovgo.String("Doe"),
        },
        FundingWalletID: moovgo.String("non"),
        Memo: moovgo.String("et"),
        Type: moovgo.String("dolorum"),
    }
    accountID := "ac366c8d-d6b1-4442-9074-74778a7bd466"

    ctx := context.Background()
    res, err := s.CardIssuing.RequestCard(ctx, requestCard, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `requestCard`                                            | [shared.RequestCard](../../models/shared/requestcard.md) | :heavy_check_mark:                                       | N/A                                                      |
| `accountID`                                              | *string*                                                 | :heavy_check_mark:                                       | ID of the account                                        |


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
    accountID := "d28c10ab-3cdc-4a42-9190-4e523c7e0bc7"
    issuedCardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.CardIssuing.GetCard(ctx, accountID, issuedCardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `issuedCardID`                                        | *string*                                              | :heavy_check_mark:                                    | ID of the issued card                                 | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
    accountID := "178e4796-f2a7-40c6-8828-2aa482562f22"
    issuedCardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.CardIssuing.GetCardFullDetails(ctx, accountID, issuedCardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.FullIssuedCard != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `issuedCardID`                                        | *string*                                              | :heavy_check_mark:                                    | ID of the issued card                                 | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
    accountID := "2e9817ee-17cb-4e61-a6b7-b95bc0ab3c20"
    count := 796392
    skip := 308286
    states := shared.IssuedCardStateClosed

    ctx := context.Background()
    res, err := s.CardIssuing.ListCards(ctx, accountID, count, skip, states)
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCards != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `accountID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | ID of the account                                                                |
| `count`                                                                          | **int64*                                                                         | :heavy_minus_sign:                                                               | Optional parameter to limit the number of results in the query                   |
| `skip`                                                                           | **int64*                                                                         | :heavy_minus_sign:                                                               | The number of items to offset before starting to collect the result set          |
| `states`                                                                         | [*shared.IssuedCardState](../../models/shared/issuedcardstate.md)                | :heavy_minus_sign:                                                               | Optional, comma-separated states to filter the Moov list issued cards response.<br/> |


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
    updateIssuedCard := shared.UpdateIssuedCard{
        AuthorizationControls: &shared.AuthorizationControls{
            SpendLimits: []shared.AuthorizationSpendLimitControl{
                shared.AuthorizationSpendLimitControl{
                    Amount: moovgo.Int64(10000),
                    Duration: moovgo.String("consectetur"),
                },
            },
        },
        AuthorizedUser: &shared.CreateAuthorizedUser{
            BirthDate: &shared.BirthDate{
                Day: 9,
                Month: 11,
                Year: 1989,
            },
            FirstName: moovgo.String("Jane"),
            LastName: moovgo.String("Doe"),
        },
        Memo: moovgo.String("esse"),
        State: shared.IssuedCardStatePendingVerification.ToPointer(),
    }
    accountID := "9fd871f9-9dd2-4efd-921a-a6f1e674bdb0"
    issuedCardID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.CardIssuing.UpdateCard(ctx, updateIssuedCard, accountID, issuedCardID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        | Example                                                            |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |                                                                    |
| `updateIssuedCard`                                                 | [shared.UpdateIssuedCard](../../models/shared/updateissuedcard.md) | :heavy_check_mark:                                                 | N/A                                                                |                                                                    |
| `accountID`                                                        | *string*                                                           | :heavy_check_mark:                                                 | ID of the account                                                  |                                                                    |
| `issuedCardID`                                                     | *string*                                                           | :heavy_check_mark:                                                 | ID of the issued card                                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                               |


### Response

**[*operations.UpdateIssuedCardResponse](../../models/operations/updateissuedcardresponse.md), error**

