# analytics_transfer_statuses

### Available Operations

* [post](#post) - Count the transfer statuses

## post

Count the transfer statuses

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
    every=dateutil.parser.isoparse('2021-04-14T09:13:11.675Z'),
    from_=dateutil.parser.isoparse('2022-07-22T03:36:34.615Z'),
    to=dateutil.parser.isoparse('2020-02-22T17:45:21.686Z'),
    tz=dateutil.parser.isoparse('2022-02-08T14:21:47.573Z'),
)

res = s.analytics_transfer_statuses.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsTransferStatusesResponse](../../models/operations/postanalyticstransferstatusesresponse.md)**

