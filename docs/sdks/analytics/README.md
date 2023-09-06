# Analytics

## Overview

You can retrieve helpful at-a-glance information about your account by getting metrics on categories such as new accounts, transfer counts, and transfer volume over different time periods. To use this endpoint, you must specify the `/analytics.read` scope.

### Available Operations

* [CountAccountsCreated](#countaccountscreated) - Count the number of profiles created by an individual or business
* [CountTransferStatuses](#counttransferstatuses) - Count the transfer statuses
* [LargestTransfer](#largesttransfer) - Return the largest number of transfers
* [SmallestTransfer](#smallesttransfer) - Return the smallest number of transfers
* [SumTransfers](#sumtransfers) - Sum all transfers across intervals

## CountAccountsCreated

Count the number of profiles created by an individual or business

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountAccountsCreated(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-10-02T04:55:20.234Z"),
        From: types.MustTimeFromString("2021-04-10T02:35:06.342Z"),
        To: types.MustTimeFromString("2022-05-29T21:42:45.399Z"),
        Tz: types.MustTimeFromString("2021-04-14T09:13:11.675Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.TimeRange](../../models/shared/timerange.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAnalyticsAccountsCreatedResponse](../../models/operations/postanalyticsaccountscreatedresponse.md), error**


## CountTransferStatuses

Count the transfer statuses

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountTransferStatuses(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-07-22T03:36:34.615Z"),
        From: types.MustTimeFromString("2020-02-22T17:45:21.686Z"),
        To: types.MustTimeFromString("2022-02-08T14:21:47.573Z"),
        Tz: types.MustTimeFromString("2022-02-26T01:27:36.152Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.TimeRange](../../models/shared/timerange.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAnalyticsTransferStatusesResponse](../../models/operations/postanalyticstransferstatusesresponse.md), error**


## LargestTransfer

Return the largest number of transfers

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.LargestTransfer(ctx, shared.LimitedTimeRange{
        Count: moov.Int64(978571),
        Every: moov.String("rerum"),
        From: types.MustTimeFromString("2022-09-14T10:27:07.590Z"),
        To: types.MustTimeFromString("2020-07-23T21:23:35.691Z"),
        Tz: moov.String("ea"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostAnalyticsTransferLargestResponse](../../models/operations/postanalyticstransferlargestresponse.md), error**


## SmallestTransfer

Return the smallest number of transfers

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SmallestTransfer(ctx, shared.LimitedTimeRange{
        Count: moov.Int64(396506),
        Every: moov.String("laborum"),
        From: types.MustTimeFromString("2022-04-02T11:21:13.260Z"),
        To: types.MustTimeFromString("2022-05-17T08:24:52.669Z"),
        Tz: moov.String("accusamus"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `request`                                                          | [shared.LimitedTimeRange](../../models/shared/limitedtimerange.md) | :heavy_check_mark:                                                 | The request object to use for the request.                         |


### Response

**[*operations.PostAnalyticsTransferSmallestResponse](../../models/operations/postanalyticstransfersmallestresponse.md), error**


## SumTransfers

Sum all transfers across intervals

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SumTransfers(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2020-12-03T16:16:10.882Z"),
        From: types.MustTimeFromString("2021-07-20T13:32:41.687Z"),
        To: types.MustTimeFromString("2021-12-31T00:47:48.012Z"),
        Tz: types.MustTimeFromString("2021-02-02T01:24:52.629Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `request`                                             | [shared.TimeRange](../../models/shared/timerange.md)  | :heavy_check_mark:                                    | The request object to use for the request.            |


### Response

**[*operations.PostAnalyticsTransferSumResponse](../../models/operations/postanalyticstransfersumresponse.md), error**

