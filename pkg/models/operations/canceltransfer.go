// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type CancelTransferRequest struct {
	// Prevents duplicate reversals from being created
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
	// ID of the Transfer
	TransferID     string                 `pathParam:"style=simple,explode=false,name=transferID"`
	CreateReversal *shared.CreateReversal `request:"mediaType=application/json"`
}

func (o *CancelTransferRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

func (o *CancelTransferRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

func (o *CancelTransferRequest) GetCreateReversal() *shared.CreateReversal {
	if o == nil {
		return nil
	}
	return o.CreateReversal
}

type CancelTransferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Successfully initiated a reversal
	CreatedReversal *sdkerrors.CreatedReversal
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *CancelTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CancelTransferResponse) GetCreatedReversal() *sdkerrors.CreatedReversal {
	if o == nil {
		return nil
	}
	return o.CreatedReversal
}

func (o *CancelTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CancelTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
