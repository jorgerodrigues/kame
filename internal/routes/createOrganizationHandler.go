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
	var body CreateOrganizationRequestBody
	err := json.NewDecoder(r.Body).Decode(&body)

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

	// todo
}
