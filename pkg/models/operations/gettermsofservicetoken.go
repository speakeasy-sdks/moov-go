// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetTermsOfServiceTokenResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The newly generated terms of service token
	TermsOfServiceToken *shared.TermsOfServiceToken
}

func (o *GetTermsOfServiceTokenResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTermsOfServiceTokenResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTermsOfServiceTokenResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetTermsOfServiceTokenResponse) GetTermsOfServiceToken() *shared.TermsOfServiceToken {
	if o == nil {
		return nil
	}
	return o.TermsOfServiceToken
}
