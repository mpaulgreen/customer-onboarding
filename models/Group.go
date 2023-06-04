package models

type GroupResult struct {
	ID          string `json:"sys_id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Email       string `json:"email,omitempty"`
}
type Group struct { // struct to sys_user_group table
	Result []GroupResult `json:"result,omitempty"`
}
