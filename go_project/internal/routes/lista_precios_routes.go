package routes

import (
	"github.com/gin-gonic/gin"

	"go_project/internal/controllers"
)

func RegisterListaPreciosRoutes(r *gin.RouterGroup, c *controllers.ListaPreciosController) {
	listaPrecios := r.Group("/listas_precios")
	{
		listaPrecios.GET("", c.GetAll)
		listaPrecios.GET("/:id", c.GetByID)
		listaPrecios.POST("", c.Create)
		listaPrecios.PUT("/:id", c.Update)
		listaPrecios.DELETE("/:id", c.Delete)
	}
}
