package controllers

import (
	"fmt"
	"net/http"

	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type ProductStController struct {
	usecase usecase.ProductStUsecase
}

// ================================================================= Find All Product =================================================
func (c *ProductStController) FindAllProductsSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindAllProductsSt()
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

// ========================================================Find Product By Id ========================================================================
func (c *ProductStController) FindProductsStById(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")
	fmt.Println(id)
	res := c.usecase.FindProductStById(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

// ======================================================== Register Product By Id ========================================================================
func (c *ProductStController) RegisterProductSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newProductSt entity.ProductSt

	if err := ctx.BindJSON(&newProductSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid input")
		fmt.Println(err)
		return
	}

	err := Validate.Struct(newProductSt)
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

	res := c.usecase.RegisterProductSt(&newProductSt)

	ctx.JSON(http.StatusOK, gin.H{
		"login with": email,
		"data":       res,
	})
}

// ========================================================Find Product By Id ========================================================================
func (c *ProductStController) EditProductSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var productSt entity.ProductSt

	if err := ctx.BindJSON(&productSt); err != nil {
		ctx.JSON(http.StatusBadRequest, "Invalid Input")
		return
	}

	res := c.usecase.EditProductSt(&productSt)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

// ========================================================Find Product By Id ========================================================================
func (c *ProductStController) UnregProductSt(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	res := c.usecase.UnregProductSt(id)
	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

// ==========================================================================================================================================
func NewProductStController(u usecase.ProductStUsecase) *ProductStController {
	controller := ProductStController{
		usecase: u,
	}

	return &controller
}
