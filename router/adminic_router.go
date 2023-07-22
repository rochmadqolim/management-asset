package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"

	// "go_inven_ctrl/middleware"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	"github.com/gin-gonic/gin"

	jwt_controller "github.com/Uchel/auth-final/controller"
	jwt_middleware "github.com/Uchel/auth-final/middleware"
	jwt_repository "github.com/Uchel/auth-final/repository"
	jwt_usecase "github.com/Uchel/auth-final/usecase"
)

func RouterAdminIc(router *gin.Engine, db *sql.DB) {

	jwtAdminIcRepo := jwt_repository.NewIcTeamLoginRepo(db)
	jwtIcUsecase := jwt_usecase.NewIcTeamUsecase(jwtAdminIcRepo)
	jwtIcCtrl := jwt_controller.NewIcTeamLoginController(jwtIcUsecase, 60)

	adminIcRepo := repository.NewAdminIcRepo(db)
	adminIcUsecase := usecase.NewAdminIcUsecase(adminIcRepo)
	adminCtrl := controllers.NewAdminIcController(adminIcUsecase)

	//============================= Warehouse/SuperAdmin akses(register dan delete ic_team/admin_id) ===============================

	adminWhRouter := router.Group("/warehouse-team/employees")
	adminWhRouter.Use(jwt_middleware.AuthMiddleware())
	adminWhRouter.POST("/admin-ic", adminCtrl.Register)
	adminWhRouter.DELETE("/admin-ic/:id", adminCtrl.Unreg)

	//=========================================================================================================================

	router.POST("/auth/login-ic", jwtIcCtrl.LoginIcTeam)

	adminIcRouter := router.Group("/admin-ic/employees")

	adminIcRouter.Use(jwt_middleware.AuthMiddleware())
	//akses admin Ic
	adminIcRouter.GET("/find-all", adminCtrl.FindAdminIc)
	adminIcRouter.PUT("/update", adminCtrl.Edit)
	adminIcRouter.GET("/:id", adminCtrl.FindAdminIcByid)
}
