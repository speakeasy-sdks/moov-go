# refund

### Available Operations

* [get](#get) - Get refund details

## get

Get details of a refund for a card transfer

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetRefundRequest(
    refund_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    transfer_id='12fde047-7177-48ff-a1d0-17476360a15d',
)

res = s.refund.get(req)

if res.get_refund is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.GetRefundRequest](../../models/operations/getrefundrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.GetRefundResponse](../../models/operations/getrefundresponse.md)**

