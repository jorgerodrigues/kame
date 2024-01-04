package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgerodrigues/upkame/internal/database"
	"github.com/jorgerodrigues/upkame/internal/routes"
	"github.com/jorgerodrigues/upkame/internal/models"
)

type Server struct {
	port int
  models *models.Model
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
  db := database.New()
	NewServer := &Server{
		port: port,
    models: models.New(db),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      routes.RegisterRoutes(NewServer.models),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
