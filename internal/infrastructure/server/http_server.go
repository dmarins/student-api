package server

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dmarins/student-api/internal/infrastructure/env"
	"github.com/labstack/echo/v4"
)

type (
	IHttpServer interface {
		Start()
		listenAndServe()
		gracefulShutdown()
		GetEcho() *echo.Echo
	}

	//	@title			Student API

	// HttpServer
	//	@host		localhost:8080
	//	@BasePath	/
	HttpServer struct {
		Echo *echo.Echo
	}
)

func NewEchoHttpServer() IHttpServer {
	httpServer := &HttpServer{
		Echo: echo.New(),
	}

	httpServer.config()

	return httpServer
}

func (h *HttpServer) config() {
	h.Echo.Use(ConfigRequestContext())
	h.Echo.Use(ConfigRequestTimeout())
	h.Echo.Use(ConfigCORS())

	h.Echo.Validator = ConfigCustomValidator()

	h.Echo.Server.Addr = fmt.Sprintf("%s:%s", env.GetEnvVar("APP_HOST"), env.GetEnvVar("APP_PORT"))
}

func (h *HttpServer) GetEcho() *echo.Echo {
	return h.Echo
}

func (h *HttpServer) Start() {
	go h.listenAndServe()

	h.gracefulShutdown()
}

func (h *HttpServer) listenAndServe() {
	if err := h.Echo.Server.ListenAndServe(); err != nil {
		log.Fatalf("fail to start the HTTP Server in: %s", env.GetEnvVar("APP_NAME"))

		os.Exit(2)
	}
}

func (h *HttpServer) gracefulShutdown() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	<-quit
	log.Fatalf("graceful shutdown started in: %s", env.GetEnvVar("APP_NAME"))

	duration, err := time.ParseDuration(env.GetEnvVar("APP_GRACEFUL_SHUTDOWN_TIMEOUT"))
	if err != nil {
		log.Fatalf("could not parse APP_GRACEFUL_SHUTDOWN_TIMEOUT: %s", err)

		duration = time.Second * 5
		log.Printf("using default APP_GRACEFUL_SHUTDOWN_TIMEOUT of 5s")
	}

	ctx, cancel := context.WithTimeout(context.Background(), duration)
	defer cancel()

	log.Printf("waiting for the application to shutdown gracefully")

	time.Sleep(time.Second * 5)

	if err := h.Echo.Server.Shutdown(ctx); err != nil {
		log.Fatalf("webserver shutdown timeout reached")
	}

	log.Printf("application shutdown completed")
}
