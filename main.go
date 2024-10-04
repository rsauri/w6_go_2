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
func parseItemId(path string) (int, error) {

	//Validate the URI
	//The URI should be /
	parts := strings.Split(path, "/")
	if len(parts) < 3 {
		return 0, fmt.Errorf("invalid Path")
	}

	//Get the ID from the URI
	i, err := strconv.Atoi(parts[2])
	if err != nil {
		return 0, fmt.Errorf("unexpected Pantry Item Id data type. Expecting integer")
	}

	return i, nil
}

// Function Name: getPantryItem
// Description: Returns the Pantry Item based on the ID in the URL path
// Returns: Item Index, Pantry Item Object, Error Object
func getPantryItem(path string) (int, PantryItem, error) {

	//Initialize Variables
	var item PantryItem

	//Get the ID from the URI
	iId, err := parseItemId(path)
	if err != nil {
		return 0, item, err
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

// PATCH Update the Pantry Item attributes that are updatable
func updatePantryItem(w http.ResponseWriter, r *http.Request) {
	// Validate if the HTTP Method is correct
	if r.Method != http.MethodPatch {
		http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	// Get the Pantry Item item
	i, item, err := getPantryItem(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Get the Pantry Item details from the request body
	var rItem PantryItem
	json.NewDecoder(r.Body).Decode(&rItem)

	//Update the Pantry Item in the Pantry Item List
	//The following attributes are not updatable:
	//     ID
	//     Buy
	rItem.ID = item.ID
	rItem.Buy = item.Buy
	pantryItems[i] = rItem

	//Return the updated item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rItem)

}

// DELETE Delete the Pantry Item By Id
func deletePantryItem(w http.ResponseWriter, r *http.Request) {

	// Validate if the HTTP Method is correct
	if r.Method != http.MethodDelete {
		http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		return
	}

	//Get the ID from the Request URI
	id, err := parseItemId(r.URL.Path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Delete the Pantry Item
	for i, item := range pantryItems {
		if item.ID == id {
			pantryItems = append(pantryItems[:i], pantryItems[i+1:]...)
			w.WriteHeader(http.StatusOK)
			return
		}
	}

	http.Error(w, "Pantry Item Id not found", http.StatusNotFound)

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
		case http.MethodPatch:
			updatePantryItem(w, r)
		case http.MethodDelete:
			deletePantryItem(w, r)
		default:
			http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})

	// Start the server
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))

}
