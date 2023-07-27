// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type ListDisputesRequest struct {
	// Optional parameter to limit the number of results in the query
	Count *int64 `queryParam:"style=form,explode=true,name=count"`
	// Optional date-time which exclusively filters all disputes with respond by before this date-time.
	RespondEndDateTime *string `queryParam:"style=form,explode=true,name=respondEndDateTime"`
	// Optional date-time which inclusively filters all disputes with respond by after this date-time.
	RespondStartDateTime *string `queryParam:"style=form,explode=true,name=respondStartDateTime"`
	// The number of items to offset before starting to collect the result set
	Skip *int64 `queryParam:"style=form,explode=true,name=skip"`
	// Optional dispute status by which to filter the disputes.
	Status *shared.DisputeStatus `queryParam:"style=form,explode=true,name=status"`
}

func (o *ListDisputesRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListDisputesRequest) GetRespondEndDateTime() *string {
	if o == nil {
		return nil
	}
	return o.RespondEndDateTime
}

func (o *ListDisputesRequest) GetRespondStartDateTime() *string {
	if o == nil {
		return nil
	}
	return o.RespondStartDateTime
}

func (o *ListDisputesRequest) GetSkip() *int64 {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *ListDisputesRequest) GetStatus() *shared.DisputeStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type ListDisputesResponse struct {
	ContentType string
	// List of dispute details
	Disputes    []shared.Dispute
	StatusCode  int
	RawResponse *http.Response
}

func (o *ListDisputesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListDisputesResponse) GetDisputes() []shared.Dispute {
	if o == nil {
		return nil
	}
	return o.Disputes
}

func (o *ListDisputesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListDisputesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}