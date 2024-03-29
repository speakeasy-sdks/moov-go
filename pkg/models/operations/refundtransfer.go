// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type RefundTransferRequest struct {
	// Prevents duplicate refunds from being created. Note that we only accept UUID v4.
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
	// ID of the Transfer
	TransferID   string               `pathParam:"style=simple,explode=false,name=transferID"`
	CreateRefund *shared.CreateRefund `request:"mediaType=application/json"`
	// Optional header that indicates whether to return a synchronous response that includes the full refund and card transaction details or an asynchronous response indicating the refund was created (this is the default response if the header is omitted).
	XWaitFor *shared.WaitFor `header:"style=simple,explode=false,name=X-Wait-For"`
}

func (o *RefundTransferRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

func (o *RefundTransferRequest) GetTransferID() string {
	if o == nil {
		return ""
	}
	return o.TransferID
}

func (o *RefundTransferRequest) GetCreateRefund() *shared.CreateRefund {
	if o == nil {
		return nil
	}
	return o.CreateRefund
}

func (o *RefundTransferRequest) GetXWaitFor() *shared.WaitFor {
	if o == nil {
		return nil
	}
	return o.XWaitFor
}

type RefundTransferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A refund was successfully created but an error occurred while waiting for a synchronous response.
	GetRefund *sdkerrors.GetRefund
	// Successfully initiated a card refund
	RefundPostResponse *shared.RefundPostResponse
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *RefundTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *RefundTransferResponse) GetGetRefund() *sdkerrors.GetRefund {
	if o == nil {
		return nil
	}
	return o.GetRefund
}

func (o *RefundTransferResponse) GetRefundPostResponse() *shared.RefundPostResponse {
	if o == nil {
		return nil
	}
	return o.RefundPostResponse
}

func (o *RefundTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *RefundTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
