package database

import (
	"heltra_gmo/docker/dev_app/pkg/model"
	"log"
)

func Migration() error {
	db, _ := Connection()
	err := db.AutoMigrate(model.User{})
	if err != nil {
		return err
	}
	log.Println("Migration successfully finished!!")
	return nil
}
