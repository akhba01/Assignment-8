package router

import (
	"assignment8/controllers"
	"assignment8/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	productRouter := r.Group("/products")
	{
		productRouter.Use(middlewares.Aunthentication())

		productRouter.POST("/", controllers.CreateProduct)

		productRouter.PUT("/:productID", middlewares.AuthorizeProductAccessOnlyAdmin(), middlewares.ProductAuthorization(), controllers.UpdateProduct)

		productRouter.GET("/:productID", middlewares.ProductAuthorization(), controllers.ReadProductByID)

		productRouter.GET("", middlewares.AuthorizeProductAccessOnlyAdmin(), controllers.ReadProducts)

		productRouter.DELETE("/:productID", middlewares.AuthorizeProductAccessOnlyAdmin(), middlewares.ProductAuthorization(), controllers.DeleteProductHandler)
	}
	return r
}
