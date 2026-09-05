package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/luizandrends/appointments/shared/utils"
)

func handleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in CreateUserIn

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		adapt := in.ToUser()

		fmt.Println(adapt)

		out := CreateUserOut{
			ID:       in.ID,
			Username: in.Username,
			CPF:      in.CPF,
			Email:    in.Email,
		}

		utils.SendJSON(w, utils.ApiResponse[CreateUserOut]{Data: out}, http.StatusCreated)
	}
}
