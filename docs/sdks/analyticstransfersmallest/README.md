# analytics_transfer_smallest

### Available Operations

* [post](#post) - Return the smallest number of transfers

## post

Return the smallest number of transfers

### Example Usage

```python
import petstore
import dateutil.parser
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.LimitedTimeRange(
    count=962189,
    every='eum',
    from_=dateutil.parser.isoparse('2022-03-31T00:30:19.135Z'),
    to=dateutil.parser.isoparse('2022-03-17T20:21:28.792Z'),
    tz='provident',
)

res = s.analytics_transfer_smallest.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[operations.PostAnalyticsTransferSmallestResponse](../../models/operations/postanalyticstransfersmallestresponse.md)**

