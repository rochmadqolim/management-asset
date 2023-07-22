package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func ProductSoRouter(router *gin.Engine, db *sql.DB) {
	pdtSoRepository := repository.NewProductSoRepo(db)
	pdtSoUsecase := usecase.NewProductSoUsecase(pdtSoRepository)
	pdtSoCtrl := controllers.NewProductSoController(pdtSoUsecase)

	adminIcRouter := router.Group("/admin-ic/employees")
	adminIcRouter.Use(jwt_middleware.AuthMiddleware())

	adminIcRouter.GET("/product-so/getall", pdtSoCtrl.FindAllProductsSo)
	adminIcRouter.GET("/product-so/lessthan", pdtSoCtrl.FindProductsSoByLessThan)

}
