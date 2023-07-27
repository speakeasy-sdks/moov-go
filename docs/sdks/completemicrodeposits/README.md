# complete_micro_deposits

### Available Operations

* [put](#put) - Complete micro-deposits

## put

Complete the micro-deposit validation process by passing the amounts of the two transfers within three tries. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PutCompleteMicroDepositsRequest(
    complete_micro_deposits_request=shared.CompleteMicroDepositsRequest(
        amounts=[
            283519,
            433439,
            379927,
            826871,
        ],
    ),
    account_id='28c10ab3-cdca-4425-9904-e523c7e0bc71',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.complete_micro_deposits.put(req)

if res.complete_micro_deposits_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                | [operations.PutCompleteMicroDepositsRequest](../../models/operations/putcompletemicrodepositsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[operations.PutCompleteMicroDepositsResponse](../../models/operations/putcompletemicrodepositsresponse.md)**

