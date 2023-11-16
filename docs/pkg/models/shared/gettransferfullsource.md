# GetTransferFullSource


## Fields

| Field                                                                                                      | Type                                                                                                       | Required                                                                                                   | Description                                                                                                | Example                                                                                                    |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `Account`                                                                                                  | [*shared.GetTransferPartialAccount](../../../pkg/models/shared/gettransferpartialaccount.md)               | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `AchDetails`                                                                                               | [*shared.ACHDetailsSource](../../../pkg/models/shared/achdetailssource.md)                                 | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `ApplePay`                                                                                                 | [*shared.GetTransferFullSourceApplePay](../../../pkg/models/shared/gettransferfullsourceapplepay.md)       | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `BankAccount`                                                                                              | [*shared.GetTransferFullSourceBankAccount](../../../pkg/models/shared/gettransferfullsourcebankaccount.md) | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `Card`                                                                                                     | [*shared.GetTransferFullSourceCard](../../../pkg/models/shared/gettransferfullsourcecard.md)               | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `CardDetails`                                                                                              | [*shared.CardDetails](../../../pkg/models/shared/carddetails.md)                                           | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |
| `PaymentMethodID`                                                                                          | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | UUID v4                                                                                                    | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                       |
| `PaymentMethodType`                                                                                        | [*shared.PaymentMethodsType](../../../pkg/models/shared/paymentmethodstype.md)                             | :heavy_minus_sign:                                                                                         | The payment method type that represents a payment rail and directionality                                  |                                                                                                            |
| `TransferID`                                                                                               | **string*                                                                                                  | :heavy_minus_sign:                                                                                         | UUID v4                                                                                                    | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                       |
| `Wallet`                                                                                                   | [*shared.GetTransferFullSourceWallet](../../../pkg/models/shared/gettransferfullsourcewallet.md)           | :heavy_minus_sign:                                                                                         | N/A                                                                                                        |                                                                                                            |