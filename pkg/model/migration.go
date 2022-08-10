package model

import (
	"log"
)

func Migration() error {
	db, _ := Connection()
	err := db.AutoMigrate(User{})
	if err != nil {
		return err
	}
	log.Println("Migration successfully finished!!")
	return nil
}
