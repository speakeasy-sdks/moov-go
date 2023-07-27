# list_bank_accounts

### Available Operations

* [get](#get) - List bank accounts

## get

List all the bank accounts associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetListBankAccountsRequest(
    account_id='88e51862-065e-4904-b3b1-194b8abf603a',
)

res = s.list_bank_accounts.get(req)

if res.bank_accounts is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetListBankAccountsRequest](../../models/operations/getlistbankaccountsrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetListBankAccountsResponse](../../models/operations/getlistbankaccountsresponse.md)**

