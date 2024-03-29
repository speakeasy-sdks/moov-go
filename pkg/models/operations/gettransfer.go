// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"net/http"
)

type GetTransferRequest struct {
	// ID of the Transfer
	TransferID string `pathParam:"style=simple,explode=false,name=transferID"`
	// ID of a connected account
	AccountID *string `queryParam:"style=form,explode=true,name=accountID"`
}

func (o *GetTransferRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

func (o *GetTransferRequest) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

type GetTransferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Details of a transfer
	GetTransferFull *sdkerrors.GetTransferFull
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransferResponse) GetGetTransferFull() *sdkerrors.GetTransferFull {
	if o == nil {
		return nil
	}
	return o.GetTransferFull
}

func (o *GetTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
