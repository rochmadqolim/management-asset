package controllers

import (
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ReportSoController struct {
	usecase usecase.ReporSoUsecase
}

func (c *ReportSoController) FindAllInterimSoReport(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindAllInterimSoReport()
	ctx.JSON(http.StatusOK, gin.H{
		"data":             res,
		"login with email": email,
	})
}

func (c *ReportSoController) FindAllReportSoDetail(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindAllReportSoDetail()
	ctx.JSON(http.StatusOK, gin.H{
		"data":             res,
		"login with email": email,
	})
}

func NewReportSoController(u usecase.ReporSoUsecase) *ReportSoController {
	controller := ReportSoController{
		usecase: u,
	}

	return &controller
}
