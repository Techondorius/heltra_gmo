package model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string    `gorm:"not null;unique" json:"userID"`
	Password  string    `gorm:"not null" json:"-"`
	Name      string    `gorm:"not null" json:"name"`
	Birthdate time.Time `gorm:"not null" json:"birthdate"`
	Sex       int       `gorm:"not null;size:2" json:"sex"`
	Height    int       `gorm:"not null" json:"height"`
	Weight    int       `gorm:"not null" json:"weight"`
	Objective int       `gorm:"not null" json:"objective"`
}

func (u *User) Create() error {
	db, _ := Connection()
	result := db.Debug().Create(&u)
	return result.Error
}

func (u *User) Read() error {
	db, _ := Connection()
	result := db.Debug().Where("user_id = ?", u.UserID).First(&u)
	return result.Error
}

func (u *User) ReadByID() error {
	db, _ := Connection()
	//result := db.Debug().Where("id = ?", u.ID).Find(&u)
	result := db.Debug().First(&u)
	return result.Error
}
