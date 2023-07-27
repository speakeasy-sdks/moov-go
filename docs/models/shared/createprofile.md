# CreateProfile

Describes the fields available when creating a profile.
If `accountType` is set to `individual`, the `individual` object should be completed. All others should populate `business`.



## Fields

| Field                                                                               | Type                                                                                | Required                                                                            | Description                                                                         |
| ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- | ----------------------------------------------------------------------------------- |
| `business`                                                                          | [Optional[CreateProfileBusiness]](../../models/shared/createprofilebusiness.md)     | :heavy_minus_sign:                                                                  | N/A                                                                                 |
| `individual`                                                                        | [Optional[CreateProfileIndividual]](../../models/shared/createprofileindividual.md) | :heavy_minus_sign:                                                                  | N/A                                                                                 |