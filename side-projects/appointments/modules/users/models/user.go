package models

type User struct {
	Username string
	ID       int64 `json:"id,string"`
	CPF      string
	Email    string
	Password string `json:"-"`
}
