package routes

import (
	"encoding/json"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/validator"
	"github.com/julienschmidt/httprouter"
)

type LoginRequestBody struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func (h *Routes) CreateUserHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var body LoginRequestBody
	// handling if the request is sent as JSON
	_ = json.NewDecoder(r.Body).Decode(&body)

	// handling if the request is sent as form data
	if body.Email == "" || body.Password == "" {
		body.Email = r.FormValue("user")
		body.Password = r.FormValue("password")
		if body.Email == "" || body.Password == "" {
			w.Write([]byte("User and Password are required"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	// check if the email is valid
	emailValid := validator.IsEmailValid(body.Email)
	if !emailValid {
		w.Write([]byte("Email is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check if the password is valid
	pwValid := validator.IsPasswordValid(body.Password)
	if !pwValid {
		w.Write([]byte("Password is not valid"))
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// check if the user exists

	user, err := h.models.User.FindByEmail(body.Email)
	if user != nil && user.Email == body.Email {
		// user already exists
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("User already exists"))
		return
	}
	// create the user
	err = h.models.User.CreateUser(body.Email, body.Firstname, body.Lastname, body.Password)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating user"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("User created"))
	return
}
