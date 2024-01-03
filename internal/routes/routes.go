package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/database"
	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	db database.Service
}

func RegisterRoutes(d database.Service) http.Handler {
	r := httprouter.New()
	h := &Routes{
		db: database.New(),
	}

	r.HandlerFunc(http.MethodGet, "/", h.HelloWorldHandler)
	r.HandlerFunc(http.MethodGet, "/health", h.healthHandler)

  // users
  r.HandlerFunc(http.MethodPost, "/api/v1/users", h.CreateUserHandler)

	return r
}

func (h *Routes) HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}

func (h *Routes) healthHandler(w http.ResponseWriter, r *http.Request) {
	jsonResp, err := json.Marshal(h.db.Health())

	if err != nil {
		log.Fatalf("error handling JSON marshal. Err: %v", err)
	}

	_, _ = w.Write(jsonResp)
}
