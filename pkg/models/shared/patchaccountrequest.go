// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type PatchAccountRequestCustomerSupportAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *PatchAccountRequestCustomerSupportAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *PatchAccountRequestCustomerSupportAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *PatchAccountRequestCustomerSupportAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PatchAccountRequestCustomerSupportAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PatchAccountRequestCustomerSupportAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PatchAccountRequestCustomerSupportAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

type PatchAccountRequestCustomerSupportPhone struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

func (o *PatchAccountRequestCustomerSupportPhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PatchAccountRequestCustomerSupportPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// PatchAccountRequestCustomerSupport - User-provided information that can be displayed on credit card transactions for customers to use when contacting a customer support team. This data is only allowed on a business account
type PatchAccountRequestCustomerSupport struct {
	Address *PatchAccountRequestCustomerSupportAddress `json:"address,omitempty"`
	// Email Address
	Email   *string                                  `json:"email,omitempty"`
	Phone   *PatchAccountRequestCustomerSupportPhone `json:"phone,omitempty"`
	Website *string                                  `json:"website,omitempty"`
}

func (o *PatchAccountRequestCustomerSupport) GetAddress() *PatchAccountRequestCustomerSupportAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *PatchAccountRequestCustomerSupport) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *PatchAccountRequestCustomerSupport) GetPhone() *PatchAccountRequestCustomerSupportPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *PatchAccountRequestCustomerSupport) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

type PatchAccountRequestProfileBusinessAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *PatchAccountRequestProfileBusinessAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *PatchAccountRequestProfileBusinessAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *PatchAccountRequestProfileBusinessAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PatchAccountRequestProfileBusinessAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PatchAccountRequestProfileBusinessAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PatchAccountRequestProfileBusinessAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

// PatchAccountRequestProfileBusinessBusinessType - The type of entity represented by this Business
type PatchAccountRequestProfileBusinessBusinessType string

const (
	PatchAccountRequestProfileBusinessBusinessTypeSoleProprietorship        PatchAccountRequestProfileBusinessBusinessType = "soleProprietorship"
	PatchAccountRequestProfileBusinessBusinessTypeUnincorporatedAssociation PatchAccountRequestProfileBusinessBusinessType = "unincorporatedAssociation"
	PatchAccountRequestProfileBusinessBusinessTypeTrust                     PatchAccountRequestProfileBusinessBusinessType = "trust"
	PatchAccountRequestProfileBusinessBusinessTypePublicCorporation         PatchAccountRequestProfileBusinessBusinessType = "publicCorporation"
	PatchAccountRequestProfileBusinessBusinessTypePrivateCorporation        PatchAccountRequestProfileBusinessBusinessType = "privateCorporation"
	PatchAccountRequestProfileBusinessBusinessTypeLlc                       PatchAccountRequestProfileBusinessBusinessType = "llc"
	PatchAccountRequestProfileBusinessBusinessTypePartnership               PatchAccountRequestProfileBusinessBusinessType = "partnership"
	PatchAccountRequestProfileBusinessBusinessTypeUnincorporatedNonProfit   PatchAccountRequestProfileBusinessBusinessType = "unincorporatedNonProfit"
	PatchAccountRequestProfileBusinessBusinessTypeIncorporatedNonProfit     PatchAccountRequestProfileBusinessBusinessType = "incorporatedNonProfit"
)

func (e PatchAccountRequestProfileBusinessBusinessType) ToPointer() *PatchAccountRequestProfileBusinessBusinessType {
	return &e
}

func (e *PatchAccountRequestProfileBusinessBusinessType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "soleProprietorship":
		fallthrough
	case "unincorporatedAssociation":
		fallthrough
	case "trust":
		fallthrough
	case "publicCorporation":
		fallthrough
	case "privateCorporation":
		fallthrough
	case "llc":
		fallthrough
	case "partnership":
		fallthrough
	case "unincorporatedNonProfit":
		fallthrough
	case "incorporatedNonProfit":
		*e = PatchAccountRequestProfileBusinessBusinessType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for PatchAccountRequestProfileBusinessBusinessType: %v", v)
	}
}

// PatchAccountRequestProfileBusinessIndustryCodes - Describes industry specific identifiers
type PatchAccountRequestProfileBusinessIndustryCodes struct {
	Mcc   *string `json:"mcc,omitempty"`
	Naics *string `json:"naics,omitempty"`
	Sic   *string `json:"sic,omitempty"`
}

