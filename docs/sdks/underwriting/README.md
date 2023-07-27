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
    accountID := "efb0b348-96c3-4ca5-acfb-e2fd57075779"

    ctx := context.Background()
    res, err := s.Underwriting.Get(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Underwriting != nil {
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

**[*operations.GetUnderwritingResponse](../../models/operations/getunderwritingresponse.md), error**


## Update

Update the account's underwriting by passing new values for one or more of the fields. <br><br> To update an account's underwriting details, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

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
    underwritingRequest := shared.UnderwritingRequest{
        AverageMonthlyTransactionVolume: moov.Int64(250000),
        AverageTransactionSize: moov.Int64(10000),
        MaxTransactionSize: moov.Int64(50000),
    }
    accountID := "29177dea-c646-4ecb-9734-09e3eb1e5a2b"

    ctx := context.Background()
    res, err := s.Underwriting.Update(ctx, underwritingRequest, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Underwriting != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                | Type                                                                     | Required                                                                 | Description                                                              |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `ctx`                                                                    | [context.Context](https://pkg.go.dev/context#Context)                    | :heavy_check_mark:                                                       | The context to use for the request.                                      |
| `underwritingRequest`                                                    | [shared.UnderwritingRequest](../../models/shared/underwritingrequest.md) | :heavy_check_mark:                                                       | N/A                                                                      |
| `accountID`                                                              | *string*                                                                 | :heavy_check_mark:                                                       | ID of the account                                                        |


### Response

**[*operations.UpdateUnderwritingResponse](../../models/operations/updateunderwritingresponse.md), error**

