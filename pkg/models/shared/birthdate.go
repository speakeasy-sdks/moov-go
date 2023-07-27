// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BirthDate - Birthdate for an individual
type BirthDate struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}

func (o *BirthDate) GetDay() int64 {
	if o == nil {
		return 0
	}
	return o.Day
}

func (o *BirthDate) GetMonth() int64 {
	if o == nil {
		return 0
	}
	return o.Month
}

func (o *BirthDate) GetYear() int64 {
	if o == nil {
		return 0
	}
	return o.Year
}