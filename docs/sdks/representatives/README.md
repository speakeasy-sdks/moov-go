# representatives

## Overview

We think of a representative as an individual who can take major actions on behalf of a business. A representative can be the business owner, or anyone who owns 25% or more of the business. Some examples of representatives are the CEO, CFO, COO, or president. We require all business accounts to have valid information for at least one representative. Moov accounts must have verified business representatives before a business account can send funds, collect money from other accounts, or store funds in a wallet. To learn more, read our guide on [business representatives](https://docs.moov.io/guides/accounts/business-representatives/).

### Available Operations

* [create](#create) - Create representative
* [delete](#delete) - Delete a representative
* [get](#get) - Get representative
* [list](#list) - List representatives
* [update](#update) - Patch representative

## create

Moov accounts associated with businesses require information regarding individuals who represent the business. You can provide this information by creating a representative. Each account is allowed a maximum of 7 representatives.<br><br> To create a representative, you must specify the `/accounts/{accountID}/representatives.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.CreateRepresentativeRequest(
    create_representative=shared.CreateRepresentative(
        address=shared.CreateRepresentativeAddress(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        birth_date=shared.CreateRepresentativeBirthDate(
            day=9,
            month=11,
            year=1989,
        ),
        email='amanda@classbooker.dev',
        government_id=shared.CreateRepresentativeGovernmentID(
            itin=shared.CreateRepresentativeGovernmentIDItin(
                full='123-45-6789',
                last_four='6789',
            ),
            ssn=shared.CreateRepresentativeGovernmentIDSsn(
                full='123-45-6789',
                last_four='6789',
            ),
        ),
        name=shared.CreateRepresentativeName(
            first_name='Amanda',
            last_name='Yang',
            middle_name='Amanda',
            suffix='Jr',
        ),
        phone=shared.CreateRepresentativePhone(
            country_code='1',
            number='8185551212',
        ),
        responsibilities=shared.CreateRepresentativeResponsibilities(
            is_controller=False,
            is_owner=True,
            job_title='CEO',
            ownership_percentage=38,
        ),
    ),
    account_id='512c1032-648d-4c2f-a151-99ebfd0e9fe6',
)

res = s.representatives.create(req)

if res.representative is not None:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.CreateRepresentativeRequest](../../models/operations/createrepresentativerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.CreateRepresentativeResponse](../../models/operations/createrepresentativeresponse.md)**


## delete

Deletes a business representative associated with a Moov account. <br><br> To use this endpoint, you'll need to specify the `/accounts/{accountID}/representatives.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.DeleteRepresentativeRequest(
    account_id='c632ca3a-ed01-4179-9631-2fde04771778',
    representative_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.representatives.delete(req)

if res.status_code == 200:
    # handle response
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `request`                                                                                        | [operations.DeleteRepresentativeRequest](../../models/operations/deleterepresentativerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[operations.DeleteRepresentativeResponse](../../models/operations/deleterepresentativeresponse.md)**


## get

Retrieve a specific representative associated with a given Moov account. <br><br> To get a representative, you'll need to specify the `/accounts/{accountID}/representatives.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetRepresentativeRequest(
    account_id='ff61d017-4763-460a-95db-6a660659a1ad',
    representative_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.representatives.get(req)

if res.representative is not None:
    # handle response
```

### Parameters

| Parameter                                                                                  | Type                                                                                       | Required                                                                                   | Description                                                                                |
| ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------ |
| `request`                                                                                  | [operations.GetRepresentativeRequest](../../models/operations/getrepresentativerequest.md) | :heavy_check_mark:                                                                         | The request object to use for the request.                                                 |


### Response

**[operations.GetRepresentativeResponse](../../models/operations/getrepresentativeresponse.md)**


## list

A Moov account may have multiple representatives depending on the associated business's ownership and management structure. You can use this method to list all the representatives for a given Moov account. Note that Moov accounts associated with an individual do not have representatives. <br><br> To list representatives, you need to specify the `/accounts/{accountID}/representatives.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListRepresentativesRequest(
    account_id='eaab5851-d6c6-445b-88b6-1891baa0fe1a',
)

res = s.representatives.list(req)

if res.representatives is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.ListRepresentativesRequest](../../models/operations/listrepresentativesrequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.ListRepresentativesResponse](../../models/operations/listrepresentativesresponse.md)**


## update

If a representative's information has changed you can patch the information associated with a specific representative ID.  
To patch a representative, you'll need to specify the `/accounts/{accountID}/representatives.write` scope and provide the changed information.

When **can** profile data be updated:  
  + For unverified representatives, all profile data can be edited.
  + During the verification process, missing or incomplete profile data can be edited.
  + Verified representatives can only add missing profile data.  

  When **can't** profile data be updated:  
  + Verified representatives cannot change any existing profile data.  

If you need to update information in a locked state, please contact Moov support.


### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.PatchRepresentativeRequest(
    patch_representative_request=shared.PatchRepresentativeRequest(
        address=shared.PatchRepresentativeRequestAddress(
            address_line1='123 Main Street',
            address_line2='Apt 302',
            city='Boulder',
            country='US',
            postal_code='80301',
            state_or_province='CO',
        ),
        birth_date=shared.PatchRepresentativeRequestBirthDate(
            day=9,
            month=11,
            year=1989,
        ),
        email='amanda@classbooker.dev',
        government_id=shared.PatchRepresentativeRequestGovernmentID(
            itin=shared.PatchRepresentativeRequestGovernmentIDItin(
                full='123-45-6789',
                last_four='6789',
            ),
            ssn=shared.PatchRepresentativeRequestGovernmentIDSsn(
                full='123-45-6789',
                last_four='6789',
            ),
        ),
        name=shared.PatchRepresentativeRequestName(
            first_name='Amanda',
            last_name='Yang',
            middle_name='Amanda',
            suffix='Jr',
        ),
        phone=shared.PatchRepresentativeRequestPhone(
            country_code='1',
            number='8185551212',
        ),
        responsibilities=shared.PatchRepresentativeRequestResponsibilities(
            is_controller=False,
            is_owner=True,
            job_title='CEO',
            ownership_percentage=38,
        ),
    ),
    account_id='de008e6f-8c5f-4350-98cd-b5a341814301',
    representative_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.representatives.update(req)

if res.representative is not None:
    # handle response
```

### Parameters

| Parameter                                                                                      | Type                                                                                           | Required                                                                                       | Description                                                                                    |
| ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------- |
| `request`                                                                                      | [operations.PatchRepresentativeRequest](../../models/operations/patchrepresentativerequest.md) | :heavy_check_mark:                                                                             | The request object to use for the request.                                                     |


### Response

**[operations.PatchRepresentativeResponse](../../models/operations/patchrepresentativeresponse.md)**

