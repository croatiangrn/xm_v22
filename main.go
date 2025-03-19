package main

import (
	"fmt"
	httpController "github.com/croatiangrn/xm_v22/internal/controller/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/database"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/repository"
	"github.com/croatiangrn/xm_v22/internal/usecase/company"
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

	// Initialize repository
	companyRepo := repository.NewCompanyRepository(db)
	companyUseCase := company.NewInteractor(companyRepo)
	companyHandler := httpController.NewCompanyHandler(companyUseCase)

	http.InitRouter(companyHandler, cfg)
}
