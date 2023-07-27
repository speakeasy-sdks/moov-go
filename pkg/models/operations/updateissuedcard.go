// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type UpdateIssuedCardRequest struct {
	UpdateIssuedCard shared.UpdateIssuedCard `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID of the issued card
	IssuedCardID string `pathParam:"style=simple,explode=false,name=issuedCardID"`
}

func (o *UpdateIssuedCardRequest) GetUpdateIssuedCard() shared.UpdateIssuedCard {
	if o == nil {
		return shared.UpdateIssuedCard{}
	}
	return o.UpdateIssuedCard
}

func (o *UpdateIssuedCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *UpdateIssuedCardRequest) GetIssuedCardID() string {
	if o == nil {
		return ""
	}
	return o.IssuedCardID
}

type UpdateIssuedCardResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *UpdateIssuedCardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateIssuedCardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateIssuedCardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
