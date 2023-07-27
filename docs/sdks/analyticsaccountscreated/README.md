# analytics_accounts_created

### Available Operations

* [post](#post) - Count the number of profiles created by an individual or business

## post

Count the number of profiles created by an individual or business

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
    every=dateutil.parser.isoparse('2022-05-31T22:08:47.731Z'),
    from_=dateutil.parser.isoparse('2022-12-17T07:42:55.593Z'),
    to=dateutil.parser.isoparse('2022-03-04T10:29:07.095Z'),
    tz=dateutil.parser.isoparse('2022-12-30T06:52:02.282Z'),
)

res = s.analytics_accounts_created.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsAccountsCreatedResponse](../../models/operations/postanalyticsaccountscreatedresponse.md)**

