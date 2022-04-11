package infra

import (
	"jojogo/server/api"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router = gin.Default()

func InitRouter() {
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/groups", api.GetGroups)
	Router.GET("/books", api.GetBooks)
	Router.GET("/books/:id", api.BookById)
	Router.POST("/books", api.CreateBook)
	Router.PATCH("/checkout", api.CheckoutBook)
	Router.PATCH("/return", api.ReturnBook)
}
