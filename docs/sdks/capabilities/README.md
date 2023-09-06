# Capabilities

## Overview

Capabilities determine what a Moov account can do. Each capability has specific [verification requirements](https://docs.moov.io/guides/accounts/identity-verification/), depending on risk and compliance standards associated with different account activities. For example, there are more information requirements for a business that wants to charge other accounts than for an individual who simply wants to receive funds. When you request a capability, we list the information requirements for that capability. Once you submit the required information, we need to verify the data. Because of this, a requested capability may not immediately become active. For more detailed information on capabilities and capability IDs, read our [capabilities guide](https://docs.moov.io/guides/accounts/capabilities/).

### Available Operations

* [Delete](#delete) - Disable a capability for an account
* [Get](#get) - Get capability for account
* [List](#list) - List capabilities for account
* [Request](#request) - Request capabilities

## Delete

Disable a specific capability that an account has requested. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.write` scope.

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
    accountID := "5e6a95d8-a0d4-446c-a2af-7a73cf3be453"
    capabilityID := shared.CapabilityIDCardIssuing

    ctx := context.Background()
    res, err := s.Capabilities.Delete(ctx, accountID, capabilityID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `accountID`                                                | *string*                                                   | :heavy_check_mark:                                         | ID of the account                                          |
| `capabilityID`                                             | [shared.CapabilityID](../../models/shared/capabilityid.md) | :heavy_check_mark:                                         | The requested capability identifier                        |


### Response

**[*operations.DeleteCapabilityResponse](../../models/operations/deletecapabilityresponse.md), error**


## Get

Retrieve a specific capability that an account has requested. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.read` scope.

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
    accountID := "870b326b-5a73-4429-8db1-a8422bb679d2"
    capabilityID := shared.CapabilityIDSendFunds

    ctx := context.Background()
    res, err := s.Capabilities.Get(ctx, accountID, capabilityID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Capability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                  | Type                                                       | Required                                                   | Description                                                |
| ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- | ---------------------------------------------------------- |
| `ctx`                                                      | [context.Context](https://pkg.go.dev/context#Context)      | :heavy_check_mark:                                         | The context to use for the request.                        |
| `accountID`                                                | *string*                                                   | :heavy_check_mark:                                         | ID of the account                                          |
| `capabilityID`                                             | [shared.CapabilityID](../../models/shared/capabilityid.md) | :heavy_check_mark:                                         | The requested capability identifier                        |


### Response

**[*operations.GetCapabilityResponse](../../models/operations/getcapabilityresponse.md), error**


## List

Retrieve all the capabilities an account has requested. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/capabilities.read` scope.

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
    accountID := "22715bf0-cbb1-4e31-b8b9-0f3443a1108e"

    ctx := context.Background()
    res, err := s.Capabilities.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Capabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


### Response

**[*operations.ListCapabilitiesResponse](../../models/operations/listcapabilitiesresponse.md), error**


## Request

Request capabilities for a specific account. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.write` scope.

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
    addCapabilityRequest := shared.AddCapabilityRequest{
        Capabilities: []shared.CapabilityID{
            shared.CapabilityIDTransfers,
        },
    }
    accountID := "adcf4b92-1879-4fce-953f-73ef7fbc7abd"

    ctx := context.Background()
    res, err := s.Capabilities.Request(ctx, addCapabilityRequest, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Capabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `addCapabilityRequest`                                                     | [shared.AddCapabilityRequest](../../models/shared/addcapabilityrequest.md) | :heavy_check_mark:                                                         | N/A                                                                        |
| `accountID`                                                                | *string*                                                                   | :heavy_check_mark:                                                         | ID of the account                                                          |


### Response

**[*operations.PostCapabilityResponse](../../models/operations/postcapabilityresponse.md), error**

