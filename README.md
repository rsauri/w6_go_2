# Pantry Item List API

## Summary
This API facilitates efficient management of household pantry items, allowing users to add, update, retrieve, and delete items while keeping track of expiration dates and restocking needs.

#### Key Features:
- **Expiration Tracking**: Automatically detects expired items based on their expiration date.
- **Restocking Alerts**: Flags items that need to be restocked when the count falls below **1 or less**.
- **Duplicate Handling**: Prevents duplicate entries for pantry items by checking for existing names.
- **Filtering**: Supports querying and filtering of pantry items based on attributes like name, type, expiration status, and restocking needs.

The API supports the following operations:
- **`POST`** `/pantryItems` 
- **`DELETE`** `/pantryItems` 
- **`GET`** `/pantryItems?name={name}&description={descr}&itemType={itemType}&isExpired={true|false}&buy={true|false}` 
- **`GET`** `/pantryItem/{id}`
- **`PATCH`** `/pantryItem/{id}`
- **`DELETE`** `/pantryItem/{id}`

## Functions

### Create Pantry Items
#### Overview
This method creates new pantry items. It accepts multiple items in the request, but if any duplicate item names are detected, the entire request will be rejected and no items will be processed.
- **Method**: POST
- **Path**: `/pantryItems`

#### Path Parameters
*No path parameter for this operation.*

#### Query Parameters
*No query parameter for this operation.*

#### Request Body
- **Content-Type**: application/json
- **Sample Request**

```json
[
{
  "name": "Toothpaste",
  "description": "Toothpaste",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05"
},
{
  "name": "Shampoo",
  "description": "Shampoo",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05"
}
]
```

#### Sample `curl` Request
```bash
curl -X POST http://localhost:8080/pantryItems \
-H "Content-Type: application/json" \
-d '[
{
  "name": "Toothpaste",
  "description": "Toothpaste",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05"
},
{
  "name": "Shampoo",
  "description": "Shampoo",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05"
}
]'
```

#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
[
{
  "id": 1,
  "name": "Toothpaste",
  "description": "Toothpaste",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05",
  "isExpired": false,
  "buy": true
},
{
  "id": 2,
  "name": "Shampoo",
  "description": "Shampoo",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05",
  "isExpired": false,
  "buy": true
}
]
```
---

### Get a Pantry Item by ID
#### Overview

This method will retrieve a pantry item by its ID.
- **Method**: GET
- **Path**: `/pantryItem/{id}`

#### Path Parameters
- **id** *(int)*: The pantry item ID.

#### Query Parameters
*No query parameter for this operation.*

#### Request Body
*No request body for this operation.*

#### Sample `curl` Request
```bash
curl -X GET http://localhost:8080/pantryItem/1
```
#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
{
  "id": 1,
  "name": "Toothpaste",
  "description": "Toothpaste",
  "itemType": "Personal Care",
  "count": 1,
  "expiryDate": "2024-10-05",
  "isExpired": false,
  "buy": true
}

```
---

### Get All Pantry Items
#### Overview

This method retrieves all pantry items and supports optional filters through query parameters.
- **Method**: GET
- **Path**: `/pantryItems`

#### Path Parameters
*No path parameter for this operation.*

#### Query Parameters
- **name** *(string)*: Filter by item name.
- **description** *(string)*: Filter by item description.
- **itemType** *(string)*: Filter by item type.
- **isExpired** *(boolean)*: Filter by expiration status (`true`/`false`).
- **buy** *(boolean)*: Filter by purchase status (`true`/`false`).

#### Request Body
*No request body for this operation.*

#### Sample `curl` Request
```bash
curl -X GET "http://localhost:8080/pantryItems?name=Toothpaste&description=Minty&itemType=Personal%20Care&isExpired=false&buy=false"
```
#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
[
  {
    "id": 1,
    "name": "Toothpaste",
    "description": "Toothpaste",
    "itemType": "Personal Care",
    "count": 1,
    "expiryDate": "2024-10-05",
    "isExpired": false,
    "buy": true
  },
  {
    "id": 2,
    "name": "Carrot",
    "description": "Carrot",
    "itemType": "Produce",
    "count": 3,
    "expiryDate": "2024-12-15",
    "isExpired": false,
    "buy": false
  }
]
```

---

### Update a Pantry Item
#### Overview

This method updates the details of an existing pantry item. If a duplicate item name is detected, the update will be rejected, and the request will not be processed.

- **Method**: PATCH
- **Path**: `/pantryItem/{id}`

#### Path Parameters
- **id** *(int)*: The pantry item ID.

#### Query Parameters
*No query parameter for this operation.*

#### Request Body
- **Content-Type**: application/json
- **Sample Request**
```json
{
  "name": "Toothpaste",
  "description": "Toothpaste",
  "itemType": "Personal Care",
  "count": 2,
  "expiryDate": "2025-01-10"
}
```

#### Sample `curl` Request
```bash
curl -X PATCH http://localhost:8080/pantryItem/1 \
-H "Content-Type: application/json" \
-d '{
    "name": "Toothpaste",
    "description": "Minty fresh toothpaste",
    "itemType": "Personal Care",
    "count": 2,
    "expiryDate": "2025-01-10"
}'
```

#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
{
  "id": 1,
  "name": "Toothpaste",
  "description": "Minty fresh toothpaste",
  "itemType": "Personal Care",
  "count": 2,
  "expiryDate": "2025-01-10",
  "isExpired": false,
  "buy": false
}
```
---

### Delete a Pantry Item
#### Overview

This method deletes a pantry item by its ID.
- **Method**: DELETE
- **Path**: `/pantryItem/{id}`

#### Path Parameters
- **id** *(int)*: The pantry item ID.

#### Query Parameters
*No query parameter for this operation.*

#### Request Body
*No request body for this operation.*

#### Sample `curl` Request
```bash
curl -X DELETE http://localhost:8080/pantryItem/1
```

#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
null
```
---

### Delete All Pantry Items
#### Overview

This method deletes all pantry items.
- **Method**: DELETE
- **Path**: `/pantryItems`

#### Path Parameters
*No path parameter for this operation.*

#### Query Parameters
*No query parameter for this operation.*

#### Request Body
*No request body for this operation.*

#### Sample `curl` Request
```bash
curl -X DELETE http://localhost:8080/pantryItems
```
#### Response Body
- **Content-Type**: application/json
- **Sample Response**
```json
null
```
---

## Conclusion
_You'll know it when you see it :-)_

