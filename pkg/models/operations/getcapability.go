// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetCapabilityRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// The requested capability identifier
	CapabilityID shared.CapabilityID `pathParam:"style=simple,explode=false,name=capabilityID"`
}

func (o *GetCapabilityRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetCapabilityRequest) GetCapabilityID() shared.CapabilityID {
	if o == nil {
		return shared.CapabilityID("")
	}
	return o.CapabilityID
}

type GetCapabilityResponse struct {
	// The requested capability
	Capability  *shared.Capability
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCapabilityResponse) GetCapability() *shared.Capability {
	if o == nil {
		return nil
	}
	return o.Capability
}

func (o *GetCapabilityResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCapabilityResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCapabilityResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
