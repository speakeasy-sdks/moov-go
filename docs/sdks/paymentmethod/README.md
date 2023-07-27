# payment_method

### Available Operations

* [get](#get) - Get payment method

## get

Get the specified payment method associated with a Moov account. <br><br> To get a payment method, you must specify the `/accounts/{accountID}/payment-methods.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetPaymentMethodRequest(
    account_id='4a3e865e-7956-4f92-91a5-a9da660ff57b',
    payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.payment_method.get(req)

if res.payment_method is not None:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.GetPaymentMethodRequest](../../models/operations/getpaymentmethodrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.GetPaymentMethodResponse](../../models/operations/getpaymentmethodresponse.md)**

