# RevokeTokenRequest

Allows clients to notify the authorization server that a previously obtained refresh or access token is no longer needed


## Fields

| Field                                                                  | Type                                                                   | Required                                                               | Description                                                            | Example                                                                |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `ClientID`                                                             | **string*                                                              | :heavy_minus_sign:                                                     | If not specified in `Authorization: Basic` it can be specified here    | 5clTR_MdVrrkgxw2                                                       |
| `ClientSecret`                                                         | **string*                                                              | :heavy_minus_sign:                                                     | If not specified in `Authorization: Basic` it can be specified here    | dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-                                       |
| `Token`                                                                | *string*                                                               | :heavy_check_mark:                                                     | String passed to the authorization server to gain access to the system | i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...                            |
| `TokenTypeHint`                                                        | [*shared.TokenTypeHint](../../../pkg/models/shared/tokentypehint.md)   | :heavy_minus_sign:                                                     | A hint about the type of the token submitted for revocation            |                                                                        |