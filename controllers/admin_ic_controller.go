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

var Validate = validator.New()

type AdminIcController struct {
	usecase usecase.AdminIcUsecase
}

func (c *AdminIcController) FindAdminIc(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	res := c.usecase.FindAdminIc()
	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

// ================================================Find By Id ==============================================================
func (c *AdminIcController) FindAdminIcByid(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	id := ctx.Param("id")

	photo := c.usecase.FindAdminIcById(id)
	filepath := fmt.Sprintf("./images/%s", photo)

	file, err := os.Open(filepath)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "file not found"})
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

func (c *AdminIcController) Register(ctx *gin.Context) {
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "wh" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	//===================================== Validation =================================
	var newAdminIc entity.AdminIc

	if err := ctx.ShouldBind(&newAdminIc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input"})
		return
	}

	err := Validate.Struct(newAdminIc)
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
	//==================================== Photo ====================================
	file, err := ctx.FormFile("photo")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Input photo"})
		return
	}

	filename := file.Filename
	newAdminIc.Photo = filename

	if errFile := ctx.SaveUploadedFile(file, fmt.Sprintf("./images/%s", filename)); errFile != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	//============================================ Response ========================
	res := c.usecase.Register(&newAdminIc)

	ctx.JSON(http.StatusCreated, gin.H{
		"data":       res,
		"login with": email,
	})

}

// =================================================== Edit ======================================
func (c *AdminIcController) Edit(ctx *gin.Context) {

	claims := ctx.MustGet("claims").(jwt.MapClaims)
	email := claims["email"].(string)
	role := claims["role"].(string)
	if role != "ic" {
		ctx.JSON(http.StatusUnauthorized, gin.H{
			"error": "you have no access to this role",
		})
		return
	}

	var adminIc entity.AdminIc

	if err := ctx.BindJSON(&adminIc); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res := c.usecase.Edit(&adminIc)

	ctx.JSON(http.StatusOK, gin.H{
		"data":  res,
		"email": email,
	})
}

func (c *AdminIcController) Unreg(ctx *gin.Context) {
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

	filename := c.usecase.FindAdminIcById(id)

	res := c.usecase.Unreg(id)
	if res == "admin not found" {
		ctx.JSON(http.StatusBadRequest, "invalid input ID")
		return
	}

	if filename != "" {
		err := os.Remove(fmt.Sprintf("./images/%s", filename))
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "deleted success"})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data":       res,
		"login with": email,
	})
}

func NewAdminIcController(c usecase.AdminIcUsecase) *AdminIcController {
	controller := AdminIcController{
		usecase: c,
	}
	return &controller
}
