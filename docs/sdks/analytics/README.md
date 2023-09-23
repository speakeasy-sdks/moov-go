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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountAccountsCreated(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-12-07T10:53:17.121Z"),
        From: types.MustTimeFromString("2021-01-23T15:47:23.464Z"),
        To: types.MustTimeFromString("2022-05-11T16:07:41.164Z"),
        Tz: types.MustTimeFromString("2022-10-12T05:44:19.260Z"),
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountTransferStatuses(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-07-30T07:35:03.817Z"),
        From: types.MustTimeFromString("2022-11-26T12:00:10.052Z"),
        To: types.MustTimeFromString("2022-01-06T19:47:24.047Z"),
        Tz: types.MustTimeFromString("2022-03-21T22:14:24.691Z"),
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.LargestTransfer(ctx, shared.LimitedTimeRange{
        Count: moovgo.Int64(806194),
        Every: moovgo.String("deleniti"),
        From: types.MustTimeFromString("2022-02-08T00:19:59.821Z"),
        To: types.MustTimeFromString("2022-11-25T15:46:28.441Z"),
        Tz: moovgo.String("repudiandae"),
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SmallestTransfer(ctx, shared.LimitedTimeRange{
        Count: moovgo.Int64(352312),
        Every: moovgo.String("expedita"),
        From: types.MustTimeFromString("2022-01-01T10:06:00.916Z"),
        To: types.MustTimeFromString("2022-07-21T08:29:53.942Z"),
        Tz: moovgo.String("saepe"),
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
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/types"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SumTransfers(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-11-20T20:56:20.791Z"),
        From: types.MustTimeFromString("2022-06-29T11:09:23.468Z"),
        To: types.MustTimeFromString("2022-09-01T04:49:52.515Z"),
        Tz: types.MustTimeFromString("2022-03-22T15:30:46.869Z"),
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

