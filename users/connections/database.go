package connections

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func Connect() *gorm.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), "disable")
	db, err := gorm.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	log.Println("Database connected successfully")
	return db
}
