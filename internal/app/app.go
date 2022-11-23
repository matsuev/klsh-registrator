package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/matsuev/klsh-registrator/internal/config"
	"github.com/matsuev/klsh-registrator/internal/logging"
	"github.com/matsuev/klsh-registrator/internal/router"
	"github.com/matsuev/klsh-registrator/internal/server"
	"github.com/matsuev/klsh-registrator/internal/service"
)

// Application struct
type Application struct {
	logger *logging.Logger
	cfg    *config.Config
	srv    *http.Server
}

// New function
func New(cfg *config.Config) (*Application, error) {
	logger, err := logging.New(cfg.Logger)
	if err != nil {
		return nil, err
	}

	service, err := service.New()
	if err != nil {
		return nil, err
	}

	handler := router.New(cfg.Router, service)

	appInstance := &Application{
		logger: logger,
		cfg:    cfg,
		srv:    server.New(cfg.Server, handler),
	}

	return appInstance, nil
}

// Run function
func (a *Application) Run() {
	defer a.Shutdown()

	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Fatal("listen error: %s\n", err)
		}
	}()

	<-doneChan
}

// Shutdown function
func (a *Application) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), a.cfg.Shutdown)
	defer func() {
		cancel()
	}()

	if err := a.srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		a.logger.Fatal("shutdown error: %s\n", err)
	}
}
