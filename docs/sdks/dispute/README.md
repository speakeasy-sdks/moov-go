# dispute

### Available Operations

* [get](#get) - Get Dispute by ID

## get

Returns a user's dispute by ID. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetDisputeRequest(
    dispute_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.dispute.get(req)

if res.dispute is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetDisputeRequest](../../models/operations/getdisputerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetDisputeResponse](../../models/operations/getdisputeresponse.md)**

