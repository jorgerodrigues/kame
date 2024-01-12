package models

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type MonitoringRun struct {
	Id           string `json:"id"`
	UrlId        string `json:"url_id"`
	Result       string `json:"result"`
	StatusCode   int    `json:"status_code"`
	ResponseTime int    `json:"response_time"`
	CreatedAt    string `json:"created_at"`
	UpdatedAt    string `json:"updated_at"`
}

type MonitoringRunModel struct {
  db *pgxpool.Pool
}

func (m *MonitoringRunModel) Create(urlId, result string, statusCode, responseTime int) (*MonitoringRun, error) {
  query := `INSERT INTO monitoring_runs (url_id, result, status_code, response_time) VALUES ($1, $2, $3, $4) RETURNING id, url_id, result, status_code, response_time, created_at, updated_at`

  item := MonitoringRun{}
  row := m.db.QueryRow(context.Background(), query, urlId, result, statusCode, responseTime)
  err := row.Scan(&item.Id, &item.UrlId, &item.Result, &item.StatusCode, &item.ResponseTime, &item.CreatedAt, &item.UpdatedAt)
  
  if err != nil {
    return nil, err
  }

  return &item, nil
}

func (m *MonitoringRunModel) Delete(urlId string) error {
  query := `DELETE FROM monitoring_runs WHERE url_id = $1`

  _, err := m.db.Exec(context.Background(), query, urlId)

  if err != nil {
    return err
  }

  return nil
}

func (m *MonitoringRunModel) RunAll(urlId string) error {

  query := `SELECT * FROM urls WHERE status = 'active'`
  rows, err := m.db.Query(context.Background(), query)

  for rows.Next() {
    item := URL{}
    err = rows.Scan(&item.ID, &item.URL, &item.Status, &item.CreatedAt, &item.UpdatedAt)
    if err != nil {
      return err
    }
    // run the monitor for each url
    // TODO: execute in parallel as a goroutine
    go m.MonitorURL(item.ID, item.URL)
  }
  
  
  if err != nil {
    return err
  }

  return nil
}

func (m *MonitoringRunModel) MonitorURL(urlId, url string) (*MonitoringRun, error) {

	client := &http.Client{}
	// Create a new GET request
	req, err := http.NewRequest("GET", "https://api.example.com/data", nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
    return nil, err
	}
  start := time.Now()

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
    return nil, err
	}

  statusCode := resp.StatusCode
  result := resp.Status

  // get the elapsed time
  elapsed := time.Since(start)
  run, err := m.Create(urlId, result, statusCode, int(elapsed.Milliseconds()))
  if err != nil {
    return nil, err
  }
  return run, nil

}

