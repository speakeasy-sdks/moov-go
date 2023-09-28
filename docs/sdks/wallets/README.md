# Wallets
(*Wallets*)

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
    accountID := "ea055b19-7cd4-44e2-b52d-82d3513bb6f4"
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
    accountID := "8b656bcd-b35f-4f2e-8b27-537a8cd9e731"
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
    accountID := "9c177d52-5f77-4b11-8eeb-52ff785fc378"

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

    ctx := context.Background()
    res, err := s.Wallets.ListTransactions(ctx, operations.ListWalletTransactionsRequest{
        AccountID: "14d4c98e-0c2b-4b89-ab75-dad636c60050",
        CompletedEndDateTime: moovgo.String("amet"),
        CompletedStartDateTime: moovgo.String("illum"),
        Count: moovgo.Int64(506863),
        CreatedEndDateTime: moovgo.String("quidem"),
        CreatedStartDateTime: moovgo.String("cum"),
        Skip: moovgo.Int64(230411),
        SourceID: moovgo.String("quasi"),
        SourceType: moovgo.String("dicta"),
        Status: moovgo.String("laudantium"),
        TransactionType: moovgo.String("doloremque"),
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

