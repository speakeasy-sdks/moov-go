# Analytics
(*Analytics*)

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
        Every: types.MustTimeFromString("2021-03-16T03:51:45.617Z"),
        From: types.MustTimeFromString("2023-04-16T05:17:21.118Z"),
        To: types.MustTimeFromString("2021-02-07T13:38:18.890Z"),
        Tz: types.MustTimeFromString("2023-01-02T16:17:39.869Z"),
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
        Every: types.MustTimeFromString("2021-07-25T11:48:02.172Z"),
        From: types.MustTimeFromString("2022-05-15T09:50:33.891Z"),
        To: types.MustTimeFromString("2023-11-15T15:56:35.850Z"),
        Tz: types.MustTimeFromString("2021-11-26T21:59:13.376Z"),
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
        Count: moovgo.Int64(672967),
        Every: moovgo.String("America/Lima"),
        From: types.MustTimeFromString("2023-08-17T01:22:07.057Z"),
        To: types.MustTimeFromString("2022-10-30T11:03:17.489Z"),
        Tz: moovgo.String("America/New_York"),
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
        Count: moovgo.Int64(16506),
        Every: moovgo.String("Pacific/Midway"),
        From: types.MustTimeFromString("2023-08-25T00:26:28.999Z"),
        To: types.MustTimeFromString("2023-10-03T14:23:09.092Z"),
        Tz: moovgo.String("Asia/Taipei"),
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
        Every: types.MustTimeFromString("2023-09-24T04:57:26.458Z"),
        From: types.MustTimeFromString("2021-07-05T09:11:47.235Z"),
        To: types.MustTimeFromString("2021-12-02T22:28:20.094Z"),
        Tz: types.MustTimeFromString("2022-10-22T17:48:41.788Z"),
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