func (o *PatchAccountRequestProfileBusinessIndustryCodes) GetMcc() *string {
	if o == nil {
		return nil
	}
	return o.Mcc
}

func (o *PatchAccountRequestProfileBusinessIndustryCodes) GetNaics() *string {
	if o == nil {
		return nil
	}
	return o.Naics
}

func (o *PatchAccountRequestProfileBusinessIndustryCodes) GetSic() *string {
	if o == nil {
		return nil
	}
	return o.Sic
}

type PatchAccountRequestProfileBusinessPhone struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

func (o *PatchAccountRequestProfileBusinessPhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PatchAccountRequestProfileBusinessPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// PatchAccountRequestProfileBusinessTaxID - An EIN (employer identification number) for the business. For sole proprietors, an SSN can be used as the EIN.
type PatchAccountRequestProfileBusinessTaxID struct {
	Ein *Ein `json:"ein,omitempty"`
}

func (o *PatchAccountRequestProfileBusinessTaxID) GetEin() *Ein {
	if o == nil {
		return nil
	}
	return o.Ein
}

// PatchAccountRequestProfileBusiness - Describes the fields available when patching a business
type PatchAccountRequestProfileBusiness struct {
	Address           *PatchAccountRequestProfileBusinessAddress       `json:"address,omitempty"`
	BusinessType      *PatchAccountRequestProfileBusinessBusinessType  `json:"businessType,omitempty"`
	Description       *string                                          `json:"description,omitempty"`
	DoingBusinessAs   *string                                          `json:"doingBusinessAs,omitempty"`
	Email             *string                                          `json:"email,omitempty"`
	IndustryCodes     *PatchAccountRequestProfileBusinessIndustryCodes `json:"industryCodes,omitempty"`
	LegalBusinessName *string                                          `json:"legalBusinessName,omitempty"`
	OwnersProvided    *bool                                            `json:"ownersProvided,omitempty"`
	Phone             *PatchAccountRequestProfileBusinessPhone         `json:"phone,omitempty"`
	TaxID             *PatchAccountRequestProfileBusinessTaxID         `json:"taxID,omitempty"`
	Website           *string                                          `json:"website,omitempty"`
}

func (o *PatchAccountRequestProfileBusiness) GetAddress() *PatchAccountRequestProfileBusinessAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *PatchAccountRequestProfileBusiness) GetBusinessType() *PatchAccountRequestProfileBusinessBusinessType {
	if o == nil {
		return nil
	}
	return o.BusinessType
}

func (o *PatchAccountRequestProfileBusiness) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *PatchAccountRequestProfileBusiness) GetDoingBusinessAs() *string {
	if o == nil {
		return nil
	}
	return o.DoingBusinessAs
}

func (o *PatchAccountRequestProfileBusiness) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *PatchAccountRequestProfileBusiness) GetIndustryCodes() *PatchAccountRequestProfileBusinessIndustryCodes {
	if o == nil {
		return nil
	}
	return o.IndustryCodes
}

func (o *PatchAccountRequestProfileBusiness) GetLegalBusinessName() *string {
	if o == nil {
		return nil
	}
	return o.LegalBusinessName
}

func (o *PatchAccountRequestProfileBusiness) GetOwnersProvided() *bool {
	if o == nil {
		return nil
	}
	return o.OwnersProvided
}

func (o *PatchAccountRequestProfileBusiness) GetPhone() *PatchAccountRequestProfileBusinessPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *PatchAccountRequestProfileBusiness) GetTaxID() *PatchAccountRequestProfileBusinessTaxID {
	if o == nil {
		return nil
	}
	return o.TaxID
}

func (o *PatchAccountRequestProfileBusiness) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

type PatchAccountRequestProfileIndividualAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *PatchAccountRequestProfileIndividualAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *PatchAccountRequestProfileIndividualAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *PatchAccountRequestProfileIndividualAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *PatchAccountRequestProfileIndividualAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *PatchAccountRequestProfileIndividualAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

// PatchAccountRequestProfileIndividualBirthDate - Birthdate for an individual
type PatchAccountRequestProfileIndividualBirthDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

