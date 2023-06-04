package handlers

import (
	"customer-onboarding/models"
	"customer-onboarding/servicenow"
	"encoding/json"
	"log"
	"net/http"
)

func IncidentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var servieNowIncident models.ServiceNowIncident
		err := json.NewDecoder(r.Body).Decode(&servieNowIncident)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		incident, error := servicenow.CreateIncident(&servieNowIncident)

		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(error.Error()))
			return
		}

		json, err := json.Marshal(incident)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(error.Error()))
			return
		}
		_, err = w.Write(json)
		if err != nil {
			log.Fatal(err)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
