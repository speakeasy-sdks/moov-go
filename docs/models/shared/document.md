# Document

Describes an uploaded file


## Fields

| Field                                                                | Type                                                                 | Required                                                             | Description                                                          | Example                                                              |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `content_type`                                                       | *str*                                                                | :heavy_check_mark:                                                   | N/A                                                                  | application/pdf                                                      |
| `document_id`                                                        | *str*                                                                | :heavy_check_mark:                                                   | A unique identifier for this document                                | e210a9d6                                                             |
| `parse_errors`                                                       | list[*str*]                                                          | :heavy_minus_sign:                                                   | Optional array of errors encountered dring automated parsing.        |                                                                      |
| `type`                                                               | [DocumentType](../../models/shared/documenttype.md)                  | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |
| `uploaded_at`                                                        | [date](https://docs.python.org/3/library/datetime.html#date-objects) | :heavy_check_mark:                                                   | N/A                                                                  |                                                                      |