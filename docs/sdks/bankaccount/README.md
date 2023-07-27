# bank_account

### Available Operations

* [post](#post) - Bank account

## post

Link a bank account to an existing Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostBankAccountRequest(
    request_body=shared.BankAccount(
        account=shared.BankAccountIntegration(
            account_number='0004321567000',
            bank_account_type=shared.BankAccountType.UNKNOWN,
            holder_name='Jules Jackson',
            holder_type=shared.HolderType.BUSINESS,
            routing_number='123456789',
        ),
    ),
    account_id='49a8d9cb-f486-4333-a3f9-b77f3a410067',
)

res = s.bank_account.post(req)

if res.bank_account_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.PostBankAccountRequest](../../models/operations/postbankaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.PostBankAccountResponse](../../models/operations/postbankaccountresponse.md)**

