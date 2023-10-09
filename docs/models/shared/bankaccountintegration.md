# BankAccountIntegration

Describes the account to link to the Moov account


## Fields

| Field                                                     | Type                                                      | Required                                                  | Description                                               | Example                                                   |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `AdditionalProperties`                                    | map[string]*interface{}*                                  | :heavy_minus_sign:                                        | N/A                                                       |                                                           |
| `AccountNumber`                                           | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | 0004321567000                                             |
| `BankAccountType`                                         | [BankAccountType](../../models/shared/bankaccounttype.md) | :heavy_check_mark:                                        | The bank account type                                     |                                                           |
| `HolderName`                                              | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | Jules Jackson                                             |
| `HolderType`                                              | [HolderType](../../models/shared/holdertype.md)           | :heavy_check_mark:                                        | The type of holder on a funding source                    |                                                           |
| `RoutingNumber`                                           | *string*                                                  | :heavy_check_mark:                                        | N/A                                                       | 123456789                                                 |