package controllers

import (
	"assignment8/database"
	"assignment8/helpers"
	"assignment8/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ReadProducts(ctx *gin.Context) {
	db := database.GetDB()
	var products []models.Product

	if err := db.Find(&products).Error; err != nil {

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func ReadProductByID(ctx *gin.Context) {
	db := database.GetDB()
	var products models.Product

	if err := db.Where("id= ?", ctx.Param("productID")).First(&products).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"products": products,
	})
}

func CreateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)

	Product := models.Product{}
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID

	err := db.Debug().Create(&Product).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, Product)
}

func UpdateProduct(ctx *gin.Context) {
	db := database.GetDB()
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(ctx)
	Product := models.Product{}

	productId, _ := strconv.Atoi(ctx.Param("productId"))
	userID := uint(userData["id"].(float64))

	if contentType == appJSON {
		ctx.ShouldBindJSON(&Product)
	} else {
		ctx.ShouldBind(&Product)
	}

	Product.UserID = userID
	Product.ID = uint(productId)

	err := db.Model(&Product).Where("id = ?", productId).Updates(models.Product{
		Title:       Product.Title,
		Description: Product.Description,
	}).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err":     "Bad Request",
			"message": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, Product)
}

func DeleteProductHandler(ctx *gin.Context) {
	productID := ctx.Param("productID")
	db := database.GetDB()

	var product models.Product

	if err := db.Where("id = ?", productID).First(&product).Error; err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": "product not found",
		})
		return
	}

	if err := db.Delete(&product).Error; err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "product deleted succesfully",
	})

}
