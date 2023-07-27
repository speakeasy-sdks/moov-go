# wallets_for_account

### Available Operations

* [get](#get) - Get wallet
* [list](#list) - List wallets

## get

Get information on a specific wallet (e.g., the available balance). <br><br> To get wallet information, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetWalletForAccountRequest(
    account_id='8b656bcd-b35f-4f2e-8b27-537a8cd9e731',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallets_for_account.get(req)

if res.wallet is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetWalletForAccountRequest](../../models/operations/getwalletforaccountrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetWalletForAccountResponse](../../models/operations/getwalletforaccountresponse.md)**


## list

List the wallets associated with a Moov account. <br><br> To list wallets, you'll need to specify the `/accounts/{accountID}/wallets.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListWalletsForAccountRequest(
    account_id='9c177d52-5f77-4b11-8eeb-52ff785fc378',
)

res = s.wallets_for_account.list(req)

if res.wallets is not None:
    # handle response
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `request`                                                                                          | [operations.ListWalletsForAccountRequest](../../models/operations/listwalletsforaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[operations.ListWalletsForAccountResponse](../../models/operations/listwalletsforaccountresponse.md)**

