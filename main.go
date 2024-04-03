package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/GO_STOCK_API_POSTGRES/router"
)

func main() {
	r := router.Router()
	fmt.Println("Starting server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
