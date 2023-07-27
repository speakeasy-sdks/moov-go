// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type ListIssuedCardsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// Optional parameter to limit the number of results in the query
	Count *int64 `queryParam:"style=form,explode=true,name=count"`
	// The number of items to offset before starting to collect the result set
	Skip *int64 `queryParam:"style=form,explode=true,name=skip"`
	// Optional, comma-separated states to filter the Moov list issued cards response.
	//
	States *shared.IssuedCardState `queryParam:"style=form,explode=true,name=states"`
}

func (o *ListIssuedCardsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListIssuedCardsRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListIssuedCardsRequest) GetSkip() *int64 {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *ListIssuedCardsRequest) GetStates() *shared.IssuedCardState {
	if o == nil {
		return nil
	}
	return o.States
}

type ListIssuedCardsResponse struct {
	ContentType string
	// Successfully retrieved cards
	IssuedCards []shared.IssuedCard
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListIssuedCardsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListIssuedCardsResponse) GetIssuedCards() []shared.IssuedCard {
	if o == nil {
		return nil
	}
	return o.IssuedCards
}

func (o *ListIssuedCardsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListIssuedCardsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
