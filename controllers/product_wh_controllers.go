package controllers

import (
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductWhController struct {
	usecase usecase.ProductWhUsecase
}

func NewProductWhController(u usecase.ProductWhUsecase) *ProductWhController {

	controller := ProductWhController{
		usecase: u,
	}

	return &controller
}

func (c *ProductWhController) FindProducts(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindProducts()

	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ProductWhController) FindProductById(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	res := c.usecase.FindProductById(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ProductWhController) Input(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var product entity.ProductWh

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input data. Please check the request body and try again."})
		return
	}

	err := Validate.Struct(product)
	if err != nil {
		for _, errfield := range err.(validator.ValidationErrors) {

			if errfield.Field() == "Id" {
				ctx.JSON(400, gin.H{"error": "required Id, must filled in"})

			}
			if errfield.Field() == "ProductName" {
				ctx.JSON(400, gin.H{"error": "Min Product Name is 5 characters"})

			}
			if errfield.Field() == "Price" {
				ctx.JSON(400, gin.H{"error": "min Price equals 100"})

			}
			if errfield.Field() == "ProductCtg" {
				ctx.JSON(400, gin.H{"error": "Min Product Category is 3 characters"})

			}
		}
		return
	}

	res := c.usecase.Input(&product)

	ctx.JSON(http.StatusCreated, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ProductWhController) Edit(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var product entity.ProductWh

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.Edit(&product)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *ProductWhController) Output(ctx *gin.Context) {
	id := ctx.Param("id")

	res := c.usecase.Output(id)
	ctx.JSON(http.StatusOK, res)
}
