package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/luizandrends/appointments/modules/users/services"
	"github.com/luizandrends/appointments/shared/utils"
)

func handleCreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var in CreateUserIn

		if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		adaptIn := in.requestToUser()

		usr, err := services.CreateUserServiceExec(adaptIn)

		if err != nil {
			panic(err)
		}

		adaptOut := userToResponse(usr)

		utils.SendJSON(w, utils.ApiResponse[CreateUserOut]{Data: adaptOut}, http.StatusCreated)
	}
}
