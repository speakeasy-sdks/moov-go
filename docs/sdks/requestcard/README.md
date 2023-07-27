# request_card

### Available Operations

* [post](#post) - Request card

## post

Request a virtual card be created
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

req = operations.PostRequestCardRequest(
    request_card=shared.RequestCard(
        authorization_controls=shared.AuthorizationControls(
            spend_limits=[
                shared.AuthorizationSpendLimitControl(
                    amount=10000,
                    duration=shared.AuthorizationSpendDuration.TRANSACTION,
                ),
                shared.AuthorizationSpendLimitControl(
                    amount=10000,
                    duration=shared.AuthorizationSpendDuration.TRANSACTION,
                ),
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
        funding_wallet_id='accusamus',
        memo='voluptatibus',
        type=shared.IssuedCardType.SINGLE_USE,
    ),
    account_id='b9f58c4d-86e6-48e4-be05-6013f59da757',
)

res = s.request_card.post(req)

if res.issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.PostRequestCardRequest](../../models/operations/postrequestcardrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.PostRequestCardResponse](../../models/operations/postrequestcardresponse.md)**

