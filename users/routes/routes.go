package routes

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/online-doctor/responses"
)

// NewRouter is router function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home).Methods("GET")
	return router
}

func home(response http.ResponseWriter, req *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	responses.TOJSON(response, struct {
		Message string `json:"message"`
	}{
		Message: "Welcome ",
	})
}
