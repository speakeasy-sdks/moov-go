# File

Describes a file linked to a Moov account


## Fields

| Field                                              | Type                                               | Required                                           | Description                                        | Example                                            |
| -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- | -------------------------------------------------- |
| `AccountID`                                        | **string*                                          | :heavy_minus_sign:                                 | UUID v4                                            | ec7e1848-dc80-4ab0-8827-dd7fc0737b43               |
| `CreatedOn`                                        | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |                                                    |
| `FileID`                                           | **string*                                          | :heavy_minus_sign:                                 | UUID v4                                            | ec7e1848-dc80-4ab0-8827-dd7fc0737b43               |
| `FileName`                                         | **string*                                          | :heavy_minus_sign:                                 | N/A                                                | logo.png                                           |
| `FilePurpose`                                      | [*FilePurpose](../../models/shared/filepurpose.md) | :heavy_minus_sign:                                 | The file purpose                                   |                                                    |
| `FileSizeBytes`                                    | **int64*                                           | :heavy_minus_sign:                                 | N/A                                                | 1024                                               |
| `FileStatus`                                       | [*FileStatus](../../models/shared/filestatus.md)   | :heavy_minus_sign:                                 | The file status                                    |                                                    |
| `UpdatedOn`                                        | [*time.Time](https://pkg.go.dev/time#Time)         | :heavy_minus_sign:                                 | N/A                                                |                                                    |