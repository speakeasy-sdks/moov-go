# enrichment_profile

### Available Operations

* [get](#get) - Get enriched profile

## get

Fetch enriched profile data. Requires a valid email address. This service is offered in collaboration with Clearbit.
<br><br> To get enriched profile information, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetEnrichmentProfileRequest(
    email='Buddy.Parker@yahoo.com',
)

res = s.enrichment_profile.get(req)

if res.enrichment_profile is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.GetEnrichmentProfileRequest](../../models/operations/getenrichmentprofilerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.GetEnrichmentProfileResponse](../../models/operations/getenrichmentprofileresponse.md)**

