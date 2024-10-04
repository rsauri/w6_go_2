package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
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

// Functions
// Function Name: getPantryItem
// Description: Returns the Pantry Item based on the ID in the URL path
// Returns: Item Index, Pantry Item Object, Error Object
func getPantryItem(path string) (int, PantryItem, error) {
	var item PantryItem

	//Validate the URI
	//The URI should be /
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, item, fmt.Errorf("invalid Path")
	}

	//Get the from the URI
	iId, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, item, fmt.Errorf("unexpected Pantry Item Id data type. Expecting integer")
	}

	//Find the Pantry Item by Id
	for i, item := range pantryItems {
		if item.ID == iId {
			return i, item, nil
		}
	}

	return 0, item, fmt.Errorf("pantry Item Id not found")
}

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

// GET Pantry Item by Id
func getPantryItemById(w http.ResponseWriter, r *http.Request) {

	// Validate if the HTTP Method is correct
	if r.Method != http.MethodGet {
		http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	// Get the pantry item
	_, item, err := getPantryItem(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(item)
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
	//     GET:  getItems
	//     POST: createPantryItem
	http.HandleFunc("/pantryItems", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		case http.MethodPost:
			createPantryItem(w, r)
		default:
			http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})

	//Handle the following /pantryItem/{id} methods
	//     GET:  getItem
	http.HandleFunc("/pantryItem/{id}", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getPantryItemById(w, r)
		default:
			http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
