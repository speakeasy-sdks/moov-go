# payment_methods

### Available Operations

* [get](#get) - List payment methods

## get

Retrieve a list of payment methods associated with a Moov account. <br><br> To list payment methods, you must specify the `/accounts/{accountID}/payment-methods.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetPaymentMethodsRequest(
    account_id='faad4f9e-fc1b-4451-ac10-32648dc2f615',
    source_id='199ebfd0-e9fe-46c6-b2ca-3aed01179963',
)

res = s.payment_methods.get(req)

if res.payment_methods is not None:
    # handle response
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `request`                                                                                  | [operations.GetPaymentMethodsRequest](../../models/operations/getpaymentmethodsrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[operations.GetPaymentMethodsResponse](../../models/operations/getpaymentmethodsresponse.md)**

