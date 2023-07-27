# SearchInstitutionResponse


## Fields

| Field                                                                         | Type                                                                          | Required                                                                      | Description                                                                   |
| ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- | ----------------------------------------------------------------------------- |
| `ContentType`                                                                 | *string*                                                                      | :heavy_check_mark:                                                            | N/A                                                                           |
| `FinancialInstitutions`                                                       | [*shared.FinancialInstitutions](../../models/shared/financialinstitutions.md) | :heavy_minus_sign:                                                            | Financial institutions returned from a search                                 |
| `StatusCode`                                                                  | *int*                                                                         | :heavy_check_mark:                                                            | N/A                                                                           |
| `RawResponse`                                                                 | [*http.Response](https://pkg.go.dev/net/http#Response)                        | :heavy_minus_sign:                                                            | N/A                                                                           |