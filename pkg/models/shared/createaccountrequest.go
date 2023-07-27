// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CreateAccountRequestCustomerSupportAddress struct {
	AddressLine1    *string `json:"addressLine1,omitempty"`
	AddressLine2    *string `json:"addressLine2,omitempty"`
	City            *string `json:"city,omitempty"`
	Country         *string `json:"country,omitempty"`
	PostalCode      *string `json:"postalCode,omitempty"`
	StateOrProvince *string `json:"stateOrProvince,omitempty"`
}

func (o *CreateAccountRequestCustomerSupportAddress) GetAddressLine1() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine1
}

func (o *CreateAccountRequestCustomerSupportAddress) GetAddressLine2() *string {
	if o == nil {
		return nil
	}
	return o.AddressLine2
}

func (o *CreateAccountRequestCustomerSupportAddress) GetCity() *string {
	if o == nil {
		return nil
	}
	return o.City
}

func (o *CreateAccountRequestCustomerSupportAddress) GetCountry() *string {
	if o == nil {
		return nil
	}
	return o.Country
}

func (o *CreateAccountRequestCustomerSupportAddress) GetPostalCode() *string {
	if o == nil {
		return nil
	}
	return o.PostalCode
}

func (o *CreateAccountRequestCustomerSupportAddress) GetStateOrProvince() *string {
	if o == nil {
		return nil
	}
	return o.StateOrProvince
}

type CreateAccountRequestCustomerSupportPhone struct {
	CountryCode *string `json:"countryCode,omitempty"`
	Number      *string `json:"number,omitempty"`
}

func (o *CreateAccountRequestCustomerSupportPhone) GetCountryCode() *string {
	if o == nil {
		return nil
	}
	return o.CountryCode
}

func (o *CreateAccountRequestCustomerSupportPhone) GetNumber() *string {
	if o == nil {
		return nil
	}
	return o.Number
}

// CreateAccountRequestCustomerSupport - User-provided information that can be displayed on credit card transactions for customers to use when contacting a customer support team. This data is only allowed on a business account
type CreateAccountRequestCustomerSupport struct {
	Address *CreateAccountRequestCustomerSupportAddress `json:"address,omitempty"`
	// Email Address
	Email   *string                                   `json:"email,omitempty"`
	Phone   *CreateAccountRequestCustomerSupportPhone `json:"phone,omitempty"`
	Website *string                                   `json:"website,omitempty"`
}

func (o *CreateAccountRequestCustomerSupport) GetAddress() *CreateAccountRequestCustomerSupportAddress {
	if o == nil {
		return nil
	}
	return o.Address
}

func (o *CreateAccountRequestCustomerSupport) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *CreateAccountRequestCustomerSupport) GetPhone() *CreateAccountRequestCustomerSupportPhone {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *CreateAccountRequestCustomerSupport) GetWebsite() *string {
	if o == nil {
		return nil
	}
	return o.Website
}

// CreateAccountRequestSettingsAchPayment - User provided settings to manage ACH payments
type CreateAccountRequestSettingsAchPayment struct {
	// The description that shows up on ACH transactions. This will default to the account's display name on account creation.
	CompanyName *string `json:"companyName,omitempty"`
}

func (o *CreateAccountRequestSettingsAchPayment) GetCompanyName() *string {
	if o == nil {
		return nil
	}
	return o.CompanyName
}

// CreateAccountRequestSettingsCardPayment - User provided settings to manage card payments. This data is only allowed on a business account
type CreateAccountRequestSettingsCardPayment struct {
	// The description that shows up on credit card transactions. This will default to the accounts display name on account creation.
	StatementDescriptor *string `json:"statementDescriptor,omitempty"`
}

func (o *CreateAccountRequestSettingsCardPayment) GetStatementDescriptor() *string {
	if o == nil {
		return nil
	}
	return o.StatementDescriptor
}

// CreateAccountRequestSettings - User provided settings to manage an account
type CreateAccountRequestSettings struct {
	AchPayment  *CreateAccountRequestSettingsAchPayment  `json:"achPayment,omitempty"`
	CardPayment *CreateAccountRequestSettingsCardPayment `json:"cardPayment,omitempty"`
}

func (o *CreateAccountRequestSettings) GetAchPayment() *CreateAccountRequestSettingsAchPayment {
	if o == nil {
		return nil
	}
	return o.AchPayment
}

func (o *CreateAccountRequestSettings) GetCardPayment() *CreateAccountRequestSettingsCardPayment {
	if o == nil {
		return nil
	}
	return o.CardPayment
}

// CreateAccountRequestTermsOfService - An encrypted value used to record acceptance of Moov's Terms of Service
type CreateAccountRequestTermsOfService struct {
	Token *string `json:"token,omitempty"`
}

func (o *CreateAccountRequestTermsOfService) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// CreateAccountRequest - Describes the fields available when creating a Moov Account
type CreateAccountRequest struct {
	// The type of entity represented by this Account
	AccountType AccountType `json:"accountType"`
	// The list of capabilities to request when the account is created.
	Capabilities    []CapabilityID                       `json:"capabilities,omitempty"`
	CustomerSupport *CreateAccountRequestCustomerSupport `json:"customerSupport,omitempty"`
	// Serves as an optional alias from a foreign/external system which can be used to reference this resource
	ForeignID *string `json:"foreignID,omitempty"`
	// Free-form key-value pair list. Useful for storing information that is not captured elsewhere.
	Metadata map[string]string `json:"metadata,omitempty"`
	// The mode this account is allowed to be used within.
	Mode *Mode `json:"mode,omitempty"`
	// Describes the fields available when creating a profile.
	// If `accountType` is set to `individual`, the `individual` object should be completed. All others should populate `business`.
	//
	Profile        CreateProfile                       `json:"profile"`
	Settings       *CreateAccountRequestSettings       `json:"settings,omitempty"`
	TermsOfService *CreateAccountRequestTermsOfService `json:"termsOfService,omitempty"`
}

func (o *CreateAccountRequest) GetAccountType() AccountType {
	if o == nil {
		return AccountType("")
	}
	return o.AccountType
}

func (o *CreateAccountRequest) GetCapabilities() []CapabilityID {
	if o == nil {
		return nil
	}
	return o.Capabilities
}

func (o *CreateAccountRequest) GetCustomerSupport() *CreateAccountRequestCustomerSupport {
	if o == nil {
		return nil
	}
	return o.CustomerSupport
}

func (o *CreateAccountRequest) GetForeignID() *string {
	if o == nil {
		return nil
	}
	return o.ForeignID
}

func (o *CreateAccountRequest) GetMetadata() map[string]string {
	if o == nil {
		return nil
	}
	return o.Metadata
}

func (o *CreateAccountRequest) GetMode() *Mode {
	if o == nil {
		return nil
	}
	return o.Mode
}

func (o *CreateAccountRequest) GetProfile() CreateProfile {
	if o == nil {
		return CreateProfile{}
	}
	return o.Profile
}

func (o *CreateAccountRequest) GetSettings() *CreateAccountRequestSettings {
	if o == nil {
		return nil
	}
	return o.Settings
}

func (o *CreateAccountRequest) GetTermsOfService() *CreateAccountRequestTermsOfService {
	if o == nil {
		return nil
	}
	return o.TermsOfService
}
