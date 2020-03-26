package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/mahmoudsror/online-doctor/models"
	"github.com/mahmoudsror/online-doctor/responses"
)

// CreateUser add new user
func CreateUser(response http.ResponseWriter, req *http.Request) {
	user := &models.User{}
	err := json.NewDecoder(req.Body).Decode(user)
	if err != nil {
		fmt.Println("Error ", err)
	}
	user.CreateUser()
	responses.TOJSON(response, struct {
		Message string
	}{
		Message: "User created successfulyy",
	})
}
