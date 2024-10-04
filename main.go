package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Start the server
	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
