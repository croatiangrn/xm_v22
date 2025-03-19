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
	"log"
)

func main() {
	cfg, err := config.Load("./")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	postgresDSN := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable", cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBName, cfg.DBPassword)
	gormDB := database.NewGormDB(postgresDSN)
	if err := database.AutoMigrate(gormDB); err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
		return
	}

	dbPgxPoolDSN := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable", cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName)
	dbPgxPool, err := pgxpool.New(context.Background(), dbPgxPoolDSN)
	if err != nil {
		log.Fatalf("Failed to create connection pool: %v", err)
	}

	kafkaProducer := kafka.NewProducer([]string{cfg.KafkaBrokers}, cfg.KafkaTopic)

	companyRepo := repository.NewCompanyRepository(dbPgxPool)
	companyUseCase := company.NewInteractor(companyRepo, kafkaProducer)
	companyHandler := httpController.NewCompanyHandler(companyUseCase)

	http.InitRouter(companyHandler, cfg)
}