func (o *PatchAccountRequestProfileIndividualBirthDate) GetDay() int64 {
	if o == nil {
		return 0
	}
	return o.Day
}

func (o *PatchAccountRequestProfileIndividualBirthDate) GetMonth() int64 {
	if o == nil {
		return 0
	}
	return o.Month
}

func (o *PatchAccountRequestProfileIndividualBirthDate) GetYear() int64 {
	if o == nil {
		return 0
	}
	return o.Year
}

type PatchAccountRequestProfileIndividualGovernmentIDItin struct {
	Full     *string `json:"full,omitempty"`
	LastFour *string `json:"lastFour,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualGovernmentIDItin) GetFull() *string {
	if o == nil {
		return nil
	}
	return o.Full
}

func (o *PatchAccountRequestProfileIndividualGovernmentIDItin) GetLastFour() *string {
	if o == nil {
		return nil
	}
	return o.LastFour
}

type PatchAccountRequestProfileIndividualGovernmentIDSsn struct {
	Full     *string `json:"full,omitempty"`
	LastFour *string `json:"lastFour,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualGovernmentIDSsn) GetFull() *string {
	if o == nil {
		return nil
	}
	return o.Full
}

func (o *PatchAccountRequestProfileIndividualGovernmentIDSsn) GetLastFour() *string {
	if o == nil {
		return nil
	}
	return o.LastFour
}

type PatchAccountRequestProfileIndividualGovernmentID struct {
	Itin *PatchAccountRequestProfileIndividualGovernmentIDItin `json:"itin,omitempty"`
	Ssn  *PatchAccountRequestProfileIndividualGovernmentIDSsn  `json:"ssn,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualGovernmentID) GetItin() *PatchAccountRequestProfileIndividualGovernmentIDItin {
	if o == nil {
		return nil
	}
	return o.Itin
}

func (o *PatchAccountRequestProfileIndividualGovernmentID) GetSsn() *PatchAccountRequestProfileIndividualGovernmentIDSsn {
	if o == nil {
		return nil
	}
	return o.Ssn
}

// PatchAccountRequestProfileIndividualName - Name for an individual
type PatchAccountRequestProfileIndividualName struct {
	// Name this person was given. This is usually the the same as first name.
	FirstName *string `json:"firstName,omitempty"`
	// Family name of this person. This is usually the the same as last name.
	LastName *string `json:"lastName,omitempty"`
	// Name this person was given. This is usually the the same as first name.
	MiddleName *string `json:"middleName,omitempty"`
	// Suffix of a given name
	Suffix *string `json:"suffix,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualName) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *PatchAccountRequestProfileIndividualName) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *PatchAccountRequestProfileIndividualName) GetMiddleName() *string {
	if o == nil {
		return nil
	}
	return o.MiddleName
}

func (o *PatchAccountRequestProfileIndividualName) GetSuffix() *string {
	if o == nil {
		return nil
	}
	return o.Suffix
}

type PatchAccountRequestProfileIndividualPhone struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

func (o *PatchAccountRequestProfileIndividualPhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *PatchAccountRequestProfileIndividualPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// PatchAccountRequestProfileIndividual - Describes the fields available when patching an individual
type PatchAccountRequestProfileIndividual struct {
	Address      *PatchAccountRequestProfileIndividualAddress      `json:"address,omitempty"`
	BirthDate    *PatchAccountRequestProfileIndividualBirthDate    `json:"birthDate,omitempty"`
	Email        *string                                           `json:"email,omitempty"`
	GovernmentID *PatchAccountRequestProfileIndividualGovernmentID `json:"governmentID,omitempty"`
	Name         *PatchAccountRequestProfileIndividualName         `json:"name,omitempty"`
	Phone        *PatchAccountRequestProfileIndividualPhone        `json:"phone,omitempty"`
}

func (o *PatchAccountRequestProfileIndividual) GetAddress() *PatchAccountRequestProfileIndividualAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *PatchAccountRequestProfileIndividual) GetBirthDate() *PatchAccountRequestProfileIndividualBirthDate {
	if o == nil {
		return nil
	}
	return o.BirthDate
}

func (o *PatchAccountRequestProfileIndividual) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *PatchAccountRequestProfileIndividual) GetGovernmentID() *PatchAccountRequestProfileIndividualGovernmentID {
	if o == nil {
		return nil
	}
	return o.GovernmentID
}

