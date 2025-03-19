package http

import (
	"github.com/croatiangrn/xm_v22/internal/usecase/company"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CompanyHandler struct {
	uc company.UseCase
}

func NewCompanyHandler(uc company.UseCase) *CompanyHandler {
	return &CompanyHandler{uc}
}

func (h *CompanyHandler) GetCompany(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	comp, err := h.uc.GetCompany(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, comp)
}
