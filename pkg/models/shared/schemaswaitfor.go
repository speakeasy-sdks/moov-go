// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SchemasWaitFor string

const (
	SchemasWaitForPaymentMethod SchemasWaitFor = "payment-method"
)

func (e SchemasWaitFor) ToPointer() *SchemasWaitFor {
	return &e
}

func (e *SchemasWaitFor) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "payment-method":
		*e = SchemasWaitFor(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SchemasWaitFor: %v", v)
	}
}
