package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) deleteURLHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("URL ID is required"))
		return
	}

	err := h.models.URLs.Delete(id)
	if err != nil {
		println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error deleting URL"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("URL deleted"))
}
