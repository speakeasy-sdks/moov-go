// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

type AccessTokenErrorResponse struct {
}

var _ error = &AccessTokenErrorResponse{}

func (e *AccessTokenErrorResponse) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
