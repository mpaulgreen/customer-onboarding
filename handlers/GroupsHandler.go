package handlers

import (
	"customer-onboarding/servicenow"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const GroupsPath = "groups"

// TODO: This will be the hanlder for /incidents
func GroupsHandler(w http.ResponseWriter, r *http.Request) {
	println("Inside Groups handler")
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", GroupsPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	case http.MethodGet:
		group, error := servicenow.GetGroupByName("", name)

		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if group == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		groupJson, err := json.Marshal(group)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(groupJson)
		if err != nil {
			log.Fatal(err)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
