// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

type RepresentativeAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *RepresentativeAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *RepresentativeAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *RepresentativeAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *RepresentativeAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *RepresentativeAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *RepresentativeAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

type RepresentativePhone struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

func (o *RepresentativePhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *RepresentativePhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// RepresentativeResponsibilities - Describes the job responsibilities of an individual
type RepresentativeResponsibilities struct {
	// Indicates whether this individual has significant management responsibilities within the business
	IsController *bool `default:"false" json:"isController"`
	// If `true`, this field indicates that this individual has an ownership stake of at least 25% in the business. If the representative does not own at least 25% of the business, this field should be `false`.
	IsOwner  *bool  `default:"false" json:"isOwner"`
	JobTitle string `json:"jobTitle"`
	// The percentage of ownership this individual has in the business (required if `isOwner` is `true`)
	OwnershipPercentage int64 `json:"ownershipPercentage"`
}

func (r RepresentativeResponsibilities) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RepresentativeResponsibilities) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RepresentativeResponsibilities) GetIsController() *bool {
	if o == nil {
		return nil
	}
	return o.IsController
}

func (o *RepresentativeResponsibilities) GetIsOwner() *bool {
	if o == nil {
		return nil
	}
	return o.IsOwner
}

func (o *RepresentativeResponsibilities) GetJobTitle() string {
	if o == nil {
		return ""
	}
	return o.JobTitle
}

func (o *RepresentativeResponsibilities) GetOwnershipPercentage() int64 {
	if o == nil {
		return 0
	}
	return o.OwnershipPercentage
}

// Representative - Describes a business representative
type Representative struct {
	Address *RepresentativeAddress `json:"address,omitempty"`
	// Indicates whether this Representative's birth date has been provided
	BirthDateProvided *bool      `default:"false" json:"birthDateProvided"`
	CreatedOn         time.Time  `json:"createdOn"`
	DisabledOn        *time.Time `json:"disabledOn,omitempty"`
	// Email Address
	Email *string `json:"email,omitempty"`
	// Indicates whether a government ID (SSN, ITIN, etc.) has been provided for this Representative
	GovernmentIDProvided *bool `default:"false" json:"governmentIDProvided"`
	// Name for an individual
	Name  Name                 `json:"name"`
	Phone *RepresentativePhone `json:"phone,omitempty"`
	// UUID v4
	RepresentativeID *string                         `json:"representativeID,omitempty"`
	Responsibilities *RepresentativeResponsibilities `json:"responsibilities,omitempty"`
	UpdatedOn        time.Time                       `json:"updatedOn"`
}

func (r Representative) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *Representative) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Representative) GetAddress() *RepresentativeAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *Representative) GetBirthDateProvided() *bool {
	if o == nil {
		return nil
	}
	return o.BirthDateProvided
}

func (o *Representative) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *Representative) GetDisabledOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.DisabledOn
}

func (o *Representative) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *Representative) GetGovernmentIDProvided() *bool {
	if o == nil {
		return nil
	}
	return o.GovernmentIDProvided
}

func (o *Representative) GetName() Name {
	if o == nil {
		return Name{}
	}
	return o.Name
}

func (o *Representative) GetPhone() *RepresentativePhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *Representative) GetRepresentativeID() *string {
	if o == nil {
		return nil
	}
	return o.RepresentativeID
}

func (o *Representative) GetResponsibilities() *RepresentativeResponsibilities {
	if o == nil {
		return nil
	}
	return o.Responsibilities
}

func (o *Representative) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}
