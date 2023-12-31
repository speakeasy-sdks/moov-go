// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

var PostLinkCardServerList = []string{
	"https://cards.moov.io/api",
}

type PostLinkCardRequest struct {
	CardRequest shared.CardRequest `request:"mediaType=application/json"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// Optional header that indicates whether to return a synchronous response or an asynchronous response.
	XWaitFor *shared.SchemasWaitFor `header:"style=simple,explode=false,name=X-Wait-For"`
}

func (o *PostLinkCardRequest) GetCardRequest() shared.CardRequest {
	if o == nil {
		return shared.CardRequest{}
	}
	return o.CardRequest
}

func (o *PostLinkCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *PostLinkCardRequest) GetXWaitFor() *shared.SchemasWaitFor {
	if o == nil {
		return nil
	}
	return o.XWaitFor
}

type PostLinkCardResponse struct {
	// Card linked
	Card *shared.Card
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *PostLinkCardResponse) GetCard() *shared.Card {
	if o == nil {
		return nil
	}
	return o.Card
}

func (o *PostLinkCardResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *PostLinkCardResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *PostLinkCardResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
