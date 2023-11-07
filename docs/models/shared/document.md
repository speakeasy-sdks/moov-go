# Document

Describes an uploaded file


## Fields

| Field                                                         | Type                                                          | Required                                                      | Description                                                   | Example                                                       |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ContentType`                                                 | *string*                                                      | :heavy_check_mark:                                            | N/A                                                           | application/pdf                                               |
| `DocumentID`                                                  | *string*                                                      | :heavy_check_mark:                                            | A unique identifier for this document                         | e210a9d6                                                      |
| `ParseErrors`                                                 | []*string*                                                    | :heavy_minus_sign:                                            | Optional array of errors encountered dring automated parsing. |                                                               |
| `Type`                                                        | [shared.Type](../../models/shared/type.md)                    | :heavy_check_mark:                                            | N/A                                                           |                                                               |
| `UploadedAt`                                                  | [time.Time](https://pkg.go.dev/time#Time)                     | :heavy_check_mark:                                            | N/A                                                           |                                                               |