package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/utils"
	"github.com/julienschmidt/httprouter"
)

type CreateOrganizationRequestBody struct {
	Name string `json:"name"`
}

func (h *Routes) createOrganizationHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	var body CreateOrganizationRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.ResponsePayload{
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	if body.Name == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.ResponsePayload{
			Message: "Invalid request body",
			Data:    nil,
		})
		return
	}

	org, err := h.models.Organizations.FindByName(body.Name)
	if err != nil && err.Error() != "Error finding organization" {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ResponsePayload{
			Message: "Error finding organization",
			Data:    nil,
		})
		return
	}

	if org != nil {
		utils.SendJSONResponse(w, http.StatusConflict, utils.ResponsePayload{
			Message: "Organization already exists",
			Data:    nil,
		})
		return

	}

	err = h.models.Organizations.Create(body.Name)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ResponsePayload{
			Message: "Error creating organization",
			Data:    nil,
		})
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, utils.ResponsePayload{
		Message: "Organization created",
		Data:    nil,
	})
}
