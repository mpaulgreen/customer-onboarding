package models

type Attachment struct {
	ID             string `json:"sys_id,omitempty"`
	IncidentNumber string `json:"number,omitempty"`
	FileName       string `json:"file_name,omitempty"`
	Content        []byte `json:"content,omitempty"`
}
