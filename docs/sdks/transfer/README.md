# transfer

### Available Operations

* [cancel](#cancel) - Cancel or refund a card transfer
* [create](#create) - Create a transfer
* [get](#get) - Get a transfer
* [patch](#patch) - Patch transfer metadata
* [refund](#refund) - Refund a transfer

## cancel

Reverses a card transfer by initiating a cancellation or refund depending on the transaction status

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.CancelTransferRequest(
    create_reversal=shared.CreateReversal(
        amount=1000,
    ),
    x_idempotency_key='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    transfer_id='a59ecfef-66ef-41ca-a338-3c2beb477373',
)

res = s.transfer.cancel(req)

if res.created_reversal is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.CancelTransferRequest](../../models/operations/canceltransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.CancelTransferResponse](../../models/operations/canceltransferresponse.md)**


## create

Move money by providing the source, destination, and amount in the request body. <br><br> To create a transfer, you must specify the `/accounts/{yourAccountID}/transfers.write` scope. <br> You can find your account id on the [Business details](https://dashboard.moov.io/settings/business) page.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.CreateTransferRequest(
    create_transfer=shared.CreateTransfer(
        amount=shared.Amount(
            currency='USD',
            value=1204,
        ),
        description='Pay Instructor for May 15 Class',
        destination=shared.CreateTransferDestination(
            ach_details=shared.CreateACHDetailsBase(
                company_entry_description='Gym Dues',
                originating_company_name='Whole Body Fit',
            ),
            payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
        ),
        facilitator_fee=shared.CreateFacilitatorFee(
            markup=778172,
            total=535468,
        ),
        metadata={
            "iure": 'odit',
            "voluptatibus": 'vel',
            "magnam": 'quibusdam',
            "inventore": 'facere',
        },
        source=shared.CreateTransferSource(
            ach_details=shared.CreateAchDetailsSource(
                company_entry_description='Gym Dues',
                debit_hold_period=shared.CreateAchDetailsSourceDebitHoldPeriod.TWO_DAYS,
                originating_company_name='Whole Body Fit',
            ),
            card_details=shared.CreateCardDetails(
                dynamic_descriptor='WhlBdy *Yoga 11-12',
                transaction_source=shared.TransactionSource.UNSCHEDULED,
            ),
            payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
            transfer_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
        ),
    ),
    x_idempotency_key='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    x_wait_for=shared.WaitFor.RAIL_RESPONSE,
)

res = s.transfer.create(req)

if res.transfer_post_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.CreateTransferResponse](../../models/operations/createtransferresponse.md)**


## get

Retrieve full transfer details such as the amount, source, and destination. Payment rail-specific details are included in the source and destination. <br><br> To get a transfer, you must specify the `/accounts/{yourAccountID}/transfers.read` scope. The accountID included must be your facilitator accountID. <br> You can find your accountID on the [Business details](https://dashboard.moov.io/settings/business) page.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetTransferRequest(
    account_id='1f2c4310-661e-4963-89e1-cf9e06e3a437',
    transfer_id='000ae6b6-bc9b-48f7-99ea-c55a9741d311',
)

res = s.transfer.get(req)

if res.get_transfer_full is not None:
    # handle response
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `request`                                                                      | [operations.GetTransferRequest](../../models/operations/gettransferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[operations.GetTransferResponse](../../models/operations/gettransferresponse.md)**


## patch

Update the metadata contained on a transfer <br><br> To patch a transfer, you must specify the `/accounts/{yourAccountID}/transfers.write` scope. The accountID included must be your facilitator accountID. <br> You can find your account ID on the [Business details](https://dashboard.moov.io/settings/business) page.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PatchTransferRequest(
    patch_transfer=shared.PatchTransfer(
        metadata={
            "ad": 'qui',
        },
    ),
    account_id='965bb8a7-2026-4114-b5e1-39dbc2259b1a',
    transfer_id='bda8c070-e108-44cb-8672-d1ad879eeb96',
)

res = s.transfer.patch(req)

if res.get_transfer_full is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.PatchTransferRequest](../../models/operations/patchtransferrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.PatchTransferResponse](../../models/operations/patchtransferresponse.md)**


## refund

<strong>Use the <a href="index.html#tag/Transfers/operation/reverseTransfer">Cancel or refund a card transfer</a> endpoint for more robust cancel and refund options.</strong> <br><br> Initiate a refund for a card transfer <br><br> To initiate a refund, you will need to specify the `/accounts/{accountID}/transfers.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.RefundTransferRequest(
    create_refund=shared.CreateRefund(
        amount=1000,
    ),
    x_idempotency_key='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    x_wait_for=shared.WaitFor.RAIL_RESPONSE,
    transfer_id='65b85efb-d02b-4ae0-be2d-782259e3ea4b',
)

res = s.transfer.refund(req)

if res.refund_post_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.RefundTransferRequest](../../models/operations/refundtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.RefundTransferResponse](../../models/operations/refundtransferresponse.md)**

