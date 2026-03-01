package dto

import "time"

type ErrorResponse struct {
	Message string    `json:"message"`
	Time    time.Time `json:"time"`
}
