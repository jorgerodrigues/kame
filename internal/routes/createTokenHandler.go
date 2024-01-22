package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) createTokenHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
  userId := r.FormValue("user_id")
  token, err := h.models.Tokens.Create(userId)
  if err != nil {
    http.Error(w, err.Error(), http.StatusInternalServerError)
    return
  }
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(http.StatusCreated)
  _, _ = w.Write([]byte(token))
}
