package http

import (
	httpController "github.com/croatiangrn/xm_v22/internal/controller/http"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/config"
	"github.com/croatiangrn/xm_v22/internal/infrastructure/http/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitRouter(
	loginHandler *httpController.LoginHandler,
	companyHandler *httpController.CompanyHandler,
	cfg config.Config) *gin.Engine {
	router := gin.Default()

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowHeaders = append(corsConfig.AllowHeaders, "Authorization")
	corsConfig.AllowOrigins = []string{"*"} // This should be replaced with an actual frontend URL for example
	router.Use(gin.Recovery(), cors.New(corsConfig))

	v1API := router.Group("/v1")
	{
		loginAPI := v1API.Group("/login")
		{
			loginAPI.POST("", loginHandler.Login(cfg.JWTSecret))
		}

		companiesAPI := v1API.Group("/companies")
		{
			companiesAPI.POST("", middleware.JWTAuthMiddleware(cfg.JWTSecret), companyHandler.CompanyCreate)
			companiesAPI.GET("/:id", companyHandler.CompanyGet)
			companiesAPI.PATCH("/:id", middleware.JWTAuthMiddleware(cfg.JWTSecret), companyHandler.CompanyUpdate)
			companiesAPI.DELETE("/:id", middleware.JWTAuthMiddleware(cfg.JWTSecret), companyHandler.CompanyDelete)
		}
	}

	return router
}
