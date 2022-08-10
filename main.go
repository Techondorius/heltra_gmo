package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"heltra_gmo/middleware"
	"heltra_gmo/pkg/controller"
	"heltra_gmo/pkg/model"
	"log"
	"net/http"
	"time"
)

func main() {
	// DBマイグレーション
	// model.Connectionがエラー発生しなくなるまで=DBが立ち上がるまで待機
	// (docker composeで立ち上げると必ずdbのほうが立ち上がり遅い)
	_, dbConErr := model.Connection()
	for dbConErr != nil {
		time.Sleep(time.Second)
		_, dbConErr = model.Connection()
	}
	if err := model.Migration(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(middleware.Logger())

	// Cors
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))

	// ルーティング
	//routing.Routing(r)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hi",
		})
	})
	r.POST("/register", controller.Register)
	r.GET("/getUser", controller.GetUser)

	var GinJWT = middleware.AuthMiddleware

	r.Use(GinJWT.MiddlewareFunc())

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := GinJWT.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	_ = r.Run()
}
