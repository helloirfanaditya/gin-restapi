package database

import (
	"fmt"
	"trawlcode/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func SetupDB() *gorm.DB {
	U := "root"
	PS := ""
	H := "localhost"
	PORT := "3306"
	DBNAME := "go_test"

	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", U, PS, H, PORT, DBNAME)

	db, err := gorm.Open("mysql", URL)
	if err != nil {
		defer db.Close()
		utils.ResError(500, err.Error())
	}
	return db
}
