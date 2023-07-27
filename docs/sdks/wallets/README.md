# Wallets

## Overview

A [Moov wallet](https://docs.moov.io/guides/wallet/) can serve as a funding source as you accumulate funds. You can also use the Moov wallet to:
- Pre-fund transfers for faster payouts
- Transfer funds between Moov wallets for instantly available funds

<em> If you've requested the `send-funds` or `collect-funds` capability, the `wallet` capability will be automatically requested as well. Read more on the [data requirements for wallets here](https://docs.moov.io/guides/accounts/capabilities/#wallet).</em>


### Available Operations

* [Get](#get) - Get wallet
* [GetTransaction](#gettransaction) - Get wallet transaction
* [List](#list) - List wallets
* [ListTransactions](#listtransactions) - List wallet transactions

## Get

Get information on a specific wallet (e.g., the available balance). <br><br> To get wallet information, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
    res, err := s.Wallets.Get(ctx, operations.GetWalletForAccountRequest{
        AccountID: "ca5acfbe-2fd5-4707-9779-29177deac646",
        WalletID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `ctx`                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                          | :heavy_check_mark:                                                                             | The context to use for the request.                                                            |
| `request`                                                                                      | [operations.GetWalletForAccountRequest](../../models/operations/getwalletforaccountrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[*operations.GetWalletForAccountResponse](../../models/operations/getwalletforaccountresponse.md), error**


## GetTransaction

Get details on a specific wallet transaction. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
    res, err := s.Wallets.GetTransaction(ctx, operations.GetWalletTransactionRequest{
        AccountID: "ecb57340-9e3e-4b1e-9a2b-12eb07f116db",
        TransactionID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
        WalletID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WalletTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetWalletTransactionRequest](../../models/operations/getwallettransactionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetWalletTransactionResponse](../../models/operations/getwallettransactionresponse.md), error**


## List

List the wallets associated with a Moov account. <br><br> To list wallets, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
    res, err := s.Wallets.List(ctx, operations.ListWalletsForAccountRequest{
        AccountID: "99545fc9-5fa8-4897-8e18-9dbb30fcb33e",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallets != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.ListWalletsForAccountRequest](../../models/operations/listwalletsforaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.ListWalletsForAccountResponse](../../models/operations/listwalletsforaccountresponse.md), error**


## ListTransactions

List all the transactions associated with a particular Moov wallet. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
    res, err := s.Wallets.ListTransactions(ctx, operations.ListWalletTransactionsRequest{
        AccountID: "a055b197-cd44-4e2f-92d8-2d3513bb6f48",
        CompletedEndDateTime: moov.String("distinctio"),
        CompletedStartDateTime: moov.String("nisi"),
        Count: moov.Int64(335977),
        CreatedEndDateTime: moov.String("nisi"),
        CreatedStartDateTime: moov.String("libero"),
        Skip: moov.Int64(794507),
        SourceID: moov.String("facere"),
        SourceType: moov.String("facilis"),
        Status: moov.String("ipsum"),
        TransactionType: moov.String("ad"),
        WalletID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WalletTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.ListWalletTransactionsRequest](../../models/operations/listwallettransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.ListWalletTransactionsResponse](../../models/operations/listwallettransactionsresponse.md), error**

