// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetTransferRequest struct {
	// ID of a connected account
	AccountID *string `queryParam:"style=form,explode=true,name=accountID"`
	// ID of the Transfer
	TransferID string `pathParam:"style=simple,explode=false,name=transferID"`
}

func (o *GetTransferRequest) GetAccountID() *string {
	if o == nil {
		return nil
	}
	return o.AccountID
}

func (o *GetTransferRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

type GetTransferResponse struct {
	ContentType string
	// Details of a transfer
	GetTransferFull *shared.GetTransferFull
	StatusCode      int
	RawResponse     *http.Response
}

func (o *GetTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetTransferResponse) GetGetTransferFull() *shared.GetTransferFull {
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