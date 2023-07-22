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

func Stteamrouter(router *gin.Engine, db *sql.DB) {

	jwtStTeamRepo := jwt_repository.NewStTeamLoginRepo(db)
	jwtStTeamUsecase := jwt_usecase.NewStTeamLoginUsecase(jwtStTeamRepo)
	jwtStTeamCrtl := jwt_controller.NewStTeamLoginController(jwtStTeamUsecase, 60)

	storeteamRepo := repository.NewStoreteamRepo(db)
	storeteamUsecase := usecase.NewStoreteamUsecase(storeteamRepo)
	storeteamCtrl := controllers.NewStoreteamController(storeteamUsecase)

	//============================= Warehouse/SuperAdmin akses(register dan delete st_team) ===============================

	adminWhRouter := router.Group("/warehouse-team/employees")
	adminWhRouter.Use(jwt_middleware.AuthMiddleware())
	adminWhRouter.POST("/admin-st", storeteamCtrl.Register)
	adminWhRouter.DELETE("/admin-st/:id", storeteamCtrl.Unreg)

	router.POST("/auth/login-st", jwtStTeamCrtl.LoginStTeam)

	storeteamRouter := router.Group("/admin-st/employees")
	storeteamRouter.Use(jwt_middleware.AuthMiddleware())
	//akses untuk st_team
	storeteamRouter.GET("/find-all", storeteamCtrl.FindSellers)
	storeteamRouter.PUT("/update", storeteamCtrl.Edit)
	storeteamRouter.GET("/:id", storeteamCtrl.FindSellerById)
}
