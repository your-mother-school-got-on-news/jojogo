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

	// login
	Router.POST("/register", api.Register)
	Router.POST("/login", api.LoginHandler)
	Router.GET("/info", middleware.Auth(), handler.GetUserInfo)

	// group ops
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/group/view", api.GetGroups)
	Router.GET("/group/get/name/:group_name", api.GetGroupByName)

	Router.POST("/group/new/:group_name", middleware.Auth(), api.CreateGroup)

	Router.PATCH("/group/update/name/:search_name/:set_name", middleware.Auth(), api.UpdateGroupName)
	Router.PATCH("/group/update/person/add/:search_name/:person_name", middleware.Auth(), api.AddToGroup)
	Router.PATCH("/group/update/person/del/:search_name/:person_name", middleware.Auth(), api.DelFromGroup)
	Router.PATCH("/group/update/state/:search_name/:state", middleware.Auth(), api.ChangeGroupState)
}
