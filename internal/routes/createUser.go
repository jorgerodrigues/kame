package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/validator"
)

type LoginRequestBody struct {
	User     string `json:"user"`
	Password string `json:"password"`
  Name     string `json:"name"`
}

func (h *Routes) CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var body LoginRequestBody
	// handling if the request is sent as JSON
	_ = json.NewDecoder(r.Body).Decode(&body)

	// handling if the request is sent as form data
	if body.User == "" || body.Password == "" {
		body.User = r.FormValue("user")
		body.Password = r.FormValue("password")
		if body.User == "" || body.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// check if the email is valid
    emailValid := validator.IsEmailValid(body.User)
    if (!emailValid) {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
		// check if the password is valid
    pwValid := validator.IsPasswordValid(body.Password)
    if (!pwValid) {
      w.WriteHeader(http.StatusBadRequest)
      return
    }
		// check if the user exists

    user := h.models.User.FindByEmail(body.User)
    if (user != nil) {
      // user already exists
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("User already exists"))
      return
    }
		// create the user
    err := h.models.User.CreateUser(body.User, body.Name, body.Password)
    if err != nil {
      w.WriteHeader(http.StatusBadRequest)
      w.Write([]byte("Error creating user"))
      return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("User created"))

		return
	}

	fmt.Printf("User %s and Password %s", body.User, body.Password)
}
