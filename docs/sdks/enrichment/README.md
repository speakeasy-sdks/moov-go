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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Enrichment.GetAddress(ctx, operations.GetAddressRequest{
        ExcludeStates: moov.String("atque"),
        IncludeCities: moov.String("fugit"),
        IncludeStates: moov.String("ut"),
        IncludeZipcodes: moov.String("fugiat"),
        MaxResults: moov.Int64(30235),
        PreferCities: moov.String("culpa"),
        PreferGeolocation: moov.String("expedita"),
        PreferRatio: moov.Int64(299643),
        PreferStates: moov.String("consequatur"),
        PreferZipcodes: moov.String("esse"),
        Search: "ipsam",
        Selected: moov.String("sit"),
        Source: moov.String("voluptatum"),
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Enrichment.GetAvatar(ctx, operations.GetAvatarRequest{
        UniqueID: "quas",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetAvatar200ImageWildcardBinaryString != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [operations.GetAvatarRequest](../../models/operations/getavatarrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


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
	"openapi"
	"openapi/pkg/models/shared"
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
	"openapi"
	"openapi/pkg/models/shared"
	"openapi/pkg/models/operations"
)

func main() {
    s := moov.New(
        moov.WithSecurity(shared.Security{
            AccessToken: moov.String(""),
        }),
    )

    ctx := context.Background()
    res, err := s.Enrichment.GetProfile(ctx, operations.GetEnrichmentProfileRequest{
        Email: "Foster.Borer@hotmail.com",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.EnrichmentProfile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetEnrichmentProfileRequest](../../models/operations/getenrichmentprofilerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetEnrichmentProfileResponse](../../models/operations/getenrichmentprofileresponse.md), error**

