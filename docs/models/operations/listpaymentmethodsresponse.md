# ListPaymentMethodsResponse


## Fields

| Field                                                          | Type                                                           | Required                                                       | Description                                                    |
| -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- | -------------------------------------------------------------- |
| `ContentType`                                                  | *string*                                                       | :heavy_check_mark:                                             | N/A                                                            |
| `PaymentMethods`                                               | [][shared.PaymentMethod](../../models/shared/paymentmethod.md) | :heavy_minus_sign:                                             | Successfully retrieved payment methods                         |
| `StatusCode`                                                   | *int*                                                          | :heavy_check_mark:                                             | N/A                                                            |
| `RawResponse`                                                  | [*http.Response](https://pkg.go.dev/net/http#Response)         | :heavy_minus_sign:                                             | N/A                                                            |