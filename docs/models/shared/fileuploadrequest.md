# FileUploadRequest

Request to attach a file to an account.


## Fields

| Field                                                                 | Type                                                                  | Required                                                              | Description                                                           |
| --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- | --------------------------------------------------------------------- |
| `file`                                                                | [FileUploadRequestFile](../../models/shared/fileuploadrequestfile.md) | :heavy_check_mark:                                                    | The file to be added. Valid types are [csv, png, jpeg, pdf].          |
| `file_purpose`                                                        | [FilePurpose](../../models/shared/filepurpose.md)                     | :heavy_check_mark:                                                    | The file purpose                                                      |