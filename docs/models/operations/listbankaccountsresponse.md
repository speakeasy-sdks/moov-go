# ListBankAccountsResponse


## Fields

| Field                                                                      | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `BankAccounts`                                                             | [][shared.BankAccountResponse](../../models/shared/bankaccountresponse.md) | :heavy_minus_sign:                                                         | Successfully retrieved bank accounts                                       |
| `ContentType`                                                              | *string*                                                                   | :heavy_check_mark:                                                         | N/A                                                                        |
| `StatusCode`                                                               | *int*                                                                      | :heavy_check_mark:                                                         | N/A                                                                        |
| `RawResponse`                                                              | [*http.Response](https://pkg.go.dev/net/http#Response)                     | :heavy_minus_sign:                                                         | N/A                                                                        |