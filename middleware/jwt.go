package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"heltra_gmo/pkg/controller"
	"heltra_gmo/pkg/model"
	"log"
	"time"
)

var identityKey = "id"

type user model.User

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "asdf",
	Key:         []byte("sec"),
	Timeout:     time.Hour * 5000,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		log.Println(data)
		if v, ok := data.(*model.User); ok {
			log.Println("here")
			return jwt.MapClaims{
				identityKey: v.UserID,
			}
		} else {
			log.Println(v)
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &user{
			UserID: claims[identityKey].(string),
		}
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var loginVals model.Login
		log.Println("Authenticator called")
		if err := c.ShouldBindJSON(&loginVals); err != nil {
			return "", jwt.ErrMissingLoginValues
		}

		if controller.ComparePassword(loginVals.UserID, loginVals.Password) {
			u := model.User{UserID: loginVals.UserID}
			_ = u.Read()
			return &u, nil
		}

		return nil, jwt.ErrFailedAuthentication
	},
	Unauthorized: func(c *gin.Context, code int, message string) {
		c.JSON(code, gin.H{
			"code":    code,
			"message": message,
		})
	},
	TokenLookup:   "header: Authorization, query: token, cookie: jwt",
	TokenHeadName: "Bearer",
	TimeFunc:      time.Now,
})
