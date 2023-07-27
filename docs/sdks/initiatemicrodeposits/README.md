# initiate_micro_deposits

### Available Operations

* [post](#post) - Initiate micro-deposits

## post

Micro-deposits help confirm bank account ownership, helping reduce fraud and the risk of unauthorized activity. Use this method to initiate the micro-deposit verification, sending two small credit transfers to the bank account you want to confirm. If you request micro-deposits before 4:15PM ET, they will appear that same day. If you request micro-deposits any time after 4:15PM ET, they will appear the next banking day. When the two credits are initiated, Moov simultaneously initiates a debit to recoup the micro-deposits.<br><br> `sandbox` - Micro-deposits initiated for a `sandbox` bank account will always be `$0.00` / `$0.00` and instantly verifiable once initiated. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostInitiateMicroDepositsRequest(
    account_id='f5f0642d-ac7a-4f51-9cc4-13aa63aae8d6',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.initiate_micro_deposits.post(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                  | [operations.PostInitiateMicroDepositsRequest](../../models/operations/postinitiatemicrodepositsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[operations.PostInitiateMicroDepositsResponse](../../models/operations/postinitiatemicrodepositsresponse.md)**

