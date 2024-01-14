# SchemasPaymentMethodBankAccountBankAccount

A method of moving money that is a bank account


## Fields

| Field                                                                          | Type                                                                           | Required                                                                       | Description                                                                    | Example                                                                        |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `BankAccount`                                                                  | [*shared.SchemasBankAccount](../../../pkg/models/shared/schemasbankaccount.md) | :heavy_minus_sign:                                                             | N/A                                                                            |                                                                                |
| `PaymentMethodID`                                                              | **string*                                                                      | :heavy_minus_sign:                                                             | UUID v4                                                                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                           |
| `PaymentMethodType`                                                            | [*shared.PaymentMethodsType](../../../pkg/models/shared/paymentmethodstype.md) | :heavy_minus_sign:                                                             | The payment method type that represents a payment rail and directionality      |                                                                                |