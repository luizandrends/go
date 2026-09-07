package handlers

import "uuid"

type CreateUserIn struct {
	Username string
	Name     string
	ID       uuid.UUID
	CPF      string
	Email    string
	Password string
}
