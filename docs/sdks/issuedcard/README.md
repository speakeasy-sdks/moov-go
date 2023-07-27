# issued_card

### Available Operations

* [get](#get) - Get issued card
* [update](#update) - Update issued card

## get

Retrieve a single issued card associated with a Moov account
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetIssuedCardRequest(
    account_id='75fd5e60-b375-4ed4-b6fb-ee41f33317fe',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.issued_card.get(req)

if res.issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.GetIssuedCardRequest](../../models/operations/getissuedcardrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.GetIssuedCardResponse](../../models/operations/getissuedcardresponse.md)**


## update

Update a Moov issued card
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.UpdateIssuedCardRequest(
    update_issued_card=shared.UpdateIssuedCard(
        authorization_controls=shared.AuthorizationControls(
            spend_limits=[
                shared.AuthorizationSpendLimitControl(
                    amount=10000,
                    duration=shared.AuthorizationSpendDuration.TRANSACTION,
                ),
            ],
        ),
        authorized_user=shared.CreateAuthorizedUser(
            birth_date=shared.BirthDate(
                day=9,
                month=11,
                year=1989,
            ),
            first_name='Jane',
            last_name='Doe',
        ),
        memo='corporis',
        state=shared.IssuedCardState.PENDING_VERIFICATION,
    ),
    account_id='60eb1ea4-2655-45ba-bc28-744ed53b88f3',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.issued_card.update(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.UpdateIssuedCardRequest](../../models/operations/updateissuedcardrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.UpdateIssuedCardResponse](../../models/operations/updateissuedcardresponse.md)**

