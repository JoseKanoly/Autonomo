package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"go_project/internal/models"
	"go_project/internal/services"
)

type ListaPreciosController struct {
	service *services.ListaPreciosService
}

func NewListaPreciosController(service *services.ListaPreciosService) *ListaPreciosController {
	return &ListaPreciosController{service: service}
}

func (c *ListaPreciosController) GetAll(ctx *gin.Context) {
	listaPrecios, err := c.service.GetAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, listaPrecios)
}

func (c *ListaPreciosController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	listaPrecio, err := c.service.GetByID(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Lista de precios not found"})
		return
	}
	ctx.JSON(http.StatusOK, listaPrecio)
}

func (c *ListaPreciosController) Create(ctx *gin.Context) {
	var listaPrecio models.ListaPrecios
	if err := ctx.ShouldBindJSON(&listaPrecio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.service.Create(&listaPrecio); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, listaPrecio)
}

func (c *ListaPreciosController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	var listaPrecio models.ListaPrecios
	if err := ctx.ShouldBindJSON(&listaPrecio); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	listaPrecio.ID = id
	if err := c.service.Update(&listaPrecio); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, listaPrecio)
}

func (c *ListaPreciosController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	if err := c.service.Delete(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Lista de precios deleted successfully"})
}
