# PlaidLinkIntegration

This is used by Moov.js with a Plaid reseller relationship. The details of a Plaid link integration for a linked funding source. <br><br> `sandbox` - (requires Plaid Reseller setup with Moov) When linking a bank account to a `sandbox` account using a Plaid public token it will utilize Plaid's sandbox environment. The Plaid public token provided must be generated from Plaid's sandbox environment. Please see <a href="https://plaid.com/docs/api/sandbox/#sandboxpublic_tokencreate" target="_blank">Plaid's sandbox documentation</a> for more details.


## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `public_token`     | *Optional[str]*    | :heavy_minus_sign: | N/A                |