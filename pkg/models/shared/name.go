// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// Name for an individual
type Name struct {
	// Name this person was given. This is usually the the same as first name.
	FirstName *string `json:"firstName,omitempty"`
	// Family name of this person. This is usually the the same as last name.
	LastName *string `json:"lastName,omitempty"`
	// Name this person was given. This is usually the the same as first name.
	MiddleName *string `json:"middleName,omitempty"`
	// Suffix of a given name
	Suffix *string `json:"suffix,omitempty"`
}

func (o *Name) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *Name) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *Name) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *Name) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}
