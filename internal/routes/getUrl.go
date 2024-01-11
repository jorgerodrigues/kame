package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) getURLHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL ID is required"))
		return
	}

	url, err := h.models.URLs.GetById(id)
	if err != nil {
    println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error retrieving URL"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(url.URL))
}
