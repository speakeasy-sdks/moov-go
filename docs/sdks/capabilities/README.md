# capabilities

## Overview

Capabilities determine what a Moov account can do. Each capability has specific [verification requirements](https://docs.moov.io/guides/accounts/identity-verification/), depending on risk and compliance standards associated with different account activities. For example, there are more information requirements for a business that wants to charge other accounts than for an individual who simply wants to receive funds. When you request a capability, we list the information requirements for that capability. Once you submit the required information, we need to verify the data. Because of this, a requested capability may not immediately become active. For more detailed information on capabilities and capability IDs, read our [capabilities guide](https://docs.moov.io/guides/accounts/capabilities/).

### Available Operations

* [delete](#delete) - Disable a capability for an account
* [get](#get) - Get capability for account
* [list](#list) - List capabilities for account
* [request](#request) - Request capabilities

## delete

Disable a specific capability that an account has requested. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteCapabilityRequest(
    account_id='7f3a4100-674e-4bf6-9280-d1ba77a89ebf',
    capability_id=shared.CapabilityID.COLLECT_FUNDS,
)

res = s.capabilities.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.DeleteCapabilityRequest](../../models/operations/deletecapabilityrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.DeleteCapabilityResponse](../../models/operations/deletecapabilityresponse.md)**


## get

Retrieve a specific capability that an account has requested. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetCapabilityRequest(
    account_id='37ae4203-ce5e-46a9-9d8a-0d446ce2af7a',
    capability_id=shared.CapabilityID.COLLECT_FUNDS,
)

res = s.capabilities.get(req)

if res.capability is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.GetCapabilityRequest](../../models/operations/getcapabilityrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.GetCapabilityResponse](../../models/operations/getcapabilityresponse.md)**


## list

Retrieve all the capabilities an account has requested. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/capabilities.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListCapabilitiesRequest(
    account_id='3cf3be45-3f87-40b3-a6b5-a73429cdb1a8',
)

res = s.capabilities.list(req)

if res.capabilities is not None:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.ListCapabilitiesRequest](../../models/operations/listcapabilitiesrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.ListCapabilitiesResponse](../../models/operations/listcapabilitiesresponse.md)**


## request

Request capabilities for a specific account. <br><br> To use this endpoint, you must specify the `/accounts/{accountID}/capabilities.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostCapabilityRequest(
    add_capability_request=shared.AddCapabilityRequest(
        capabilities=[
            shared.CapabilityID.TRANSFERS,
            shared.CapabilityID.TRANSFERS,
        ],
    ),
    account_id='bb679d23-2271-45bf-8cbb-1e31b8b90f34',
)

res = s.capabilities.request(req)

if res.capabilities is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.PostCapabilityRequest](../../models/operations/postcapabilityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.PostCapabilityResponse](../../models/operations/postcapabilityresponse.md)**

