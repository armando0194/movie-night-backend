package server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/armando0194/movie-night-backend/pkg/utl/middleware/secure"
	"github.com/gin-gonic/gin"
)

// New instantates new Echo server
func New() *gin.Engine {
	e := gin.New()
	e.Use(gin.Logger(), gin.Recovery(), secure.CORS(), secure.Headers())
	e.GET("/", healthCheck)
	return e
}

func healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, "OK")
}

// Config represents server specific config
type Config struct {
	Port                string
	ReadTimeoutSeconds  int
	WriteTimeoutSeconds int
	Debug               bool
}

// Start starts echo server
func Start(e *gin.Engine, cfg *Config) {
	s := &http.Server{
		Addr:         cfg.Port,
		Handler:      e,
		ReadTimeout:  time.Duration(cfg.ReadTimeoutSeconds) * time.Second,
		WriteTimeout: time.Duration(cfg.WriteTimeoutSeconds) * time.Second,
	}

	// Start server
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
