// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Requirement - Represents individual and business data necessary to facilitate the enabling of a capability for an account
type Requirement struct {
	CurrentlyDue []RequirementID    `json:"currentlyDue,omitempty"`
	Errors       []RequirementError `json:"errors,omitempty"`
}

func (o *Requirement) GetCurrentlyDue() []RequirementID {
	if o == nil {
		return nil
	}
	return o.CurrentlyDue
}

func (o *Requirement) GetErrors() []RequirementError {
	if o == nil {
		return nil
	}
	return o.Errors
}