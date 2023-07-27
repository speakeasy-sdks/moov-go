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
    accountID := "78a64584-273a-4841-8d16-2309fb092992"
    count := 81369
    skip := 686362
    status := shared.IssuedCardTransactionStatusVoided

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

