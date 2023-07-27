// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type UpdateCardRequest struct {
	CardUpdateRequest shared.CardUpdateRequest `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID of the card
	CardID string `pathParam:"style=simple,explode=false,name=cardID"`
}

func (o *UpdateCardRequest) GetCardUpdateRequest() shared.CardUpdateRequest {
	if o == nil {
		return shared.CardUpdateRequest{}
	}
	return o.CardUpdateRequest
}

func (o *UpdateCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *UpdateCardRequest) GetCardID() string {
	if o == nil {
		return ""
	}
	return o.CardID
}

type UpdateCardResponse struct {
	// Card updated
	Card        *shared.Card
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The supplied card data appeared invalid or was declined by the issuer
	UpdateCard422ApplicationJSONObject map[string]interface{}
}

func (o *UpdateCardResponse) GetCard() *shared.Card {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *UpdateCardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCardResponse) GetUpdateCard422ApplicationJSONObject() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.UpdateCard422ApplicationJSONObject
}
