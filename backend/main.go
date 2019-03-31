package main

import (
	"backend/api"
	"backend/gen"
	"backend/api/models"
)

func main() {

	db := models.Connect()
	db.Debug().DropTableIfExists(&models.Wallet{}, &models.Owner{}, &models.Log{})
	db.Debug().AutoMigrate(&models.Owner{}, &models.Wallet{}, &models.Log{}) 
	db.Close()
	gen.GenerateData()

	api.Run()
}