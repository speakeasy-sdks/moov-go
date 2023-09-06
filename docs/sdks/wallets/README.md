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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "9177deac-646e-4cb5-b340-9e3eb1e5a2b1"
    walletID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Wallets.Get(ctx, accountID, walletID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `walletID`                                            | *string*                                              | :heavy_check_mark:                                    | ID of the wallet                                      | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "2eb07f11-6db9-4954-9fc9-5fa88970e189"
    transactionID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"
    walletID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Wallets.GetTransaction(ctx, accountID, transactionID, walletID)
    if err != nil {
        log.Fatal(err)
    }

    if res.WalletTransaction != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `transactionID`                                       | *string*                                              | :heavy_check_mark:                                    | ID associated with the wallet transaction.            | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |
| `walletID`                                            | *string*                                              | :heavy_check_mark:                                    | ID of the wallet                                      | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    accountID := "dbb30fcb-33ea-4055-b197-cd44e2f52d82"

    ctx := context.Background()
    res, err := s.Wallets.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallets != nil {
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

**[*operations.ListWalletsForAccountResponse](../../models/operations/listwalletsforaccountresponse.md), error**


## ListTransactions

List all the transactions associated with a particular Moov wallet. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.ListTransactions(ctx, operations.ListWalletTransactionsRequest{
        AccountID: "d3513bb6-f48b-4656-bcdb-35ff2e4b2753",
        CompletedEndDateTime: moov.String("iusto"),
        CompletedStartDateTime: moov.String("est"),
        Count: moov.Int64(522176),
        CreatedEndDateTime: moov.String("eligendi"),
        CreatedStartDateTime: moov.String("fugiat"),
        Skip: moov.Int64(604078),
        SourceID: moov.String("officiis"),
        SourceType: moov.String("ducimus"),
        Status: moov.String("dolor"),
        TransactionType: moov.String("dicta"),
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

