# wallet_transaction

### Available Operations

* [get](#get) - Get wallet transaction

## get

Get details on a specific wallet transaction. <br><br> To access this endpoint, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetWalletTransactionRequest(
    account_id='2eb07f11-6db9-4954-9fc9-5fa88970e189',
    transaction_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallet_transaction.get(req)

if res.wallet_transaction is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.GetWalletTransactionRequest](../../models/operations/getwallettransactionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.GetWalletTransactionResponse](../../models/operations/getwallettransactionresponse.md)**

