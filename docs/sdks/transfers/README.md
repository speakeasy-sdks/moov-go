# transfers

## Overview

A [transfer](https://docs.moov.io/guides/money-movement/) is the movement of money between Moov accounts, from source to destination. Provided you have linked a bank account which has been verified, you can initiate a transfer to another Moov account. All you need to do is note a [paymentMethod](#tag/Payment-methods), the $ amount of the transfer, and a brief description. With Moov, you can also implement transfer groups, allowing you to associate multiple transfers together and run them sequentially. To learn more, read our guide on [transfer groups](https://docs.moov.io/guides/money-movement/transfer-groups/#transfer-statuses).

### Available Operations

* [cancel](#cancel) - Cancel or refund a card transfer
* [create](#create) - Create a transfer
* [generate_options](#generate_options) - Generate transfer options
* [get](#get) - Get a transfer
* [get_refund](#get_refund) - Get refund details
* [list_refunds](#list_refunds) - Get a list of refunds for a card transfer
* [refund](#refund) - Refund a transfer
* [update](#update) - Patch transfer metadata

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
    transfer_id='5e16deab-3fec-4957-8a64-584273a8418d',
)

res = s.transfers.cancel(req)

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
            markup=117380,
            total=395544,
        ),
        metadata={
            "consectetur": 'aperiam',
        },
        source=shared.CreateTransferSource(
            ach_details=shared.CreateAchDetailsSource(
                company_entry_description='Gym Dues',
                debit_hold_period=shared.CreateAchDetailsSourceDebitHoldPeriod.TWO_DAYS,
                originating_company_name='Whole Body Fit',
            ),
            card_details=shared.CreateCardDetails(
                dynamic_descriptor='WhlBdy *Yoga 11-12',
                transaction_source=shared.TransactionSource.RECURRING,
            ),
            payment_method_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
            transfer_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
        ),
    ),
    x_idempotency_key='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    x_wait_for=shared.WaitFor.RAIL_RESPONSE,
)

res = s.transfers.create(req)

if res.transfer_post_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.CreateTransferRequest](../../models/operations/createtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.CreateTransferResponse](../../models/operations/createtransferresponse.md)**


## generate_options

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

res = s.transfers.generate_options(req)

if res.created_transfer_options is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [shared.CreateTransferOptions](../../models/shared/createtransferoptions.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.CreateTransferOptionsResponse](../../models/operations/createtransferoptionsresponse.md)**


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
    account_id='fb092992-1aef-4b9f-98c4-d86e68e4be05',
    transfer_id='6013f59d-a757-4a59-acfe-f66ef1caa338',
)

res = s.transfers.get(req)

if res.get_transfer_full is not None:
    # handle response
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `request`                                                                      | [operations.GetTransferRequest](../../models/operations/gettransferrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[operations.GetTransferResponse](../../models/operations/gettransferresponse.md)**


## get_refund

Get details of a refund for a card transfer

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetRefundRequest(
    refund_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
    transfer_id='3c2beb47-7373-4c8d-b2f6-4d1db1f2c431',
)

res = s.transfers.get_refund(req)

if res.get_refund is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.GetRefundRequest](../../models/operations/getrefundrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.GetRefundResponse](../../models/operations/getrefundresponse.md)**


## list_refunds

Get a list of refunds for a card transfer

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListRefundsRequest(
    transfer_id='0661e963-49e1-4cf9-a06e-3a437000ae6b',
)

res = s.transfers.list_refunds(req)

if res.get_refunds is not None:
    # handle response
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `request`                                                                      | [operations.ListRefundsRequest](../../models/operations/listrefundsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[operations.ListRefundsResponse](../../models/operations/listrefundsresponse.md)**


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
    transfer_id='6bc9b8f7-59ea-4c55-a974-1d311352965b',
)

res = s.transfers.refund(req)

if res.refund_post_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.RefundTransferRequest](../../models/operations/refundtransferrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.RefundTransferResponse](../../models/operations/refundtransferresponse.md)**


## update

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
            "rem": 'dolorum',
            "odio": 'fugit',
            "alias": 'magni',
        },
    ),
    account_id='611435e1-39db-4c22-99b1-abda8c070e10',
    transfer_id='84cb0672-d1ad-4879-aeb9-665b85efbd02',
)

res = s.transfers.update(req)

if res.get_transfer_full is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.PatchTransferRequest](../../models/operations/patchtransferrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.PatchTransferResponse](../../models/operations/patchtransferresponse.md)**

