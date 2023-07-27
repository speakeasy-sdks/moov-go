# apple_pay_merchant_domains

### Available Operations

* [get](#get) - Get Apple Pay domains
* [post](#post) - Register Apple Pay domains
* [update](#update) - Update Apple Pay domains

## get

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
    account_id='66ae395e-fb9b-4a88-b3a6-6997074ba446',
)

res = s.apple_pay_merchant_domains.get(req)

if res.apple_pay_merchant_domains is not None:
    # handle response
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                    | [operations.GetApplePayMerchantDomainsRequest](../../models/operations/getapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[operations.GetApplePayMerchantDomainsResponse](../../models/operations/getapplepaymerchantdomainsresponse.md)**


## post

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
            'nobis',
            'eum',
            'vero',
        ],
    ),
    account_id='21419598-90af-4a56-be25-16fe4c8b711e',
)

res = s.apple_pay_merchant_domains.post(req)

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
            'expedita',
            'nihil',
        ],
        remove_domains=[
            'quibusdam',
            'sed',
            'saepe',
            'pariatur',
        ],
    ),
    account_id='028921cd-dc69-4260-9fb5-76b0d5f0d30c',
)

res = s.apple_pay_merchant_domains.update(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `request`                                                                                                          | [operations.UpdateApplePayMerchantDomainsRequest](../../models/operations/updateapplepaymerchantdomainsrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[operations.UpdateApplePayMerchantDomainsResponse](../../models/operations/updateapplepaymerchantdomainsresponse.md)**

