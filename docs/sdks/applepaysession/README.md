# apple_pay_session

### Available Operations

* [post](#post) - Create Apple Pay session

## post

Create a session with Apple Pay to facilitate a payment.
A successful response from this endpoint should be passed through to Apple Pay unchanged.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostApplePaySessionRequest(
    create_apple_pay_session=shared.CreateApplePaySession(
        display_name='Example Merchant',
        domain='checkout.classbooker.dev',
    ),
    account_id='5fbb2587-0532-402c-b3d5-fe9b90c28909',
)

res = s.apple_pay_session.post(req)

if res.apple_pay_session is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.PostApplePaySessionRequest](../../models/operations/postapplepaysessionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.PostApplePaySessionResponse](../../models/operations/postapplepaysessionresponse.md)**

