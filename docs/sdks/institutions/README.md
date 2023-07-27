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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Institutions.Search(ctx, operations.SearchInstitutionRequest{
        Limit: petstore.Int64(499),
        Name: petstore.String("Arnold Ferry"),
        Rail: shared.RailAch,
        RoutingNumber: petstore.String("fugit"),
        State: petstore.String("id"),
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

