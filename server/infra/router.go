package infra

import (
	"jojogo/server/handler"
	"jojogo/server/infra/api"
	"jojogo/server/middleware"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

var Router = gin.Default()

func InitRouter() {
	api.Init()

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/group/view", api.GetGroups)
	Router.GET("/group/update", api.UpdateGroup)
	Router.GET("/group/get_by_name/:group_name", api.GetGroupByName)

	Router.POST("/group/create/:group_name", api.CreateGroup)

	Router.GET("/books", api.GetBooks)
	Router.GET("/books/:id", api.BookById)
	Router.POST("/login", api.LoginHandler)
	Router.GET("/info", middleware.Auth(), handler.GetUserInfo)

	Router.POST("/books", api.CreateBook)

	Router.PATCH("/checkout", api.CheckoutBook)
	Router.PATCH("/return", api.ReturnBook)
}
