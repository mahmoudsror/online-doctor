package controllers

import (
	"net/http"

	"github.com/mahmoudsror/online-doctor/responses"
	"github.com/mahmoudsror/online-doctor/services"
)

// CreateUser add new user
func CreateUser(response http.ResponseWriter, req *http.Request) {
	services.createUser()
	responses.TOJSON(response, struct {
		Message string
	}{
		Message: "User created successfulyy",
	})
}
