// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type WireLocation struct {
	City  *string `json:"city,omitempty"`
	State *string `json:"state,omitempty"`
}

func (o *WireLocation) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *WireLocation) GetState() *string {
	if o == nil {
		return nil
	}
	return o.State
}