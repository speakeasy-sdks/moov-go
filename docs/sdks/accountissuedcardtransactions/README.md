# account_issued_card_transactions

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
    account_id='71b5e6e1-3b99-4d48-8e1e-91e450ad2abd',
    count=289406,
    skip=264730,
    status=shared.IssuedCardTransactionStatus.PENDING,
)

res = s.account_issued_card_transactions.list(req)

if res.issued_card_transactions is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                                  | Type                                                                                                                       | Required                                                                                                                   | Description                                                                                                                |
| -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                                  | [operations.ListAccountIssuedCardTransactionsRequest](../../models/operations/listaccountissuedcardtransactionsrequest.md) | :heavy_check_mark:                                                                                                         | The request object to use for the request.                                                                                 |


### Response

**[operations.ListAccountIssuedCardTransactionsResponse](../../models/operations/listaccountissuedcardtransactionsresponse.md)**

