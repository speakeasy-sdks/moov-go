# PaymentMethods
(*PaymentMethods*)

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
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"context"
	"log"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var paymentMethodID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.PaymentMethods.Get(ctx, accountID, paymentMethodID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethod != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `paymentMethodID`                                     | *string*                                              | :heavy_check_mark:                                    | ID of the payment method                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.GetPaymentMethodResponse](../../pkg/models/operations/getpaymentmethodresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Retrieve a list of payment methods associated with a Moov account. <br><br> To list payment methods, you must specify the `/accounts/{accountID}/payment-methods.read` scope.

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
            AccessToken: moovgo.String(""),
        }),
    )


    var accountID string = "c184a429-302e-4aca-80db-f1718b882a50"

    var sourceID *string = "80555741-9e79-40e2-b205-5dd402eb66ec"

    ctx := context.Background()
    res, err := s.PaymentMethods.List(ctx, accountID, sourceID)
    if err != nil {
        log.Fatal(err)
    }

    if res.PaymentMethods != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `accountID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | ID of the account                                                        |
| `sourceID`                                                               | **string*                                                                | :heavy_minus_sign:                                                       | Optional parameter to filter the account's payment methods by source ID. |


### Response

**[*operations.ListPaymentMethodsResponse](../../pkg/models/operations/listpaymentmethodsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
