// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package moovgo

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/moov-go/internal/hooks"
	"github.com/speakeasy-sdks/moov-go/pkg/models/operations"
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/moov-go/pkg/models/shared"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"io"
	"net/http"
)

// BankAccounts - To transfer money with Moov, you’ll need to link a bank account to your Moov account, then verify that account. You can link a bank account to a Moov account by providing the bank account number, routing number, and Moov account ID.
//
// We require micro-deposit verification to reduce the risk of fraud or unauthorized activity. You can verify a bank account by initiating [micro-deposits](https://docs.moov.io/guides/sources/bank-accounts/#micro-deposit-verification), sending two small credit transfers to the bank account you want to confirm. Note that there is no way to initiate a micro-deposit from your bank of choice.
//
// Alternatively, you can link and verify a bank account in one step through an instant account verification token from a third party provider like [Plaid](https://docs.moov.io/guides/sources/bank-accounts/plaid) or [MX](https://docs.moov.io/guides/sources/bank-accounts/mx/). Bank accounts can have the following statuses: `new`, `pending`, `verified`, `verificationFailed`, `errored`. Learn more by reading our guide on [bank accounts](https://docs.moov.io/guides/sources/bank-accounts/).
type BankAccounts struct {
	sdkConfiguration sdkConfiguration
}

func newBankAccounts(sdkConfig sdkConfiguration) *BankAccounts {
	return &BankAccounts{
		sdkConfiguration: sdkConfig,
	}
}

// InitiateMicroDeposits - Initiate micro-deposits
// Micro-deposits help confirm bank account ownership, helping reduce fraud and the risk of unauthorized activity. Use this method to initiate the micro-deposit verification, sending two small credit transfers to the bank account you want to confirm. If you request micro-deposits before 4:15PM ET, they will appear that same day. If you request micro-deposits any time after 4:15PM ET, they will appear the next banking day. When the two credits are initiated, Moov simultaneously initiates a debit to recoup the micro-deposits.<br><br> `sandbox` - Micro-deposits initiated for a `sandbox` bank account will always be `$0.00` / `$0.00` and instantly verifiable once initiated. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.
func (s *BankAccounts) InitiateMicroDeposits(ctx context.Context, accountID string, bankAccountID string) (*operations.PostInitiateMicroDepositsResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "postInitiateMicroDeposits"}

	request := operations.PostInitiateMicroDepositsRequest{
		AccountID:     accountID,
		BankAccountID: bankAccountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts/{bankAccountID}/micro-deposits", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "422", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PostInitiateMicroDepositsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode == 204:
		fallthrough
	default:
	}

	return res, nil
}

// CompleteMicroDeposits - Complete micro-deposits
// Complete the micro-deposit validation process by passing the amounts of the two transfers within three tries. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.
func (s *BankAccounts) CompleteMicroDeposits(ctx context.Context, completeMicroDepositsRequest shared.CompleteMicroDepositsRequest, accountID string, bankAccountID string) (*operations.PutCompleteMicroDepositsResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "putCompleteMicroDeposits"}

	request := operations.PutCompleteMicroDepositsRequest{
		CompleteMicroDepositsRequest: completeMicroDepositsRequest,
		AccountID:                    accountID,
		BankAccountID:                bankAccountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts/{bankAccountID}/micro-deposits", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "CompleteMicroDepositsRequest", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "PUT", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "409", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.PutCompleteMicroDepositsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.CompleteMicroDepositsResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.CompleteMicroDepositsResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
	}

	return res, nil
}

// Delete bank account
// Discontinue using a specified bank account linked to a Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.
func (s *BankAccounts) Delete(ctx context.Context, accountID string, bankAccountID string) (*operations.DeleteBankAccountResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "deleteBankAccount"}

	request := operations.DeleteBankAccountRequest{
		AccountID:     accountID,
		BankAccountID: bankAccountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts/{bankAccountID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "DELETE", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.DeleteBankAccountResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	case httpRes.StatusCode == 204:
		fallthrough
	default:
	}

	return res, nil
}

// Get bank account
// Retrieve bank account details (i.e. routing number or account type) associated with a specific Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.
func (s *BankAccounts) Get(ctx context.Context, accountID string, bankAccountID string) (*operations.GetBankAccountResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "getBankAccount"}

	request := operations.GetBankAccountRequest{
		AccountID:     accountID,
		BankAccountID: bankAccountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts/{bankAccountID}", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetBankAccountResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.BankAccountResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BankAccountResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
	}

	return res, nil
}

// Link - Bank account
// Link a bank account to an existing Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.write` scope.
func (s *BankAccounts) Link(ctx context.Context, bankAccountPayload shared.BankAccountPayload, accountID string) (*operations.LinkBankAccountResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "linkBankAccount"}

	request := operations.LinkBankAccountRequest{
		BankAccountPayload: bankAccountPayload,
		AccountID:          accountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	bodyReader, reqContentType, err := utils.SerializeRequestBody(ctx, request, false, false, "BankAccountPayload", "json", `request:"mediaType=application/json"`)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", opURL, bodyReader)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)
	req.Header.Set("Content-Type", reqContentType)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "409", "422", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.LinkBankAccountResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.BankAccountResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BankAccountResponse = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 409:
		fallthrough
	case httpRes.StatusCode == 422:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
	}

	return res, nil
}

// List bank accounts
// List all the bank accounts associated with a particular Moov account. <br><br> To use this endpoint, you need to specify the `/accounts/{accountID}/bank-accounts.read` scope.
func (s *BankAccounts) List(ctx context.Context, accountID string) (*operations.ListBankAccountsResponse, error) {
	hookCtx := hooks.HookContext{OperationID: "listBankAccounts"}

	request := operations.ListBankAccountsRequest{
		AccountID: accountID,
	}

	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	opURL, err := utils.GenerateURL(ctx, baseURL, "/accounts/{accountID}/bank-accounts", request, nil)
	if err != nil {
		return nil, fmt.Errorf("error generating URL: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "GET", opURL, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("User-Agent", s.sdkConfiguration.UserAgent)

	req, err = s.sdkConfiguration.Hooks.BeforeRequest(hooks.BeforeRequestContext{hookCtx}, req)
	if err != nil {
		return nil, err
	}

	client := s.sdkConfiguration.SecurityClient

	httpRes, err := client.Do(req)
	if err != nil || httpRes == nil {
		if err != nil {
			err = fmt.Errorf("error sending request: %w", err)
		} else {
			err = fmt.Errorf("error sending request: no response")
		}

		_, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, nil, err)
		return nil, err
	} else if utils.MatchStatusCodes([]string{"404", "429", "4XX", "5XX"}, httpRes.StatusCode) {
		httpRes, err = s.sdkConfiguration.Hooks.AfterError(hooks.AfterErrorContext{hookCtx}, httpRes, nil)
		if err != nil {
			return nil, err
		}
	} else {
		httpRes, err = s.sdkConfiguration.Hooks.AfterSuccess(hooks.AfterSuccessContext{hookCtx}, httpRes)
		if err != nil {
			return nil, err
		}
	}
	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.ListBankAccountsResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))

	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out []shared.BankAccountResponse
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.BankAccounts = out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 404:
		fallthrough
	case httpRes.StatusCode == 429:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	default:
	}

	return res, nil
}
