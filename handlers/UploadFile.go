package handlers

import (
	"customer-onboarding/models"
	"customer-onboarding/servicenow"
	"io"
	"log"
	"net/http"
)

func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		content, err := io.ReadAll(r.Body)
		fileName := r.URL.Query().Get("file_name")
		incNumber := r.URL.Query().Get("number")
		if err != nil {
			log.Println("Error Reading the file")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		error := servicenow.AddAttachment(&models.Attachment{
			IncidentNumber: incNumber,
			Content:        content,
			FileName:       fileName,
		})

		if error != nil {
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(error.Error())
			return
		}

		w.WriteHeader(http.StatusCreated)
	default:
		w.WriteHeader(http.StatusNotImplemented)
	}

}
