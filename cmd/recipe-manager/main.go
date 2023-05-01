package main

import (
	"context"
	"fmt"
	"github.com/recipe-manager/cmd/recipe-manager/rest"
	"github.com/recipe-manager/internal/config"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func main() {

	//initialize zerologger instance
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	cfg, err := config.Get()
	if err != nil {
		logger.Fatal().Err(err).Msg("missing env vars")
	}

	restClient, err := rest.New(&logger)
	if err != nil {
		logger.Fatal().Err(err).Msg("error creating rest service")
	}

	logger.Info().
		Str("host", cfg.Host).
		Uint("port", cfg.Port).
		Msg("Starting HTTP listener")

	port := strconv.FormatUint(uint64(cfg.Port), 10)
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: restClient.Router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal().Msg(fmt.Sprintf("error connecting the server: %v", err))
		}
	}()

	//Init shutting down gracefully
	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)
	<-shutdown
	logger.Info().Msg("shutting down server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logger.Fatal().Msg(fmt.Sprintf("failed to server shutdown due: %v", err))
	}
	//catching ctx.Done(). Timeout of 5 seconds.
	select {
	case <-ctx.Done():
		logger.Info().Msg("timeout of 5 seconds")
	}
	logger.Info().Msg("the server has been turned off gracefully")
}
