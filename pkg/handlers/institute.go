package handlers

import (
	"net/http"
	"strconv"

	"github.com/blanc42/mooca/pkg/models"
	"github.com/blanc42/mooca/pkg/usecases"
	"github.com/gin-gonic/gin"
)

type InstituteHandler struct {
	instituteUseCase usecases.InstituteUseCase
}

func NewInstituteHandler(instituteUseCase usecases.InstituteUseCase) *InstituteHandler {
	return &InstituteHandler{instituteUseCase: instituteUseCase}
}

func (h *InstituteHandler) CreateInstitute(c *gin.Context) {
	var input models.Institute
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.instituteUseCase.CreateInstitute(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *InstituteHandler) GetInstituteByID(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid institute ID"})
		return
	}

	institute, err := h.instituteUseCase.GetInstituteByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Institute not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institute})
}

func (h *InstituteHandler) GetAllInstitutes(c *gin.Context) {
	institutes, err := h.instituteUseCase.GetAllInstitutes()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": institutes})
}

func (h *InstituteHandler) UpdateInstitute(c *gin.Context) {
	var input models.Institute
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.instituteUseCase.UpdateInstitute(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": input})
}

func (h *InstituteHandler) DeleteInstitute(c *gin.Context) {
	id, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid institute ID"})
		return
	}

	if err := h.instituteUseCase.DeleteInstitute(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Institute deleted"})
}
