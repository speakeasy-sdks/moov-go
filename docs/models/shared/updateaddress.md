# UpdateAddress

Provide address fields as necessary to patch the currently saved address. 
Omit any fields that should not be changed.



## Fields

| Field                    | Type                     | Required                 | Description              | Example                  |
| ------------------------ | ------------------------ | ------------------------ | ------------------------ | ------------------------ |
| `AdditionalProperties`   | map[string]*interface{}* | :heavy_minus_sign:       | N/A                      |                          |
| `AddressLine1`           | **string*                | :heavy_minus_sign:       | N/A                      | 123 Main Street          |
| `AddressLine2`           | **string*                | :heavy_minus_sign:       | N/A                      | Apt 302                  |
| `City`                   | **string*                | :heavy_minus_sign:       | N/A                      | Boulder                  |
| `Country`                | **string*                | :heavy_minus_sign:       | N/A                      | US                       |
| `PostalCode`             | **string*                | :heavy_minus_sign:       | N/A                      | 80301                    |
| `StateOrProvince`        | **string*                | :heavy_minus_sign:       | N/A                      | CO                       |