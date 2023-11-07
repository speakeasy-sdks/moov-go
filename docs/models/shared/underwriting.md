# Underwriting

Describes underwriting values (in USD) used for card payment acceptance


## Fields

| Field                                                                   | Type                                                                    | Required                                                                | Description                                                             | Example                                                                 |
| ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- | ----------------------------------------------------------------------- |
| `AverageMonthlyTransactionVolume`                                       | **int64*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     | 250000                                                                  |
| `AverageTransactionSize`                                                | **int64*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     | 10000                                                                   |
| `MaxTransactionSize`                                                    | **int64*                                                                | :heavy_minus_sign:                                                      | N/A                                                                     | 50000                                                                   |
| `Status`                                                                | [*shared.UnderwritingStatus](../../models/shared/underwritingstatus.md) | :heavy_minus_sign:                                                      | The status of underwriting for an account                               |                                                                         |