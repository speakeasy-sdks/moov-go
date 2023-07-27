// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetFileDetailsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID of the file
	FileID string `pathParam:"style=simple,explode=false,name=fileID"`
}

func (o *GetFileDetailsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetFileDetailsRequest) GetFileID() string {
	if o == nil {
		return ""
	}
	return o.FileID
}

type GetFileDetailsResponse struct {
	ContentType string
	// Successfully retrieved file Details
	File        *shared.File
	StatusCode  int
	RawResponse *http.Response
}

func (o *GetFileDetailsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetFileDetailsResponse) GetFile() *shared.File {
	if o == nil {
		return nil
	}
	return o.File
}

func (o *GetFileDetailsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetFileDetailsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}