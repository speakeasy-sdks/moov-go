// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type ClientCredentialsGrantToAccessTokenRequestGrantType string

const (
	ClientCredentialsGrantToAccessTokenRequestGrantTypeClientCredentials ClientCredentialsGrantToAccessTokenRequestGrantType = "client_credentials"
	ClientCredentialsGrantToAccessTokenRequestGrantTypeRefreshToken      ClientCredentialsGrantToAccessTokenRequestGrantType = "refresh_token"
)

func (e ClientCredentialsGrantToAccessTokenRequestGrantType) ToPointer() *ClientCredentialsGrantToAccessTokenRequestGrantType {
	return &e
}

func (e *ClientCredentialsGrantToAccessTokenRequestGrantType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "client_credentials":
		fallthrough
	case "refresh_token":
		*e = ClientCredentialsGrantToAccessTokenRequestGrantType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for ClientCredentialsGrantToAccessTokenRequestGrantType: %v", v)
	}
}

// ClientCredentialsGrantToAccessTokenRequest - Allows the use of `Client Credentials Grant` per the RFC 6749 of (OAuth 2.0 Authorization Framework)[https://tools.ietf.org/html/rfc6749#section-4.4]. Following this specification will allow any tooling to be able to use this API to get an `access_token`.
type ClientCredentialsGrantToAccessTokenRequest struct {
	// If not specified in `Authorization: Basic` it can be specified here
	ClientID *string `json:"client_id,omitempty" form:"name=client_id"`
	// If not specified in `Authorization: Basic` it can be specified here
	ClientSecret *string                                             `json:"client_secret,omitempty" form:"name=client_secret"`
	GrantType    ClientCredentialsGrantToAccessTokenRequestGrantType `json:"grant_type" form:"name=grant_type"`
	// String passed to the authorization server to gain access to the system
	RefreshToken *string `json:"refresh_token,omitempty" form:"name=refresh_token"`
	// A space-delimited list of [scopes](https://docs.moov.io/guides/developer-tools/api-keys/scopes/) that are allowed
	Scope *string `json:"scope,omitempty" form:"name=scope"`
}

func (o *ClientCredentialsGrantToAccessTokenRequest) GetClientID() *string {
	if o == nil {
		return nil
	}
	return o.ClientID
}

func (o *ClientCredentialsGrantToAccessTokenRequest) GetClientSecret() *string {
	if o == nil {
		return nil
	}
	return o.ClientSecret
}

func (o *ClientCredentialsGrantToAccessTokenRequest) GetGrantType() ClientCredentialsGrantToAccessTokenRequestGrantType {
	if o == nil {
		return ClientCredentialsGrantToAccessTokenRequestGrantType("")
	}
	return o.GrantType
}

func (o *ClientCredentialsGrantToAccessTokenRequest) GetRefreshToken() *string {
	if o == nil {
		return nil
	}
	return o.RefreshToken
}

func (o *ClientCredentialsGrantToAccessTokenRequest) GetScope() *string {
	if o == nil {
		return nil
	}
	return o.Scope
}