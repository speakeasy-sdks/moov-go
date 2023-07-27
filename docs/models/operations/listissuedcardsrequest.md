# ListIssuedCardsRequest


## Fields

| Field                                                                            | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `AccountID`                                                                      | *string*                                                                         | :heavy_check_mark:                                                               | ID of the account                                                                |
| `Count`                                                                          | **int64*                                                                         | :heavy_minus_sign:                                                               | Optional parameter to limit the number of results in the query                   |
| `Skip`                                                                           | **int64*                                                                         | :heavy_minus_sign:                                                               | The number of items to offset before starting to collect the result set          |
| `States`                                                                         | [*shared.IssuedCardState](../../models/shared/issuedcardstate.md)                | :heavy_minus_sign:                                                               | Optional, comma-separated states to filter the Moov list issued cards response.<br/> |