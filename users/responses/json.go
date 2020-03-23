package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// TOJSON format response to json
func TOJSON(response http.ResponseWriter, data interface{}) {
	response.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(response).Encode(data)
	if err != nil {
		log.Fatal(err)
	}
}

// ErrorResponse format error response
func ErrorResponse(response http.ResponseWriter, err error, statusCode int) {
	response.WriteHeader(statusCode)
	TOJSON(response, struct {
		Message string `json:"message"`
	}{
		Message: err.Error(),
	})
}
