package handlers

import (
	"customer-onboarding/servicenow"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

const UsersPath = "users"

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", UsersPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	name := urlPathSegments[len(urlPathSegments)-1]

	switch r.Method {
	case http.MethodGet:
		user, error := servicenow.GetUserByName("", name)

		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if user == nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		userJson, err := json.Marshal(user)

		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		_, err = w.Write(userJson)
		if err != nil {
			log.Fatal(err)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}
}
