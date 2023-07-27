# analytics

## Overview

You can retrieve helpful at-a-glance information about your account by getting metrics on categories such as new accounts, transfer counts, and transfer volume over different time periods. To use this endpoint, you must specify the `/analytics.read` scope.

### Available Operations

* [count_accounts_created](#count_accounts_created) - Count the number of profiles created by an individual or business
* [count_transfer_statuses](#count_transfer_statuses) - Count the transfer statuses
* [largest_transfer](#largest_transfer) - Return the largest number of transfers
* [smallest_transfer](#smallest_transfer) - Return the smallest number of transfers
* [sum_transfers](#sum_transfers) - Sum all transfers across intervals

## count_accounts_created

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
    every=dateutil.parser.isoparse('2021-10-04T09:10:06.610Z'),
    from_=dateutil.parser.isoparse('2022-10-22T18:12:12.288Z'),
    to=dateutil.parser.isoparse('2022-04-23T05:56:38.936Z'),
    tz=dateutil.parser.isoparse('2021-11-13T09:08:33.009Z'),
)

res = s.analytics.count_accounts_created(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsAccountsCreatedResponse](../../models/operations/postanalyticsaccountscreatedresponse.md)**


## count_transfer_statuses

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
    every=dateutil.parser.isoparse('2022-03-16T09:33:50.291Z'),
    from_=dateutil.parser.isoparse('2021-12-15T00:41:38.329Z'),
    to=dateutil.parser.isoparse('2022-09-20T03:14:35.704Z'),
    tz=dateutil.parser.isoparse('2021-04-10T08:07:33.561Z'),
)

res = s.analytics.count_transfer_statuses(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsTransferStatusesResponse](../../models/operations/postanalyticstransferstatusesresponse.md)**


## largest_transfer

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
    count=83112,
    every='itaque',
    from_=dateutil.parser.isoparse('2022-09-06T17:20:08.756Z'),
    to=dateutil.parser.isoparse('2022-05-02T09:29:06.042Z'),
    tz='quibusdam',
)

res = s.analytics.largest_transfer(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[operations.PostAnalyticsTransferLargestResponse](../../models/operations/postanalyticstransferlargestresponse.md)**


## smallest_transfer

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
    count=131797,
    every='deserunt',
    from_=dateutil.parser.isoparse('2021-04-26T18:54:54.344Z'),
    to=dateutil.parser.isoparse('2022-09-26T08:57:48.803Z'),
    tz='qui',
)

res = s.analytics.smallest_transfer(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[operations.PostAnalyticsTransferSmallestResponse](../../models/operations/postanalyticstransfersmallestresponse.md)**


## sum_transfers

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
    every=dateutil.parser.isoparse('2022-05-31T22:08:47.731Z'),
    from_=dateutil.parser.isoparse('2022-12-17T07:42:55.593Z'),
    to=dateutil.parser.isoparse('2022-03-04T10:29:07.095Z'),
    tz=dateutil.parser.isoparse('2022-12-30T06:52:02.282Z'),
)

res = s.analytics.sum_transfers(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                            | Type                                                 | Required                                             | Description                                          |
| ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- | ---------------------------------------------------- |
| `request`                                            | [shared.TimeRange](../../models/shared/timerange.md) | :heavy_check_mark:                                   | The request object to use for the request.           |


### Response

**[operations.PostAnalyticsTransferSumResponse](../../models/operations/postanalyticstransfersumresponse.md)**

