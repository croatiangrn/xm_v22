package main

import (
	"fmt"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/database"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/http"
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

	http.InitRouter(cfg)
}
