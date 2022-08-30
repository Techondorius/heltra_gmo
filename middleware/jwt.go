package middleware

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"heltra_gmo/pkg/controller"
	"heltra_gmo/pkg/model"
	"heltra_gmo/pkg/model/dao"
	"time"
)

var identityKey = "id"

var AuthMiddleware, _ = jwt.New(&jwt.GinJWTMiddleware{
	Realm:       "asdf",
	Key:         []byte("sec"),
	Timeout:     time.Hour * 5000,
	MaxRefresh:  time.Hour,
	IdentityKey: identityKey,
	PayloadFunc: func(data interface{}) jwt.MapClaims {
		if v, ok := data.(*model.User); ok {
			return jwt.MapClaims{
				identityKey: v.ID,
			}
		}
		return jwt.MapClaims{}
	},
	IdentityHandler: func(c *gin.Context) interface{} {
		claims := jwt.ExtractClaims(c)
		return &model.User{
			ID: int(claims[identityKey].(float64)),
		}
	},
	Authenticator: func(c *gin.Context) (interface{}, error) {
		var loginVals dao.LoginReq
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
