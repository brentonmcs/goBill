package models

import "time"

// JobResult - model for saving the result of a job check
type JobResult struct {
	StatusCode   string    `json:"statusCode"`
	Started      time.Time `json:"started"`
	ResponseTime int64     `json:"responsetime"`
	Error        string    `json:"error"`
}
