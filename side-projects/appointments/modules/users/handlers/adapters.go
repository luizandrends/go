package handlers

import "github.com/luizandrends/appointments/modules/users/models"

func (in CreateUserIn) ToUser() models.User {
	return models.User{
		ID:       in.ID,
		Username: in.Username,
		CPF:      in.CPF,
		Email:    in.Email,
		Password: in.Password,
	}
}
