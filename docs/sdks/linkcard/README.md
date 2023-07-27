# link_card

### Available Operations

* [post](#post) - Link card

## post

Link a card to an existing Moov account. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostLinkCardRequest(
    card_request=shared.CardRequest(
        billing_address=shared.Address(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        card_cvv='0123',
        card_number='illo',
        card_on_file=False,
        expiration=shared.CardExpiration(
            month='01',
            year='21',
        ),
        holder_name='Jules Jackson',
        merchant_account_id='f08f4294-e369-48f4-87f6-03e8b445e80c',
    ),
    x_wait_for=shared.SchemasWaitFor.PAYMENT_METHOD,
    account_id='a55efd20-e457-4e18-98b6-a89fbe3a5aa8',
)

res = s.link_card.post(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.PostLinkCardRequest](../../models/operations/postlinkcardrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `server_url`                                                                     | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | An optional server URL to use.                                                   |


### Response

**[operations.PostLinkCardResponse](../../models/operations/postlinkcardresponse.md)**

