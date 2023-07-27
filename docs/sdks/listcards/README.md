# list_cards

### Available Operations

* [get](#get) - List cards

## get

List all the cards associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetListCardsRequest(
    account_id='689eee95-26f8-4d98-ae88-1ead4f0e1012',
)

res = s.list_cards.get(req)

if res.cards is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.GetListCardsRequest](../../models/operations/getlistcardsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.GetListCardsResponse](../../models/operations/getlistcardsresponse.md)**

