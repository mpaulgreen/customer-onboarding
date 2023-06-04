package models

type ServiceNowIncident struct {
	ShortDescription  string `json:"short_description,omitempty"`
	Urgency           string `json:"urgency,omitempty"`
	Impact            string `json:"impact,omitempty"`
	Description       string `json:"description,omitempty"`
	ContactType       string `json:"contact_type,omitempty"`
	CallerName        string `json:"caller_id,omitempty"`        // optional
	AssignedGroupName string `json:"assignment_group,omitempty"` // optional
}
