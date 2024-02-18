package routes

import (
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/utils"
	"github.com/julienschmidt/httprouter"
)

func (h *Routes) getOrganizationByIdHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	//get id parameter
	id := p.ByName("id")

	if id == "" {
		utils.SendJSONResponse(w, http.StatusBadRequest, utils.ResponsePayload{
			Message: "Invalid request",
			Data:    nil,
		})
		return
	}

	org, err := h.models.Organizations.FindById(id)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, utils.ResponsePayload{
			Message: "Error finding organization",
			Data:    nil,
		})
		return
	}

	utils.SendJSONResponse(w, http.StatusOK, utils.ResponsePayload{
		Message: "Organization found",
		Data:    org,
	})

}
