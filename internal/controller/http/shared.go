package http

import (
	"errors"
	customErrors "github.com/croatiangrn/xm_v22/internal/pkg/errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HandleError(c *gin.Context, err error) {
	var internalServerErr *customErrors.InternalServerError
	var badRequestErr *customErrors.BadRequestError
	var notFoundErr *customErrors.NotFoundError

	switch {
	case errors.As(err, &internalServerErr):
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	case errors.As(err, &badRequestErr):
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	case errors.As(err, &notFoundErr):
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
	default:
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unknown error"})
	}
}
