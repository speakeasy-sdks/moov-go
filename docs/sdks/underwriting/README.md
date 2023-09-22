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
    accountID := "6ecb5734-09e3-4eb1-a5a2-b12eb07f116d"

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
    underwritingRequest := shared.UnderwritingRequest{
        AverageMonthlyTransactionVolume: moovgo.Int64(250000),
        AverageTransactionSize: moovgo.Int64(10000),
        MaxTransactionSize: moovgo.Int64(50000),
    }
    accountID := "b99545fc-95fa-4889-b0e1-89dbb30fcb33"

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

