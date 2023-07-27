# network_i_ds

### Available Operations

* [get](#get) - Get network IDs of an account

## get

Retrieves network IDs for the account of the specified ID. 
<br><br> To get an account, you will need to specify the `/accounts/{accountID}/profile.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetNetworkIDsRequest(
    account_id='2b6e3ab8-845f-4059-ba60-ff2a54a31e94',
)

res = s.network_i_ds.get(req)

if res.network_i_ds is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.GetNetworkIDsRequest](../../models/operations/getnetworkidsrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.GetNetworkIDsResponse](../../models/operations/getnetworkidsresponse.md)**

