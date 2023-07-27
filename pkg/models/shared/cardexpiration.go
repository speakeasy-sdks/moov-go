// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CardExpiration - The expiration date of the linked card or token
type CardExpiration struct {
	Month *string `json:"month,omitempty"`
	Year  *string `json:"year,omitempty"`
}

func (o *CardExpiration) GetMonth() *string {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *CardExpiration) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}