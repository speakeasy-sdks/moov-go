# CancelTransferRequest


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `XIdempotencyKey`                                                      | *string*                                                               | :heavy_check_mark:                                                     | Prevents duplicate reversals from being created                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                   |
| `TransferID`                                                           | *string*                                                               | :heavy_check_mark:                                                     | ID of the Transfer                                                     |                                                                        |
| `CreateReversal`                                                       | [*shared.CreateReversal](../../../pkg/models/shared/createreversal.md) | :heavy_minus_sign:                                                     | N/A                                                                    |                                                                        |