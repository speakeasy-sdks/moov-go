# cards

## Overview

You can link credit or debit cards to Moov accounts. You can use a card as a source for making transfers, which charges the card. To link a card to a Moov account and avoid some of the burden of PCI compliance, use the [card link Moov Drop](https://docs.moov.io/moovjs/drops/card-link). You cannot add a card via the Dashboard. If you're linking a card via API, you must provide Moov with a copy of your PCI attestation of compliance. When testing cards, use the designated [card numbers for test mode](https://docs.moov.io/guides/set-up-your-account/test-mode/#cards). You must contact Moov before going live in production with cards. Read our guide on [cards](https://docs.moov.io/guides/sources/cards/) for more information.

### Available Operations

* [link_apple_pay_token](#link_apple_pay_token) - Link Apple Pay token
* [link_card](#link_card) - Link card
* [list_cards](#list_cards) - List cards
* [create_apple_pay_session](#create_apple_pay_session) - Create Apple Pay session
* [delete](#delete) - Disable card
* [get](#get) - Get card
* [list_apple_pay_domains](#list_apple_pay_domains) - Get Apple Pay domains
* [register_apple_pay_domain](#register_apple_pay_domain) - Register Apple Pay domains
* [update](#update) - Update card
* [update_apple_pay_domains](#update_apple_pay_domains) - Update Apple Pay domains

## link_apple_pay_token

Connect an Apple Pay token to the specified account.
The `token` data is defined by Apple Pay and should be passed through from Apple Pay's response unmodified.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostLinkApplePayTokenRequest(
    link_apple_pay=shared.LinkApplePay(
        token=shared.LinkApplePayToken(
            payment_data=shared.LinkApplePayTokenPaymentData(
                data='3+f4oOTwPa6f1UZ6tG...CE=',
                header=shared.LinkApplePayTokenPaymentDataHeader(
                    ephemeral_public_key='MFkwEK...Md==',
                    public_key_hash='l0CnXdMv...D1I=',
                    transaction_id='32b...4f3',
                ),
                signature='MIAGCSqGSIb3DQ.AAAA==',
                version='EC_v1',
            ),
            payment_method=shared.LinkApplePayTokenPaymentMethod(
                display_name='Visa 1234',
                network='Visa',
                type='debit',
            ),
            transaction_identifier='32b...4f3',
        ),
    ),
    account_id='c7e0bc71-78e4-4796-b2a7-0c688282aa48',
)

res = s.cards.link_apple_pay_token(req)

if res.linked_apple_pay_payment_method is not None:
    # handle response
```

### Parameters

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `request`                                                                                          | [operations.PostLinkApplePayTokenRequest](../../models/operations/postlinkapplepaytokenrequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[operations.PostLinkApplePayTokenResponse](../../models/operations/postlinkapplepaytokenresponse.md)**


## link_card

Link a card to an existing Moov account. Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostLinkCardRequest(
    card_request=shared.CardRequest(
        billing_address=shared.Address(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        card_cvv='0123',
        card_number='explicabo',
        card_on_file=False,
        expiration=shared.CardExpiration(
            month='01',
            year='21',
        ),
        holder_name='Jules Jackson',
        merchant_account_id='562f222e-9817-4ee1-bcbe-61e6b7b95bc0',
    ),
    x_wait_for=shared.SchemasWaitFor.PAYMENT_METHOD,
    account_id='ab3c20c4-f378-49fd-871f-99dd2efd121a',
)

res = s.cards.link_card(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.PostLinkCardRequest](../../models/operations/postlinkcardrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |
| `server_url`                                                                     | *Optional[str]*                                                                  | :heavy_minus_sign:                                                               | An optional server URL to use.                                                   |


### Response

**[operations.PostLinkCardResponse](../../models/operations/postlinkcardresponse.md)**


## list_cards

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
    account_id='a6f1e674-bdb0-44f1-9756-082d68ea19f1',
)

res = s.cards.list_cards(req)

if res.cards is not None:
    # handle response
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `request`                                                                        | [operations.GetListCardsRequest](../../models/operations/getlistcardsrequest.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[operations.GetListCardsResponse](../../models/operations/getlistcardsresponse.md)**


## create_apple_pay_session

Create a session with Apple Pay to facilitate a payment.
A successful response from this endpoint should be passed through to Apple Pay unchanged.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostApplePaySessionRequest(
    create_apple_pay_session=shared.CreateApplePaySession(
        display_name='Example Merchant',
        domain='checkout.classbooker.dev',
    ),
    account_id='d1705133-9d08-4086-a184-0394c26071f9',
)

res = s.cards.create_apple_pay_session(req)

if res.apple_pay_session is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.PostApplePaySessionRequest](../../models/operations/postapplepaysessionrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.PostApplePaySessionResponse](../../models/operations/postapplepaysessionresponse.md)**


## delete

Disables a card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteCardRequest(
    account_id='3f5f0642-dac7-4af5-95cc-413aa63aae8d',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.cards.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.DeleteCardRequest](../../models/operations/deletecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.DeleteCardResponse](../../models/operations/deletecardresponse.md)**


## get

Fetch a specific card associated with a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetCardRequest(
    account_id='67864dbb-675f-4d5e-a0b3-75ed4f6fbee4',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.cards.get(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                              | Type                                                                   | Required                                                               | Description                                                            |
| ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- | ---------------------------------------------------------------------- |
| `request`                                                              | [operations.GetCardRequest](../../models/operations/getcardrequest.md) | :heavy_check_mark:                                                     | The request object to use for the request.                             |


### Response

**[operations.GetCardResponse](../../models/operations/getcardresponse.md)**


## list_apple_pay_domains

Get domains registered with Apple Pay.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetApplePayMerchantDomainsRequest(
    account_id='1f33317f-e35b-460e-b1ea-426555ba3c28',
)

res = s.cards.list_apple_pay_domains(req)

if res.apple_pay_merchant_domains is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                    | [operations.GetApplePayMerchantDomainsRequest](../../models/operations/getapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[operations.GetApplePayMerchantDomainsResponse](../../models/operations/getapplepaymerchantdomainsresponse.md)**


## register_apple_pay_domain

Add domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PostApplePayMerchantDomainsRequest(
    register_apple_pay_merchant_domains=shared.RegisterApplePayMerchantDomains(
        display_name='Example Merchant',
        domains=[
            'dolore',
            'aliquam',
        ],
    ),
    account_id='ed53b88f-3a8d-48f5-80b2-f2fb7b194a27',
)

res = s.cards.register_apple_pay_domain(req)

if res.apple_pay_merchant_domains is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `request`                                                                                                      | [operations.PostApplePayMerchantDomainsRequest](../../models/operations/postapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[operations.PostApplePayMerchantDomainsResponse](../../models/operations/postapplepaymerchantdomainsresponse.md)**


## update

Update a Linked Card and/or resubmit it for verification. If a value is provided for CVV, 
a new verification ($0 authorization) will be submitted for the card. Updating the expiration date or 
address will update the information stored on file for the card but will not be verified. 
Only use this endpoint if you have provided Moov with a copy of your PCI attestation of compliance. 
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/cards.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.UpdateCardRequest(
    card_update_request=shared.CardUpdateRequest(
        billing_address=shared.UpdateAddress(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        card_cvv='123',
        card_on_file=False,
        expiration=shared.UpdateCardExpiration(
            month='01',
            year='21',
        ),
    ),
    account_id='6b26916f-e1f0-48f4-a94e-3698f447f603',
    card_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.cards.update(req)

if res.card is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.UpdateCardRequest](../../models/operations/updatecardrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.UpdateCardResponse](../../models/operations/updatecardresponse.md)**


## update_apple_pay_domains

Add or remove domains to be registered with Apple Pay.
<br><br> Any domains that will be used to accept payments must first be [verified](https://docs.moov.io/guides/money-movement/cards/apple-pay/#step-1-register-your-domains) with Apple.
<br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/apple-pay.write` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.UpdateApplePayMerchantDomainsRequest(
    update_apple_pay_merchant_domains=shared.UpdateApplePayMerchantDomains(
        add_domains=[
            'praesentium',
            'facilis',
            'quaerat',
            'incidunt',
        ],
        remove_domains=[
            'debitis',
            'rem',
        ],
    ),
    account_id='0ca55efd-20e4-457e-9858-b6a89fbe3a5a',
)

res = s.cards.update_apple_pay_domains(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                          | [operations.UpdateApplePayMerchantDomainsRequest](../../models/operations/updateapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[operations.UpdateApplePayMerchantDomainsResponse](../../models/operations/updateapplepaymerchantdomainsresponse.md)**

