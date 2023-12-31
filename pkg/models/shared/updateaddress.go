// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateAddress - Provide address fields as necessary to patch the currently saved address.
// Omit any fields that should not be changed.
type UpdateAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *UpdateAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *UpdateAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *UpdateAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *UpdateAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *UpdateAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *UpdateAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}
