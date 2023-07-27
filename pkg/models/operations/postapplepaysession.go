// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type PostApplePaySessionRequest struct {
	CreateApplePaySession shared.CreateApplePaySession `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *PostApplePaySessionRequest) GetCreateApplePaySession() shared.CreateApplePaySession {
	if o == nil {
		return shared.CreateApplePaySession{}
	}
	return o.CreateApplePaySession
}

func (o *PostApplePaySessionRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type PostApplePaySessionResponse struct {
	// Session created with Apple Pay
	ApplePaySession interface{}
	ContentType     string
	StatusCode      int
	RawResponse     *http.Response
}

func (o *PostApplePaySessionResponse) GetApplePaySession() interface{} {
	if o == nil {
		return nil
	}
	return o.ApplePaySession
}

func (o *PostApplePaySessionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostApplePaySessionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostApplePaySessionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
