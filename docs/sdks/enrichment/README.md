# enrichment

### Available Operations

* [get_address](#get_address) - Get address suggestions
* [get_avatar](#get_avatar) - Get avatar
* [get_industries](#get_industries) - List all industries
* [get_profile](#get_profile) - Get enriched profile

## get_address

Fetch enriched address suggestions. Requires a partial address. 
<br><br> You must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAddressRequest(
    exclude_states='atque',
    include_cities='fugit',
    include_states='ut',
    include_zipcodes='fugiat',
    max_results=30235,
    prefer_cities='culpa',
    prefer_geolocation='expedita',
    prefer_ratio=299643,
    prefer_states='consequatur',
    prefer_zipcodes='esse',
    search='ipsam',
    selected='sit',
    source='voluptatum',
)

res = s.enrichment.get_address(req)

if res.enrichment_addresses is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.GetAddressRequest](../../models/operations/getaddressrequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.GetAddressResponse](../../models/operations/getaddressresponse.md)**


## get_avatar

Get avatar image for an account using a unique ID.
<br><br> To get an avatar, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetAvatarRequest(
    unique_id='quas',
)

res = s.enrichment.get_avatar(req)

if res.get_avatar_200_image_wildcard_binary_string is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.GetAvatarRequest](../../models/operations/getavatarrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.GetAvatarResponse](../../models/operations/getavatarresponse.md)**


## get_industries

Returns a list of all industry titles and their corresponding MCC/SIC/NAICS codes. Results are ordered by title.
<br><br> To list industries, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)


res = s.enrichment.get_industries()

if res.industries is not None:
    # handle response
```


### Response

**[operations.GetIndustriesResponse](../../models/operations/getindustriesresponse.md)**


## get_profile

Fetch enriched profile data. Requires a valid email address. This service is offered in collaboration with Clearbit.
<br><br> To get enriched profile information, you must specify the `/profile-enrichment.read` scope.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetEnrichmentProfileRequest(
    email='Foster.Borer@hotmail.com',
)

res = s.enrichment.get_profile(req)

if res.enrichment_profile is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.GetEnrichmentProfileRequest](../../models/operations/getenrichmentprofilerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.GetEnrichmentProfileResponse](../../models/operations/getenrichmentprofileresponse.md)**

