# Plaid

Describes the account to link to the Moov account using a Plaid processor token.


## Fields

| Field                                                                                                                                                                                                                                                                                                                                                                                                                                             | Type                                                                                                                                                                                                                                                                                                                                                                                                                                              | Required                                                                                                                                                                                                                                                                                                                                                                                                                                          | Description                                                                                                                                                                                                                                                                                                                                                                                                                                       |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `Plaid`                                                                                                                                                                                                                                                                                                                                                                                                                                           | [*shared.PlaidIntegration](../../../pkg/models/shared/plaidintegration.md)                                                                                                                                                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                                                                                                                                                                                                                                | The details of a Plaid processor integration for a linked funding source. <br><br> `sandbox` - When linking a bank account to a `sandbox` account using a Plaid processor token a default bank account response will be used. The following default data will be used to generate the bank account in this flow:<br/>```<br/>  RoutingNumber: "011401533",<br/>  AccountNumber: "1111222233330000",<br/>  AccountType:   "checking",<br/>  Mask:          "0000"<br/>```<br/> |