package http

import (
	"fmt"
	v1 "leonardodelira/go-clean-template/cmd/http/routes/v1"
	"leonardodelira/go-clean-template/pkg/httpserver"
	"leonardodelira/go-clean-template/pkg/logger"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("API_PORT")
	logLevel := os.Getenv("LOG_LEVEL")

	l := logger.New(logLevel)

	//HTTP Server
	handler := gin.New()
	handler.Use(gin.Logger())
	handler.Use(gin.Recovery())

	hv1 := handler.Group("/v1")
	{
		v1.NewRouterTranslations(hv1)
	}

	httpServer := httpserver.New(handler, httpserver.Port(port))

	//Waiting Signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))
	}

	err := httpServer.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
	}
}
