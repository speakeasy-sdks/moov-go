# file_details

### Available Operations

* [get](#get) - Get File Details

## get

Retrieve file details associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetFileDetailsRequest(
    account_id='b3c20c4f-3789-4fd8-b1f9-9dd2efd121aa',
    file_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.file_details.get(req)

if res.file is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.GetFileDetailsRequest](../../models/operations/getfiledetailsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.GetFileDetailsResponse](../../models/operations/getfiledetailsresponse.md)**

