package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go_project/internal/models"
	"go_project/internal/services"
)

type ResiduosController struct {
	service *services.ResiduosService
}

func NewResiduosController(service *services.ResiduosService) *ResiduosController {
	return &ResiduosController{service: service}
}

func (c *ResiduosController) GetAll(ctx *gin.Context) {
	residuos, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, residuos)
}

func (c *ResiduosController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	residuo, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Residuo not found"})
		return
	}
	ctx.JSON(http.StatusOK, residuo)
}

func (c *ResiduosController) Create(ctx *gin.Context) {
	var residuo models.Residuo
	if err := ctx.ShouldBindJSON(&residuo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(&residuo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, residuo)
}

func (c *ResiduosController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var residuo models.Residuo
	if err := ctx.ShouldBindJSON(&residuo); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	residuo.ID = id
	if err := c.service.Update(&residuo); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, residuo)
}

func (c *ResiduosController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Residuo deleted successfully"})
}
