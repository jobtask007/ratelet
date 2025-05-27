package server

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	engine *gin.Engine
	port   uint
}

func New(port uint, devMode bool) *Server {
	mode := gin.ReleaseMode
	if devMode {
		mode = gin.DebugMode
	}

	gin.SetMode(mode)

	e := gin.New()

	return &Server{
		engine: e,
		port:   port,
	}
}

func (s *Server) Run() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", s.port),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Handler:      s.engine,
	}

	go func() {
		err := srv.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Failed to start HTTP server", "err", err)
			os.Exit(1)
		}

		slog.Info("HTTP server is shutting down")
	}()

	slog.Info("HTTP server is up and running", "port", s.port)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, cancelFunc := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancelFunc()

	err := srv.Shutdown(shutdownCtx)
	if err != nil {
		slog.Error("Failed to shutdown HTTP server", "err", err)
		os.Exit(1)
	}

	slog.Info("HTTP Server is gracefully stopped")
}
