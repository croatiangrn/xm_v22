package http

import (
	"github.com/croatiangrn/xm_v22/internal/controller/http/dto"
	"github.com/croatiangrn/xm_v22/internal/usecase/company"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx := c.Request.Context()

	comp, err := h.uc.GetCompany(ctx, idAsUUID)
	if err != nil {
		HandleError(c, err)
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
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, companyObj)
}

func (h *CompanyHandler) CompanyUpdate(c *gin.Context) {
	id := c.Param("id")

	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	ctx := c.Request.Context()

	var req dto.UpdateCompanyRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	companyObj, err := h.uc.UpdateCompany(ctx, req, idAsUUID)
	if err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusCreated, companyObj)
}

func (h *CompanyHandler) CompanyDelete(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	idAsUUID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.uc.DeleteCompany(ctx, idAsUUID); err != nil {
		HandleError(c, err)
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
