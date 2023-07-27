// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

// LimitedTimeRange - Return `count` number of results within time range between two timestamps and then the interval duration for each result in the specific `tz` timezone
type LimitedTimeRange struct {
	Count *int64     `json:"count,omitempty"`
	Every *string    `json:"every,omitempty"`
	From  *time.Time `json:"from,omitempty"`
	To    *time.Time `json:"to,omitempty"`
	Tz    *string    `json:"tz,omitempty"`
}

func (o *LimitedTimeRange) GetCount() *int64 {
	if o == nil {
		return nil
	}
	return o.Count
}

func (o *LimitedTimeRange) GetEvery() *string {
	if o == nil {
		return nil
	}
	return o.Every
}

func (o *LimitedTimeRange) GetFrom() *time.Time {
	if o == nil {
		return nil
	}
	return o.From
}

func (o *LimitedTimeRange) GetTo() *time.Time {
	if o == nil {
		return nil
	}
	return o.To
}

func (o *LimitedTimeRange) GetTz() *string {
	if o == nil {
		return nil
	}
	return o.Tz
}