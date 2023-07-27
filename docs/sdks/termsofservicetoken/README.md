# terms_of_service_token

### Available Operations

* [get](#get) - Get terms of service token

## get

Generates a non-expiring token that can then be used to accept Moov's terms of service. This token can only be generated via API. Any Moov account requesting the `collect-funds`, `send-funds`, `wallet`, or `card-issuing` capabilities must accept Moov's terms of service, then have the generated terms of service token patched to the account. Read more on our [quick start guide](https://docs.moov.io/guides/quick-start/#request-capabilities).

### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)


res = s.terms_of_service_token.get()

if res.terms_of_service_token is not None:
    # handle response
```


### Response

**[operations.GetTermsOfServiceTokenResponse](../../models/operations/gettermsofservicetokenresponse.md)**

