package main

import (
	"bytes"
	"encoding/json"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"heltra_gmo/docker/dev_app/pkg/model/database"
	"heltra_gmo/middlware"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// DBマイグレーション
	// model.Connectionがエラー発生しなくなるまで=DBが立ち上がるまで待機
	// (docker composeで立ち上げると必ずdbのほうが立ち上がり遅い)
	_, dbConErr := database.Connection()
	for dbConErr != nil {
		time.Sleep(time.Second)
		_, dbConErr = database.Connection()
	}
	if err := database.Migration(); err != nil {
		panic(err)
	}

	r := gin.Default()

	r.Use(logger())

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

	var GinJWT = middlware.AuthMiddleware

	r.Use(GinJWT.MiddlewareFunc())

	// When you use jwt.New(), the function is already automatically called for checking,
	// which means you don't need to call it again.
	errInit := GinJWT.MiddlewareInit()

	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}

	_ = r.Run()
}

func logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("Endpoint: " + c.Request.URL.Path)
		q := c.Request.URL.Query()
		j, _ := json.Marshal(q)
		log.Println("Query Params: " + string(j))

		ByteBody, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(ByteBody))
		log.Println("Body: " + string(ByteBody))

		c.Next()
	}
}
