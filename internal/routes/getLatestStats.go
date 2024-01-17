package routes

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (h *Routes) getLatestStatsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
  url := r.URL.Query().Get("urlId")
  today := time.Now()
  threeDaysAgo := today.AddDate(0, 0, -4)

  monitoringRuns, err := h.models.MonitoringRuns.GetRunsForPeriod(url, threeDaysAgo, today)
  if err != nil {
    http.Error(w, "Error getting monitoring runs", http.StatusInternalServerError)
    return
  }

  for _, run := range monitoringRuns {
    println(run.CreatedAt.String(), run.ResponseTime, run.StatusCode, run.Result)
  }

}
