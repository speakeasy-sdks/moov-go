# institutions

## Overview

Lookup ACH and wire participating financial institutions. We recommend using this endpoint when an end-user enters a routing number to confirm their bank or credit union.

### Available Operations

* [search](#search) - Search institutions

## search

Search for institutions by their routing number or name. <br><br> To use this endpoint, you need to specify the `/fed.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.SearchInstitutionsRequest(
    limit=499,
    name='Billie Keebler',
    rail=shared.Rail.WIRE,
    routing_number='cum',
    state='commodi',
)

res = s.institutions.search(req)

if res.financial_institutions is not None:
    # handle response
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `request`                                                                                    | [operations.SearchInstitutionsRequest](../../models/operations/searchinstitutionsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[operations.SearchInstitutionsResponse](../../models/operations/searchinstitutionsresponse.md)**

