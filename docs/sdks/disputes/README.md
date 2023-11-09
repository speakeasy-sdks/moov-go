# Disputes
(*Disputes*)

## Overview

A [dispute](https://docs.moov.io/guides/money-movement/cards/disputes/) is a situation where a cardholder formally questions a transaction on their account with their card issuer. This could be for a number of reasons ranging from billing errors to fraudulent activity or dissatisfactory goods/services. If the dispute is recognized as legitimate, the issuer will reverse the payment (otherwise known as a chargeback).

### Available Operations

* [Get](#get) - Get Dispute by ID
* [List](#list) - List of all disputes

## Get

Returns a user's dispute by ID. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

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


    var disputeID string = "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Disputes.Get(ctx, disputeID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Dispute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `disputeID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of dispute                                         | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.GetDisputeResponse](../../pkg/models/operations/getdisputeresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |

## List

Returns the list of disputes. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

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

    ctx := context.Background()
    res, err := s.Disputes.List(ctx, operations.ListDisputesRequest{
        Status: shared.DisputeStatusResponseNeeded.ToPointer(),
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

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.ListDisputesRequest](../../pkg/models/operations/listdisputesrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.ListDisputesResponse](../../pkg/models/operations/listdisputesresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 400-600            | */*                |
