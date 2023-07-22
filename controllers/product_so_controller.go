package controllers

import (
	"fmt"
	"net/http"

	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"

	// "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type ProductSoController struct {
	usecase usecase.ProductSoUsecase
}

func (c *ProductSoController) FindAllProductsSo(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}
	res := c.usecase.FindAllProductSo()
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func (c *ProductSoController) FindProductsSoByLessThan(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}
	var productSo entity.ProductSo

	if err := ctx.BindJSON(&productSo); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		fmt.Println(err)
		return
	}
	res := c.usecase.FindByLessThan(productSo.Stock)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func NewProductSoController(u usecase.ProductSoUsecase) *ProductSoController {
	controller := ProductSoController{
		usecase: u,
	}

	return &controller
}
