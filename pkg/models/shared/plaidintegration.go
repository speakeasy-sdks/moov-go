// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// PlaidIntegration - The details of a Plaid processor integration for a linked funding source. <br><br> `sandbox` - When linking a bank account to a `sandbox` account using a Plaid processor token a default bank account response will be used. The following default data will be used to generate the bank account in this flow:
// ```
//
//	RoutingNumber: "011401533",
//	AccountNumber: "1111222233330000",
//	AccountType:   "checking",
//	Mask:          "0000"
//
// ```
type PlaidIntegration struct {
	Token *string `json:"token,omitempty"`
}

func (o *PlaidIntegration) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}
