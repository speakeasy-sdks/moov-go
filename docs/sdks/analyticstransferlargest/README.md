# analytics_transfer_largest

### Available Operations

* [post](#post) - Return the largest number of transfers

## post

Return the largest number of transfers

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
    count=146441,
    every='dolorum',
    from_=dateutil.parser.isoparse('2022-06-17T21:27:36.672Z'),
    to=dateutil.parser.isoparse('2021-07-13T07:24:02.478Z'),
    tz='labore',
)

res = s.analytics_transfer_largest.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[operations.PostAnalyticsTransferLargestResponse](../../models/operations/postanalyticstransferlargestresponse.md)**

