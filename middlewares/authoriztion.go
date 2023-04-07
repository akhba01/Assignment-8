package middlewares

import (
	"assignment8/database"
	"assignment8/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func ProductAuthorization() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		db := database.GetDB()

		productId, err := strconv.Atoi(ctx.Param("productID"))
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"error":   "Bad Request",
				"message": "Invalid paramater",
			})
			return
		}
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint(userData["id"].(float64))
		Product := models.Product{}
		err = db.Select("user_id").First(&Product, uint(productId)).Error

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"error":   "Data Not Found",
				"message": "data doesn't exist",
			})
			return
		}

		if Product.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "you are not allowed to access this data",
			})
			return
		}
		ctx.Next()
	}

}

func AuthorizeProductAccessOnlyAdmin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		userData := ctx.MustGet("userData").(jwt.MapClaims)
		UserRole := userData["role"]
		if UserRole != "admin" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error":   "Unauthorized",
				"message": "Non admin user do not have permission to access this data",
			})
			ctx.Abort()
			return
		}
	}
}
