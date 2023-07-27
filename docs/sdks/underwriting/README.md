# Underwriting

## Overview

[Underwriting](https://docs.moov.io/guides/accounts/underwriting) is a tool Moov uses to understand a business’s profile before allowing them to collect funds on our platform. This profile includes information like a description of the company or the merchant’s business model, the industry they operate in, and transaction volume. Through underwriting, we can understand and prevent unnecessary financial risk for Moov and those transacting on our platform. Note that underwriting can be instant, but in some cases make take around 2 business days before approval.


### Available Operations

* [Get](#get) - Retrieve underwriting details
* [Update](#update) - Update underwriting details

## Get

Retrieve underwriting associated with a given Moov account. <br><br> To get an account's underwriting details, you'll need to specify the `/accounts/{accountID}/underwriting.read` scope.

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
    res, err := s.Underwriting.Get(ctx, operations.GetUnderwritingRequest{
        AccountID: "bae0be2d-7822-459e-bea4-b5197f92443d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Underwriting != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.GetUnderwritingRequest](../../models/operations/getunderwritingrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.GetUnderwritingResponse](../../models/operations/getunderwritingresponse.md), error**


## Update

Update the account's underwriting by passing new values for one or more of the fields. <br><br> To update an account's underwriting details, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

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
    res, err := s.Underwriting.Update(ctx, operations.UpdateUnderwritingRequest{
        UnderwritingRequest: shared.UnderwritingRequest{
            AverageMonthlyTransactionVolume: moov.Int64(250000),
            AverageTransactionSize: moov.Int64(10000),
            MaxTransactionSize: moov.Int64(50000),
        },
        AccountID: "a7ce52b8-95c5-437c-a454-efb0b34896c3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Underwriting != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `ctx`                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                        | :heavy_check_mark:                                                                           | The context to use for the request.                                                          |
| `request`                                                                                    | [operations.UpdateUnderwritingRequest](../../models/operations/updateunderwritingrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[*operations.UpdateUnderwritingResponse](../../models/operations/updateunderwritingresponse.md), error**

