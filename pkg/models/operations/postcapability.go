// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type PostCapabilityRequest struct {
	AddCapabilityRequest shared.AddCapabilityRequest `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *PostCapabilityRequest) GetAddCapabilityRequest() shared.AddCapabilityRequest {
	if o == nil {
		return shared.AddCapabilityRequest{}
	}
	return o.AddCapabilityRequest
}

func (o *PostCapabilityRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type PostCapabilityResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The list of capabilities for the account
	Classes []shared.Capability
}

func (o *PostCapabilityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostCapabilityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostCapabilityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *PostCapabilityResponse) GetClasses() []shared.Capability {
	if o == nil {
		return nil
	}
	return o.Classes
}
