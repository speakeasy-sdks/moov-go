// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

type CreatedTransfer struct {
	CreatedOn *time.Time `json:"createdOn,omitempty"`
	// ID of Transfer
	TransferID *string `json:"transferID,omitempty"`
}

func (c CreatedTransfer) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreatedTransfer) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, true); err != nil {
		return err
	}
	return nil
}

func (o *CreatedTransfer) GetCreatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedOn
}

func (o *CreatedTransfer) GetTransferID() *string {
	if o == nil {
		return nil
	}
	return o.TransferID
}
