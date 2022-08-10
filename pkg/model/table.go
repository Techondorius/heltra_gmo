package model

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"intID"`
	UserID   string `gorm:"not null;unique" json:"userID"`
	Password string `gorm:"not null" json:"-"`
}
