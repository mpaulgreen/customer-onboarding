package models

type UserResult struct {
	ID        string `json:"sys_id,omitempty"`
	UserName  string `json:"user_name,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Email     string `json:"email,omitempty"`
}

type User struct { // struct to sys_user table
	Result []UserResult `json:"result,omitempty"`
}
