package models

import (
	"context"
	"log/slog"
	"net/http"
	"sync"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MonitoringRun struct {
	Id           string    `json:"id"`
	UrlId        string    `json:"url_id"`
	Result       string    `json:"result"`
	StatusCode   int       `json:"status_code"`
	ResponseTime int       `json:"response_time"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MonitoringRunModel struct {
	DB     *pgxpool.Pool
	logger *slog.Logger
}

func (m *MonitoringRunModel) Create(urlId, result string, statusCode, responseTime int) (*MonitoringRun, error) {
	m.logger.Info("Creating monitoring run", urlId, result)
	query := `INSERT INTO monitoring_results (url_id, result, status_code, response_time) VALUES ($1, $2, $3, $4) RETURNING id, url_id, result, status_code, response_time, created_at, updated_at`

	item := MonitoringRun{}
	row := m.DB.QueryRow(context.Background(), query, urlId, result, statusCode, responseTime)
	err := row.Scan(&item.Id, &item.UrlId, &item.Result, &item.StatusCode, &item.ResponseTime, &item.CreatedAt, &item.UpdatedAt)

	if err != nil {
		m.logger.Error("Error creating monitoring run", err)
		return nil, err
	}

	return &item, nil
}

func (m *MonitoringRunModel) Delete(urlId string) error {
	query := `DELETE FROM monitoring_runs WHERE url_id = $1`

	_, err := m.DB.Exec(context.Background(), query, urlId)

	if err != nil {
		return err
	}

	return nil
}

func (m *MonitoringRunModel) RunAll() error {

	query := `SELECT id, url, status, created_at, updated_at FROM urls WHERE status = 'active'`
	rows, err := m.DB.Query(context.Background(), query)
	if err != nil {
		return err
	}

	println("Running all monitoring jobs")
	var urlsToMonitor []URL
	for rows.Next() {
		item := URL{}
		err = rows.Scan(&item.ID, &item.URL, &item.Status, &item.CreatedAt, &item.UpdatedAt)
		if err != nil {
			return err
		}
		urlsToMonitor = append(urlsToMonitor, item)
		// run the monitor for each url
		// TODO: execute in parallel as a goroutine
	}

	var wg sync.WaitGroup
	for _, url := range urlsToMonitor {
		wg.Add(1)
		go m.MonitorURL(url.ID, url.URL, &wg)
	}
	wg.Wait()
	if err != nil {
		return err
	}

	return nil
}

func (m *MonitoringRunModel) MonitorURL(urlId, url string, wg *sync.WaitGroup) (*MonitoringRun, error) {

	defer wg.Done()
	client := &http.Client{
		Timeout: 4 * time.Second,
	}
	// Create a new GET request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		m.logger.Error("Error creating request", err.Error(), url)
		return nil, err
	}
	start := time.Now()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		m.logger.Error("Error sending request", err.Error(), url)
		m.Create(urlId, "500 Internal Server Error", 500, 2000)
		return nil, err
	}

	statusCode := resp.StatusCode
	result := resp.Status

	// get the elapsed time
	elapsed := time.Since(start)
	m.logger.Debug("Request completed", url, statusCode, result, elapsed.Milliseconds())
	println("=====================================")
	run, err := m.Create(urlId, result, statusCode, int(elapsed.Milliseconds()))
	if err != nil {
		return nil, err
	}
	return run, nil

}
