# files

## Overview

Files can be used for a multitude of different use cases including but not limited to, individual identity verification and business underwriting. You may need to provide documentation to enable capabilities or to keep capabilities enabled for an account. The maximum file size is 10MB. Each account is allowed a maximum of 10 files. Acceptable file types include csv, jpg, pdf, and png. To learn about uploading files in the Moov Dashboard, read our [file upload guide](https://docs.moov.io/guides/dashboard/accounts/#file-upload).

### Available Operations

* [get](#get) - Get File Details
* [list](#list) - List files
* [upload](#upload) - Upload File

## get

Retrieve file details associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.GetFileDetailsRequest(
    account_id='2065e904-f3b1-4194-b8ab-f603a79f9dfe',
    file_id='ec7e1848-dc80-4ab0-8827-dd7fc0737b43',
)

res = s.files.get(req)

if res.file is not None:
    # handle response
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `request`                                                                            | [operations.GetFileDetailsRequest](../../models/operations/getfiledetailsrequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[operations.GetFileDetailsResponse](../../models/operations/getfiledetailsresponse.md)**


## list

List all the files associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.read` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.ListFilesRequest(
    account_id='0ab7da8a-50ce-4187-b86b-c173d689eee9',
)

res = s.files.list(req)

if res.files is not None:
    # handle response
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `request`                                                                  | [operations.ListFilesRequest](../../models/operations/listfilesrequest.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[operations.ListFilesResponse](../../models/operations/listfilesresponse.md)**


## upload

Upload a file and link it to the provided Moov account. The maximum file size is 10MB. Each account is allowed a maximum of 10 files. Acceptable file types include csv, jpg, pdf, and png. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.write` scope.

### Example Usage

```python
import petstore
from petstore.models import operations, shared

s = petstore.Petstore(
    security=shared.Security(
        access_token="",
    ),
)

req = operations.UploadFileRequest(
    file_upload_request=shared.FileUploadRequest(
        file=shared.FileUploadRequestFile(
            content='minima'.encode(),
            file='aspernatur',
        ),
        file_purpose=shared.FilePurpose.BUSINESS_VERIFICATION,
    ),
    account_id='f8d986e8-81ea-4d4f-8e10-12563f94e29e',
)

res = s.files.upload(req)

if res.file is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.UploadFileRequest](../../models/operations/uploadfilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.UploadFileResponse](../../models/operations/uploadfileresponse.md)**

