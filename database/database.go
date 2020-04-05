package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root@(127.0.0.1)/jide")
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
}
