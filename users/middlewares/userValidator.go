package middlewares

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/mahmoudsror/online-doctor/models"

	"gopkg.in/thedevsaddam/govalidator.v1"
)

// SignupValidator validate signup request
func SignupValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			http.Error(response, err.Error(), http.StatusInternalServerError)
			return
		}

		// you can reassign the body if you need to parse it as multipart
		req.Body = ioutil.NopCloser(bytes.NewReader(body))
		user := models.User{}
		rules := govalidator.MapData{
			"fname":    []string{"required", "between:3,20"},
			"lname":    []string{"required", "between:3,20"},
			"email":    []string{"required", "min:4", "max:60", "email"},
			"password": []string{"required", "between:8,20"},
			"mobile":   []string{"digits:11"},
			"usertype": []string{"required"},
		}
		opts := govalidator.Options{
			Request: req,
			Rules:   rules,
			Data:    &user,
		}

		validator := govalidator.New(opts)
		validationErrors := validator.ValidateJSON()
		if len(validationErrors) > 0 {
			//fmt.Println("error :  ", validationErrorsMap)
			response.Header().Set("Content-type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(validationErrors)
			return
		}

		// create a new url from the raw RequestURI sent by the client
		next(response, req)
	}
}

// SigninValidator validate login request

func SigninValidator(next http.HandlerFunc) http.HandlerFunc {
	return func(response http.ResponseWriter, req *http.Request) {
		user := models.User{}
		rules := govalidator.MapData{
			"email":       []string{"required", "min:4", "max:60", "email"},
			"password":    []string{"required", "between:8,20"},
			"deviceToken": []string{"required"},
		}
		opts := govalidator.Options{
			Request: req,
			Rules:   rules,
			Data:    &user,
		}

		validator := govalidator.New(opts)
		validationErrors := validator.ValidateJSON()
		if len(validationErrors) > 0 {
			fmt.Println("I am here  ")
			response.Header().Set("Content-type", "application/json")
			response.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(response).Encode(validationErrors)
			return
		}
		fmt.Println("I am clean  ")

		next(response, req)
	}
}
