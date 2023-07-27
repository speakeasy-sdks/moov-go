// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

var PostLinkCardServerList = []string{
	"https://cards.moov.io/api",
}

type PostLinkCardRequest struct {
	CardRequest shared.CardRequest `request:"mediaType=application/json"`
	// Optional header that indicates whether to return a synchronous response or an asynchronous response.
	XWaitFor *shared.SchemasWaitFor `header:"style=simple,explode=false,name=X-Wait-For"`
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
}

func (o *PostLinkCardRequest) GetCardRequest() shared.CardRequest {
	if o == nil {
		return shared.CardRequest{}
	}
	return o.CardRequest
}

func (o *PostLinkCardRequest) GetXWaitFor() *shared.SchemasWaitFor {
	if o == nil {
		return nil
	}
	return o.XWaitFor
}

func (o *PostLinkCardRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

type PostLinkCardResponse struct {
	// Card linked
	Card        *shared.Card
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// The supplied card data appeared invalid or was declined by the issuer
	PostLinkCard422ApplicationJSONObject map[string]interface{}
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

func (o *PostLinkCardResponse) GetPostLinkCard422ApplicationJSONObject() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.PostLinkCard422ApplicationJSONObject
}
