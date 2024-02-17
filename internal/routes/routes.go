package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/jorgerodrigues/upkame/internal/models"
	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	models         *models.Model
	urls           *models.URLModel
	montioringRuns *models.MonitoringRunModel
}

func RegisterRoutes(m *models.Model) http.Handler {
	r := httprouter.New()
	h := &Routes{
		models: m,
	}

	r.HandlerFunc(http.MethodGet, "/", h.HelloWorldHandler)
	r.HandlerFunc(http.MethodGet, "/health", h.healthHandler)

	// users
	r.POST("/api/v1/users", h.CreateUserHandler)
  // r.GET("/api/v1/users/:id", h.GetUserHandler)
  // r.DELETE("/api/v1/users/:id", h.DeleteUserHandler)
  // r.PATCH("/api/v1/users/:id", h.UpdateUserHandler)

	// urls
	r.POST("/api/v1/urls", h.createURLHandler)
	r.GET("/api/v1/urls/:id", h.getURLHandler)
	r.DELETE("/api/v1/urls/:id", h.deleteURLHandler)
	r.PATCH("/api/v1/urls/:id", h.updateURLHandler)

	// monitoring
	r.POST("/api/v1/runAll", h.postRunAllMonitoringJobsHandler)

  // monitoring_stats
  r.GET("/api/v1/monitoring_stats/latest_stats", h.getLatestStatsHandler)

  // tokens
  r.POST("/api/v1/tokens", h.createTokenHandler)
  r.POST("/api/v1/tokens/validate", h.validateTokenHandler)

	return h.enableCORS(r)
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
	//	jsonResp, err := json.Marshal(database.Health(h.db))
	//
	//	if err != nil {
	//		log.Fatalf("error handling JSON marshal. Err: %v", err)
	//	}

	_, _ = w.Write([]byte("It's healthy, baby!"))
}
