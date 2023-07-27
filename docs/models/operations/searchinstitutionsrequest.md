# SearchInstitutionsRequest


## Fields

| Field                                      | Type                                       | Required                                   | Description                                | Example                                    |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `limit`                                    | *Optional[int]*                            | :heavy_minus_sign:                         | Maximum results returned by a search       | 499                                        |
| `name`                                     | *Optional[str]*                            | :heavy_minus_sign:                         | Name of the financial institution          |                                            |
| `rail`                                     | [shared.Rail](../../models/shared/rail.md) | :heavy_check_mark:                         | Payment rail to search on                  |                                            |
| `routing_number`                           | *Optional[str]*                            | :heavy_minus_sign:                         | Routing number for a financial institution |                                            |
| `state`                                    | *Optional[str]*                            | :heavy_minus_sign:                         | Optional parameters to filter results      |                                            |