package routes

import (
	"github.com/gorilla/mux"
	"github.com/mahmoudsror/online-doctor/controllers"
	"github.com/mahmoudsror/online-doctor/middlewares"
)

// NewRouter is router function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", controllers.Home).Methods("GET")
	router.HandleFunc("/users", middlewares.SignupValidator(controllers.CreateUser)).Methods("POST")
	return router
}

// func home(response http.ResponseWriter, req *http.Request) {
// 	response.Header().Set("Content-Type", "application/json")
// 	responses.TOJSON(response, struct {
// 		Message string `json:"message"`
// 	}{
// 		Message: "Welcome ",
// 	})
// }
