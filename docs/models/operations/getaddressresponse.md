# GetAddressResponse


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ContentType`                                                          | *string*                                                               | :heavy_check_mark:                                                     | N/A                                                                    |
| `EnrichmentAddresses`                                                  | [][shared.EnrichmentAddress](../../models/shared/enrichmentaddress.md) | :heavy_minus_sign:                                                     | Address suggestions                                                    |
| `StatusCode`                                                           | *int*                                                                  | :heavy_check_mark:                                                     | N/A                                                                    |
| `RawResponse`                                                          | [*http.Response](https://pkg.go.dev/net/http#Response)                 | :heavy_minus_sign:                                                     | N/A                                                                    |