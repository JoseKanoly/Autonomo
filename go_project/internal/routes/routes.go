package routes

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	"go_project/internal/controllers"
	"go_project/internal/repository"
	"go_project/internal/services"
)

func RegisterRoutes(r *gin.RouterGroup, db *sql.DB) {
	listaPreciosRepo := repository.NewListaPreciosRepository(db)
	listaPreciosService := services.NewListaPreciosService(listaPreciosRepo)
	listaPreciosController := controllers.NewListaPreciosController(listaPreciosService)

	residuosRepo := repository.NewResiduosRepository(db)
	residuosService := services.NewResiduosService(residuosRepo)
	residuosController := controllers.NewResiduosController(residuosService)

	r.POST("/login", controllers.Login)

	RegisterListaPreciosRoutes(r, listaPreciosController)
	RegisterResiduosRoutes(r, residuosController)
}
