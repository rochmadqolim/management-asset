package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func ProducStRoutes(router *gin.Engine, db *sql.DB) {

	//Controller ProductSt ProductSt
	productStRepo := repository.NewProductStRepo(db)
	productStUsecase := usecase.NewProductStUsecase(productStRepo)
	productStctrl := controllers.NewProductStController(productStUsecase)

	storeteamRouter := router.Group("/admin-st/employees")
	storeteamRouter.Use(jwt_middleware.AuthMiddleware())

	// Akses admin st untuk stok produk
	storeteamRouter.GET("/product-st", productStctrl.FindAllProductsSt)
	storeteamRouter.GET("/product-st/:id", productStctrl.FindProductsStById)
	storeteamRouter.POST("/product-st/register", productStctrl.RegisterProductSt)
	storeteamRouter.PUT("/product-st/edit", productStctrl.EditProductSt)
	storeteamRouter.DELETE("/product-st/:id", productStctrl.UnregProductSt)

}
