// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"net/http"
)

type ListAccountIssuedCardTransactionsRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// Optional parameter to limit the number of results in the query
	Count *int64 `default:"20" queryParam:"style=form,explode=true,name=count"`
	// The number of items to offset before starting to collect the result set
	Skip *int64 `queryParam:"style=form,explode=true,name=skip"`
	// Optional parameters to filter results IssuedCardTransactions.
	Status *shared.IssuedCardTransactionStatus `queryParam:"style=form,explode=true,name=status"`
}

func (l ListAccountIssuedCardTransactionsRequest) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(l, "", false)
}

func (l *ListAccountIssuedCardTransactionsRequest) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &l, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ListAccountIssuedCardTransactionsRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *ListAccountIssuedCardTransactionsRequest) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *ListAccountIssuedCardTransactionsRequest) GetSkip() *int64 {
	if o == nil {
		return nil
	}
	return o.Skip
}

func (o *ListAccountIssuedCardTransactionsRequest) GetStatus() *shared.IssuedCardTransactionStatus {
	if o == nil {
		return nil
	}
	return o.Status
}

type ListAccountIssuedCardTransactionsResponse struct {
	ContentType string
	// Successfully retrieved issued card transactions
	IssuedCardTransactions []shared.IssuedCardTransaction
	StatusCode             int
	RawResponse            *http.Response
}

func (o *ListAccountIssuedCardTransactionsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListAccountIssuedCardTransactionsResponse) GetIssuedCardTransactions() []shared.IssuedCardTransaction {
	if o == nil {
		return nil
	}
	return o.IssuedCardTransactions
}

func (o *ListAccountIssuedCardTransactionsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListAccountIssuedCardTransactionsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
