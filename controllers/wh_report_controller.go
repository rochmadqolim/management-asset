package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ReportTrxWhController struct {
	usecase usecase.ReportTrxWhUsecase
}

func (c *ReportTrxWhController) FindAllReportTrxWh(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindAllReportTrxWh()
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *ReportTrxWhController) FindByIdReportTrxWh(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var reportsTrxWh entity.ReportWh

	if err := ctx.BindJSON(&reportsTrxWh); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.FindByIdReportTrxWh(reportsTrxWh.ProductWhId)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *ReportTrxWhController) FindByDateReportTrxWh(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var reportsTrxWh entity.ReportWh

	if err := ctx.BindJSON(&reportsTrxWh); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.FindByDateReportTrxWh(reportsTrxWh.CreatedAt)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func NewReportTrxWhController(u usecase.ReportTrxWhUsecase) *ReportTrxWhController {
	controller := ReportTrxWhController{
		usecase: u,
	}

	return &controller
}
