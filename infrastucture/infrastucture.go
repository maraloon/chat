package infrastructure

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func LoadEnv() {
	err := godotenv.Load("../.env")
	// err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load .env file")
	}
}

// NewDatabase : intializes and returns mysql db
func NewDatabase() *gorm.DB {
    // TODO: глупый пакет по загрузке env, потом найти альтернативу
    // а пока хардкодим значения

	// LoadEnv()
	// USER := os.Getenv("MYSQL_USER")
	// PASS := os.Getenv("MYSQL_PASSWORD")
	// HOST := os.Getenv("DB_HOST")
	// DBNAME := os.Getenv("MYSQL_DATABASE")

	USER := "user"
	PASS := "chatpass"
	HOST := "db"
	DBNAME := "chat"

	URL := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		USER, PASS, HOST, DBNAME)

	fmt.Println(URL)
	db, err := gorm.Open(mysql.Open(URL), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(fmt.Sprintf(
			"Failed to connect to database: %s", err))

	}
	fmt.Println("Database connection established")
	return db
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
