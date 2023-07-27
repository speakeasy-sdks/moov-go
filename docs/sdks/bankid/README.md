# bank_id

### Available Operations

* [delete](#delete) - Delete bank account
* [get](#get) - Get bank account

## delete

Discontinue using a specified bank account linked to a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteBankIDRequest(
    account_id='4ebf6928-0d1b-4a77-a89e-bf737ae4203c',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_id.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.DeleteBankIDRequest](../../models/operations/deletebankidrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.DeleteBankIDResponse](../../models/operations/deletebankidresponse.md)**


## get

Retrieve bank account details (i.e. routing number or account type) associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetBankIDRequest(
    account_id='e5e6a95d-8a0d-4446-8e2a-f7a73cf3be45',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_id.get(req)

if res.bank_account_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.GetBankIDRequest](../../models/operations/getbankidrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.GetBankIDResponse](../../models/operations/getbankidresponse.md)**

