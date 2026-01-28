package model

import (
	"encoding/json"
	"time"
)

type ErrorResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}

func (e ErrorResponse) ToString() string {
	value, err := json.MarshalIndent(e, "", "	")
	if err != nil {
		panic(err)
	}

	return string(value)
}
