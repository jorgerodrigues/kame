package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) postRunAllMonitoringJobsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  err := h.models.MonitoringRuns.RunAll()
  if err != nil {
    println(err.Error())
    w.WriteHeader(http.StatusInternalServerError)
    return
  }
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("OK"))
}
