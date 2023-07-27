<!-- Start SDK Example Usage -->


```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.ClientCredentialsGrantToAccessTokenRequest(
    client_id='5clTR_MdVrrkgxw2',
    client_secret='dNC-hg7sVm22jc3g_Eogtyu0_1Mqh_4-',
    grant_type=shared.ClientCredentialsGrantToAccessTokenRequestGrantType.REFRESH_TOKEN,
    refresh_token='i1qxz68gu50zp4i8ceyxqogmq7y0yienm52351c6...',
    scope='/accounts.write',
)

res = s.access_token.create(req)

if res.access_token_response is not None:
    # handle response
```
<!-- End SDK Example Usage -->