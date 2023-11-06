# BankAccounts
(*BankAccounts*)

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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "f991ae31-67b4-4f05-b976-44ffd0cfd682"

    var bankAccountID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.BankAccounts.InitiateMicroDeposits(ctx, accountID, bankAccountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `bankAccountID`                                       | *string*                                              | :heavy_check_mark:                                    | ID of the bank account                                | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    completeMicroDepositsRequest := shared.CompleteMicroDepositsRequest{
        Amounts: []int64{
            18,
            21,
        },
    }

    var accountID string = "edb3fec3-2575-43fd-9928-10dbca09c27d"

    var bankAccountID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.BankAccounts.CompleteMicroDeposits(ctx, completeMicroDepositsRequest, accountID, bankAccountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.CompleteMicroDepositsResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                | Example                                                                                    |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |                                                                                            |
| `completeMicroDepositsRequest`                                                             | [shared.CompleteMicroDepositsRequest](../../models/shared/completemicrodepositsrequest.md) | :heavy_check_mark:                                                                         | N/A                                                                                        |                                                                                            |
| `accountID`                                                                                | *string*                                                                                   | :heavy_check_mark:                                                                         | ID of the account                                                                          |                                                                                            |
| `bankAccountID`                                                                            | *string*                                                                                   | :heavy_check_mark:                                                                         | ID of the bank account                                                                     | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                       |


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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"

    var bankAccountID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.BankAccounts.Delete(ctx, accountID, bankAccountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `bankAccountID`                                       | *string*                                              | :heavy_check_mark:                                    | ID of the bank account                                | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var bankAccountID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.BankAccounts.Get(ctx, accountID, bankAccountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `bankAccountID`                                       | *string*                                              | :heavy_check_mark:                                    | ID of the bank account                                | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var bankAccountPayload shared.BankAccountPayload = shared.CreateBankAccountPayloadPlaid(
            shared.Plaid{
                Plaid: &shared.PlaidIntegration{},
            },
    )

    var accountID string = "800f9f9c-1823-4596-ae95-1f9387e7ad1d"

    ctx := context.Background()
    res, err := s.BankAccounts.Link(ctx, bankAccountPayload, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccountResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ctx`                                                                  | [context.Context](https://pkg.go.dev/context#Context)                  | :heavy_check_mark:                                                     | The context to use for the request.                                    |
| `bankAccountPayload`                                                   | [shared.BankAccountPayload](../../models/shared/bankaccountpayload.md) | :heavy_check_mark:                                                     | N/A                                                                    |
| `accountID`                                                            | *string*                                                               | :heavy_check_mark:                                                     | ID of the account                                                      |


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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "c184a429-302e-4aca-80db-f1718b882a50"

    ctx := context.Background()
    res, err := s.BankAccounts.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.BankAccounts != nil {
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

**[*operations.ListBankAccountsResponse](../../models/operations/listbankaccountsresponse.md), error**

