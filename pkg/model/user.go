package model

type User struct {
	ID       int    `gorm:"primaryKey;autoIncrement" json:"intID"`
	UserID   string `gorm:"not null;unique" json:"userID"`
	Password string `gorm:"not null" json:"-"`
}

func (u *User) Create() error {
	db, _ := Connection()
	result := db.Debug().Create(&u)
	return result.Error
}

func (u *User) Read() error {
	db, _ := Connection()
	result := db.Debug().Find(&u)
	return result.Error
}
