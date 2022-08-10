package controller

import (
	"golang.org/x/crypto/bcrypt"
	"heltra_gmo/pkg/model"
)

func EncodePassword(password string) string {
	ep, _ := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(ep)
}

func ComparePassword(userID, password string) bool {
	user := model.User{
		UserID: userID,
	}
	if err := user.Read(); err != nil {
		panic(err)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return false
	}
	return true

}
