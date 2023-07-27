# transactions

## Overview

A transaction is a record of a card's activity on a particular Moov account.

### Available Operations

* [list](#list) - Get account transactions

## list

List issued card transactions associated with a Moov account

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListAccountIssuedCardTransactionsRequest(
    account_id='0421813d-5208-4ece-be25-3b668451c6c6',
    count=927212,
    skip=160393,
    status=shared.IssuedCardTransactionStatus.PENDING,
)

res = s.transactions.list(req)

if res.issued_card_transactions is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                                  | [operations.ListAccountIssuedCardTransactionsRequest](../../models/operations/listaccountissuedcardtransactionsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[operations.ListAccountIssuedCardTransactionsResponse](../../models/operations/listaccountissuedcardtransactionsresponse.md)**

