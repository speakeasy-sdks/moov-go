# card_issuing

## Overview

A Moov wallet can serve as a funding source for issuing virtual cards. Note that we currently only issue Discover cards. Virtual cards can then be used to spend funds from the wallet.

<em> The `card-issuing` and `wallet` capabilities are required to be enabled before any card issuing functionality is available. Moov is in a private beta with select customers for card issuing.</em>


### Available Operations

* [request_card](#request_card) - Request card
* [get_card](#get_card) - Get issued card
* [get_card_full_details](#get_card_full_details) - Get full card details
* [list_cards](#list_cards) - List issued cards
* [update_card](#update_card) - Update issued card

## request_card

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
        funding_wallet_id='adipisci',
        memo='dolorum',
        type=shared.IssuedCardType.SINGLE_USE,
    ),
    account_id='1108e0ad-cf4b-4921-879f-ce953f73ef7f',
)

res = s.card_issuing.request_card(req)

if res.issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.PostRequestCardRequest](../../models/operations/postrequestcardrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.PostRequestCardResponse](../../models/operations/postrequestcardresponse.md)**


## get_card

Retrieve a single issued card associated with a Moov account
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

req = operations.GetIssuedCardRequest(
    account_id='bc7abd74-dd39-4c0f-9d2c-ff7c70a45626',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card_issuing.get_card(req)

if res.issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `request`                                                                          | [operations.GetIssuedCardRequest](../../models/operations/getissuedcardrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[operations.GetIssuedCardResponse](../../models/operations/getissuedcardresponse.md)**


## get_card_full_details

Get issued card with PAN, CVV, and expiration. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/issued-cards.read-secure` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetFullIssuedCardRequest(
    account_id='d436813f-16d9-4f5f-8e6c-556146c3e250',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card_issuing.get_card_full_details(req)

if res.full_issued_card is not None:
    # handle response
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `request`                                                                                  | [operations.GetFullIssuedCardRequest](../../models/operations/getfullissuedcardrequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[operations.GetFullIssuedCardResponse](../../models/operations/getfullissuedcardresponse.md)**


## list_cards

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

req = operations.ListIssuedCardsRequest(
    account_id='fb008c42-e141-4aac-b66c-8dd6b1442907',
    count=301598,
    skip=487935,
    states=shared.IssuedCardState.INACTIVE,
)

res = s.card_issuing.list_cards(req)

if res.issued_cards is not None:
    # handle response
```

### Parameters

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `request`                                                                              | [operations.ListIssuedCardsRequest](../../models/operations/listissuedcardsrequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[operations.ListIssuedCardsResponse](../../models/operations/listissuedcardsresponse.md)**


## update_card

Update a Moov issued card
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

req = operations.UpdateIssuedCardRequest(
    update_issued_card=shared.UpdateIssuedCard(
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
        memo='esse',
        state=shared.IssuedCardState.PENDING_VERIFICATION,
    ),
    account_id='a7bd466d-28c1-40ab-bcdc-a4251904e523',
    issued_card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.card_issuing.update_card(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `request`                                                                                | [operations.UpdateIssuedCardRequest](../../models/operations/updateissuedcardrequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[operations.UpdateIssuedCardResponse](../../models/operations/updateissuedcardresponse.md)**

