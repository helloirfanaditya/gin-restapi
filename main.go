package main

import (
	"trawlcode/database"
	"trawlcode/models"
	"trawlcode/routes"
)

func main() {
	db := database.SetupDB()
	db.AutoMigrate(&models.User{})
	r := routes.SetupRoute(db)
	r.Run(":1234")
}
