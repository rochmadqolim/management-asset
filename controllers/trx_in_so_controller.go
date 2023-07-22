package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TrxInSoController struct {
	usecase usecase.TrxInSoUsecase
}

func (c *TrxInSoController) EnrollInsertTrxSo(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newtrxInSo entity.TrxInSo

	if err := ctx.BindJSON(&newtrxInSo); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.EnrollInsertTrxSo(&newtrxInSo)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *TrxInSoController) ReportInterim(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.ReportInterim()
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *TrxInSoController) ReportConfirmation(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.ReportConfirmation()
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func NewTrxInSoController(u usecase.TrxInSoUsecase) *TrxInSoController {
	controller := TrxInSoController{
		usecase: u,
	}

	return &controller
}
