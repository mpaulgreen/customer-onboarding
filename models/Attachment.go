package models

type Attachment struct {
	ID             string `json:"sys_id,omitempty"`
	IncidentNumber string
	FileName       string
	Content        []byte
}
