# UpdateIssuedCard


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `AuthorizationControls`                                                | [*AuthorizationControls](../../models/shared/authorizationcontrols.md) | :heavy_minus_sign:                                                     | N/A                                                                    |
| `AuthorizedUser`                                                       | [*CreateAuthorizedUser](../../models/shared/createauthorizeduser.md)   | :heavy_minus_sign:                                                     | Fields to identify a human                                             |
| `Memo`                                                                 | **string*                                                              | :heavy_minus_sign:                                                     | Optional descriptive name                                              |
| `State`                                                                | [*IssuedCardState](../../models/shared/issuedcardstate.md)             | :heavy_minus_sign:                                                     | State of a Moov issued card                                            |