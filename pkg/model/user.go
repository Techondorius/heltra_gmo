package model

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
