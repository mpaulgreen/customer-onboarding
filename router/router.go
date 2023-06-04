package router

import (
	"customer-onboarding/handlers"
	"net/http"
)

func SetUpRoutes() {
	incidentHandler := http.HandlerFunc(handlers.IncidentHandler)
	http.Handle("/incidents", incidentHandler)
}
