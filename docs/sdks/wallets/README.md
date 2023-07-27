# wallets

## Overview

A [Moov wallet](https://docs.moov.io/guides/wallet/) can serve as a funding source as you accumulate funds. You can also use the Moov wallet to:
- Pre-fund transfers for faster payouts
- Transfer funds between Moov wallets for instantly available funds

<em> If you've requested the `send-funds` or `collect-funds` capability, the `wallet` capability will be automatically requested as well. Read more on the [data requirements for wallets here](https://docs.moov.io/guides/accounts/capabilities/#wallet).</em>


### Available Operations

* [get](#get) - Get wallet
* [get_transaction](#get_transaction) - Get wallet transaction
* [list](#list) - List wallets
* [list_transactions](#list_transactions) - List wallet transactions

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
    account_id='ca5acfbe-2fd5-4707-9779-29177deac646',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallets.get(req)

if res.wallet is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetWalletForAccountRequest](../../models/operations/getwalletforaccountrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetWalletForAccountResponse](../../models/operations/getwalletforaccountresponse.md)**


## get_transaction

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
    account_id='ecb57340-9e3e-4b1e-9a2b-12eb07f116db',
    transaction_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallets.get_transaction(req)

if res.wallet_transaction is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.GetWalletTransactionRequest](../../models/operations/getwallettransactionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.GetWalletTransactionResponse](../../models/operations/getwallettransactionresponse.md)**


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
    account_id='99545fc9-5fa8-4897-8e18-9dbb30fcb33e',
)

res = s.wallets.list(req)

if res.wallets is not None:
    # handle response
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `request`                                                                                          | [operations.ListWalletsForAccountRequest](../../models/operations/listwalletsforaccountrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[operations.ListWalletsForAccountResponse](../../models/operations/listwalletsforaccountresponse.md)**


## list_transactions

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
    account_id='a055b197-cd44-4e2f-92d8-2d3513bb6f48',
    completed_end_date_time='distinctio',
    completed_start_date_time='nisi',
    count=335977,
    created_end_date_time='nisi',
    created_start_date_time='libero',
    skip=794507,
    source_id='facere',
    source_type='facilis',
    status='ipsum',
    transaction_type='ad',
    wallet_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.wallets.list_transactions(req)

if res.wallet_transactions is not None:
    # handle response
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `request`                                                                                            | [operations.ListWalletTransactionsRequest](../../models/operations/listwallettransactionsrequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[operations.ListWalletTransactionsResponse](../../models/operations/listwallettransactionsresponse.md)**

