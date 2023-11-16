# ClientCredentialsGrantToAccessTokenRequest

Allows the use of `Client Credentials Grant` per the RFC 6749 of (OAuth 2.0 Authorization Framework)[https://tools.ietf.org/html/rfc6749#section-4.4]. Following this specification will allow any tooling to be able to use this API to get an `access_token`.


## Fields

| Field                                                                                                             | Type                                                                                                              | Required                                                                                                          | Description                                                                                                       | Example                                                                                                           |
| ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------- |
| `ClientID`                                                                                                        | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | If not specified in `Authorization: Basic` it can be specified here                                               | 5clTR_MdVrrkgxw2                                                                                                  |
| `ClientSecret`                                                                                                    | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | If not specified in `Authorization: Basic` it can be specified here                                               | dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-                                                                                  |
| `GrantType`                                                                                                       | [shared.GrantType](../../../pkg/models/shared/granttype.md)                                                       | :heavy_check_mark:                                                                                                | N/A                                                                                                               |                                                                                                                   |
| `RefreshToken`                                                                                                    | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | String passed to the authorization server to gain access to the system                                            | i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...                                                                       |
| `Scope`                                                                                                           | **string*                                                                                                         | :heavy_minus_sign:                                                                                                | A space-delimited list of [scopes](https://docs.moov.io/guides/developer-tools/api-keys/scopes/) that are allowed | /accounts.write                                                                                                   |