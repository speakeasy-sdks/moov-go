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
    accountID := "e4be0560-13f5-49da-b57a-59ecfef66ef1"
    count := 791880
    skip := 685478
    status := shared.IssuedCardTransactionStatusDeclined

    ctx := context.Background()
    res, err := s.Transactions.List(ctx, accountID, count, skip, status)
    if err != nil {
        log.Fatal(err)
    }

    if res.IssuedCardTransactions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                 | Type                                                                                      | Required                                                                                  | Description                                                                               |
| ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------- |
| `ctx`                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                     | :heavy_check_mark:                                                                        | The context to use for the request.                                                       |
| `accountID`                                                                               | *string*                                                                                  | :heavy_check_mark:                                                                        | ID of the account                                                                         |
| `count`                                                                                   | **int64*                                                                                  | :heavy_minus_sign:                                                                        | Optional parameter to limit the number of results in the query                            |
| `skip`                                                                                    | **int64*                                                                                  | :heavy_minus_sign:                                                                        | The number of items to offset before starting to collect the result set                   |
| `status`                                                                                  | [*shared.IssuedCardTransactionStatus](../../models/shared/issuedcardtransactionstatus.md) | :heavy_minus_sign:                                                                        | Optional parameters to filter results IssuedCardTransactions.                             |


### Response

**[*operations.ListAccountIssuedCardTransactionsResponse](../../models/operations/listaccountissuedcardtransactionsresponse.md), error**

