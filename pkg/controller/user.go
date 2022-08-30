package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"heltra_gmo/pkg/model"
	"heltra_gmo/pkg/model/dao"
	"heltra_gmo/pkg/view"
	"io/ioutil"
	"log"
)

func Register(c *gin.Context) {
	var req dao.RegisterReq
	c.Set("responded", "responded")
	if err := c.ShouldBindJSON(&req); err != nil {
		view.BadRequest(c, "Body is not complete!", nil)
		return
	}
	thingsWriteToBody := dao.LoginReq{
		UserID:   req.UserID,
		Password: req.Password,
	}
	req.Password = EncodePassword(req.Password)

	u := model.User{}
	req.RegisterToUser(&u)
	if err := u.Create(); err != nil {
		log.Println(err)
		view.BadRequest(c, "SQL error", nil)
		return
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
