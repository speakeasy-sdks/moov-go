# Transactions

## Overview

A transaction is a record of a card's activity on a particular Moov account.

### Available Operations

* [List](#list) - Get account transactions

## List

List issued card transactions associated with a Moov account

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
    res, err := s.Transactions.List(ctx, operations.ListAccountIssuedCardTransactionsRequest{
        AccountID: "0421813d-5208-4ece-be25-3b668451c6c6",
        Count: petstore.Int64(927212),
        Skip: petstore.Int64(160393),
        Status: shared.IssuedCardTransactionStatusPending.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCardTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                                      | :heavy_check_mark:                                                                                                         | The context to use for the request.                                                                                        |
| `request`                                                                                                                  | [operations.ListAccountIssuedCardTransactionsRequest](../../models/operations/listaccountissuedcardtransactionsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[*operations.ListAccountIssuedCardTransactionsResponse](../../models/operations/listaccountissuedcardtransactionsresponse.md), error**

