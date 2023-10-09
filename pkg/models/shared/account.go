// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

type AccountCustomerSupportAddress struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	AddressLine1         *string                `json:"addressLine1,omitempty"`
	AddressLine2         *string                `json:"addressLine2,omitempty"`
	City                 *string                `json:"city,omitempty"`
	Country              *string                `json:"country,omitempty"`
	PostalCode           *string                `json:"postalCode,omitempty"`
	StateOrProvince      *string                `json:"stateOrProvince,omitempty"`
}

func (a AccountCustomerSupportAddress) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountCustomerSupportAddress) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountCustomerSupportAddress) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountCustomerSupportAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *AccountCustomerSupportAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *AccountCustomerSupportAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *AccountCustomerSupportAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *AccountCustomerSupportAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *AccountCustomerSupportAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

type AccountCustomerSupportPhone struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	CountryCode          *string                `json:"countryCode,omitempty"`
	Number               *string                `json:"number,omitempty"`
}

func (a AccountCustomerSupportPhone) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountCustomerSupportPhone) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountCustomerSupportPhone) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountCustomerSupportPhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *AccountCustomerSupportPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// AccountCustomerSupport - User-provided information that can be displayed on credit card transactions for customers to use when contacting a customer support team. This data is only allowed on a business account
type AccountCustomerSupport struct {
	AdditionalProperties map[string]interface{}         `additionalProperties:"true" json:"-"`
	Address              *AccountCustomerSupportAddress `json:"address,omitempty"`
	// Email Address
	Email   *string                      `json:"email,omitempty"`
	Phone   *AccountCustomerSupportPhone `json:"phone,omitempty"`
	Website *string                      `json:"website,omitempty"`
}

func (a AccountCustomerSupport) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountCustomerSupport) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountCustomerSupport) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountCustomerSupport) GetAddress() *AccountCustomerSupportAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *AccountCustomerSupport) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *AccountCustomerSupport) GetPhone() *AccountCustomerSupportPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *AccountCustomerSupport) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

// AccountSettingsAchPayment - User provided settings to manage ACH payments
type AccountSettingsAchPayment struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The description that shows up on ACH transactions. This will default to the account's display name on account creation.
	CompanyName *string `json:"companyName,omitempty"`
}

func (a AccountSettingsAchPayment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountSettingsAchPayment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountSettingsAchPayment) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountSettingsAchPayment) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

// AccountSettingsCardPayment - User provided settings to manage card payments. This data is only allowed on a business account
type AccountSettingsCardPayment struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The description that shows up on credit card transactions. This will default to the accounts display name on account creation.
	StatementDescriptor *string `json:"statementDescriptor,omitempty"`
}

func (a AccountSettingsCardPayment) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountSettingsCardPayment) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountSettingsCardPayment) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountSettingsCardPayment) GetStatementDescriptor() *string {
	if o == nil {
		return nil
	}
	return o.StatementDescriptor
}

// AccountSettings - User provided settings to manage an account
type AccountSettings struct {
	AdditionalProperties map[string]interface{}      `additionalProperties:"true" json:"-"`
	AchPayment           *AccountSettingsAchPayment  `json:"achPayment,omitempty"`
	CardPayment          *AccountSettingsCardPayment `json:"cardPayment,omitempty"`
}

func (a AccountSettings) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountSettings) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountSettings) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountSettings) GetAchPayment() *AccountSettingsAchPayment {
	if o == nil {
		return nil
	}
	return o.AchPayment
}

func (o *AccountSettings) GetCardPayment() *AccountSettingsCardPayment {
	if o == nil {
		return nil
	}
	return o.CardPayment
}

// AccountTermsOfService - Describes the acceptance of the Terms of Service
type AccountTermsOfService struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	AcceptedDate         time.Time              `json:"acceptedDate"`
	AcceptedIP           string                 `json:"acceptedIP"`
}

func (a AccountTermsOfService) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *AccountTermsOfService) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *AccountTermsOfService) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *AccountTermsOfService) GetAcceptedDate() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.AcceptedDate
}

func (o *AccountTermsOfService) GetAcceptedIP() string {
	if o == nil {
		return ""
	}
	return o.AcceptedIP
}

// Account - Describes a Moov account
type Account struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// UUID v4
	AccountID string `json:"accountID"`
	// The type of entity represented by this Account
	AccountType     AccountType             `json:"accountType"`
	CreatedOn       time.Time               `json:"createdOn"`
	CustomerSupport *AccountCustomerSupport `json:"customerSupport,omitempty"`
	DisconnectedOn  *time.Time              `json:"disconnectedOn,omitempty"`
	DisplayName     *string                 `json:"displayName,omitempty"`
	// Serves as an optional alias from a foreign/external system which can be used to reference this resource
	ForeignID *string `json:"foreignID,omitempty"`
	// Free-form key-value pair list. Useful for storing information that is not captured elsewhere.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The mode this account is allowed to be used within.
	Mode *Mode `json:"mode,omitempty"`
	// Describes a Moov Account Profile
	Profile        Profile                `json:"profile"`
	Settings       *AccountSettings       `json:"settings,omitempty"`
	TermsOfService *AccountTermsOfService `json:"termsOfService,omitempty"`
	UpdatedOn      time.Time              `json:"updatedOn"`
	// Describes identity verification status and relevant identity verification documents
	Verification *Verification `json:"verification,omitempty"`
}

func (a Account) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(a, "", false)
}

func (a *Account) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &a, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Account) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *Account) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *Account) GetAccountType() AccountType {
	if o == nil {
		return AccountType("")
	}
	return o.AccountType
}

func (o *Account) GetCreatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.CreatedOn
}

func (o *Account) GetCustomerSupport() *AccountCustomerSupport {
	if o == nil {
		return nil
	}
	return o.CustomerSupport
}

func (o *Account) GetDisconnectedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.DisconnectedOn
}

func (o *Account) GetDisplayName() *string {
	if o == nil {
		return nil
	}
	return o.DisplayName
}

func (o *Account) GetForeignID() *string {
	if o == nil {
		return nil
	}
	return o.ForeignID
}

func (o *Account) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *Account) GetMode() *Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *Account) GetProfile() Profile {
	if o == nil {
		return Profile{}
	}
	return o.Profile
}

func (o *Account) GetSettings() *AccountSettings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *Account) GetTermsOfService() *AccountTermsOfService {
	if o == nil {
		return nil
	}
	return o.TermsOfService
}

func (o *Account) GetUpdatedOn() time.Time {
	if o == nil {
		return time.Time{}
	}
	return o.UpdatedOn
}

func (o *Account) GetVerification() *Verification {
	if o == nil {
		return nil
	}
	return o.Verification
}
