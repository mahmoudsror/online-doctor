package responses

import (
	"encoding/json"
	"log"
	"net/http"
)
func TOJSON(response http.ResponseWriter, data interface{})  {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err !=nil {
		log.Fatal(err)
	}
}