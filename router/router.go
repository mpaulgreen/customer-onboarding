package router

import (
	"customer-onboarding/handlers"
	"net/http"
)

func SetUpRoutes() {
	incidentHandler := http.HandlerFunc(handlers.IncidentHandler)
	uploadFile := http.HandlerFunc(handlers.UploadFile)
	http.Handle("/incidents", incidentHandler)
	http.Handle("/upload", uploadFile)
}
