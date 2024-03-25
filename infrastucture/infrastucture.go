package infrastructure

import (
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDatabase : intializes and returns mysql db
func NewDatabase() *gorm.DB {
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
    return db;
}

type Chat struct {
	gorm.Model
	Name  string
	Users []User `gorm:"many2many:chat_users;"`
}

type User struct {
	gorm.Model
	Nickname string
	Chats    []Chat `gorm:"many2many:chat_users;"`
}

type Message struct {
	gorm.Model
	ChatId uint
	UserId uint
	Text   string
}

func Migrate() {
	db := NewDatabase()
	db.AutoMigrate(&Chat{}, &User{}, &Message{})
}
