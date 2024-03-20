package infrastructure

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// Database struct
type Database struct {
	DB *gorm.DB
}

// NewDatabase : intializes and returns mysql db
func NewDatabase() Database {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("DB_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	URL := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, DBNAME)
	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL))

	if err != nil {
		panic("Failed to connect to database!")

	}
	fmt.Println("Database connection established")
	return Database{
		DB: db,
	}

}
