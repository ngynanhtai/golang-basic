package middleware

import (
	"github.com/gin-gonic/gin"
	"go-demo/common"
	"net/http"
)

func Recovery() func(c *gin.Context) {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				if err, ok := r.(error); ok {
					c.AbortWithStatusJSON(http.StatusInternalServerError, common.ErrInternal(err))
				}
				panic(r)
			}
		}()

		c.Next()
	}
}
