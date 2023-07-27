# GetTermsOfServiceTokenResponse


## Fields

| Field                                                                     | Type                                                                      | Required                                                                  | Description                                                               |
| ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- | ------------------------------------------------------------------------- |
| `ContentType`                                                             | *string*                                                                  | :heavy_check_mark:                                                        | N/A                                                                       |
| `StatusCode`                                                              | *int*                                                                     | :heavy_check_mark:                                                        | N/A                                                                       |
| `RawResponse`                                                             | [*http.Response](https://pkg.go.dev/net/http#Response)                    | :heavy_minus_sign:                                                        | N/A                                                                       |
| `TermsOfServiceToken`                                                     | [*shared.TermsOfServiceToken](../../models/shared/termsofservicetoken.md) | :heavy_minus_sign:                                                        | The newly generated terms of service token                                |