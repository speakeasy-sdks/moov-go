// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/moov-go/pkg/utils"
	"time"
)

// CardAccountUpdater - The results of the most recent card update request
type CardAccountUpdater struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	// The results of the card update request
	UpdateType *CardUpdateReason `json:"updateType,omitempty"`
	UpdatedOn  *time.Time        `json:"updatedOn,omitempty"`
}

func (c CardAccountUpdater) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CardAccountUpdater) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CardAccountUpdater) GetAdditionalProperties() map[string]interface{} {
	if o == nil {
		return nil
	}
	return o.AdditionalProperties
}

func (o *CardAccountUpdater) GetUpdateType() *CardUpdateReason {
	if o == nil {
		return nil
	}
	return o.UpdateType
}

func (o *CardAccountUpdater) GetUpdatedOn() *time.Time {
	if o == nil {
		return nil
	}
	return o.UpdatedOn
}
