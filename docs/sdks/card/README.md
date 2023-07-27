# card

### Available Operations

* [delete](#delete) - Disable card
* [get](#get) - Get card
* [update](#update) - Update card

## delete

Disables a card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteCardRequest(
    account_id='d39c0f5d-2cff-47c7-8a45-626d436813f1',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.DeleteCardRequest](../../models/operations/deletecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.DeleteCardResponse](../../models/operations/deletecardresponse.md)**


## get

Fetch a specific card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetCardRequest(
    account_id='6d9f5fce-6c55-4614-ac3e-250fb008c42e',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card.get(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `request`                                                              | [operations.GetCardRequest](../../models/operations/getcardrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[operations.GetCardResponse](../../models/operations/getcardresponse.md)**


## update

Update a Linked Card and/or resubmit it for verification. If a value is provided for CVV, 
a new verification ($0 authorization) will be submitted for the card. Updating the expiration date or 
address will update the information stored on file for the card but will not be verified. 
Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
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

req = operations.UpdateCardRequest(
    card_update_request=shared.CardUpdateRequest(
        billing_address=shared.UpdateAddress(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        card_cvv='123',
        card_on_file=False,
        expiration=shared.UpdateCardExpiration(
            month='01',
            year='21',
        ),
    ),
    account_id='141aac36-6c8d-4d6b-9442-907474778a7b',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card.update(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.UpdateCardRequest](../../models/operations/updatecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.UpdateCardResponse](../../models/operations/updatecardresponse.md)**

