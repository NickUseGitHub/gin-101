package modules

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	bookEntity "github.com/NickUseGitHub/gin-101/src/modules/book"
)

var DB *gorm.DB

func ConnectDatabase() *gorm.DB {
	dsn := "host=localhost user=go_db password=go_db dbname=go_db port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	database, err := gorm.Open("postgres", dsn)

	if err != nil {
		panic("Failed to connect to database!")
	}

	database.AutoMigrate(&bookEntity.Book{})

	return database
}
