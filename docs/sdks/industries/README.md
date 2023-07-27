# industries

## Overview

Information about industries and their merchant codes.

### Available Operations

* [list](#list) - List all industries

## list

Returns a list of all industry titles and their corresponding MCC/SIC/NAICS codes. Results are ordered by title.
<br><br> To list industries, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)


res = s.industries.list()

if res.industries is not None:
    # handle response
```


### Response

**[operations.ListIndustriesResponse](../../models/operations/listindustriesresponse.md)**

