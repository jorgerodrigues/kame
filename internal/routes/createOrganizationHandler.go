package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type CreateOrganizationRequestBody struct {
	Name string `json:"name"`
}

func (h *Routes) createOrganizationHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	var body CreateOrganizationRequestBody

	err := json.NewDecoder(r.Body).Decode(&body)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}

	if body.Name == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Name is required"))
		return
	}

	org, err := h.models.Organizations.FindByName(body.Name)
	if err != nil && err.Error() != "Error finding organization" {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error finding organization"))
		return
	}

	if org != nil {
		w.WriteHeader(http.StatusConflict)
		w.Write([]byte("Organization already exists"))
		return
	}

	err = h.models.Organizations.Create(body.Name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error creating organization"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("ok"))
}
