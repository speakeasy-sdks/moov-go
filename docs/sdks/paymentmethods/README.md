# payment_methods

## Overview

[Payment methods](https://docs.moov.io/guides/money-movement/payment-methods/) represent all of the ways an account can move funds to another Moov account. Payment methods are generated programmatically when a card or bank account is added or the status is updated e.g., `ach-debit-fund` will be added as a payment method once the bank account is verified.

<em>RTP® is not yet generally available on Moov. Contact us for more information.</em>


### Available Operations

* [get](#get) - Get payment method
* [list](#list) - List payment methods

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
    account_id='57a15be3-e060-4807-a2b6-e3ab8845f059',
    payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.payment_methods.get(req)

if res.payment_method is not None:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.GetPaymentMethodRequest](../../models/operations/getpaymentmethodrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.GetPaymentMethodResponse](../../models/operations/getpaymentmethodresponse.md)**


## list

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

req = operations.ListPaymentMethodsRequest(
    account_id='7a60ff2a-54a3-41e9-8764-a3e865e7956f',
    source_id='9251a5a9-da66-40ff-97bf-aad4f9efc1b4',
)

res = s.payment_methods.list(req)

if res.payment_methods is not None:
    # handle response
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `request`                                                                                    | [operations.ListPaymentMethodsRequest](../../models/operations/listpaymentmethodsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[operations.ListPaymentMethodsResponse](../../models/operations/listpaymentmethodsresponse.md)**

