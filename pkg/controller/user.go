package controller

import (
	"github.com/gin-gonic/gin"
	"heltra_gmo/pkg/model"
	"heltra_gmo/pkg/view"
	"log"
)

func Register(c *gin.Context) {
	var req model.Register
	if err := c.ShouldBindJSON(&req); err != nil {
		view.BadRequest(c, "Body is not complete!", nil)
		return
	}

	u := model.User{
		UserID:   req.UserID,
		Password: EncodePassword(req.Password),
	}
	if err := u.Create(); err != nil {
		log.Println(err)
		view.BadRequest(c, "SQL error", nil)
		return
	}
	view.StatusOK(c, "OK", u)
}

func GetUser(c *gin.Context) {
	u := model.User{
		UserID: c.Query("userID"),
	}
	if err := u.Read(); err != nil {
		view.BadRequest(c, "Fail", nil)
		return
	}
	view.StatusOK(c, "OK", u)

}