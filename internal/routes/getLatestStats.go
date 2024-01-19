package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	m "github.com/jorgerodrigues/upkame/internal/models"
	"github.com/julienschmidt/httprouter"
)

type MonitoringStats struct {
	UrlId           string            `json:"url_id"`
	StartDate       time.Time         `json:"start_date"`
	EndDate         time.Time         `json:"end_date"`
	HealtIndex      int               `json:"summary"`
	AvgResponseTime int               `json:"avg_response_time"`
	Runs            []m.MonitoringRun `json:"runs"`
}

func (h *Routes) getLatestStatsHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	url := r.URL.Query().Get("urlId")
	today := time.Now()
	threeDaysAgo := today.AddDate(0, 0, -3)

	days := map[string][]m.MonitoringRun{}
	monitoringRuns, err := h.models.MonitoringRuns.GetRunsForPeriod(url, threeDaysAgo, today)
	if err != nil {
		http.Error(w, "Error getting monitoring runs", http.StatusInternalServerError)
		return
	}

	for _, run := range monitoringRuns {
		year := run.CreatedAt.Year()
		month := run.CreatedAt.Month()
		day := run.CreatedAt.Day()
		dateString := strconv.Itoa(year) + "-" + strconv.Itoa(int(month)) + "-" + strconv.Itoa(day)

		// aggregate the runs
		if days[dateString] == nil {
			days[string(dateString)] = []m.MonitoringRun{run}
		} else {
			days[string(dateString)] = append(days[string(dateString)], run)
		}
	}

  var resultingStats []MonitoringStats

	for _, value := range days {
		var stats MonitoringStats
		amountOfRuns := len(value)
		amountOfSuccesses := 0
		amountOfFailures := 0
		sumOfResponseTimes := 0
		for _, run := range value {
			strStatusCode := strconv.Itoa(run.StatusCode)
			sumOfResponseTimes += run.ResponseTime
			if run.StatusCode == 200 {
				amountOfSuccesses++
			}
			if strings.HasPrefix(strStatusCode, "5") {
				amountOfFailures++
			}
			if strings.HasPrefix(strStatusCode, "4") {
				amountOfFailures++
			}
		}
		stats.HealtIndex = 100 - (amountOfFailures/amountOfRuns)*100
		stats.AvgResponseTime = sumOfResponseTimes / amountOfRuns
		stats.Runs = value
		stats.UrlId = url
		stats.StartDate = today
		stats.EndDate = threeDaysAgo
    resultingStats = append(resultingStats, stats)
	}

	jsonResp, err := json.Marshal(resultingStats)
	if err != nil {
		http.Error(w, "Error marshalling json", http.StatusInternalServerError)
		return
	}
	w.Write(jsonResp)

}
