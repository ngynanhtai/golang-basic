package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"go-demo/biz"
	"go-demo/initializers"
	"go-demo/storage"
	"net/http"
	"os"
	"time"
)

func RequireAuth() func(c *gin.Context) {
	return func(c *gin.Context) {
		// GET cookie of request
		tokenJwt, err := c.Cookie("Authorization")
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		// Decode JWT and validate
		token, err := jwt.Parse(tokenJwt, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Check the exp
			if float64(time.Now().Unix()) > claims["exp"].(float64) {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Find user with token sub
			store := storage.NewSQLStore(initializers.DB)
			business := biz.NewUserBiz(store)

			data, err := business.GetUserById(c.Request.Context(), int(claims["sub"].(float64)))
			if err != nil || data == nil {
				c.AbortWithStatus(http.StatusUnauthorized)
				return
			}

			// Attach to request
			c.Set("user", data)

			// Continue
			c.Next()
		} else {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	}
}
