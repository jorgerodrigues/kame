package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) validateTokenHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	token := r.FormValue("token")
	valid, err := h.models.Tokens.Validate(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !valid {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte("Valid token"))
}
