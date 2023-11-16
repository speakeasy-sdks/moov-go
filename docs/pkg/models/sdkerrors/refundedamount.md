# RefundedAmount

A representation of money containing an integer value and it's currency.


## Fields

| Field                                                                                                                   | Type                                                                                                                    | Required                                                                                                                | Description                                                                                                             | Example                                                                                                                 |
| ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------- |
| `Currency`                                                                                                              | *string*                                                                                                                | :heavy_check_mark:                                                                                                      | A 3-letter ISO 4217 currency code                                                                                       | USD                                                                                                                     |
| `Value`                                                                                                                 | *int64*                                                                                                                 | :heavy_check_mark:                                                                                                      | Quantity in the smallest unit of the specified currency. In USD this is cents, so $12.04 is 1204 and $0.99 would be 99. | 1204                                                                                                                    |