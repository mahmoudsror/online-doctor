package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahmoudsror/online-doctor/users/routes"
)

func main() {
	port := "3000"
	fmt.Printf("Api running on port %s\n", port)
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), routes))
}
