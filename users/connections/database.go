package connections

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {
	URL := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_HOST"), "disable")
	db, err := sql.Open("postgres", URL)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	return db
}
func TestConnection() {
	con := Connect()
	defer con.Close()
	err := con.Ping()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Database connected successfully")
}
