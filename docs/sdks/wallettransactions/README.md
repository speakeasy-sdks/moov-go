# wallet_transactions

### Available Operations

* [list](#list) - List wallet transactions

## list

List all the transactions associated with a particular Moov wallet. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListWalletTransactionsRequest(
    account_id='dbb30fcb-33ea-4055-b197-cd44e2f52d82',
    completed_end_date_time='pariatur',
    completed_start_date_time='amet',
    count=347698,
    created_end_date_time='ab',
    created_start_date_time='velit',
    skip=705710,
    source_id='tempore',
    source_type='nisi',
    status='voluptatibus',
    transaction_type='quaerat',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallet_transactions.list(req)

if res.wallet_transactions is not None:
    # handle response
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `request`                                                                                            | [operations.ListWalletTransactionsRequest](../../models/operations/listwallettransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[operations.ListWalletTransactionsResponse](../../models/operations/listwallettransactionsresponse.md)**

