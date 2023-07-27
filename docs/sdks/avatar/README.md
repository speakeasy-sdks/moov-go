# avatar

### Available Operations

* [get](#get) - Get avatar

## get

Get avatar image for an account using a unique ID.
<br><br> To get an avatar, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAvatarRequest(
    unique_id='rerum',
)

res = s.avatar.get(req)

if res.get_avatar_200_image_wildcard_binary_string is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.GetAvatarRequest](../../models/operations/getavatarrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.GetAvatarResponse](../../models/operations/getavatarresponse.md)**

