package routes

import (
	"net/http"
	"time"

	m "github.com/jorgerodrigues/upkame/internal/models"
	"github.com/julienschmidt/httprouter"
)

type MonitoringStats struct {
	UrlId     string            `json:"url_id"`
	StartDate time.Time         `json:"start_date"`
	EndDate   time.Time         `json:"end_date"`
	Summary   int               `json:"summary"`
	Runs      []m.MonitoringRun `json:"runs"`
}

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
		// calculate the stats
		println(run.CreatedAt.String(), run.ResponseTime, run.StatusCode, run.Result)
	}

}
