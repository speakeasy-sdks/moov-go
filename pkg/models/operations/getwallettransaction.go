// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"net/http"
	"openapi/pkg/models/shared"
)

type GetWalletTransactionRequest struct {
	// ID of the account
	AccountID string `pathParam:"style=simple,explode=false,name=accountID"`
	// ID associated with the wallet transaction.
	TransactionID string `pathParam:"style=simple,explode=false,name=transactionID"`
	// ID of the wallet
	WalletID string `pathParam:"style=simple,explode=false,name=walletID"`
}

func (o *GetWalletTransactionRequest) GetAccountID() string {
	if o == nil {
		return ""
	}
	return o.AccountID
}

func (o *GetWalletTransactionRequest) GetTransactionID() string {
	if o == nil {
		return ""
	}
	return o.TransactionID
}

func (o *GetWalletTransactionRequest) GetWalletID() string {
	if o == nil {
		return ""
	}
	return o.WalletID
}

type GetWalletTransactionResponse struct {
	ContentType string
	StatusCode  int
	RawResponse *http.Response
	// Transaction associated with the wallet
	WalletTransaction *shared.WalletTransaction
}

func (o *GetWalletTransactionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetWalletTransactionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetWalletTransactionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetWalletTransactionResponse) GetWalletTransaction() *shared.WalletTransaction {
	if o == nil {
		return nil
	}
	return o.WalletTransaction
}
