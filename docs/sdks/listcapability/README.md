# list_capability

### Available Operations

* [get](#get) - List capabilities for account

## get

Retrieve all the capabilities an account has requested. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/capabilities.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetListCapabilitiesRequest(
    account_id='79f9dfe0-ab7d-4a8a-90ce-187f86bc173d',
)

res = s.list_capability.get(req)

if res.capabilities is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetListCapabilitiesRequest](../../models/operations/getlistcapabilitiesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetListCapabilitiesResponse](../../models/operations/getlistcapabilitiesresponse.md)**

