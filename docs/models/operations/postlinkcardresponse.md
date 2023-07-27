# PostLinkCardResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Card`                                                                | [*shared.Card](../../models/shared/card.md)                           | :heavy_minus_sign:                                                    | Card linked                                                           |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | N/A                                                                   |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | N/A                                                                   |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | N/A                                                                   |
| `PostLinkCard422ApplicationJSONObject`                                | map[string]*interface{}*                                              | :heavy_minus_sign:                                                    | The supplied card data appeared invalid or was declined by the issuer |