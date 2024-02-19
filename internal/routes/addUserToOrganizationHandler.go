package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/utils"
	"github.com/julienschmidt/httprouter"
)

type addUserBody struct {
	UserId string `json:"userId"`
}

func (h *Routes) addUserToOrganizationHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// get org id from param
	orgID := p.ByName("id")
	var userID addUserBody

	err := json.NewDecoder(r.Body).Decode(&userID)
	defer r.Body.Close()

	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.ResponsePayload{
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	if orgID == "" || userID.UserId == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.ResponsePayload{
			Message: "Invalid request body",
			Data:    nil,
		})
	}

	err = h.models.Organizations.AddUserToOrganization(userID.UserId, orgID, "admin")

	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ResponsePayload{
			Message: "Error adding user to organization",
			Data:    nil,
		})
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, utils.ResponsePayload{
		Message: "User added to organization",
		Data:    nil,
	})

}
