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
    res, err := s.Capabilities.Delete(ctx, operations.DeleteCapabilityRequest{
        AccountID: "b77f3a41-0067-44eb-b692-80d1ba77a89e",
        CapabilityID: shared.CapabilityIDWallet,
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

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteCapabilityRequest](../../models/operations/deletecapabilityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Capabilities.Get(ctx, operations.GetCapabilityRequest{
        AccountID: "f737ae42-03ce-45e6-a95d-8a0d446ce2af",
        CapabilityID: shared.CapabilityIDCollectFunds,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Capability != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCapabilityRequest](../../models/operations/getcapabilityrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


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
    res, err := s.Capabilities.List(ctx, operations.ListCapabilitiesRequest{
        AccountID: "a73cf3be-453f-4870-b326-b5a73429cdb1",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Capabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListCapabilitiesRequest](../../models/operations/listcapabilitiesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


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
    res, err := s.Capabilities.Request(ctx, operations.PostCapabilityRequest{
        AddCapabilityRequest: shared.AddCapabilityRequest{
            Capabilities: []shared.CapabilityID{
                shared.CapabilityIDCollectFunds,
                shared.CapabilityIDSendFunds,
                shared.CapabilityIDTransfers,
            },
        },
        AccountID: "2bb679d2-3227-415b-b0cb-b1e31b8b90f3",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Capabilities != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.PostCapabilityRequest](../../models/operations/postcapabilityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.PostCapabilityResponse](../../models/operations/postcapabilityresponse.md), error**

