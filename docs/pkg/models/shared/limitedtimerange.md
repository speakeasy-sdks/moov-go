# LimitedTimeRange

Return `count` number of results within time range between two timestamps and then the interval duration for each result in the specific `tz` timezone


## Fields

| Field                                      | Type                                       | Required                                   | Description                                |
| ------------------------------------------ | ------------------------------------------ | ------------------------------------------ | ------------------------------------------ |
| `Count`                                    | **int64*                                   | :heavy_minus_sign:                         | N/A                                        |
| `Every`                                    | **string*                                  | :heavy_minus_sign:                         | N/A                                        |
| `From`                                     | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `To`                                       | [*time.Time](https://pkg.go.dev/time#Time) | :heavy_minus_sign:                         | N/A                                        |
| `Tz`                                       | **string*                                  | :heavy_minus_sign:                         | N/A                                        |