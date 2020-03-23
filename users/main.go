package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mahmoudsror/online-doctor/connections"
	"github.com/mahmoudsror/online-doctor/routes"
)

func main() {
	port := "3002"
	fmt.Printf("Api running on port %s\n", port)
	connections.TestConnection()
	routes := routes.NewRouter()
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), routes))
}
