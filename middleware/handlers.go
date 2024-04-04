package middleware

import (
	"GO_STOCK_API_POSTGRES/models"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

// Creating connection to the DB server
func CreateConnection() *sql.DB {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("POSTGRES_URL"))

	if err != nil {
		panic(err)
	}

	err = db.Ping()

	if err != nil {
		panic(err)
	}

	// If no error returned
	fmt.Println("Connection to postgres db established")
	return db
}

func CreateStock(w http.ResponseWriter, r *http.Request) {
	var stock models.Stock

	err := json.NewDecoder(r.Body).Decode(&stock)

	if err != nil {
		log.Fatalf("Failed to decode body: %v", err)
	}

	insertID := insertStock(stock)

	res := response{
		ID: insertID,
		Message: "stock created successfully",
	}

	json.NewDecoder(w).Encode(res)
}

func GetStock() {

}

func GetAllStock() {

}

func UpdateStock() {

}

func DeleteStock() {

}
