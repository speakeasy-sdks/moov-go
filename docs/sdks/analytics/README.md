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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountAccountsCreated(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2021-10-04T09:10:06.610Z"),
        From: types.MustTimeFromString("2022-10-22T18:12:12.288Z"),
        To: types.MustTimeFromString("2022-04-23T05:56:38.936Z"),
        Tz: types.MustTimeFromString("2021-11-13T09:08:33.009Z"),
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.CountTransferStatuses(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-03-16T09:33:50.291Z"),
        From: types.MustTimeFromString("2021-12-15T00:41:38.329Z"),
        To: types.MustTimeFromString("2022-09-20T03:14:35.704Z"),
        Tz: types.MustTimeFromString("2021-04-10T08:07:33.561Z"),
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.LargestTransfer(ctx, shared.LimitedTimeRange{
        Count: petstore.Int64(83112),
        Every: petstore.String("itaque"),
        From: types.MustTimeFromString("2022-09-06T17:20:08.756Z"),
        To: types.MustTimeFromString("2022-05-02T09:29:06.042Z"),
        Tz: petstore.String("quibusdam"),
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SmallestTransfer(ctx, shared.LimitedTimeRange{
        Count: petstore.Int64(131797),
        Every: petstore.String("deserunt"),
        From: types.MustTimeFromString("2021-04-26T18:54:54.344Z"),
        To: types.MustTimeFromString("2022-09-26T08:57:48.803Z"),
        Tz: petstore.String("qui"),
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/types"
)

func main() {
    s := petstore.New(
        petstore.WithSecurity(shared.Security{
            AccessToken: petstore.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Analytics.SumTransfers(ctx, shared.TimeRange{
        Every: types.MustTimeFromString("2022-05-31T22:08:47.731Z"),
        From: types.MustTimeFromString("2022-12-17T07:42:55.593Z"),
        To: types.MustTimeFromString("2022-03-04T10:29:07.095Z"),
        Tz: types.MustTimeFromString("2022-12-30T06:52:02.282Z"),
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

