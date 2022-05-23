package middleware

import (
	"fmt"
	"jojogo/server/jwt"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	"strings"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Auth Auth
func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 取得token
		log.Info("Get in middleware auth")
		token, ok := getToken(c)
		fmt.Println(token)
		if !ok {
			log.Error("Auth: Unauthorized")
			template.UnauthorityError(c, template.ErrUnauthorizedCode, "Auth: Unauthorized")
			c.Abort()
			return
		}

		// 解析token 取得會員的資料
		userID, userName, err := jwt.ParseToken(token)
		fmt.Println(userID, userName)
		if err != nil || userID == "" || userName == "" {
			template.UnauthorityError(c, template.ErrUnauthorizedCode, "jwt parse error")
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
	log.Info("Get Authorization Bearer token")
	authValue := c.GetHeader("Authorization")
	fmt.Println("authValue", authValue)
	arr := strings.Split(authValue, " ")
	if len(arr) != 2 {
		return "", false
	}
	fmt.Println(arr)
	authType := strings.Trim(arr[0], "\n\r\t")
	fmt.Println(authType)
	if !strings.EqualFold(strings.ToLower(authType), strings.ToLower("Bearer")) {
		return "", false
	}
	return strings.Trim(arr[1], "\n\t\r"), true
}
