package models

type IncidentResult struct {
	ID               string `json:"sys_id,omitempty"`
	Number           string `json:"number,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
	Urgency          string `json:"urgency,omitempty"`
	Impact           string `json:"impact,omitempty"`
	Description      string `json:"description,omitempty"`
	ContactType      string `json:"contact_type,omitempty"`
}
type Incident struct { // struct to incident table
	Result []IncidentResult `json:"result,omitempty"`
}

type IncidentResponse struct {
	Result IncidentResult `json:"result,omitempty"`
}
