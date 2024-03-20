package main

import (
	"renderin-html/initializers"
	"renderin-html/models"
)

func main() {
	db, err := initializers.ConnectToDb()
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&models.Delivery{})
	if err != nil {
		panic(err)
	}
}
