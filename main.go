package main

import (
	"trawlcode/database"
	"trawlcode/routes"
)

func main() {
	database.SetupDB()
	defer database.Close()
	// db.AutoMigrate(&models.User{})
	r := routes.SetupRoute()
	r.Run(":1234")
}
