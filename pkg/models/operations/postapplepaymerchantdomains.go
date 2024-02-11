// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type PostApplePayMerchantDomainsRequest struct {
	RegisterApplePayMerchantDomains shared.RegisterApplePayMerchantDomains `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *PostApplePayMerchantDomainsRequest) GetRegisterApplePayMerchantDomains() shared.RegisterApplePayMerchantDomains {
	if o == nil {
		return shared.RegisterApplePayMerchantDomains{}
	}
	return o.RegisterApplePayMerchantDomains
}

func (o *PostApplePayMerchantDomainsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type PostApplePayMerchantDomainsResponse struct {
	// Domains registered with Apple Pay
	ApplePayMerchantDomains *shared.ApplePayMerchantDomains
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostApplePayMerchantDomainsResponse) GetApplePayMerchantDomains() *shared.ApplePayMerchantDomains {
	if o == nil {
		return nil
	}
	return o.ApplePayMerchantDomains
}

func (o *PostApplePayMerchantDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostApplePayMerchantDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostApplePayMerchantDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
