package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func TrxSoRoutes(router *gin.Engine, db *sql.DB) {
	trxSoRepo := repository.NewTrxInSoRepo(db)
	trxSoUsecase := usecase.NewTrxInSoUsecase(trxSoRepo)
	trxSoCtrl := controllers.NewTrxInSoController(trxSoUsecase)

	reportSoRepo := repository.NewReportSoRepo(db)
	reportSoUsecase := usecase.NewReportSoUseCase(reportSoRepo)
	reportSoCtrl := controllers.NewReportSoController(reportSoUsecase)

	adminIcRouter := router.Group("/admin-ic/employees")

	adminIcRouter.Use(jwt_middleware.AuthMiddleware())

	adminIcRouter.GET("/trx/so_interim_report", reportSoCtrl.FindAllInterimSoReport)
	adminIcRouter.GET("/trx/so_detail_report", reportSoCtrl.FindAllReportSoDetail)

	adminIcRouter.POST("/trx/trx_so", trxSoCtrl.EnrollInsertTrxSo)
	adminIcRouter.GET("/trx/trx_so_inconfirm", trxSoCtrl.ReportInterim)
	adminIcRouter.GET("/trx/trx_so_confirm", trxSoCtrl.ReportConfirmation)
}
