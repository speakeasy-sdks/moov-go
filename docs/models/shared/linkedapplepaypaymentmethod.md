# LinkedApplePayPaymentMethod

Apple Pay token linked


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `apple_pay`                                                               | [ApplePayResponse](../../models/shared/applepayresponse.md)               | :heavy_check_mark:                                                        | Describes an Apple Pay token on a Moov account.                           |
| `payment_method_id`                                                       | *str*                                                                     | :heavy_check_mark:                                                        | ID of the payment method                                                  |
| `payment_method_type`                                                     | [PaymentMethodType](../../models/shared/paymentmethodtype.md)             | :heavy_check_mark:                                                        | The payment method type that represents a payment rail and directionality |