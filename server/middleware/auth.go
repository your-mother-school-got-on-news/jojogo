package middleware

import (
	"jojogo/server/jwt"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Auth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		token, ok := getToken(c)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": template.ErrUnauthorizedCode,
			})
			log.Error("Auth: Unauthorized")
			c.Abort()
			return
		}

		// 解析token 取得會員的資料
		userID, userName, err := jwt.ParseToken(token)
		if err != nil || userID == "" || userName == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"result":     false,
				"error_code": template.ErrUnauthorizedCode,
			})
			log.Error("jwt parse error", zap.Any("Error", err))
			c.Abort()
			return
		}

		// 把值傳到下一層
		c.Set("user_id", userID)
		c.Set("user_name", userName)
		c.Writer.Header().Set("Authorization", "Bearer "+token)
		c.Next()
	}
}

func getToken(c *gin.Context) (string, bool) {
	authValue := c.GetHeader("Authorization")
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	authType := strings.Trim(arr[0], "\n\r\t")
	if !strings.EqualFold(strings.ToLower(authType), strings.ToLower("Bearer")) {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
