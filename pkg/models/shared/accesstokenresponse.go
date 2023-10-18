// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

type AccessTokenResponseType string

const (
	AccessTokenResponseTypeClientCredentialsGrantToAccessTokenResponse AccessTokenResponseType = "ClientCredentialsGrantToAccessTokenResponse"
)

type AccessTokenResponse struct {
	ClientCredentialsGrantToAccessTokenResponse *ClientCredentialsGrantToAccessTokenResponse

	Type AccessTokenResponseType
}

func CreateAccessTokenResponseClientCredentialsGrantToAccessTokenResponse(clientCredentialsGrantToAccessTokenResponse ClientCredentialsGrantToAccessTokenResponse) AccessTokenResponse {
	typ := AccessTokenResponseTypeClientCredentialsGrantToAccessTokenResponse

	return AccessTokenResponse{
		ClientCredentialsGrantToAccessTokenResponse: &clientCredentialsGrantToAccessTokenResponse,
		Type: typ,
	}
}

func (u *AccessTokenResponse) UnmarshalJSON(data []byte) error {

	clientCredentialsGrantToAccessTokenResponse := new(ClientCredentialsGrantToAccessTokenResponse)
	if err := utils.UnmarshalJSON(data, &clientCredentialsGrantToAccessTokenResponse, "", true, true); err == nil {
		u.ClientCredentialsGrantToAccessTokenResponse = clientCredentialsGrantToAccessTokenResponse
		u.Type = AccessTokenResponseTypeClientCredentialsGrantToAccessTokenResponse
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AccessTokenResponse) MarshalJSON() ([]byte, error) {
	if u.ClientCredentialsGrantToAccessTokenResponse != nil {
		return utils.MarshalJSON(u.ClientCredentialsGrantToAccessTokenResponse, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}