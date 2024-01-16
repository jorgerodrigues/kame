package main

import (
	"github.com/jorgerodrigues/upkame/internal/database"
	"github.com/jorgerodrigues/upkame/internal/logger"
	"github.com/jorgerodrigues/upkame/internal/models"
	"github.com/jorgerodrigues/upkame/internal/server"
)

func main() {

	db := database.New()
	logger := logger.NewLogger()
  models := models.New(db, logger)
  server.StartServer(models)
}
