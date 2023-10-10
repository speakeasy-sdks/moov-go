# Transfers
(*Transfers*)

## Overview

A [transfer](https://docs.moov.io/guides/money-movement/) is the movement of money between Moov accounts, from source to destination. Provided you have linked a bank account which has been verified, you can initiate a transfer to another Moov account. All you need to do is note a [paymentMethod](#tag/Payment-methods), the $ amount of the transfer, and a brief description. With Moov, you can also implement transfer groups, allowing you to associate multiple transfers together and run them sequentially. To learn more, read our guide on [transfer groups](https://docs.moov.io/guides/money-movement/transfer-groups/#transfer-statuses).

### Available Operations

* [Cancel](#cancel) - Cancel or refund a card transfer
* [Create](#create) - Create a transfer
* [GenerateOptions](#generateoptions) - Generate transfer options
* [Get](#get) - Get a transfer
* [GetRefund](#getrefund) - Get refund details
* [ListRefunds](#listrefunds) - Get a list of refunds for a card transfer
* [Refund](#refund) - Refund a transfer
* [Update](#update) - Patch transfer metadata

## Cancel

Reverses a card transfer by initiating a cancellation or refund depending on the transaction status

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


    var xIdempotencyKey string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    var transferID string = "03fa7112-315a-4072-a9f2-43f3f1ec962e"

    createReversal := &shared.CreateReversal{
        Amount: moovgo.Int64(1000),
    }

    ctx := context.Background()
    res, err := s.Transfers.Cancel(ctx, xIdempotencyKey, transferID, createReversal)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatedReversal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                       | Type                                                            | Required                                                        | Description                                                     | Example                                                         |
| --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- | --------------------------------------------------------------- |
| `ctx`                                                           | [context.Context](https://pkg.go.dev/context#Context)           | :heavy_check_mark:                                              | The context to use for the request.                             |                                                                 |
| `xIdempotencyKey`                                               | *string*                                                        | :heavy_check_mark:                                              | Prevents duplicate reversals from being created                 | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                            |
| `transferID`                                                    | *string*                                                        | :heavy_check_mark:                                              | ID of the Transfer                                              |                                                                 |
| `createReversal`                                                | [*shared.CreateReversal](../../models/shared/createreversal.md) | :heavy_minus_sign:                                              | N/A                                                             |                                                                 |


### Response

**[*operations.CancelTransferResponse](../../models/operations/canceltransferresponse.md), error**


## Create

Move money by providing the source, destination, and amount in the request body. <br><br> To create a transfer, you must specify the `/accounts/{yourAccountID}/transfers.write` scope. <br> You can find your account id on the [Business details](https://dashboard.moov.io/settings/business) page.

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


    createTransfer := shared.CreateTransfer{
        Amount: &shared.Amount{
            Currency: "USD",
            Value: 1204,
        },
        Description: moovgo.String("Pay Instructor for May 15 Class"),
        Destination: &shared.CreateTransferDestination{
            AchDetails: &shared.CreateACHDetailsBase{
                CompanyEntryDescription: moovgo.String("Gym Dues"),
                OriginatingCompanyName: moovgo.String("Whole Body Fit"),
            },
            PaymentMethodID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
        },
        FacilitatorFee: &shared.CreateFacilitatorFee{},
        Metadata: map[string]string{
            "online": "Configuration",
        },
        Source: &shared.CreateTransferSource{
            AchDetails: &shared.CreateAchDetailsSource{
                CompanyEntryDescription: moovgo.String("Gym Dues"),
                DebitHoldPeriod: shared.CreateAchDetailsSourceDebitHoldPeriodTwoDays.ToPointer(),
                OriginatingCompanyName: moovgo.String("Whole Body Fit"),
            },
            CardDetails: &shared.CreateCardDetails{
                DynamicDescriptor: moovgo.String("WhlBdy *Yoga 11-12"),
            },
            PaymentMethodID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            TransferID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
        },
    }

    var xIdempotencyKey string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    var xWaitFor *shared.WaitFor = shared.WaitForRailResponse

    ctx := context.Background()
    res, err := s.Transfers.Create(ctx, createTransfer, xIdempotencyKey, xWaitFor)
    if err != nil {
        log.Fatal(err)
    }

    if res.TransferPostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                              | Type                                                                                                                                                                                                                                                   | Required                                                                                                                                                                                                                                               | Description                                                                                                                                                                                                                                            | Example                                                                                                                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                                                                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                     | The context to use for the request.                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                        |
| `createTransfer`                                                                                                                                                                                                                                       | [shared.CreateTransfer](../../models/shared/createtransfer.md)                                                                                                                                                                                         | :heavy_check_mark:                                                                                                                                                                                                                                     | N/A                                                                                                                                                                                                                                                    |                                                                                                                                                                                                                                                        |
| `xIdempotencyKey`                                                                                                                                                                                                                                      | *string*                                                                                                                                                                                                                                               | :heavy_check_mark:                                                                                                                                                                                                                                     | Prevents duplicate transfers from being created. Note that we only accept UUID v4.                                                                                                                                                                     | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                                                                                                                                                                   |
| `xWaitFor`                                                                                                                                                                                                                                             | [*shared.WaitFor](../../models/shared/waitfor.md)                                                                                                                                                                                                      | :heavy_minus_sign:                                                                                                                                                                                                                                     | Optional header that indicates whether to return a synchronous response that includes full transfer and rail-specific details or an asynchronous response indicating the transfer was created (this is the default response if the header is omitted). |                                                                                                                                                                                                                                                        |


### Response

**[*operations.CreateTransferResponse](../../models/operations/createtransferresponse.md), error**


## GenerateOptions

Generate available payment method options for one or multiple transfer participants depending on the accountID or paymentMethodID you supply in the request. <br><br> To get transfer options, you must specify the `/accounts/{yourAccountID}/transfers.read` scope. The accountID you include should be associated with the facilitator account. <br> You can find your accountID on the [Business details](https://dashboard.moov.io/settings/business) page.

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

    ctx := context.Background()
    res, err := s.Transfers.GenerateOptions(ctx, shared.CreateTransferOptions{
        Amount: shared.Amount{
            Currency: "USD",
            Value: 1204,
        },
        Destination: &shared.CreateTransferOptionsDestination{
            AccountID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            PaymentMethodID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
        },
        Source: &shared.CreateTransferOptionsSource{
            AccountID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            PaymentMethodID: moovgo.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatedTransferOptions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [shared.CreateTransferOptions](../../models/shared/createtransferoptions.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.CreateTransferOptionsResponse](../../models/operations/createtransferoptionsresponse.md), error**


## Get

Retrieve full transfer details such as the amount, source, and destination. Payment rail-specific details are included in the source and destination. <br><br> To get a transfer, you must specify the `/accounts/{yourAccountID}/transfers.read` scope. The accountID included must be your facilitator accountID. <br> You can find your accountID on the [Business details](https://dashboard.moov.io/settings/business) page.

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


    var transferID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var accountID *string = "58ccc65b-c928-4154-952e-30c048b8c2b5"

    ctx := context.Background()
    res, err := s.Transfers.Get(ctx, transferID, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransferFull != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `transferID`                                          | *string*                                              | :heavy_check_mark:                                    | ID of the Transfer                                    |
| `accountID`                                           | **string*                                             | :heavy_minus_sign:                                    | ID of a connected account                             |


### Response

**[*operations.GetTransferResponse](../../models/operations/gettransferresponse.md), error**


## GetRefund

Get details of a refund for a card transfer

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


    var refundID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    var transferID string = "dcb7cb00-9cc1-4b92-8e6f-29c9d50583d5"

    ctx := context.Background()
    res, err := s.Transfers.GetRefund(ctx, refundID, transferID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRefund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `refundID`                                            | *string*                                              | :heavy_check_mark:                                    | ID of the refund                                      | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |
| `transferID`                                          | *string*                                              | :heavy_check_mark:                                    | ID of the Transfer                                    |                                                       |


### Response

**[*operations.GetRefundResponse](../../models/operations/getrefundresponse.md), error**


## ListRefunds

Get a list of refunds for a card transfer

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


    var transferID string = "7fdf7d4f-a689-4415-8f2e-71877eebee60"

    ctx := context.Background()
    res, err := s.Transfers.ListRefunds(ctx, transferID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRefunds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `transferID`                                          | *string*                                              | :heavy_check_mark:                                    | ID of the Transfer                                    |


### Response

**[*operations.ListRefundsResponse](../../models/operations/listrefundsresponse.md), error**


## Refund

<strong>Use the <a href="index.html#tag/Transfers/operation/reverseTransfer">Cancel or refund a card transfer</a> endpoint for more robust cancel and refund options.</strong> <br><br> Initiate a refund for a card transfer <br><br> To initiate a refund, you will need to specify the `/accounts/{accountID}/transfers.write` scope.

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


    var xIdempotencyKey string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    var transferID string = "94e35bc6-be02-470f-a21a-a226b055b594"

    createRefund := &shared.CreateRefund{
        Amount: moovgo.Int64(1000),
    }

    var xWaitFor *shared.WaitFor = shared.WaitForRailResponse

    ctx := context.Background()
    res, err := s.Transfers.Refund(ctx, xIdempotencyKey, transferID, createRefund, xWaitFor)
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundPostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                                                                                                                                                 | Type                                                                                                                                                                                                                                                      | Required                                                                                                                                                                                                                                                  | Description                                                                                                                                                                                                                                               | Example                                                                                                                                                                                                                                                   |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                                                                                                                                                     | [context.Context](https://pkg.go.dev/context#Context)                                                                                                                                                                                                     | :heavy_check_mark:                                                                                                                                                                                                                                        | The context to use for the request.                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                           |
| `xIdempotencyKey`                                                                                                                                                                                                                                         | *string*                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                        | Prevents duplicate refunds from being created. Note that we only accept UUID v4.                                                                                                                                                                          | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                                                                                                                                                                      |
| `transferID`                                                                                                                                                                                                                                              | *string*                                                                                                                                                                                                                                                  | :heavy_check_mark:                                                                                                                                                                                                                                        | ID of the Transfer                                                                                                                                                                                                                                        |                                                                                                                                                                                                                                                           |
| `createRefund`                                                                                                                                                                                                                                            | [*shared.CreateRefund](../../models/shared/createrefund.md)                                                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                                                                                                        | N/A                                                                                                                                                                                                                                                       |                                                                                                                                                                                                                                                           |
| `xWaitFor`                                                                                                                                                                                                                                                | [*shared.WaitFor](../../models/shared/waitfor.md)                                                                                                                                                                                                         | :heavy_minus_sign:                                                                                                                                                                                                                                        | Optional header that indicates whether to return a synchronous response that includes the full refund and card transaction details or an asynchronous response indicating the refund was created (this is the default response if the header is omitted). |                                                                                                                                                                                                                                                           |


### Response

**[*operations.RefundTransferResponse](../../models/operations/refundtransferresponse.md), error**


## Update

Update the metadata contained on a transfer <br><br> To patch a transfer, you must specify the `/accounts/{yourAccountID}/transfers.write` scope. The accountID included must be your facilitator accountID. <br> You can find your account ID on the [Business details](https://dashboard.moov.io/settings/business) page.

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


    patchTransfer := shared.PatchTransfer{
        Metadata: map[string]string{
            "Van": "East",
        },
    }

    var transferID string = "bf4aa77f-204e-4775-8c35-2acfe54077ca"

    var accountID *string = "bf6805c5-ca71-4871-8355-ad7d4e1b5845"

    ctx := context.Background()
    res, err := s.Transfers.Update(ctx, patchTransfer, transferID, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransferFull != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                    | Type                                                         | Required                                                     | Description                                                  |
| ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| `ctx`                                                        | [context.Context](https://pkg.go.dev/context#Context)        | :heavy_check_mark:                                           | The context to use for the request.                          |
| `patchTransfer`                                              | [shared.PatchTransfer](../../models/shared/patchtransfer.md) | :heavy_check_mark:                                           | N/A                                                          |
| `transferID`                                                 | *string*                                                     | :heavy_check_mark:                                           | ID of the Transfer                                           |
| `accountID`                                                  | **string*                                                    | :heavy_minus_sign:                                           | ID of a connected account                                    |


### Response

**[*operations.PatchTransferResponse](../../models/operations/patchtransferresponse.md), error**

