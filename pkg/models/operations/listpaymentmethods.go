// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"net/http"
)

type ListPaymentMethodsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// Optional parameter to filter the account's payment methods by source ID.
	SourceID *string `queryParam:"style=form,explode=true,name=sourceID"`
}

func (o *ListPaymentMethodsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListPaymentMethodsRequest) GetSourceID() *string {
	if o == nil {
		return nil
	}
	return o.SourceID
}

type ListPaymentMethodsResponse struct {
	ContentType string
	// Successfully retrieved payment methods
	PaymentMethods []shared.PaymentMethod
	StatusCode     int
	RawResponse    *http.Response
}

func (o *ListPaymentMethodsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListPaymentMethodsResponse) GetPaymentMethods() []shared.PaymentMethod {
	if o == nil {
		return nil
	}
	return o.PaymentMethods
}

func (o *ListPaymentMethodsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListPaymentMethodsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
