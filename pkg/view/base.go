package view

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func StatusOK(c *gin.Context, message string, detail any) {
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"detail":  detail,
		"message": message,
	})
}

func BadRequest(c *gin.Context, message string, detail any) {
	c.JSON(http.StatusBadRequest, gin.H{
		"code":    200,
		"detail":  detail,
		"message": message,
	})
}
