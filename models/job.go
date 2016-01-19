package models

// Job - details for running the synthentic transaction on the service
type Job struct {
	URI      string                 `json:"uri"`
	Name     string                 `json:"name"`
	Headers  map[string]interface{} `json:"header"`
	HTTPVerb string                 `json:"httpVerb"`
	Timeout  int                    `json:"timeout"`
}
