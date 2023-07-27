# transfers

## Overview

A [transfer](https://docs.moov.io/guides/money-movement/) is the movement of money between Moov accounts, from source to destination. Provided you have linked a bank account which has been verified, you can initiate a transfer to another Moov account. All you need to do is note a [paymentMethod](#tag/Payment-methods), the $ amount of the transfer, and a brief description. With Moov, you can also implement transfer groups, allowing you to associate multiple transfers together and run them sequentially. To learn more, read our guide on [transfer groups](https://docs.moov.io/guides/money-movement/transfer-groups/#transfer-statuses).

### Available Operations

* [get](#get) - Get a list of refunds for a card transfer

## get

Get a list of refunds for a card transfer

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetRefundsRequest(
    transfer_id='5197f924-43da-47ce-92b8-95c537c6454e',
)

res = s.transfers.get(req)

if res.get_refunds is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetRefundsRequest](../../models/operations/getrefundsrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetRefundsResponse](../../models/operations/getrefundsresponse.md)**

