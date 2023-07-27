# GetIndustriesResponse


## Fields

| Field                                                   | Type                                                    | Required                                                | Description                                             |
| ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- | ------------------------------------------------------- |
| `ContentType`                                           | *string*                                                | :heavy_check_mark:                                      | N/A                                                     |
| `Industries`                                            | [*shared.Industries](../../models/shared/industries.md) | :heavy_minus_sign:                                      | A list of industries and their codes                    |
| `StatusCode`                                            | *int*                                                   | :heavy_check_mark:                                      | N/A                                                     |
| `RawResponse`                                           | [*http.Response](https://pkg.go.dev/net/http#Response)  | :heavy_minus_sign:                                      | N/A                                                     |