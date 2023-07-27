# full_issued_card

### Available Operations

* [get](#get) - Get full card details

## get

Get issued card with PAN, CVV, and expiration. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read-secure` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetFullIssuedCardRequest(
    account_id='17051339-d080-486a-9840-394c26071f93',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.full_issued_card.get(req)

if res.full_issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `request`                                                                                  | [operations.GetFullIssuedCardRequest](../../models/operations/getfullissuedcardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[operations.GetFullIssuedCardResponse](../../models/operations/getfullissuedcardresponse.md)**

