package jwt

import (
	"errors"
	"fmt"
	"jojogo/server/config"
	"jojogo/server/template"
	"jojogo/server/utils/log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.uber.org/zap"
)

// Key cookie key
const Key = "token"

// Claims Token的結構，裡面放你要的資訊s
type Claims struct {
	UserID       string `json:"user_id"`
	UserName     string `json:"user_name"`
	UserPassword string `json:"user_password"`

	jwt.StandardClaims
}

// GenerateToken 產生Token
func GenerateToken(User template.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(config.Val.JWTTokenLife) * time.Second) // Token有效時間

	claims := Claims{
		UserID: User.ID.Hex(),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Subject:   User.Name,
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString([]byte(config.Val.JWTSecret))

	return token, err
}

// ParseToken 驗證Token對不對，如果對就回傳user info
func ParseToken(token string) (userID, userName string, err error) {
	var claims Claims
	tokenClaims, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Val.JWTSecret), nil
	})

	if err != nil {
		return "", "", err
	}
	if !tokenClaims.Valid {
		return "", "", errors.New("invalid token")
	}
	id := claims.UserID
	username := claims.Subject
	log.Info("user, ", zap.String("UserName", username), zap.String("UserID", id))
	return id, username, nil
}
