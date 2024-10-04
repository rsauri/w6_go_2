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
var nextId int = 1

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

// POST Create a new pantry item
func createPantryItem(w http.ResponseWriter, r *http.Request) {
	// Validate if the HTTP Method is correct
	if r.Method != http.MethodPost {
		http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	//Initialize Variable
	var item PantryItem

	//Get the item information from the request body and insert it to the list of items
	json.NewDecoder(r.Body).Decode(&item)

	//Add the new item to the Pantry Item List
	item.ID = nextId
	pantryItems = append(pantryItems, item)

	//Increment the Id for the next Pantry Item
	nextId++

	//Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)

}

func main() {

	//Handle the following /pantryItems methods
	//     GET: getItems
	http.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		case http.MethodPost:
			createPantryItem(w, r)
		default:
			http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server is running on port 8081...")
	log.Fatal(http.ListenAndServe(":8081", nil))

}
