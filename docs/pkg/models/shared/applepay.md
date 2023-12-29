# ApplePay

A method of moving money using an Apple Pay token.


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ApplePay`                                                                     | [*shared.SchemasApplePay](../../../pkg/models/shared/schemasapplepay.md)       | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `PaymentMethodID`                                                              | **string*                                                                      | :heavy_minus_sign:                                                             | UUID v4                                                                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                           |
| `PaymentMethodType`                                                            | [*shared.PaymentMethodsType](../../../pkg/models/shared/paymentmethodstype.md) | :heavy_minus_sign:                                                             | The payment method type that represents a payment rail and directionality      |                                                                                |