package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type TrxInStController struct {
	usecase usecase.TrxInStUsecase
}

func (c *TrxInStController) EnrollInsertTrxInSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newtrxInSt entity.TrxInST

	if err := ctx.BindJSON(&newtrxInSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	res := c.usecase.EnrollInsertTrx(&newtrxInSt)

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func NewTrxInStController(u usecase.TrxInStUsecase) *TrxInStController {
	controller := TrxInStController{
		usecase: u,
	}

	return &controller
}
