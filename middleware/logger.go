package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

func Logger() gin.HandlerFunc {
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
