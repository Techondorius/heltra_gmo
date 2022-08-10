package model

type Register struct {
	UserID   string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type Login struct {
	UserID   string `json:"userid" binding:"required"`
	Password string `json:"password" binding:"required"`
}
