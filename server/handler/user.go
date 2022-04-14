package handler

import (
	"fmt"
	"jojogo/server/template"
	"jojogo/server/utils/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetUserInfo GetUserInfo
func GetUserInfo(c *gin.Context) {
	id, ok := c.Get("user_id")
	if !ok {
		c.JSON(http.StatusUnauthorized, gin.H{})
		return
	}

	user, err := user.FindUserByID(id.(string))
	if err != nil {
		template.SystemError(c, template.ErrSystemCode, fmt.Sprintf("%v\n", err))
		return
	}
	template.Success(c, user)
}
