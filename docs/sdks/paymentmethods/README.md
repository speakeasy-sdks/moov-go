# PaymentMethods

## Overview

[Payment methods](https://docs.moov.io/guides/money-movement/payment-methods/) represent all of the ways an account can move funds to another Moov account. Payment methods are generated programmatically when a card or bank account is added or the status is updated e.g., `ach-debit-fund` will be added as a payment method once the bank account is verified.

<em>RTP® is not yet generally available on Moov. Contact us for more information.</em>


### Available Operations

* [Get](#get) - Get payment method
* [List](#list) - List payment methods

## Get

Get the specified payment method associated with a Moov account. <br><br> To get a payment method, you must specify the `/accounts/{accountID}/payment-methods.read` scope.

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
    res, err := s.PaymentMethods.Get(ctx, operations.GetPaymentMethodRequest{
        AccountID: "e865e795-6f92-451a-9a9d-a660ff57bfaa",
        PaymentMethodID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.GetPaymentMethodRequest](../../models/operations/getpaymentmethodrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.GetPaymentMethodResponse](../../models/operations/getpaymentmethodresponse.md), error**


## List

Retrieve a list of payment methods associated with a Moov account. <br><br> To list payment methods, you must specify the `/accounts/{accountID}/payment-methods.read` scope.

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
    res, err := s.PaymentMethods.List(ctx, operations.ListPaymentMethodsRequest{
        AccountID: "d4f9efc1-b451-42c1-8326-48dc2f615199",
        SourceID: moov.String("ebfd0e9f-e6c6-432c-a3ae-d0117996312f"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethods != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.ListPaymentMethodsRequest](../../models/operations/listpaymentmethodsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.ListPaymentMethodsResponse](../../models/operations/listpaymentmethodsresponse.md), error**

