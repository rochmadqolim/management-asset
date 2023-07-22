package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouterProductWh(router *gin.Engine, db *sql.DB) {
	// Dependencies Product Warehouse
	productWhRepo := repository.NewProductWhRepo(db)
	productWhUsecase := usecase.NewProductWhUsecase(productWhRepo)
	productWhController := controllers.NewProductWhController(productWhUsecase)

	// group products
	adminWhRouter := router.Group("/warehouse-team/employees")
	adminWhRouter.Use(jwt_middleware.AuthMiddleware())

	// routes (GET, POST, PUT, DELETE)
	adminWhRouter.GET("/product-wh", productWhController.FindProducts)
	adminWhRouter.GET("/product-wh/:id", productWhController.FindProductById)
	adminWhRouter.POST("/product-wh/register", productWhController.Input)
	adminWhRouter.PUT("/product-wh/update", productWhController.Edit)
	adminWhRouter.DELETE("product-wh/:id", productWhController.Output)
}
