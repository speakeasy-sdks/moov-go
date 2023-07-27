# underwriting

## Overview

[Underwriting](https://docs.moov.io/guides/accounts/underwriting) is a tool Moov uses to understand a business’s profile before allowing them to collect funds on our platform. This profile includes information like a description of the company or the merchant’s business model, the industry they operate in, and transaction volume. Through underwriting, we can understand and prevent unnecessary financial risk for Moov and those transacting on our platform. Note that underwriting can be instant, but in some cases make take around 2 business days before approval.


### Available Operations

* [get](#get) - Retrieve underwriting details
* [update](#update) - Update underwriting details

## get

Retrieve underwriting associated with a given Moov account. <br><br> To get an account's underwriting details, you'll need to specify the `/accounts/{accountID}/underwriting.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetUnderwritingRequest(
    account_id='bae0be2d-7822-459e-bea4-b5197f92443d',
)

res = s.underwriting.get(req)

if res.underwriting is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.GetUnderwritingRequest](../../models/operations/getunderwritingrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.GetUnderwritingResponse](../../models/operations/getunderwritingresponse.md)**


## update

Update the account's underwriting by passing new values for one or more of the fields. <br><br> To update an account's underwriting details, you'll need to specify the `/accounts/{accountID}/profile.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.UpdateUnderwritingRequest(
    underwriting_request=shared.UnderwritingRequest(
        average_monthly_transaction_volume=250000,
        average_transaction_size=10000,
        max_transaction_size=50000,
    ),
    account_id='a7ce52b8-95c5-437c-a454-efb0b34896c3',
)

res = s.underwriting.update(req)

if res.underwriting is not None:
    # handle response
```

### Parameters

| Parameter                                                                                    | Type                                                                                         | Required                                                                                     | Description                                                                                  |
| -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------- |
| `request`                                                                                    | [operations.UpdateUnderwritingRequest](../../models/operations/updateunderwritingrequest.md) | :heavy_check_mark:                                                                           | The request object to use for the request.                                                   |


### Response

**[operations.UpdateUnderwritingResponse](../../models/operations/updateunderwritingresponse.md)**

