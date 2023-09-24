# Files

## Overview

Files can be used for a multitude of different use cases including but not limited to, individual identity verification and business underwriting. You may need to provide documentation to enable capabilities or to keep capabilities enabled for an account. The maximum file size is 10MB. Each account is allowed a maximum of 10 files. Acceptable file types include csv, jpg, pdf, and png. To learn about uploading files in the Moov Dashboard, read our [file upload guide](https://docs.moov.io/guides/dashboard/accounts/#file-upload).

### Available Operations

* [Get](#get) - Get File Details
* [List](#list) - List files
* [Upload](#upload) - Upload File

## Get

Retrieve file details associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "e973e922-a57a-415b-a3e0-60807e2b6e3a"
    fileID := "ec7e1848-dc80-4ab0-8827-dd7fc0737b43"

    ctx := context.Background()
    res, err := s.Files.Get(ctx, accountID, fileID)
    if err != nil {
        log.Fatal(err)
    }

    if res.File != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           | Example                                               |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |                                                       |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |                                                       |
| `fileID`                                              | *string*                                              | :heavy_check_mark:                                    | ID of the file                                        | ec7e1848-dc80-4ab0-8827-dd7fc0737b43                  |


### Response

**[*operations.GetFileDetailsResponse](../../models/operations/getfiledetailsresponse.md), error**


## List

List all the files associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.read` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    accountID := "b8845f05-97a6-40ff-aa54-a31e94764a3e"

    ctx := context.Background()
    res, err := s.Files.List(ctx, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.Files != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `accountID`                                           | *string*                                              | :heavy_check_mark:                                    | ID of the account                                     |


### Response

**[*operations.ListFilesResponse](../../models/operations/listfilesresponse.md), error**


## Upload

Upload a file and link it to the provided Moov account. The maximum file size is 10MB. Each account is allowed a maximum of 10 files. Acceptable file types include csv, jpg, pdf, and png. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/files.write` scope.

### Example Usage

```go
package main

import(
	"context"
	"log"
	moovgo "github.com/speakeasy-sdks/moov-go"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
)

func main() {
    s := moovgo.New(
        moovgo.WithSecurity(shared.Security{
            AccessToken: moovgo.String(""),
        }),
    )
    fileUploadRequest := shared.FileUploadRequest{
        File: shared.FileUploadRequestFile{
            Content: []byte("laudantium"),
            File: "eum",
        },
        FilePurpose: shared.FilePurposeBusinessVerification,
    }
    accountID := "e7956f92-51a5-4a9d-a660-ff57bfaad4f9"

    ctx := context.Background()
    res, err := s.Files.Upload(ctx, fileUploadRequest, accountID)
    if err != nil {
        log.Fatal(err)
    }

    if res.File != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `fileUploadRequest`                                                  | [shared.FileUploadRequest](../../models/shared/fileuploadrequest.md) | :heavy_check_mark:                                                   | N/A                                                                  |
| `accountID`                                                          | *string*                                                             | :heavy_check_mark:                                                   | ID of the account                                                    |


### Response

**[*operations.UploadFileResponse](../../models/operations/uploadfileresponse.md), error**

