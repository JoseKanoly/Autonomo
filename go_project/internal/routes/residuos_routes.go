package routes

import (
	"github.com/gin-gonic/gin"

	"go_project/internal/controllers"
)

func RegisterResiduosRoutes(r *gin.RouterGroup, c *controllers.ResiduosController) {
	residuos := r.Group("/residuos")
	{
		residuos.GET("", c.GetAll)
		residuos.GET("/:id", c.GetByID)
		residuos.POST("", c.Create)
		residuos.PUT("/:id", c.Update)
		residuos.DELETE("/:id", c.Delete)
	}
}
