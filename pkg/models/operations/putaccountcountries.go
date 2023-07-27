// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type PutAccountCountriesRequest struct {
	Countries shared.Countries `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *PutAccountCountriesRequest) GetCountries() shared.Countries {
	if o == nil {
		return shared.Countries{}
	}
	return o.Countries
}

func (o *PutAccountCountriesRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type PutAccountCountriesResponse struct {
	ContentType string
	// countries assigned to account
	Countries   *shared.Countries
	StatusCode  int
	RawResponse *http.Response
}

func (o *PutAccountCountriesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PutAccountCountriesResponse) GetCountries() *shared.Countries {
	if o == nil {
		return nil
	}
	return o.Countries
}

func (o *PutAccountCountriesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PutAccountCountriesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}