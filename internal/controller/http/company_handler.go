package http

import (
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
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

func (h *CompanyHandler) CompanyGet(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	comp, err := h.uc.GetCompany(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}
	c.JSON(http.StatusOK, comp)
}

func (h *CompanyHandler) CompanyCreate(c *gin.Context) {
	var req dto.CreateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	companyObj, err := h.uc.CreateCompany(ctx, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, companyObj)
}

func (h *CompanyHandler) CompanyUpdate(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	companyObj, err := h.uc.GetCompany(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Company not found"})
		return
	}

	var req dto.UpdateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.UpdateCompany(ctx, companyObj, req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, companyObj)
}

func (h *CompanyHandler) CompanyDelete(c *gin.Context) {

}
