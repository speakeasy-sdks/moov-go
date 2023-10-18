// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

type AccessTokenErrorResponseType string

const (
	AccessTokenErrorResponseTypeClientCredentialsGrantToAccessTokenErrorResponse AccessTokenErrorResponseType = "ClientCredentialsGrantToAccessTokenErrorResponse"
)

type AccessTokenErrorResponse struct {
	ClientCredentialsGrantToAccessTokenErrorResponse *ClientCredentialsGrantToAccessTokenErrorResponse

	Type AccessTokenErrorResponseType
}

func CreateAccessTokenErrorResponseClientCredentialsGrantToAccessTokenErrorResponse(clientCredentialsGrantToAccessTokenErrorResponse ClientCredentialsGrantToAccessTokenErrorResponse) AccessTokenErrorResponse {
	typ := AccessTokenErrorResponseTypeClientCredentialsGrantToAccessTokenErrorResponse

	return AccessTokenErrorResponse{
		ClientCredentialsGrantToAccessTokenErrorResponse: &clientCredentialsGrantToAccessTokenErrorResponse,
		Type: typ,
	}
}

func (u *AccessTokenErrorResponse) UnmarshalJSON(data []byte) error {

	clientCredentialsGrantToAccessTokenErrorResponse := new(ClientCredentialsGrantToAccessTokenErrorResponse)
	if err := utils.UnmarshalJSON(data, &clientCredentialsGrantToAccessTokenErrorResponse, "", true, true); err == nil {
		u.ClientCredentialsGrantToAccessTokenErrorResponse = clientCredentialsGrantToAccessTokenErrorResponse
		u.Type = AccessTokenErrorResponseTypeClientCredentialsGrantToAccessTokenErrorResponse
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AccessTokenErrorResponse) MarshalJSON() ([]byte, error) {
	if u.ClientCredentialsGrantToAccessTokenErrorResponse != nil {
		return utils.MarshalJSON(u.ClientCredentialsGrantToAccessTokenErrorResponse, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}