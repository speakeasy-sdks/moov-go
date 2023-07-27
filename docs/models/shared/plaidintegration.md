# PlaidIntegration

The details of a Plaid processor integration for a linked funding source. <br><br> `sandbox` - When linking a bank account to a `sandbox` account using a Plaid processor token a default bank account response will be used. The following default data will be used to generate the bank account in this flow:
```
  RoutingNumber: "011401533",
  AccountNumber: "1111222233330000",
  AccountType:   "checking",
  Mask:          "0000"
```



## Fields

| Field              | Type               | Required           | Description        |
| ------------------ | ------------------ | ------------------ | ------------------ |
| `token`            | *Optional[str]*    | :heavy_minus_sign: | N/A                |