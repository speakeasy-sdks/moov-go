// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"net/http"
)

type ListAccountsRequest struct {
	// Optional parameter to limit the number of results in the query
	Count *int64 `default:"20" queryParam:"style=form,explode=true,name=count"`
	// Filter connected accounts by email address.<br><br>
	// Provide the full email address to filter by email.
	//
	Email *string `queryParam:"style=form,explode=true,name=email"`
	// Serves as an optional alias from a foreign/external system which can be used to reference this resource
	//
	ForeignID *string `queryParam:"style=form,explode=true,name=foreignID"`
	// Filter disconnected accounts.<br><br>
	// If true, the response will include disconnected accounts.
	//
	IncludeDisconnected *bool `queryParam:"style=form,explode=true,name=includeDisconnected"`
	// Filter connected accounts by name.<br><br>
	// If provided, this query will attempt to find matches against the following Account and Profile fields:
	// <ul>
	//   <li>Account `displayName`</li>
	//   <li>Individual Profile `firstName`, `middleName`, and `lastName`</li>
	//   <li>Business Profile `legalBusinessName`</li>
	// </ul>
	//
	Name *string `queryParam:"style=form,explode=true,name=name"`
	// The number of items to offset before starting to collect the result set
	Skip *int64 `queryParam:"style=form,explode=true,name=skip"`
	// Filter connected accounts by AccountType.<br><br>
	// If the `type` parameter is used in combination with `name`, only the corresponding type's name fields will be searched.
	// For example, if `type=business` and `name=moov`, the search will attempt to find matches against the display name and Business Profile name fields (`legalBusinessName`, and `doingBusinessAs`).
	//
	Type *shared.AccountType `queryParam:"style=form,explode=true,name=type"`
	// Filter by the `verificationStatus` of accounts.
	//
	VerificationStatus *shared.AccountVerificationStatus `queryParam:"style=form,explode=true,name=verification_status"`
}

func (l ListAccountsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountsRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListAccountsRequest) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *ListAccountsRequest) GetForeignID() *string {
	if o == nil {
		return nil
	}
	return o.ForeignID
}

func (o *ListAccountsRequest) GetIncludeDisconnected() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeDisconnected
}

func (o *ListAccountsRequest) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *ListAccountsRequest) GetSkip() *int64 {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *ListAccountsRequest) GetType() *shared.AccountType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ListAccountsRequest) GetVerificationStatus() *shared.AccountVerificationStatus {
	if o == nil {
		return nil
	}
	return o.VerificationStatus
}

type ListAccountsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// All connected accounts matching the filter parameters
	Classes []shared.Account
}

func (o *ListAccountsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListAccountsResponse) GetClasses() []shared.Account {
	if o == nil {
		return nil
	}
	return o.Classes
}
