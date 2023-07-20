package main

import (
	"trawlcode/database"
	"trawlcode/routes"
)

func main() {
	database.SetupDB()
	defer database.Close()
	r := routes.SetupRoute()
	r.Run(":1234")
}
