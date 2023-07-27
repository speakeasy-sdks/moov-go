# transfer_options

### Available Operations

* [create](#create) - Generate transfer options

## create

Generate available payment method options for one or multiple transfer participants depending on the accountID or paymentMethodID you supply in the request. <br><br> To get transfer options, you must specify the `/accounts/{yourAccountID}/transfers.read` scope. The accountID you include should be associated with the facilitator account. <br> You can find your accountID on the [Business details](https://dashboard.moov.io/settings/business) page.

### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = shared.CreateTransferOptions(
    amount=shared.Amount(
        currency='USD',
        value=1204,
    ),
    destination=shared.CreateTransferOptionsDestination(
        account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
        payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    ),
    source=shared.CreateTransferOptionsSource(
        account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
        payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    ),
)

res = s.transfer_options.create(req)

if res.created_transfer_options is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [shared.CreateTransferOptions](../../models/shared/createtransferoptions.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.CreateTransferOptionsResponse](../../models/operations/createtransferoptionsresponse.md)**

