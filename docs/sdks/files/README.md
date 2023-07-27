# files

## Overview

Files can be used for a multitude of different use cases including but not limited to, individual identity verification and business underwriting. You may need to provide documentation to enable capabilities or to keep capabilities enabled for an account. The maximum file size is 10MB. Each account is allowed a maximum of 10 files. Acceptable file types include csv, jpg, pdf, and png. To learn about uploading files in the Moov Dashboard, read our [file upload guide](https://docs.moov.io/guides/dashboard/accounts/#file-upload).

### Available Operations

* [list](#list) - List files

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
    account_id='6f1e674b-db04-4f15-b560-82d68ea19f1d',
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

