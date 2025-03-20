package main

import (
	"context"
	"fmt"
	httpController "github.com/croatiangrn/xm_v22/internal/controller/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/database"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/http"
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

	dbPgxPoolDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	if err := database.RunMigrations(dbPgxPoolDSN, "file://migrations"); err != nil {
		log.Printf("Failed to run migrations: %v", err)
	}

	dbPgxPool, err := pgxpool.New(context.Background(), dbPgxPoolDSN)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}

	kafkaProducer := kafka.NewProducer([]string{cfg.KafkaBrokers}, cfg.KafkaTopic)

	companyRepo := repository.NewCompanyRepository(dbPgxPool)
	companyUseCase := company.NewInteractor(companyRepo, kafkaProducer)
	companyHandler := httpController.NewCompanyHandler(companyUseCase, Logger)

	loginHandler := httpController.NewLoginHandler()

	http.InitRouter(
		loginHandler,
		companyHandler,
		cfg)
}
