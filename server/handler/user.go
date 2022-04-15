package handler

import (
	"fmt"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	"jojogo/server/utils/user"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// GetUserInfo GetUserInfo
func GetUserInfo(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		log.Error("Cannot get user_id in cookie")
		template.SystemError(c, template.ErrSystemCode, "Cannot get user_id in cookie")
		return
	}

	user, err := user.FindUserByID(id.(string))
	if err != nil {
		log.Error("Cannot get find th id in the db", zap.Any("Error", err))
		template.SystemError(c, template.ErrSystemCode, fmt.Sprintf("%v\n", err))
		return
	}
	template.Success(c, user)
}
