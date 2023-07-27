// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type GetListCardsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *GetListCardsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type GetListCardsResponse struct {
	// Successfully retrieved cards
	Cards       []shared.Card
	ContentType string
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetListCardsResponse) GetCards() []shared.Card {
	if o == nil {
		return nil
	}
	return o.Cards
}

func (o *GetListCardsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetListCardsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetListCardsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
