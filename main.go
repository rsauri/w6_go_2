package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// Create new struct
type PantryItem struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	ItemType    string `json:"itemType"`
	Count       int    `json:"count"`
	ExpiryDate  string `json:"expiryDate"`
	IsExpired   bool   `json:"isExpired"`
	Buy         bool   `json:"buy"`
}

// Initialize the Global Variables
var pantryItems []PantryItem
var nextId int = 1

// Setters
func (i *PantryItem) SetID() {
	i.ID = nextId
	nextId++
}

func (i *PantryItem) SetIsExpired() bool {
	expiry, _ := time.Parse(time.DateOnly, i.ExpiryDate)
	i.IsExpired = expiry.After(time.Now())
	return i.IsExpired

}

func (i *PantryItem) SetBuy() bool {
	if i.Count <= 1 {
		i.Buy = true
	} else {
		i.Buy = false
	}

	return i.Buy

}

// Functions
// Function Name: parseItemId
// Description: Returns the value in the {id} from the URI /pantryItem/{id}
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
// The Pantry Items can be filtered. The queryable attributes are
//
//	---name
//	---description
//	---itemType
//	---isExpired
//	---buy
func getItems(w http.ResponseWriter, r *http.Request) {

	//Get the query parameters
	q := r.URL.Query()
	name := q.Get("name")
	descr := q.Get("description")
	itemType := q.Get("itemType")

	var isExpired bool
	var err error
	if q.Has("isExpired") {
		isExpired, err = strconv.ParseBool(q.Get("isExpired"))
		if err != nil {
			http.Error(w, "Query Parameter isExpired has an invalid value. Expected Value is true or false", http.StatusBadRequest)
			return
		}
	}

	var buy bool
	if q.Has("buy") {
		buy, err = strconv.ParseBool(q.Get("buy"))
		if err != nil {
			http.Error(w, "Query Parameter buy has an invalid value. Expected Value is true or false", http.StatusBadRequest)
			return
		}
	}

	//Filter the items based on the query parameters
	var items []PantryItem
	for _, item := range pantryItems {
		if strings.Contains(strings.ToLower(item.Name), strings.ToLower(name)) && strings.Contains(strings.ToLower(item.Description), strings.ToLower(descr)) && strings.Contains(strings.ToLower(item.ItemType), strings.ToLower(itemType)) && (item.IsExpired == isExpired || !q.Has("isExpired")) && (item.Buy == buy || !q.Has("buy")) {
			items = append(items, item)
		}
	}

	//Respond with filtered Pantry Item list
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// GET Pantry Item by Id
func getPantryItemById(w http.ResponseWriter, r *http.Request) {

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
// If a duplicate is found, the whole request body will be ignored
func createPantryItem(w http.ResponseWriter, r *http.Request) {

	//Initialize Variable
	var items []PantryItem

	//Get the item information from the request body and insert it to the list of items
	json.NewDecoder(r.Body).Decode(&items)

	//Find if there are duplicates in the request
	for i, item1 := range items {
		for j, item2 := range items {
			if item1.Name == item2.Name && i != j {
				http.Error(w, fmt.Sprintf("Duplicate item %s found", item1.Name), http.StatusNotAcceptable)
				return
			}
		}
	}

	//Find if there are duplicates in the existing item list
	for _, item1 := range items {
		for _, item2 := range pantryItems {
			if item1.Name == item2.Name {
				http.Error(w, fmt.Sprintf("Duplicate item %s found", item1.Name), http.StatusNotAcceptable)
				return
			}
		}
	}

	//Add the new items to the Pantry Item List
	for i := 0; i < len(items); i++ {
		items[i].SetID()
		items[i].SetIsExpired()
		items[i].SetBuy()
		pantryItems = append(pantryItems, items[i])

	}

	//Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&items)

}

// PATCH Update the Pantry Item attributes that are updatable
// The following attributes are not updatable:
//
//	---ID
//	---IsExpired
//	---Buy
func updatePantryItem(w http.ResponseWriter, r *http.Request) {

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
	rItem.ID = item.ID
	rItem.SetIsExpired()
	rItem.SetBuy()
	pantryItems[i] = rItem

	//Return the updated item
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rItem)

}

// DELETE Delete the Pantry Item By Id
func deletePantryItem(w http.ResponseWriter, r *http.Request) {

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

// DELETE Delete all pantry items
func deleteAllPantryItems(w http.ResponseWriter) {
	pantryItems = nil

	//Write the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&pantryItems)
}

func main() {

	//Handle the following /pantryItems methods
	//     GET:    getItems
	//     POST:   createPantryItem
	//     DELETE: deleteAllPantryItems
	http.HandleFunc("/pantryItems", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getItems(w, r)
		case http.MethodPost:
			createPantryItem(w, r)
		case http.MethodDelete:
			deleteAllPantryItems(w)
		default:
			http.Error(w, fmt.Sprintf("Invalid request method %s", r.Method), http.StatusMethodNotAllowed)
		}
	})

	//Handle the following /pantryItem/{id} methods
	//     GET:    getItem
	//     PATCH:  updatePantryItem
	//     DELETE: deletePantryItem
	// If you're reading this, congratulations! If you missed it the first time,
	//the next drink's on you. But if you spotted it, the drink's on me!
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
