# PostLinkCardResponse


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `Card`                                                                | [*shared.Card](../../models/shared/card.md)                           | :heavy_minus_sign:                                                    | Card linked                                                           |
| `ContentType`                                                         | *string*                                                              | :heavy_check_mark:                                                    | HTTP response content type for this operation                         |
| `StatusCode`                                                          | *int*                                                                 | :heavy_check_mark:                                                    | HTTP response status code for this operation                          |
| `RawResponse`                                                         | [*http.Response](https://pkg.go.dev/net/http#Response)                | :heavy_minus_sign:                                                    | Raw HTTP response; suitable for custom response parsing               |
| `PostLinkCard422ApplicationJSONObject`                                | map[string]*interface{}*                                              | :heavy_minus_sign:                                                    | The supplied card data appeared invalid or was declined by the issuer |