// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type GetAccountRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *GetAccountRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type GetAccountResponse struct {
	// The created account
	Account     *shared.Account
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetAccountResponse) GetAccount() *shared.Account {
	if o == nil {
		return nil
	}
	return o.Account
}

func (o *GetAccountResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetAccountResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetAccountResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
