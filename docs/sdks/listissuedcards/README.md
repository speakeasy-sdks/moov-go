# list_issued_cards

### Available Operations

* [get](#get) - List issued cards

## get

List Moov issued cards existing for the account.
<br><br> All supported query parameters are optional.
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

req = operations.GetListIssuedCardsRequest(
    account_id='563f94e2-9e97-43e9-a2a5-7a15be3e0608',
    count=61078,
    skip=474668,
    states=shared.IssuedCardState.CLOSED,
)

res = s.list_issued_cards.get(req)

if res.issued_cards is not None:
    # handle response
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `request`                                                                                    | [operations.GetListIssuedCardsRequest](../../models/operations/getlistissuedcardsrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[operations.GetListIssuedCardsResponse](../../models/operations/getlistissuedcardsresponse.md)**

