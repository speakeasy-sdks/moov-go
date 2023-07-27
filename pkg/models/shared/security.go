// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SchemeOAuth2Credentials struct {
	Password string `security:"name=password"`
	Username string `security:"name=username"`
}

func (o *SchemeOAuth2Credentials) GetPassword() string {
	if o == nil {
		return ""
	}
	return o.Password
}

func (o *SchemeOAuth2Credentials) GetUsername() string {
	if o == nil {
		return ""
	}
	return o.Username
}

type Security struct {
	AccessToken       *string                  `security:"scheme,type=oauth2,name=Authorization"`
	GatewayAuth       *string                  `security:"scheme,type=http,subtype=bearer,name=Authorization"`
	OAuth2Credentials *SchemeOAuth2Credentials `security:"scheme,type=http,subtype=basic"`
}

func (o *Security) GetAccessToken() *string {
	if o == nil {
		return nil
	}
	return o.AccessToken
}

func (o *Security) GetGatewayAuth() *string {
	if o == nil {
		return nil
	}
	return o.GatewayAuth
}

func (o *Security) GetOAuth2Credentials() *SchemeOAuth2Credentials {
	if o == nil {
		return nil
	}
	return o.OAuth2Credentials
}
