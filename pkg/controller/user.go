package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"heltra_gmo/pkg/model"
	"heltra_gmo/pkg/view"
	"io/ioutil"
	"log"
)

func Register(c *gin.Context) {
	var req model.Register
	c.Set("responded", "responded")
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

	thingsWriteToBody := model.Login{
		UserID:   req.UserID,
		Password: req.Password,
	}
	byteBody, _ := json.Marshal(thingsWriteToBody)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(byteBody))
	c.Set("responded", nil)

	return
}

func GetUser(c *gin.Context) {
	u := model.User{
		UserID: c.Query("userid"),
	}
	if err := u.Read(); err != nil {
		view.BadRequest(c, "Fail", nil)
		return
	}
	view.StatusOK(c, "OK", u)
}

func GetMyself(c *gin.Context) {
	id, _ := c.Get("id")
	u := id.(*model.User)
	log.Println(u)
	if err := u.ReadByID(); err != nil {
		log.Println(err)
		view.BadRequest(c, "Fail", nil)
		return
	}
	view.StatusOK(c, "OK", u)
}
