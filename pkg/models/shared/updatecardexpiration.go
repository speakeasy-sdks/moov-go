// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// UpdateCardExpiration - Provide expiration date fields as necessary to patch the currently saved date.
// Omit any fields that should not be changed.
type UpdateCardExpiration struct {
	Month *string `json:"month,omitempty"`
	Year  *string `json:"year,omitempty"`
}

func (o *UpdateCardExpiration) GetMonth() *string {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *UpdateCardExpiration) GetYear() *string {
	if o == nil {
		return nil
	}
	return o.Year
}
