package handlers

import "github.com/luizandrends/appointments/modules/users/models"

func (in CreateUserIn) requestToUser() models.User {
	return models.User{
		ID:       in.ID,
		Username: in.Username,
		CPF:      in.CPF,
		Email:    in.Email,
		Password: in.Password,
	}
}

func userToResponse(usr models.User) CreateUserOut {
	return CreateUserOut{
		ID:       usr.ID,
		Username: usr.Username,
		CPF:      usr.CPF,
		Email:    usr.Email,
	}
}
