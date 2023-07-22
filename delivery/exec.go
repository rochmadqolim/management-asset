package delivery

import (
	"go_inven_ctrl/config"
	"go_inven_ctrl/router"
	"go_inven_ctrl/utils"

	"github.com/gin-gonic/gin"
)

func Exec() {
	r := gin.Default()
	db := config.ConnectDB()

	//Wh Team Router
	router.InitRouterEmployee(r, db)
	router.InitRouterProductWh(r, db)
	router.InitRouterTrxReportWh(r, db)

	//St Team Role
	router.Stteamrouter(r, db)
	router.ProducStRoutes(r, db)
	router.TrxStRoutes(r, db)

	//Ic Team Role
	router.RouterAdminIc(r, db)
	router.ProductSoRouter(r, db)
	router.TrxSoRoutes(r, db)

	r.Run(":" + utils.DotEnv("PORT"))
}
