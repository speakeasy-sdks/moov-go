# link_apple_pay_token

### Available Operations

* [post](#post) - Link Apple Pay token

## post

Connect an Apple Pay token to the specified account.
The `token` data is defined by Apple Pay and should be passed through from Apple Pay's response unmodified.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostLinkApplePayTokenRequest(
    link_apple_pay=shared.LinkApplePay(
        token=shared.LinkApplePayToken(
            payment_data=shared.LinkApplePayTokenPaymentData(
                data='3+f4oOTwPa6f1UZ6tG...CE=',
                header=shared.LinkApplePayTokenPaymentDataHeader(
                    ephemeral_public_key='MFkwEK...Md==',
                    public_key_hash='l0CnXdMv...D1I=',
                    transaction_id='32b...4f3',
                ),
                signature='MIAGCSqGSIb3DQ.AAAA==',
                version='EC_v1',
            ),
            payment_method=shared.LinkApplePayTokenPaymentMethod(
                display_name='Visa 1234',
                network='Visa',
                type='debit',
            ),
            transaction_identifier='32b...4f3',
        ),
    ),
    account_id='a8d8f5c0-b2f2-4fb7-b194-a276b26916fe',
)

res = s.link_apple_pay_token.post(req)

if res.linked_apple_pay_payment_method is not None:
    # handle response
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `request`                                                                                          | [operations.PostLinkApplePayTokenRequest](../../models/operations/postlinkapplepaytokenrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[operations.PostLinkApplePayTokenResponse](../../models/operations/postlinkapplepaytokenresponse.md)**

