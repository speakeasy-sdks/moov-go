# account_countries

### Available Operations

* [get](#get) - Get Account Countries
* [put](#put) - Assign Account Countries

## get

Retrieve the specified countries of operation for an account. <br><br> To get the list of countries, you'll need to specify the `/accounts/{accountID}/profile.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAccountCountriesRequest(
    account_id='2c595590-7aff-41a3-a2fa-9467739251aa',
)

res = s.account_countries.get(req)

if res.countries is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.GetAccountCountriesRequest](../../models/operations/getaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.GetAccountCountriesResponse](../../models/operations/getaccountcountriesresponse.md)**


## put

Assign the countries of operation for an account. This endpoint will always overwrite the previously assigned values. <br><br> To update the account countries, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PutAccountCountriesRequest(
    countries=shared.Countries(
        countries=[
            'United States',
            'United States',
        ],
    ),
    account_id='2c3f5ad0-19da-41ff-a78f-097b0074f154',
)

res = s.account_countries.put(req)

if res.countries is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.PutAccountCountriesRequest](../../models/operations/putaccountcountriesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.PutAccountCountriesResponse](../../models/operations/putaccountcountriesresponse.md)**

