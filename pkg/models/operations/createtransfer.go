// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type CreateTransferRequest struct {
	CreateTransfer shared.CreateTransfer `request:"mediaType=application/json"`
	// Prevents duplicate transfers from being created. Note that we only accept UUID v4.
	XIdempotencyKey string `header:"style=simple,explode=false,name=X-Idempotency-Key"`
	// Optional header that indicates whether to return a synchronous response that includes full transfer and rail-specific details or an asynchronous response indicating the transfer was created (this is the default response if the header is omitted).
	XWaitFor *shared.WaitFor `header:"style=simple,explode=false,name=X-Wait-For"`
}

func (o *CreateTransferRequest) GetCreateTransfer() shared.CreateTransfer {
	if o == nil {
		return shared.CreateTransfer{}
	}
	return o.CreateTransfer
}

func (o *CreateTransferRequest) GetXIdempotencyKey() string {
	if o == nil {
		return ""
	}
	return o.XIdempotencyKey
}

func (o *CreateTransferRequest) GetXWaitFor() *shared.WaitFor {
	if o == nil {
		return nil
	}
	return o.XWaitFor
}

type CreateTransferResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// A transfer was successfully created but an error occurred while generating the synchronous response. The asynchronous response object will be returned.
	CreatedTransfer *shared.CreatedTransfer
	// A transfer was successfully created but a timeout occurred while waiting for a synchronous response. Rail-specific details may be missing from the response object.
	GetTransferFull *shared.GetTransferFull
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created a transfer
	TransferPostResponse *shared.TransferPostResponse
}

func (o *CreateTransferResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTransferResponse) GetCreatedTransfer() *shared.CreatedTransfer {
	if o == nil {
		return nil
	}
	return o.CreatedTransfer
}

func (o *CreateTransferResponse) GetGetTransferFull() *shared.GetTransferFull {
	if o == nil {
		return nil
	}
	return o.GetTransferFull
}

func (o *CreateTransferResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTransferResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTransferResponse) GetTransferPostResponse() *shared.TransferPostResponse {
	if o == nil {
		return nil
	}
	return o.TransferPostResponse
}
