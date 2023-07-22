package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouterTrxReportWh(router *gin.Engine, db *sql.DB) {
	// Dependencies Warehouse Transaction
	trxWhRepo := repository.NewTrxWhRepo(db)
	trxWhUsecase := usecase.NewTrxWhUsecase(trxWhRepo)
	trxWhController := controllers.NewTrxWhController(trxWhUsecase)

	// Dependencies Warehouse Report
	reportTrxWhRepo := repository.NewReportTrxWhRepo(db)
	reportTrxWhUsecase := usecase.NewReportTrxWhUsecase(reportTrxWhRepo)
	reportTrxWhController := controllers.NewReportTrxWhController(reportTrxWhUsecase)

	// Warehouse Routes
	adminWhRouter := router.Group("/warehouse-team/employees")
	adminWhRouter.Use(jwt_middleware.AuthMiddleware())

	adminWhRouter.POST("/trx/transaction", trxWhController.EnrollInsertTrxWh)
	adminWhRouter.GET("/trx/allreport", reportTrxWhController.FindAllReportTrxWh)
	adminWhRouter.GET("/trx/reportwhid", reportTrxWhController.FindByIdReportTrxWh)
	adminWhRouter.GET("/trx/reportwhdate", reportTrxWhController.FindByDateReportTrxWh)
}
