# RequestCard


## Fields

| Field                                                                                | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `AuthorizationControls`                                                              | [*shared.AuthorizationControls](../../../pkg/models/shared/authorizationcontrols.md) | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `AuthorizedUser`                                                                     | [*shared.CreateAuthorizedUser](../../../pkg/models/shared/createauthorizeduser.md)   | :heavy_minus_sign:                                                                   | Fields to identify a human                                                           |
| `FundingWalletID`                                                                    | **string*                                                                            | :heavy_minus_sign:                                                                   | N/A                                                                                  |
| `Memo`                                                                               | **string*                                                                            | :heavy_minus_sign:                                                                   | Optional descriptive name                                                            |
| `Type`                                                                               | [*shared.IssuedCardType](../../../pkg/models/shared/issuedcardtype.md)               | :heavy_minus_sign:                                                                   | Type of a Moov issued card                                                           |