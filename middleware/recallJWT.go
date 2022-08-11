package middleware

import (
	"github.com/gin-gonic/gin"
)

func RecallJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		a, _ := c.Get("responded")
		if a == nil {
			AuthMiddleware.LoginHandler(c)
		}
	}
}
