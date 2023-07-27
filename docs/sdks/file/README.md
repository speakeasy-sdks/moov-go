# file

### Available Operations

* [upload](#upload) - Upload File

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
            content='atque'.encode(),
            file='explicabo',
        ),
        file_purpose=shared.FilePurpose.BUSINESS_VERIFICATION,
    ),
    account_id='62f222e9-817e-4e17-8be6-1e6b7b95bc0a',
)

res = s.file.upload(req)

if res.file is not None:
    # handle response
```

### Parameters

| Parameter                                                                    | Type                                                                         | Required                                                                     | Description                                                                  |
| ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- | ---------------------------------------------------------------------------- |
| `request`                                                                    | [operations.UploadFileRequest](../../models/operations/uploadfilerequest.md) | :heavy_check_mark:                                                           | The request object to use for the request.                                   |


### Response

**[operations.UploadFileResponse](../../models/operations/uploadfileresponse.md)**

