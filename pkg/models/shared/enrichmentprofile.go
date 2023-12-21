// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// EnrichmentProfile - Describes an enriched business profile
type EnrichmentProfile struct {
	// Describes a company
	Business *EnrichedBusiness `json:"business,omitempty"`
}

func (o *EnrichmentProfile) GetBusiness() *EnrichedBusiness {
	if o == nil {
		return nil
	}
	return o.Business
}
