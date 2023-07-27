# bank_accounts

## Overview

To transfer money with Moov, you’ll need to link a bank account to your Moov account, then verify that account. You can link a bank account to a Moov account by providing the bank account number, routing number, and Moov account ID. 

We require micro-deposit verification to reduce the risk of fraud or unauthorized activity. You can verify a bank account by initiating [micro-deposits](https://docs.moov.io/guides/sources/bank-accounts/#micro-deposit-verification), sending two small credit transfers to the bank account you want to confirm. Note that there is no way to initiate a micro-deposit from your bank of choice. 

Alternatively, you can link and verify a bank account in one step through an instant account verification token from a third party provider like [Plaid](https://docs.moov.io/guides/sources/bank-accounts/plaid) or [MX](https://docs.moov.io/guides/sources/bank-accounts/mx/). Bank accounts can have the following statuses: `new`, `pending`, `verified`, `verificationFailed`, `errored`. Learn more by reading our guide on [bank accounts](https://docs.moov.io/guides/sources/bank-accounts/).


### Available Operations

* [initiate_micro_deposits](#initiate_micro_deposits) - Initiate micro-deposits
* [complete_micro_deposits](#complete_micro_deposits) - Complete micro-deposits
* [delete](#delete) - Delete bank account
* [get](#get) - Get bank account
* [link](#link) - Bank account
* [list](#list) - List bank accounts

## initiate_micro_deposits

Micro-deposits help confirm bank account ownership, helping reduce fraud and the risk of unauthorized activity. Use this method to initiate the micro-deposit verification, sending two small credit transfers to the bank account you want to confirm. If you request micro-deposits before 4:15PM ET, they will appear that same day. If you request micro-deposits any time after 4:15PM ET, they will appear the next banking day. When the two credits are initiated, Moov simultaneously initiates a debit to recoup the micro-deposits.<br><br> `sandbox` - Micro-deposits initiated for a `sandbox` bank account will always be `$0.00` / `$0.00` and instantly verifiable once initiated. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostInitiateMicroDepositsRequest(
    account_id='2a94bb4f-63c9-469e-9a3e-fa77dfb14cd6',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_accounts.initiate_micro_deposits(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                  | [operations.PostInitiateMicroDepositsRequest](../../models/operations/postinitiatemicrodepositsrequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[operations.PostInitiateMicroDepositsResponse](../../models/operations/postinitiatemicrodepositsresponse.md)**


## complete_micro_deposits

Complete the micro-deposit validation process by passing the amounts of the two transfers within three tries. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PutCompleteMicroDepositsRequest(
    complete_micro_deposits_request=shared.CompleteMicroDepositsRequest(
        amounts=[
            675439,
            881104,
        ],
    ),
    account_id='395efb9b-a88f-43a6-a997-074ba4469b6e',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_accounts.complete_micro_deposits(req)

if res.complete_micro_deposits_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                | [operations.PutCompleteMicroDepositsRequest](../../models/operations/putcompletemicrodepositsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[operations.PutCompleteMicroDepositsResponse](../../models/operations/putcompletemicrodepositsresponse.md)**


## delete

Discontinue using a specified bank account linked to a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteBankAccountRequest(
    account_id='21419598-90af-4a56-be25-16fe4c8b711e',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_accounts.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `request`                                                                                  | [operations.DeleteBankAccountRequest](../../models/operations/deletebankaccountrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[operations.DeleteBankAccountResponse](../../models/operations/deletebankaccountresponse.md)**


## get

Retrieve bank account details (i.e. routing number or account type) associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetBankAccountRequest(
    account_id='5b7fd2ed-0289-421c-9dc6-92601fb576b0',
    bank_account_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.bank_accounts.get(req)

if res.bank_account_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.GetBankAccountRequest](../../models/operations/getbankaccountrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.GetBankAccountResponse](../../models/operations/getbankaccountresponse.md)**


## link

Link a bank account to an existing Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.LinkBankAccountRequest(
    request_body=shared.Mx(
        mx=shared.MXAuthorizationCode(
            authorization_code='nemo',
        ),
    ),
    account_id='f0d30c5f-bb25-4870-9320-2c73d5fe9b90',
)

res = s.bank_accounts.link(req)

if res.bank_account_response is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.LinkBankAccountRequest](../../models/operations/linkbankaccountrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.LinkBankAccountResponse](../../models/operations/linkbankaccountresponse.md)**


## list

List all the bank accounts associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListBankAccountsRequest(
    account_id='c28909b3-fe49-4a8d-9cbf-48633323f9b7',
)

res = s.bank_accounts.list(req)

if res.bank_accounts is not None:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.ListBankAccountsRequest](../../models/operations/listbankaccountsrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.ListBankAccountsResponse](../../models/operations/listbankaccountsresponse.md)**

