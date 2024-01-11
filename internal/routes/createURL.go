package routes

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type URLRequestBody struct {
	URL         string `json:"url"`
	Name        string `json:"name"`
	OwnerId     string `json:"ownerId"`
	CreatedById string `json:"createdById"`
}

func (h *Routes) createURLHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

	var url URLRequestBody

	err := json.NewDecoder(r.Body).Decode(&url)
	if err != nil {
		// handle error
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error decoding request body"))
		return
	}

	err = h.models.URLs.Create(url.URL, url.Name, url.OwnerId, url.CreatedById)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating URL"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("URL created"))
}
