package main

import (
	"heltra_gmo/pkg/model"
	"heltra_gmo/pkg/router"
	"time"
)

func main() {
	// DBマイグレーション
	// model.Connectionがエラー発生しなくなるまで=DBが立ち上がるまで待機
	// (docker composeで立ち上げると必ずdbのほうが立ち上がり遅い)
	_, dbConErr := model.Connection()
	for dbConErr != nil {
		time.Sleep(time.Second)
		_, dbConErr = model.Connection()
	}
	if err := model.Migration(); err != nil {
		panic(err)
	}

	r := router.Router()
	_ = r.Run()
}
