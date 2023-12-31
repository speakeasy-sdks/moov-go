// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type PatchTransfer struct {
	// Free-form key-value pair list. Useful for storing information that is not captured elsewhere.
	Metadata map[string]string `json:"metadata,omitempty"`
}

func (o *PatchTransfer) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}
