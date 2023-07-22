package controllers

import (
	"fmt"
	"go_inven_ctrl/entity"
	"go_inven_ctrl/usecase"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type StoreteamController struct {
	usecase usecase.StoreteamUsecase
}

// ============================================================ Find All =========================================================
func (c *StoreteamController) FindSellers(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindSellers()

	ctx.JSON(200, gin.H{
		"data":        res,
		"login email": email,
	})
}

// ============================================================= Find photo by Id==============================================================
func (c *StoreteamController) FindSellerById(ctx *gin.Context) {

	//Get Authorization
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	photo := c.usecase.FindSellerById(id)

	filePath := fmt.Sprintf("./images/%s", photo)
	file, err := os.Open(filePath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "error found",
		})

		return
	}

	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	fileBytes := make([]byte, fileInfo.Size())
	_, err = file.Read(fileBytes)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	ctx.Data(http.StatusOK, "image/jpeg", fileBytes)

}

// ==================================================== Register ===============================================================
func (c *StoreteamController) Register(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var newSeller entity.Storeteam

	if err := ctx.ShouldBind(&newSeller); err != nil {
		fmt.Println(err)

		ctx.JSON(http.StatusBadRequest, "invalid input")
		return
	}

	err := Validate.Struct(newSeller)
	if err != nil {
		for _, errfield := range err.(validator.ValidationErrors) {
			if errfield.Field() == "Name" {
				ctx.JSON(400, gin.H{"error": "Min 5 characters for name"})

			}
			if errfield.Field() == "Password" {
				ctx.JSON(400, gin.H{"error": "Min 8 characters for password"})

			}
			if errfield.Field() == "Email" {
				ctx.JSON(400, gin.H{"error": "It was not an email"})

			}
		}
		return
	}

	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid input photo",
		})

		return
	}
	filename := file.Filename
	newSeller.Photo = filename

	if err := ctx.SaveUploadedFile(file, fmt.Sprintf("./images/%s", filename)); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "error server",
		})
	}

	res := c.usecase.Register(&newSeller)
	ctx.JSON(201, gin.H{
		"data":       res,
		"login with": email,
	})
}

// ======================================================== Edit ==================================================================================
func (c *StoreteamController) Edit(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "st" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var seller entity.Storeteam

	if err := ctx.BindJSON(&seller); err != nil {
		ctx.JSON(400, "invalid input")
		return
	}

	res := c.usecase.Edit(&seller)
	ctx.JSON(200, gin.H{
		"data":             res,
		"login with email": email,
	})
}

// ======================================================== Delete ==================================================================================
func (c *StoreteamController) Unreg(ctx *gin.Context) {
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

	photo := c.usecase.FindSellerById(id)

	filePath := fmt.Sprintf("./images/%s", photo)

	err := os.Remove(filePath)
	if err != nil {
		log.Println(err)
	}

	res := c.usecase.Unreg(id)

	ctx.JSON(http.StatusOK, gin.H{
		"data":             res,
		"login with email": email,
	})
}

func NewStoreteamController(u usecase.StoreteamUsecase) *StoreteamController {
	controller := StoreteamController{
		usecase: u,
	}

	return &controller
}
