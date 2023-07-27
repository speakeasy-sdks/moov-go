# disputes

## Overview

A [dispute](https://docs.moov.io/guides/money-movement/cards/disputes/) is a situation where a cardholder formally questions a transaction on their account with their card issuer. This could be for a number of reasons ranging from billing errors to fraudulent activity or dissatisfactory goods/services. If the dispute is recognized as legitimate, the issuer will reverse the payment (otherwise known as a chargeback).

### Available Operations

* [get](#get) - Get Dispute by ID
* [list](#list) - List of all disputes

## get

Returns a user's dispute by ID. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetDisputeRequest(
    dispute_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.disputes.get(req)

if res.dispute is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetDisputeRequest](../../models/operations/getdisputerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetDisputeResponse](../../models/operations/getdisputeresponse.md)**


## list

Returns the list of disputes. <br><br> To use this endpoint, you need to specify the `/accounts/{your-account-id}/transfers.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListDisputesRequest(
    count=676243,
    respond_end_date_time='corrupti',
    respond_start_date_time='accusamus',
    skip=272683,
    status=shared.DisputeStatus.RESPONSE_NEEDED,
)

res = s.disputes.list(req)

if res.disputes is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.ListDisputesRequest](../../models/operations/listdisputesrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.ListDisputesResponse](../../models/operations/listdisputesresponse.md)**

