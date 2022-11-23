package app

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/matsuev/klsh-registrator/internal/config"
	"github.com/matsuev/klsh-registrator/internal/logging"
	"github.com/matsuev/klsh-registrator/internal/server"
)

// Application struct
type Application struct {
	logger *logging.Logger
	cfg    *config.Config
	srv    *http.Server
}

// New function
func New(cfg *config.Config, logger *logging.Logger) (*Application, error) {
	var handler http.Handler

	appInstance := &Application{
		logger: logger,
		cfg:    cfg,
		srv:    server.New(cfg.Server, handler),
	}

	return appInstance, nil
}

// Run function
func (a *Application) Run() error {
	doneChan := make(chan os.Signal, 1)
	signal.Notify(doneChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := a.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			a.logger.Infof("listen: %s\n", err)
		}
	}()

	<-doneChan

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err := a.srv.Shutdown(ctx); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
