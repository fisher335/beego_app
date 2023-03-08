package models

type User struct {
	Id       int    `json:"id,omitempty" `
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Phone    string `json:"phone,omitempty"`
}
