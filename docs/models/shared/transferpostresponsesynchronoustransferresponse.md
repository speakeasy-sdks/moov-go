# TransferPostResponseSynchronousTransferResponse

Transfer details


## Fields

| Field                                                                                                                                                  | Type                                                                                                                                                   | Required                                                                                                                                               | Description                                                                                                                                            | Example                                                                                                                                                |
| ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Amount`                                                                                                                                               | [*Amount](../../models/shared/amount.md)                                                                                                               | :heavy_minus_sign:                                                                                                                                     | A representation of money containing an integer value and it's currency.                                                                               |                                                                                                                                                        |
| `CreatedOn`                                                                                                                                            | [*time.Time](https://pkg.go.dev/time#Time)                                                                                                             | :heavy_minus_sign:                                                                                                                                     | N/A                                                                                                                                                    |                                                                                                                                                        |
| `Description`                                                                                                                                          | **string*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                     | A description of the transfer                                                                                                                          | Pay Instructor for May 15 Class                                                                                                                        |
| `Destination`                                                                                                                                          | [*GetTransferFullSourceDestination](../../models/shared/gettransferfullsourcedestination.md)                                                           | :heavy_minus_sign:                                                                                                                                     | N/A                                                                                                                                                    |                                                                                                                                                        |
| `DisputedAmount`                                                                                                                                       | [*TransferPostResponseSynchronousTransferResponseDisputedAmount](../../models/shared/transferpostresponsesynchronoustransferresponsedisputedamount.md) | :heavy_minus_sign:                                                                                                                                     | The total disputed amount for a card transfer                                                                                                          |                                                                                                                                                        |
| `Disputes`                                                                                                                                             | [][GetDispute](../../models/shared/getdispute.md)                                                                                                      | :heavy_minus_sign:                                                                                                                                     | A list of disputes for a card transfer                                                                                                                 |                                                                                                                                                        |
| `FacilitatorFee`                                                                                                                                       | [*GetFacilitatorFee](../../models/shared/getfacilitatorfee.md)                                                                                         | :heavy_minus_sign:                                                                                                                                     | Fee you charged your customer for the transfer                                                                                                         |                                                                                                                                                        |
| `FailureReason`                                                                                                                                        | [*FailureReason](../../models/shared/failurereason.md)                                                                                                 | :heavy_minus_sign:                                                                                                                                     | Transfer failure reason                                                                                                                                | wallet-insufficient-funds                                                                                                                              |
| `GroupID`                                                                                                                                              | **string*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                     | N/A                                                                                                                                                    | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                                                                   |
| `Metadata`                                                                                                                                             | map[string]*string*                                                                                                                                    | :heavy_minus_sign:                                                                                                                                     | Free-form key-value pair list. Useful for storing information that is not captured elsewhere.                                                          |                                                                                                                                                        |
| `MoovFee`                                                                                                                                              | **int64*                                                                                                                                               | :heavy_minus_sign:                                                                                                                                     | Fee charged to your platform account for card transfers                                                                                                |                                                                                                                                                        |
| `RefundedAmount`                                                                                                                                       | [*TransferPostResponseSynchronousTransferResponseRefundedAmount](../../models/shared/transferpostresponsesynchronoustransferresponserefundedamount.md) | :heavy_minus_sign:                                                                                                                                     | The total refunded amount for a card transfer                                                                                                          |                                                                                                                                                        |
| `Refunds`                                                                                                                                              | [][GetRefund](../../models/shared/getrefund.md)                                                                                                        | :heavy_minus_sign:                                                                                                                                     | A list of refunds for a card transfer                                                                                                                  |                                                                                                                                                        |
| `Source`                                                                                                                                               | [*GetTransferFullSource](../../models/shared/gettransferfullsource.md)                                                                                 | :heavy_minus_sign:                                                                                                                                     | N/A                                                                                                                                                    |                                                                                                                                                        |
| `Status`                                                                                                                                               | [*TransferStatus](../../models/shared/transferstatus.md)                                                                                               | :heavy_minus_sign:                                                                                                                                     | Current status of a transfer                                                                                                                           | pending                                                                                                                                                |
| `TransferID`                                                                                                                                           | **string*                                                                                                                                              | :heavy_minus_sign:                                                                                                                                     | UUID v4                                                                                                                                                | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                                                                   |