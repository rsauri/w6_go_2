package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Create new struct
type PantryItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemType    string `json:"itemType"`
	Count       int    `json:"count"`
	ExpiryDate  string `json:"expiryDate"`
	Buy         bool   `json:"buy"`
}

// Initialize Variables
var pantryItems []PantryItem

// Method Handlers
// GET All Pantry Items
func getItems(w http.ResponseWriter, r *http.Request) {

	//Validate if the HTTP Method is correct
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	//Respond with all items in the Pantry Item List
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(pantryItems)
}
func main() {

	//Handle the following /pantryItems methods
	//     GET: getItems
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
