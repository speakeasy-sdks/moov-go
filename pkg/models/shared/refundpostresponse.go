// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"errors"
	"github.com/speakeasy-sdks/moov-go/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
)

type RefundPostResponseType string

const (
	RefundPostResponseTypeCreatedRefund RefundPostResponseType = "CreatedRefund"
	RefundPostResponseTypeGetRefund     RefundPostResponseType = "GetRefund"
)

type RefundPostResponse struct {
	CreatedRefund *CreatedRefund
	GetRefund     *sdkerrors.GetRefund

	Type RefundPostResponseType
}

func CreateRefundPostResponseCreatedRefund(createdRefund CreatedRefund) RefundPostResponse {
	typ := RefundPostResponseTypeCreatedRefund

	return RefundPostResponse{
		CreatedRefund: &createdRefund,
		Type:          typ,
	}
}

func CreateRefundPostResponseGetRefund(getRefund sdkerrors.GetRefund) RefundPostResponse {
	typ := RefundPostResponseTypeGetRefund

	return RefundPostResponse{
		GetRefund: &getRefund,
		Type:      typ,
	}
}

func (u *RefundPostResponse) UnmarshalJSON(data []byte) error {

	createdRefund := CreatedRefund{}
	if err := utils.UnmarshalJSON(data, &createdRefund, "", true, true); err == nil {
		u.CreatedRefund = &createdRefund
		u.Type = RefundPostResponseTypeCreatedRefund
		return nil
	}

	getRefund := sdkerrors.GetRefund{}
	if err := utils.UnmarshalJSON(data, &getRefund, "", true, true); err == nil {
		u.GetRefund = &getRefund
		u.Type = RefundPostResponseTypeGetRefund
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u RefundPostResponse) MarshalJSON() ([]byte, error) {
	if u.CreatedRefund != nil {
		return utils.MarshalJSON(u.CreatedRefund, "", true)
	}

	if u.GetRefund != nil {
		return utils.MarshalJSON(u.GetRefund, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}
