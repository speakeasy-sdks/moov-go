// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
	"fmt"
)

// CancellationStatus - Cancellation status
type CancellationStatus string

const (
	CancellationStatusPending   CancellationStatus = "pending"
	CancellationStatusCompleted CancellationStatus = "completed"
)

func (e CancellationStatus) ToPointer() *CancellationStatus {
	return &e
}

func (e *CancellationStatus) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "pending":
		fallthrough
	case "completed":
		*e = CancellationStatus(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CancellationStatus: %v", v)
	}
}
