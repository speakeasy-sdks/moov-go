# PaymentMethodCard

A method of moving money that is a credit or debit card


## Fields

| Field                                                                           | Type                                                                            | Required                                                                        | Description                                                                     | Example                                                                         |
| ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- | ------------------------------------------------------------------------------- |
| `card`                                                                          | [Optional[PaymentMethodCardCard]](../../models/shared/paymentmethodcardcard.md) | :heavy_minus_sign:                                                              | N/A                                                                             |                                                                                 |
| `payment_method_id`                                                             | *Optional[str]*                                                                 | :heavy_minus_sign:                                                              | UUID v4                                                                         | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                            |
| `payment_method_type`                                                           | [Optional[PaymentMethodType]](../../models/shared/paymentmethodtype.md)         | :heavy_minus_sign:                                                              | The payment method type that represents a payment rail and directionality       |                                                                                 |