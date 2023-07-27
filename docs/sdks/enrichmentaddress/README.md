# enrichment_address

### Available Operations

* [get](#get) - Get address suggestions

## get

Fetch enriched address suggestions. Requires a partial address. 
<br><br> You must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetEnrichmentAddressRequest(
    exclude_states='odio',
    include_cities='occaecati',
    include_states='commodi',
    include_zipcodes='sapiente',
    max_results=174112,
    prefer_cities='deserunt',
    prefer_geolocation='molestiae',
    prefer_ratio=35362,
    prefer_states='porro',
    prefer_zipcodes='eum',
    search='quas',
    selected='praesentium',
    source='consequuntur',
)

res = s.enrichment_address.get(req)

if res.enrichment_addresses is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.GetEnrichmentAddressRequest](../../models/operations/getenrichmentaddressrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.GetEnrichmentAddressResponse](../../models/operations/getenrichmentaddressresponse.md)**

