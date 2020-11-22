package db

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

func InitDB() (*gorm.DB, error) {
	dsn := "root:????@tcp(127.0.0.1:3306)/user?charset=utf8mb4&parseTime=True&loc=UTC"

	database, err := gorm.Open("mysql", dsn)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	database.LogMode(true)

	return database, nil
}
