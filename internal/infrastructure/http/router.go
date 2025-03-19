package http

import (
	httpController "github.com/croatiangrn/xm_v22/internal/controller/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"

	"github.com/croatiangrn/xm_v22/internal/infrastructure/http/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(companyHandler *httpController.CompanyHandler, cfg config.Config) {
	router := gin.Default()

	v1API := router.Group("/v1")
	{
		companiesAPI := v1API.Group("/companies", middleware.JWTAuthMiddleware(cfg.JWTSecret))
		{
			companiesAPI.POST("", companyHandler.CompanyCreate)
			companiesAPI.GET("/:id", companyHandler.CompanyGet)
			companiesAPI.PUT("/:id", companyHandler.CompanyUpdate)
			companiesAPI.DELETE("/:id", companyHandler.CompanyDelete)
		}
	}

	// Run the server
	router.Run(cfg.ServerPort)
}
