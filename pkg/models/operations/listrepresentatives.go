// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type ListRepresentativesRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *ListRepresentativesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type ListRepresentativesResponse struct {
	ContentType string
	// Successfully retrieved representatives
	Representatives []shared.Representative
	StatusCode      int
	RawResponse     *http.Response
}

func (o *ListRepresentativesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListRepresentativesResponse) GetRepresentatives() []shared.Representative {
	if o == nil {
		return nil
	}
	return o.Representatives
}

func (o *ListRepresentativesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListRepresentativesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
