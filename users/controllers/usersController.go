package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mahmoudsror/online-doctor/models"
	"github.com/mahmoudsror/online-doctor/responses"
)

// CreateUser add new user
func CreateUser(response http.ResponseWriter, req *http.Request) {

	user := &models.User{}
	body, err := ioutil.ReadAll(req.Body)
	fmt.Println("In controller", body)

	if err != nil {
		fmt.Println("Error ", err)
	}
	user.CreateUser()
	// if createError != nil {
	// 	fmt.Println("controller error : ", createError)
	// 	//responses.ErrorResponse(response, createError, 500)
	// }
	responses.TOJSON(response, struct {
		Message string
	}{
		Message: "User created successfulyy",
	})
}
