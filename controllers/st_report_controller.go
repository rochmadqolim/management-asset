package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ReportTrxStController struct {
	usecase usecase.ReportTrxStUsecase
}

func (c *ReportTrxStController) FindAllReportrxSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}
	res := c.usecase.FindAllReportTrxSt()
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *ReportTrxStController) FindByReportTrxProductStId(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var reportsTrxSt entity.ReportTrxSt
	if err := ctx.BindJSON(&reportsTrxSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}
	res := c.usecase.FindByReportTrxProductStId(reportsTrxSt.ProductStId)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *ReportTrxStController) FindByDateReportTrxSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var reportsTrxSt entity.ReportTrxSt
	if err := ctx.BindJSON(&reportsTrxSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}
	res := c.usecase.FindByDateReportTrxSt(reportsTrxSt.CreatedAt)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func NewReportTrxStController(u usecase.ReportTrxStUsecase) *ReportTrxStController {
	controller := ReportTrxStController{
		usecase: u,
	}

	return &controller
}
