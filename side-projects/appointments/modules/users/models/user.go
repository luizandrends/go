package models

import "uuid"

type User struct {
	Name     string
	Username string
	ID       uuid.UUID
	CPF      string
	Email    string
	Password string `json:"-"`
}
