package handlers

import "uuid"

type CreateUserOut struct {
	Username string
	Name     string
	ID       uuid.UUID
	CPF      string
	Email    string
}
