# ListIssuedCardsResponse


## Fields

| Field                                                    | Type                                                     | Required                                                 | Description                                              |
| -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- | -------------------------------------------------------- |
| `ContentType`                                            | *string*                                                 | :heavy_check_mark:                                       | N/A                                                      |
| `IssuedCards`                                            | [][shared.IssuedCard](../../models/shared/issuedcard.md) | :heavy_minus_sign:                                       | Successfully retrieved cards                             |
| `StatusCode`                                             | *int*                                                    | :heavy_check_mark:                                       | N/A                                                      |
| `RawResponse`                                            | [*http.Response](https://pkg.go.dev/net/http#Response)   | :heavy_minus_sign:                                       | N/A                                                      |