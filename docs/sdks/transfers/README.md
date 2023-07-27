# Transfers

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
    res, err := s.Transfers.Cancel(ctx, operations.CancelTransferRequest{
        CreateReversal: &shared.CreateReversal{
            Amount: moov.Int64(1000),
        },
        XIdempotencyKey: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
        TransferID: "58c4d86e-68e4-4be0-9601-3f59da757a59",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatedReversal != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CancelTransferRequest](../../models/operations/canceltransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Transfers.Create(ctx, operations.CreateTransferRequest{
        CreateTransfer: shared.CreateTransfer{
            Amount: &shared.Amount{
                Currency: "USD",
                Value: 1204,
            },
            Description: moov.String("Pay Instructor for May 15 Class"),
            Destination: &shared.CreateTransferDestination{
                AchDetails: &shared.CreateACHDetailsBase{
                    CompanyEntryDescription: moov.String("Gym Dues"),
                    OriginatingCompanyName: moov.String("Whole Body Fit"),
                },
                PaymentMethodID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            },
            FacilitatorFee: &shared.CreateFacilitatorFee{
                Markup: moov.Int64(880679),
                Total: moov.Int64(774684),
            },
            Metadata: map[string]string{
                "necessitatibus": "asperiores",
                "ex": "voluptas",
                "debitis": "delectus",
                "quae": "minus",
            },
            Source: &shared.CreateTransferSource{
                AchDetails: &shared.CreateAchDetailsSource{
                    CompanyEntryDescription: moov.String("Gym Dues"),
                    DebitHoldPeriod: shared.CreateAchDetailsSourceDebitHoldPeriodTwoDays.ToPointer(),
                    OriginatingCompanyName: moov.String("Whole Body Fit"),
                },
                CardDetails: &shared.CreateCardDetails{
                    DynamicDescriptor: moov.String("WhlBdy *Yoga 11-12"),
                    TransactionSource: shared.TransactionSourceUnscheduled.ToPointer(),
                },
                PaymentMethodID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
                TransferID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            },
        },
        XIdempotencyKey: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
        XWaitFor: shared.WaitForRailResponse.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.TransferPostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
	"openapi"
	"openapi/pkg/models/shared"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Transfers.GenerateOptions(ctx, shared.CreateTransferOptions{
        Amount: shared.Amount{
            Currency: "USD",
            Value: 1204,
        },
        Destination: &shared.CreateTransferOptionsDestination{
            AccountID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            PaymentMethodID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
        },
        Source: &shared.CreateTransferOptionsSource{
            AccountID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
            PaymentMethodID: moov.String("ec7e1848-dc80-4ab0-8827-dd7fc0737b43"),
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
    res, err := s.Transfers.Get(ctx, operations.GetTransferRequest{
        AccountID: moov.String("a3383c2b-eb47-4737-bc8d-72f64d1db1f2"),
        TransferID: "c4310661-e963-449e-9cf9-e06e3a437000",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransferFull != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetTransferRequest](../../models/operations/gettransferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.Transfers.GetRefund(ctx, operations.GetRefundRequest{
        RefundID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
        TransferID: "ae6b6bc9-b8f7-459e-ac55-a9741d311352",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRefund != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetRefundRequest](../../models/operations/getrefundrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
    res, err := s.Transfers.ListRefunds(ctx, operations.ListRefundsRequest{
        TransferID: "965bb8a7-2026-4114-b5e1-39dbc2259b1a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetRefunds != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListRefundsRequest](../../models/operations/listrefundsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


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
    res, err := s.Transfers.Refund(ctx, operations.RefundTransferRequest{
        CreateRefund: &shared.CreateRefund{
            Amount: moov.Int64(1000),
        },
        XIdempotencyKey: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
        XWaitFor: shared.WaitForRailResponse.ToPointer(),
        TransferID: "bda8c070-e108-44cb-8672-d1ad879eeb96",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.RefundPostResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.RefundTransferRequest](../../models/operations/refundtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


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
    res, err := s.Transfers.Update(ctx, operations.PatchTransferRequest{
        PatchTransfer: shared.PatchTransfer{
            Metadata: map[string]string{
                "ipsam": "rerum",
                "laudantium": "corporis",
            },
        },
        AccountID: moov.String("efbd02ba-e0be-42d7-8225-9e3ea4b5197f"),
        TransferID: "92443da7-ce52-4b89-9c53-7c6454efb0b3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetTransferFull != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.PatchTransferRequest](../../models/operations/patchtransferrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.PatchTransferResponse](../../models/operations/patchtransferresponse.md), error**

