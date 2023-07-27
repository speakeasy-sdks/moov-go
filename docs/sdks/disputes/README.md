# Disputes

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
    res, err := s.Disputes.Get(ctx, operations.GetDisputeRequest{
        DisputeID: "ec7e1848-dc80-4ab0-8827-dd7fc0737b43",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Dispute != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetDisputeRequest](../../models/operations/getdisputerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetDisputeResponse](../../models/operations/getdisputeresponse.md), error**


## List

Returns the list of disputes. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

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
    res, err := s.Disputes.List(ctx, operations.ListDisputesRequest{
        Count: moov.Int64(676243),
        RespondEndDateTime: moov.String("corrupti"),
        RespondStartDateTime: moov.String("accusamus"),
        Skip: moov.Int64(272683),
        Status: shared.DisputeStatusResponseNeeded.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Disputes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [operations.ListDisputesRequest](../../models/operations/listdisputesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.ListDisputesResponse](../../models/operations/listdisputesresponse.md), error**