func (o *PatchAccountRequestProfileIndividual) GetName() *PatchAccountRequestProfileIndividualName {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *PatchAccountRequestProfileIndividual) GetPhone() *PatchAccountRequestProfileIndividualPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

// PatchAccountRequestProfile - Describes the fields available when patching a profile.
// Each object can be patched independent of patching the other fields.
type PatchAccountRequestProfile struct {
	Business   *PatchAccountRequestProfileBusiness   `json:"business,omitempty"`
	Individual *PatchAccountRequestProfileIndividual `json:"individual,omitempty"`
}

func (o *PatchAccountRequestProfile) GetBusiness() *PatchAccountRequestProfileBusiness {
	if o == nil {
		return nil
	}
	return o.Business
}

func (o *PatchAccountRequestProfile) GetIndividual() *PatchAccountRequestProfileIndividual {
	if o == nil {
		return nil
	}
	return o.Individual
}

// PatchAccountRequestSettingsAchPayment - User provided settings to manage ACH payments
type PatchAccountRequestSettingsAchPayment struct {
	// The description that shows up on ACH transactions. This will default to the account's display name on account creation.
	CompanyName *string `json:"companyName,omitempty"`
}

func (o *PatchAccountRequestSettingsAchPayment) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

// PatchAccountRequestSettingsCardPayment - User provided settings to manage card payments. This data is only allowed on a business account
type PatchAccountRequestSettingsCardPayment struct {
	// The description that shows up on credit card transactions. This will default to the accounts display name on account creation.
	StatementDescriptor *string `json:"statementDescriptor,omitempty"`
}

func (o *PatchAccountRequestSettingsCardPayment) GetStatementDescriptor() *string {
	if o == nil {
		return nil
	}
	return o.StatementDescriptor
}

// PatchAccountRequestSettings - User provided settings to manage an account
type PatchAccountRequestSettings struct {
	AchPayment  *PatchAccountRequestSettingsAchPayment  `json:"achPayment,omitempty"`
	CardPayment *PatchAccountRequestSettingsCardPayment `json:"cardPayment,omitempty"`
}

func (o *PatchAccountRequestSettings) GetAchPayment() *PatchAccountRequestSettingsAchPayment {
	if o == nil {
		return nil
	}
	return o.AchPayment
}

func (o *PatchAccountRequestSettings) GetCardPayment() *PatchAccountRequestSettingsCardPayment {
	if o == nil {
		return nil
	}
	return o.CardPayment
}

// PatchAccountRequestTermsOfService - An encrypted value used to record acceptance of Moov's Terms of Service
type PatchAccountRequestTermsOfService struct {
	Token *string `json:"token,omitempty"`
}

func (o *PatchAccountRequestTermsOfService) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// PatchAccountRequest - Describes the fields available when patching a Moov Account
type PatchAccountRequest struct {
	CustomerSupport *PatchAccountRequestCustomerSupport `json:"customerSupport,omitempty"`
	ForeignID       *string                             `json:"foreignID,omitempty"`
	// Free-form key-value pair list. Useful for storing information that is not captured elsewhere.
	Metadata       map[string]string                  `json:"metadata,omitempty"`
	Profile        *PatchAccountRequestProfile        `json:"profile,omitempty"`
	Settings       *PatchAccountRequestSettings       `json:"settings,omitempty"`
	TermsOfService *PatchAccountRequestTermsOfService `json:"termsOfService,omitempty"`
}

func (o *PatchAccountRequest) GetCustomerSupport() *PatchAccountRequestCustomerSupport {
	if o == nil {
		return nil
	}
	return o.CustomerSupport
}

func (o *PatchAccountRequest) GetForeignID() *string {
	if o == nil {
		return nil
	}
	return o.ForeignID
}

func (o *PatchAccountRequest) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *PatchAccountRequest) GetProfile() *PatchAccountRequestProfile {
	if o == nil {
		return nil
	}
	return o.Profile
}

func (o *PatchAccountRequest) GetSettings() *PatchAccountRequestSettings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *PatchAccountRequest) GetTermsOfService() *PatchAccountRequestTermsOfService {
	if o == nil {
		return nil
	}
	return o.TermsOfService
}
