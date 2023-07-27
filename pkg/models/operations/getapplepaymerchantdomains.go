// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type GetApplePayMerchantDomainsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *GetApplePayMerchantDomainsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type GetApplePayMerchantDomainsResponse struct {
	// Domains registered with Apple Pay
	ApplePayMerchantDomains *shared.ApplePayMerchantDomains
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
}

func (o *GetApplePayMerchantDomainsResponse) GetApplePayMerchantDomains() *shared.ApplePayMerchantDomains {
	if o == nil {
		return nil
	}
	return o.ApplePayMerchantDomains
}

func (o *GetApplePayMerchantDomainsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetApplePayMerchantDomainsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetApplePayMerchantDomainsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
