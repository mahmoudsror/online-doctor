package controllers

import (
	"net/http"

	"github.com/mahmoudsror/online-doctor/responses"
)

// Home returns the root /
func Home(response http.ResponseWriter, req *http.Request) {
	responses.TOJSON(response, struct {
		Message string `json:"message"`
	}{
		Message: "Welcome in home",
	})
}
