package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginRequestBody struct {
	User     string `json:"user"`
	Password string `json:"password"`
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

		// check if the user exists
		// check if the email is valid
		// check if the password is valid
		// create the user

		return
	}

	fmt.Printf("User %s and Password %s", body.User, body.Password)
}
