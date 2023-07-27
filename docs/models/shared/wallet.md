# Wallet

A Moov wallet to store funds for transfers.


## Fields

| Field                                                                    | Type                                                                     | Required                                                                 | Description                                                              | Example                                                                  |
| ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ | ------------------------------------------------------------------------ |
| `available_balance`                                                      | [Optional[Amount]](../../models/shared/amount.md)                        | :heavy_minus_sign:                                                       | A representation of money containing an integer value and it's currency. |                                                                          |
| `wallet_id`                                                              | *Optional[str]*                                                          | :heavy_minus_sign:                                                       | UUID v4                                                                  | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                                     |