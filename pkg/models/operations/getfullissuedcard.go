// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type GetFullIssuedCardRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID of the issued card
	IssuedCardID string `pathParam:"style=simple,explode=false,name=issuedCardID"`
}

func (o *GetFullIssuedCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetFullIssuedCardRequest) GetIssuedCardID() string {
	if o == nil {
		return ""
	}
	return o.IssuedCardID
}

type GetFullIssuedCardResponse struct {
	ContentType string
	// Successfully retrieved card details
	FullIssuedCard *shared.FullIssuedCard
	StatusCode     int
	RawResponse    *http.Response
}

func (o *GetFullIssuedCardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFullIssuedCardResponse) GetFullIssuedCard() *shared.FullIssuedCard {
	if o == nil {
		return nil
	}
	return o.FullIssuedCard
}

func (o *GetFullIssuedCardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFullIssuedCardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
