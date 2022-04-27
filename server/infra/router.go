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
	Router.Use(middleware.Cors())

	// login
	Router.POST("/login", api.LoginHandler)
	Router.GET("/info", middleware.Auth(), handler.GetUserInfo)

	// group ops
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	Router.GET("/group/view", api.GetGroups)
	Router.GET("/group/get/name/:group_name", api.GetGroupByName)

	Router.POST("/group/new/:group_name", api.CreateGroup)

	Router.PATCH("/group/update/name/:search_name/:set_name", api.UpdateGroupName)
	Router.PATCH("/group/update/person/add/:search_name/:person_name", api.AddToGroup)
	Router.PATCH("/group/update/person/del/:search_name/:person_name", api.DelFromGroup)
	Router.PATCH("/group/update/state/:search_name/:state", api.ChangeGroupState)
}
