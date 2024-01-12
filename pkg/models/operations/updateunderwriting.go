// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type UpdateUnderwritingRequest struct {
	UnderwritingRequest shared.UnderwritingRequest `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *UpdateUnderwritingRequest) GetUnderwritingRequest() shared.UnderwritingRequest {
	if o == nil {
		return shared.UnderwritingRequest{}
	}
	return o.UnderwritingRequest
}

func (o *UpdateUnderwritingRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type UpdateUnderwritingResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// The underwriting details were successfully updated
	Underwriting *shared.Underwriting
}

func (o *UpdateUnderwritingResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateUnderwritingResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateUnderwritingResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateUnderwritingResponse) GetUnderwriting() *shared.Underwriting {
	if o == nil {
		return nil
	}
	return o.Underwriting
}
