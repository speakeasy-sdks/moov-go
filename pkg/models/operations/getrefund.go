// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"net/http"
)

type GetRefundRequest struct {
	// ID of the refund
	RefundID string `pathParam:"style=simple,explode=false,name=refundID"`
	// ID of the Transfer
	TransferID string `pathParam:"style=simple,explode=false,name=transferID"`
}

func (o *GetRefundRequest) GetRefundID() string {
	if o == nil {
		return ""
	}
	return o.RefundID
}

func (o *GetRefundRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

type GetRefundResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Refund details
	GetRefund *sdkerrors.GetRefund
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRefundResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRefundResponse) GetGetRefund() *sdkerrors.GetRefund {
	if o == nil {
		return nil
	}
	return o.GetRefund
}

func (o *GetRefundResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRefundResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
