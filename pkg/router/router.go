package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"heltra_gmo/middleware"
	"heltra_gmo/pkg/controller"
	"net/http"
)

func Router() *gin.Engine {
	r := gin.Default()

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	r.Use(middleware.Logger())

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hii",
		})
	})
	r.POST("/api/login", middleware.AuthMiddleware.LoginHandler)

	register := r.Group("/api/register")
	register.Use(middleware.RecallJWT())
	register.POST("/", controller.Register)

	auth := r.Group("/api/auth")
	auth.Use(middleware.AuthMiddleware.MiddlewareFunc())
	user := auth.Group("/user")
	{
		user.GET("/getMyself", controller.GetMyself)
		user.GET("/getUser", controller.GetUser)
	}
	return r
}
