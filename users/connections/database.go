package connections

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var db *gorm.DB

func Connect() {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), "disable")
	conn, err := gorm.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
	}
	db = conn
	log.Println("Database connected successfully")

}
func GetDB() *gorm.DB {
	return db
}
