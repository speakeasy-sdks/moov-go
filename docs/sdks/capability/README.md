# capability

### Available Operations

* [delete](#delete) - Disable a capability for an account
* [get](#get) - Get capability for account
* [post](#post) - Request capabilities

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
    account_id='3f870b32-6b5a-4734-a9cd-b1a8422bb679',
    capability_id=shared.CapabilityID.CARD_ISSUING,
)

res = s.capability.delete(req)

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
    account_id='2322715b-f0cb-4b1e-b1b8-b90f3443a110',
    capability_id=shared.CapabilityID.COLLECT_FUNDS,
)

res = s.capability.get(req)

if res.capability is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.GetCapabilityRequest](../../models/operations/getcapabilityrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.GetCapabilityResponse](../../models/operations/getcapabilityresponse.md)**


## post

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
            shared.CapabilityID.WALLET,
            shared.CapabilityID.CARD_ISSUING,
            shared.CapabilityID.WALLET,
        ],
    ),
    account_id='f4b92187-9fce-4953-b73e-f7fbc7abd74d',
)

res = s.capability.post(req)

if res.capabilities is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.PostCapabilityRequest](../../models/operations/postcapabilityrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.PostCapabilityResponse](../../models/operations/postcapabilityresponse.md)**

