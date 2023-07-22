package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TrxWhController struct {
	usecase usecase.TrxWhUsecase
}

func (c *TrxWhController) EnrollInsertTrxWh(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newTrxWh entity.TrxWh

	if err := ctx.BindJSON(&newTrxWh); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.EnrollInsertTrxWh(&newTrxWh)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})

}

func NewTrxWhController(u usecase.TrxWhUsecase) *TrxWhController {
	controller := TrxWhController{
		usecase: u,
	}

	return &controller
}
