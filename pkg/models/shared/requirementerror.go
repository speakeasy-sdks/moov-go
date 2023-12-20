// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// RequirementError - Describes an error fulfilling a Requirement
type RequirementError struct {
	ErrorCode *RequirementErrorCode `json:"errorCode,omitempty"`
	// The unique ID of what the requirement is asking to be filled out.
	Requirement *RequirementID `json:"requirement,omitempty"`
}

func (o *RequirementError) GetErrorCode() *RequirementErrorCode {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *RequirementError) GetRequirement() *RequirementID {
	if o == nil {
		return nil
	}
	return o.Requirement
}
