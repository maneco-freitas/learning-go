package model

import "fmt"

type Result struct {
	Data   string
	Err    error
	Period Period
}

func (r Result) String() string {
	if r.Err != nil {
		return fmt.Sprintf("Error processing period %v: %v", r.Period, r.Err)
	}
	return r.Data
}
