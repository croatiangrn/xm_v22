package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	httpController "github.com/croatiangrn/xm_v22/internal/controller/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/database"
	appHttp "github.com/croatiangrn/xm_v22/internal/infrastructure/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/kafka"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/repository"
	"github.com/croatiangrn/xm_v22/internal/usecase/company"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"log"
)

var Logger zerolog.Logger

func initLogger() zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(zerolog.ConsoleWriter{Out: log.Writer()}).With().Timestamp().Logger()
	return logger
}

func main() {
	Logger = initLogger()

	cfg, err := config.Load("./")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	dbPgxPoolDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err := database.RunMigrations(dbPgxPoolDSN, "file://migrations"); err != nil {
		Logger.Error().Err(err).Msg("Failed to run migrations")
	}

	dbPgxPool, err := pgxpool.New(ctx, dbPgxPoolDSN)
	if err != nil {
		Logger.Fatal().Err(err).Msg("Failed to create connection pool")
	}
	defer dbPgxPool.Close()

	kafkaProducer := kafka.NewProducer([]string{cfg.KafkaBrokers}, cfg.KafkaTopic)
	defer kafkaProducer.Close()

	companyRepo := repository.NewCompanyRepository(dbPgxPool)
	companyUseCase := company.NewInteractor(companyRepo, kafkaProducer)
	companyHandler := httpController.NewCompanyHandler(companyUseCase, Logger)
	loginHandler := httpController.NewLoginHandler()

	router := appHttp.InitRouter(loginHandler, companyHandler, cfg)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.ServerPort),
		Handler: router,
	}

	go func() {
		Logger.Info().Msgf("Starting server on port %s", cfg.ServerPort)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			Logger.Fatal().Err(err).Msg("Server failed to start")
		}
	}()

	<-ctx.Done()

	stop()
	Logger.Info().Msg("Shutting down gracefully...")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		Logger.Error().Err(err).Msg("Server forced to shutdown")
	}

	if err := kafkaProducer.Close(); err != nil {
		Logger.Error().Err(err).Msg("Failed to flush Kafka producer")
	}

	Logger.Info().Msg("Server exited properly")
}
