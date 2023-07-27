# SearchInstitutionRequest


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Limit`                                    | **int64*                                   | :heavy_minus_sign:                         | Maximum results returned by a search       | 499                                        |
| `Name`                                     | **string*                                  | :heavy_minus_sign:                         | Name of the financial institution          |                                            |
| `Rail`                                     | [shared.Rail](../../models/shared/rail.md) | :heavy_check_mark:                         | Payment rail to search on                  |                                            |
| `RoutingNumber`                            | **string*                                  | :heavy_minus_sign:                         | Routing number for a financial institution |                                            |
| `State`                                    | **string*                                  | :heavy_minus_sign:                         | Optional parameters to filter results      |                                            |