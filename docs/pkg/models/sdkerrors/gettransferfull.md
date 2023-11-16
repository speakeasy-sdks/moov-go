# GetTransferFull

Transfer details


## Fields

| Field                                                                                                            | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      | Example                                                                                                          |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `Amount`                                                                                                         | [*sdkerrors.Amount](../../../pkg/models/sdkerrors/amount.md)                                                     | :heavy_minus_sign:                                                                                               | A representation of money containing an integer value and it's currency.                                         |                                                                                                                  |
| `CreatedOn`                                                                                                      | [*time.Time](https://pkg.go.dev/time#Time)                                                                       | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `Description`                                                                                                    | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | A description of the transfer                                                                                    | Pay Instructor for May 15 Class                                                                                  |
| `Destination`                                                                                                    | [*sdkerrors.GetTransferFullSourceDestination](../../../pkg/models/sdkerrors/gettransferfullsourcedestination.md) | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `DisputedAmount`                                                                                                 | [*sdkerrors.DisputedAmount](../../../pkg/models/sdkerrors/disputedamount.md)                                     | :heavy_minus_sign:                                                                                               | The total disputed amount for a card transfer                                                                    |                                                                                                                  |
| `Disputes`                                                                                                       | [][sdkerrors.GetDispute](../../../pkg/models/sdkerrors/getdispute.md)                                            | :heavy_minus_sign:                                                                                               | A list of disputes for a card transfer                                                                           |                                                                                                                  |
| `FacilitatorFee`                                                                                                 | [*sdkerrors.GetFacilitatorFee](../../../pkg/models/sdkerrors/getfacilitatorfee.md)                               | :heavy_minus_sign:                                                                                               | Fee you charged your customer for the transfer                                                                   |                                                                                                                  |
| `FailureReason`                                                                                                  | [*sdkerrors.FailureReason](../../../pkg/models/sdkerrors/failurereason.md)                                       | :heavy_minus_sign:                                                                                               | Transfer failure reason                                                                                          | wallet-insufficient-funds                                                                                        |
| `GroupID`                                                                                                        | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | N/A                                                                                                              | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                             |
| `Metadata`                                                                                                       | map[string]*string*                                                                                              | :heavy_minus_sign:                                                                                               | Free-form key-value pair list. Useful for storing information that is not captured elsewhere.                    |                                                                                                                  |
| `MoovFee`                                                                                                        | **int64*                                                                                                         | :heavy_minus_sign:                                                                                               | Fee charged to your platform account for card transfers                                                          |                                                                                                                  |
| `RefundedAmount`                                                                                                 | [*sdkerrors.RefundedAmount](../../../pkg/models/sdkerrors/refundedamount.md)                                     | :heavy_minus_sign:                                                                                               | The total refunded amount for a card transfer                                                                    |                                                                                                                  |
| `Refunds`                                                                                                        | [][sdkerrors.GetRefund](../../../pkg/models/sdkerrors/getrefund.md)                                              | :heavy_minus_sign:                                                                                               | A list of refunds for a card transfer                                                                            |                                                                                                                  |
| `Source`                                                                                                         | [*sdkerrors.GetTransferFullSource](../../../pkg/models/sdkerrors/gettransferfullsource.md)                       | :heavy_minus_sign:                                                                                               | N/A                                                                                                              |                                                                                                                  |
| `Status`                                                                                                         | [*sdkerrors.TransferStatus](../../../pkg/models/sdkerrors/transferstatus.md)                                     | :heavy_minus_sign:                                                                                               | Current status of a transfer                                                                                     | pending                                                                                                          |
| `TransferID`                                                                                                     | **string*                                                                                                        | :heavy_minus_sign:                                                                                               | UUID v4                                                                                                          | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                                                             |