# Enrichment

### Available Operations

* [GetAddress](#getaddress) - Get address suggestions
* [GetAvatar](#getavatar) - Get avatar
* [GetIndustries](#getindustries) - List all industries
* [GetProfile](#getprofile) - Get enriched profile

## GetAddress

Fetch enriched address suggestions. Requires a partial address. 
<br><br> You must specify the `/profile-enrichment.read` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Enrichment.GetAddress(ctx, operations.GetAddressRequest{
        ExcludeStates: moov.String("est"),
        IncludeCities: moov.String("quidem"),
        IncludeStates: moov.String("reprehenderit"),
        IncludeZipcodes: moov.String("facere"),
        MaxResults: moov.Int64(685092),
        PreferCities: moov.String("praesentium"),
        PreferGeolocation: moov.String("mollitia"),
        PreferRatio: moov.Int64(333965),
        PreferStates: moov.String("voluptatem"),
        PreferZipcodes: moov.String("quisquam"),
        Search: "repudiandae",
        Selected: moov.String("quasi"),
        Source: moov.String("atque"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichmentAddresses != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `ctx`                                                                        | [context.Context](https://pkg.go.dev/context#Context)                        | :heavy_check_mark:                                                           | The context to use for the request.                                          |
| `request`                                                                    | [operations.GetAddressRequest](../../models/operations/getaddressrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[*operations.GetAddressResponse](../../models/operations/getaddressresponse.md), error**


## GetAvatar

Get avatar image for an account using a unique ID.
<br><br> To get an avatar, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    uniqueID := "reprehenderit"

    ctx := context.Background()
    res, err := s.Enrichment.GetAvatar(ctx, uniqueID)
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAvatar200ImageWildcardBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `uniqueID`                                                                                               | *string*                                                                                                 | :heavy_check_mark:                                                                                       | Any unique ID associated with an account such as AccountID, RepresentativeID, Routing Number, or User ID |


### Response

**[*operations.GetAvatarResponse](../../models/operations/getavatarresponse.md), error**


## GetIndustries

Returns a list of all industry titles and their corresponding MCC/SIC/NAICS codes. Results are ordered by title.
<br><br> To list industries, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Enrichment.GetIndustries(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.Industries != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetIndustriesResponse](../../models/operations/getindustriesresponse.md), error**


## GetProfile

Fetch enriched profile data. Requires a valid email address. This service is offered in collaboration with Clearbit.
<br><br> To get enriched profile information, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )
    email := "asperiores"

    ctx := context.Background()
    res, err := s.Enrichment.GetProfile(ctx, email)
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichmentProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ctx`                                                    | [context.Context](https://pkg.go.dev/context#Context)    | :heavy_check_mark:                                       | The context to use for the request.                      |
| `email`                                                  | *string*                                                 | :heavy_check_mark:                                       | Valid email address belonging to the profile of interest |


### Response

**[*operations.GetEnrichmentProfileResponse](../../models/operations/getenrichmentprofileresponse.md), error**

