// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetCardRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID of the card
	CardID string `pathParam:"style=simple,explode=false,name=cardID"`
}

func (o *GetCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetCardRequest) GetCardID() string {
	if o == nil {
		return ""
	}
	return o.CardID
}

type GetCardResponse struct {
	// Successfully retrieved card
	Card        *shared.Card
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetCardResponse) GetCard() *shared.Card {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *GetCardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
