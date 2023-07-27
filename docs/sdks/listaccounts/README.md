# list_accounts

### Available Operations

* [get](#get) - List accounts

## get

List or search accounts to which the caller is connected.<br><br>
All supported query parameters are optional. If none are provided
the response will include all connected accounts. Pagination is
supported via the `skip` and `count` query parameters.<br><br>
Searching by name and email will overlap and return results based on relevance.
<br><br> To list connected accounts, you must specify the `/accounts.read` scope when retrieving the authentication token.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetListAccountsRequest(
    count=879235,
    email='Karina85@gmail.com',
    foreign_id='4528aba-b9a1-11eb-8529-0242ac13003',
    include_disconnected=False,
    name='Cecilia Quitzon IV',
    skip=372679,
    type=shared.AccountType.BUSINESS,
    verification_status=shared.AccountVerificationStatus.UNVERIFIED,
)

res = s.list_accounts.get(req)

if res.accounts is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.GetListAccountsRequest](../../models/operations/getlistaccountsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.GetListAccountsResponse](../../models/operations/getlistaccountsresponse.md)**

