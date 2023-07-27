// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
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
	var d *json.Decoder

	clientCredentialsGrantToAccessTokenErrorResponse := new(ClientCredentialsGrantToAccessTokenErrorResponse)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&clientCredentialsGrantToAccessTokenErrorResponse); err == nil {
		u.ClientCredentialsGrantToAccessTokenErrorResponse = clientCredentialsGrantToAccessTokenErrorResponse
		u.Type = AccessTokenErrorResponseTypeClientCredentialsGrantToAccessTokenErrorResponse
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u AccessTokenErrorResponse) MarshalJSON() ([]byte, error) {
	if u.ClientCredentialsGrantToAccessTokenErrorResponse != nil {
		return json.Marshal(u.ClientCredentialsGrantToAccessTokenErrorResponse)
	}

	return nil, nil
}
