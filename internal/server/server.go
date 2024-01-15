package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/jorgerodrigues/upkame/internal/database"
	"github.com/jorgerodrigues/upkame/internal/logger"
	"github.com/jorgerodrigues/upkame/internal/models"
	"github.com/jorgerodrigues/upkame/internal/routes"
)

type Server struct {
	port   int
	models *models.Model
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	db := database.New()
	logger := logger.NewLogger()

	NewServer := &Server{
		port:   port,
		models: models.New(db, logger),
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

func StartServer() {
	logger := logger.NewLogger()
	// Create a new server
	server := NewServer()

	// Use a WaitGroup to wait for the server to start

	// Start the server
	go func() {
		logger.Info("🔥Server is starting on port :", server.Addr, "")
		err := server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %s", err)
		}
	}()

	// Create a channel to listen for interrupt signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM, syscall.SIGABRT)

	// Block until a signal is received
	sig := <-quit
	logger.Warn("Shutting down server... Reason: %v", sig)

	// Create a deadline for the shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Attempt to gracefully shut down the server
	err := server.Shutdown(ctx)
	if err != nil {
		logger.Error("Server forced to shutdown: %v", err)
	}

	logger.Info("Server exiting")
}
