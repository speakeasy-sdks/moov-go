// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type UpdateApplePayMerchantDomains struct {
	// A list of fully qualified top-level or sub-domain names to be added. Items must be unique.
	//
	AddDomains []string `json:"addDomains,omitempty"`
	// A list of previously added fully qualified top-level or sub-domain names to be removed.
	//
	RemoveDomains []string `json:"removeDomains,omitempty"`
}

func (o *UpdateApplePayMerchantDomains) GetAddDomains() []string {
	if o == nil {
		return nil
	}
	return o.AddDomains
}

func (o *UpdateApplePayMerchantDomains) GetRemoveDomains() []string {
	if o == nil {
		return nil
	}
	return o.RemoveDomains
}