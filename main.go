package main

import (
	"customer-onboarding/router"
	"log"
	"net/http"
)

func main() {
	router.SetUpRoutes()
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
