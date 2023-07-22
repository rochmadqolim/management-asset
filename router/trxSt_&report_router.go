package router

import (
	"database/sql"
	"go_inven_ctrl/controllers"
	"go_inven_ctrl/repository"
	"go_inven_ctrl/usecase"

	jwt_middleware "github.com/Uchel/auth-final/middleware"
	"github.com/gin-gonic/gin"
)

func TrxStRoutes(router *gin.Engine, db *sql.DB) {

	//controller Trx ProductSt

	trxInstRepo := repository.NewTrxInStRepo(db)
	trxInStUsecase := usecase.NewTrxInStUsecase(trxInstRepo)
	trxStctrl := controllers.NewTrxInStController(trxInStUsecase)

	//cotroller Report Trx Product St
	reportTrxStRepo := repository.NewReportTrxStRepo(db)
	reportTrxStUsecase := usecase.NewReportTrxStUsecase(reportTrxStRepo)
	reportTrxStctrl := controllers.NewReportTrxStController(reportTrxStUsecase)

	//Akses admin-st untuk input, sold, retur serta reportnya
	storeteamRouter := router.Group("/admin-st/employees")
	storeteamRouter.Use(jwt_middleware.AuthMiddleware())

	storeteamRouter.POST("/trx/transaction", trxStctrl.EnrollInsertTrxInSt)
	storeteamRouter.GET("/trx/report-st-id", reportTrxStctrl.FindByReportTrxProductStId)
	storeteamRouter.GET("/trx/report-st-date", reportTrxStctrl.FindByDateReportTrxSt)
	storeteamRouter.GET("/trx/all-report", reportTrxStctrl.FindAllReportrxSt)
}
