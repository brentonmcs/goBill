package httpRequestor

import (
	"encoding/json"

	"github.com/parnurzeal/gorequest"
)

var request = gorequest.New()

// GetBody - Http Get Call - Reads string from the body
func GetBody(URI string, result interface{}) error {

	_, body, errs := request.Get(URI).End()

	if errs != nil {
		panic(errs)
	}

	return json.Unmarshal([]byte(body), result)
}
