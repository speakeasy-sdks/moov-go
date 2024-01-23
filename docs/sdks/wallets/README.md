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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var walletID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

**[*operations.GetWalletForAccountResponse](../../pkg/models/operations/getwalletforaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetTransaction

Get details on a specific wallet transaction. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var accountID string = "7e52eb20-20b9-4198-9c56-1f8d538908ad"

    var transactionID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    var walletID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

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

**[*operations.GetWalletTransactionResponse](../../pkg/models/operations/getwallettransactionresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

List the wallets associated with a Moov account. <br><br> To list wallets, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

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
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )


    var accountID string = "c184a429-302e-4aca-80db-f1718b882a50"

    ctx := context.Background()
    res, err := s.Wallets.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
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

**[*operations.ListWalletsForAccountResponse](../../pkg/models/operations/listwalletsforaccountresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## ListTransactions

List all the transactions associated with a particular Moov wallet. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String("Bearer <YOUR_ACCESS_TOKEN_HERE>"),
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.ListTransactions(ctx, operations.ListWalletTransactionsRequest{
        AccountID: "80a7198d-fdbe-4bfe-9fa3-36d451c2c158",
        WalletID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.ListWalletTransactionsRequest](../../pkg/models/operations/listwallettransactionsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.ListWalletTransactionsResponse](../../pkg/models/operations/listwallettransactionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
