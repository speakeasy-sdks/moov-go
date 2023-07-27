# analytics_transfer_sum

### Available Operations

* [post](#post) - Sum all transfers across intervals

## post

Sum all transfers across intervals

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

req = shared.TimeRange(
    every=dateutil.parser.isoparse('2022-02-26T01:27:36.152Z'),
    from_=dateutil.parser.isoparse('2020-11-26T01:41:04.216Z'),
    to=dateutil.parser.isoparse('2022-09-14T10:27:07.590Z'),
    tz=dateutil.parser.isoparse('2020-07-23T21:23:35.691Z'),
)

res = s.analytics_transfer_sum.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsTransferSumResponse](../../models/operations/postanalyticstransfersumresponse.md)**

