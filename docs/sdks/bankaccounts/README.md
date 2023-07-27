# BankAccounts

## Overview

To transfer money with Moov, you’ll need to link a bank account to your Moov account, then verify that account. You can link a bank account to a Moov account by providing the bank account number, routing number, and Moov account ID. 

We require micro-deposit verification to reduce the risk of fraud or unauthorized activity. You can verify a bank account by initiating [micro-deposits](https://docs.moov.io/guides/sources/bank-accounts/#micro-deposit-verification), sending two small credit transfers to the bank account you want to confirm. Note that there is no way to initiate a micro-deposit from your bank of choice. 

Alternatively, you can link and verify a bank account in one step through an instant account verification token from a third party provider like [Plaid](https://docs.moov.io/guides/sources/bank-accounts/plaid) or [MX](https://docs.moov.io/guides/sources/bank-accounts/mx/). Bank accounts can have the following statuses: `new`, `pending`, `verified`, `verificationFailed`, `errored`. Learn more by reading our guide on [bank accounts](https://docs.moov.io/guides/sources/bank-accounts/).


### Available Operations

* [InitiateMicroDeposits](#initiatemicrodeposits) - Initiate micro-deposits
* [CompleteMicroDeposits](#completemicrodeposits) - Complete micro-deposits
* [Delete](#delete) - Delete bank account
* [Get](#get) - Get bank account
* [Link](#link) - Bank account
* [List](#list) - List bank accounts

## InitiateMicroDeposits

Micro-deposits help confirm bank account ownership, helping reduce fraud and the risk of unauthorized activity. Use this method to initiate the micro-deposit verification, sending two small credit transfers to the bank account you want to confirm. If you request micro-deposits before 4:15PM ET, they will appear that same day. If you request micro-deposits any time after 4:15PM ET, they will appear the next banking day. When the two credits are initiated, Moov simultaneously initiates a debit to recoup the micro-deposits.<br><br> `sandbox` - Micro-deposits initiated for a `sandbox` bank account will always be `$0.00` / `$0.00` and instantly verifiable once initiated. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

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
    res, err := s.BankAccounts.InitiateMicroDeposits(ctx, operations.PostInitiateMicroDepositsRequest{
        AccountID: "2a94bb4f-63c9-469e-9a3e-fa77dfb14cd6",
        BankAccountID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
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

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.PostInitiateMicroDepositsRequest](../../models/operations/postinitiatemicrodepositsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.PostInitiateMicroDepositsResponse](../../models/operations/postinitiatemicrodepositsresponse.md), error**


## CompleteMicroDeposits

Complete the micro-deposit validation process by passing the amounts of the two transfers within three tries. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

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
    res, err := s.BankAccounts.CompleteMicroDeposits(ctx, operations.PutCompleteMicroDepositsRequest{
        CompleteMicroDepositsRequest: shared.CompleteMicroDepositsRequest{
            Amounts: []int64{
                675439,
                881104,
            },
        },
        AccountID: "395efb9b-a88f-43a6-a997-074ba4469b6e",
        BankAccountID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CompleteMicroDepositsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.PutCompleteMicroDepositsRequest](../../models/operations/putcompletemicrodepositsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.PutCompleteMicroDepositsResponse](../../models/operations/putcompletemicrodepositsresponse.md), error**


## Delete

Discontinue using a specified bank account linked to a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

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
    res, err := s.BankAccounts.Delete(ctx, operations.DeleteBankAccountRequest{
        AccountID: "21419598-90af-4a56-be25-16fe4c8b711e",
        BankAccountID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
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

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.DeleteBankAccountRequest](../../models/operations/deletebankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.DeleteBankAccountResponse](../../models/operations/deletebankaccountresponse.md), error**


## Get

Retrieve bank account details (i.e. routing number or account type) associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

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
    res, err := s.BankAccounts.Get(ctx, operations.GetBankAccountRequest{
        AccountID: "5b7fd2ed-0289-421c-9dc6-92601fb576b0",
        BankAccountID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetBankAccountRequest](../../models/operations/getbankaccountrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetBankAccountResponse](../../models/operations/getbankaccountresponse.md), error**


## Link

Link a bank account to an existing Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

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
    res, err := s.BankAccounts.Link(ctx, operations.LinkBankAccountRequest{
        BankAccountPayload: shared.BankAccountPayload{},
        AccountID: "d5f0d30c-5fbb-4258-b053-202c73d5fe9b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.LinkBankAccountRequest](../../models/operations/linkbankaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.LinkBankAccountResponse](../../models/operations/linkbankaccountresponse.md), error**


## List

List all the bank accounts associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

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
    res, err := s.BankAccounts.List(ctx, operations.ListBankAccountsRequest{
        AccountID: "90c28909-b3fe-449a-8d9c-bf48633323f9",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccounts != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListBankAccountsRequest](../../models/operations/listbankaccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListBankAccountsResponse](../../models/operations/listbankaccountsresponse.md), error**

