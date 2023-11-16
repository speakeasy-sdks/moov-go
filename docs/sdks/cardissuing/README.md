# CardIssuing
(*CardIssuing*)

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


    requestCard := shared.RequestCard{
        AuthorizationControls: &shared.AuthorizationControls{
            SpendLimits: []shared.AuthorizationSpendLimitControl{
                shared.AuthorizationSpendLimitControl{
                    Amount: moovgo.Int64(10000),
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
    }

    var accountID string = "c6261876-780e-4c10-84b2-e8a3520f192b"

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

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |
| `requestCard`                                                   | [shared.RequestCard](../../../pkg/models/shared/requestcard.md) | :heavy_check_mark:                                              | N/A                                                             |
| `accountID`                                                     | *string*                                                        | :heavy_check_mark:                                              | ID of the account                                               |


### Response

**[*operations.PostRequestCardResponse](../../pkg/models/operations/postrequestcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetCard

Retrieve a single issued card associated with a Moov account
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.


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


    var accountID string = "2614fd0a-d685-42ae-8d5c-198b8e58f04c"

    var issuedCardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

**[*operations.GetIssuedCardResponse](../../pkg/models/operations/getissuedcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## GetCardFullDetails

Get issued card with PAN, CVV, and expiration. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read-secure` scope.


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


    var accountID string = "9dab2fa4-7710-412f-86a8-8f88e10db45e"

    var issuedCardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

**[*operations.GetFullIssuedCardResponse](../../pkg/models/operations/getfullissuedcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## ListCards

List Moov issued cards existing for the account.
<br><br> All supported query parameters are optional.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.


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


    var accountID string = "59888d1e-1919-498b-96e3-fc470952bb6c"

    var count *int64 = 632256

    var skip *int64 = 174785

    var states *shared.IssuedCardState = shared.IssuedCardStatePendingVerification

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
| `states`                                                                         | [*shared.IssuedCardState](../../../pkg/models/shared/issuedcardstate.md)         | :heavy_minus_sign:                                                               | Optional, comma-separated states to filter the Moov list issued cards response.<br/> |


### Response

**[*operations.ListIssuedCardsResponse](../../pkg/models/operations/listissuedcardsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## UpdateCard

Update a Moov issued card
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.write` scope.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
	"net/http"
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
    }

    var accountID string = "a292610d-0364-41fd-8f8b-865d1912e989"

    var issuedCardID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

| Parameter                                                                 | Type                                                                      | Required                                                                  | Description                                                               | Example                                                                   |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ctx`                                                                     | [context.Context](https://pkg.go.dev/context#Context)                     | :heavy_check_mark:                                                        | The context to use for the request.                                       |                                                                           |
| `updateIssuedCard`                                                        | [shared.UpdateIssuedCard](../../../pkg/models/shared/updateissuedcard.md) | :heavy_check_mark:                                                        | N/A                                                                       |                                                                           |
| `accountID`                                                               | *string*                                                                  | :heavy_check_mark:                                                        | ID of the account                                                         |                                                                           |
| `issuedCardID`                                                            | *string*                                                                  | :heavy_check_mark:                                                        | ID of the issued card                                                     | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                      |


### Response

**[*operations.UpdateIssuedCardResponse](../../pkg/models/operations/updateissuedcardresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
