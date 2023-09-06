# Institutions

## Overview

Lookup ACH and wire participating financial institutions. We recommend using this endpoint when an end-user enters a routing number to confirm their bank or credit union.

### Available Operations

* [Search](#search) - Search institutions

## Search

Search for institutions by their routing number or name. <br><br> To use this endpoint, you need to specify the `/fed.read` scope.

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

    ctx := context.Background()
    res, err := s.Institutions.Search(ctx, operations.SearchInstitutionRequest{
        Limit: moov.Int64(499),
        Name: moov.String("Kelley Dibbert"),
        Rail: shared.RailWire,
        RoutingNumber: moov.String("rem"),
        State: moov.String("aliquam"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FinancialInstitutions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `ctx`                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                      | :heavy_check_mark:                                                                         | The context to use for the request.                                                        |
| `request`                                                                                  | [operations.SearchInstitutionRequest](../../models/operations/searchinstitutionrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[*operations.SearchInstitutionResponse](../../models/operations/searchinstitutionresponse.md), error**

