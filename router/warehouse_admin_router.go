package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_controller "github.com/Uchel/auth-final/controller"
	jwt_middleware "github.com/Uchel/auth-final/middleware"
	jwt_repository "github.com/Uchel/auth-final/repository"
	jwt_usecase "github.com/Uchel/auth-final/usecase"

	"github.com/gin-gonic/gin"
)

func InitRouterEmployee(router *gin.Engine, db *sql.DB) {
	// auth jwt login ExampleAdmin
	jwtAdminWhRepo := jwt_repository.NewAdminWhLoginRepo(db)
	jwtAdminWhUsecase := jwt_usecase.NewAdminWhUsecase(jwtAdminWhRepo)
	jwtAdminWhCtrl := jwt_controller.NewAdminLoginController(jwtAdminWhUsecase, 60)

	// Dependencies Warehouse Team
	warehouseTeamRepo := repository.NewWarehouseTeamRepo(db)
	warehouseTeamUsecase := usecase.NewWarehouseTeamUsecase(warehouseTeamRepo)
	warehouseTeamController := controllers.NewWarehouseTeamController(warehouseTeamUsecase)

	// Login session
	router.POST("/auth/login-wh", jwtAdminWhCtrl.LoginAdmin)
	router.POST("/register/admin-wh", warehouseTeamController.Register)

	// Route group after authentication
	adminWhRouter := router.Group("/warehouse-team/employees")

	adminWhRouter.Use(jwt_middleware.AuthMiddleware())
	// routes (GET, POST, PUT, DELETE)
	adminWhRouter.GET("/find-all", warehouseTeamController.FindEmployees)
	adminWhRouter.GET("/:id", warehouseTeamController.FindEmployeeById)
	adminWhRouter.PUT("/update", warehouseTeamController.Edit)
	adminWhRouter.DELETE("/:id", warehouseTeamController.Unreg)

}
