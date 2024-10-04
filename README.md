# Pantry Item List API

## Summary
This API enables the management of the pantry items in a household. 

This API has the following operations:
- **`POST`** `/pantryItems` 
- **`GET`** `/pantryItems?name={name}&description={descr}&itemType={itemType}&isExpired={true|false}&buy={true|false}` 
- **`GET`** `/pantryItem/{id}`
- **`PATCH`** `/pantryItem/{id}`
- **`DELETE`** `/pantryItem/{id}`

## Functions

### Create a Pantry Item
#### Overview

This method will create pantry items 
- **Method**: POST
- **Path**: `/pantryItems`

#### Path Parameters
*There's no path parameter for this operation.*

#### Query Parameters
*There's no query parameter for this operation.*

#### Request Body
- **Content-Type**: application/json
- **Sample Request**

>{
><br>"name": "Toothpaste",
><br>    "description": "Toothpaste",
><br>    "itemType": "Personal Care",
><br>    "count": 1,
><br>    "expiryDate": "2024-10-05",
><br>    "buy": false
><br>}

#### Response Body
- **Content-Type**: application/json
- **Sample Response**

>{
><br>    "id": 1,
><br>    "name": "Toothpaste",
><br>    "description": "Toothpaste",
><br>    "itemType": "Personal Care",
><br>    "count": 1,
><br>    "expiryDate": "2024-10-05",
><br>    "isExpired": false,
><br>    "buy": true
>}

---
### Get a Pantry Item
#### Overview

This method will get a pantry item by Id
- **Method**: GET
- **Path**: `/pantryItem/{id}`

#### Path Parameters
- **id** *(int)*: Pantry Item id

#### Query Parameters
*There's no query parameter for this operation.*

#### Request Body
*There's no request body for this operation.*


#### Response Body
- **Content-Type**: application/json
- **Sample Response**

>{
><br>    "id": 1,
><br>    "name": "Toothpaste",
><br>    "description": "Toothpaste",
><br>    "itemType": "Personal Care",
><br>    "count": 1,
><br>    "expiryDate": "2024-10-05",
><br>    "isExpired": false,
><br>    "buy": true
><br>}

---
### Get all Pantry Items
#### Overview

This method will get a pantry item by Id
- **Method**: GET
- **Path**: `/pantryItem/`

#### Path Parameters
*There's no path parameter for this operation.*

#### Query Parameters
- **name** *(int)*: Name
- **description** *(string)*: Description
- **itemType** *(string)*: Item Type
- **isExpired** *(boolean)*: Expired Flag
- **buy** *(boolean)*: Buy flag

*Example*: 
<br>`/pantryItem/?name=test&description=descr`

#### Request Body
*There's no request body for this operation.*


#### Response Body
- **Content-Type**: application/json
- **Sample Response**

>[
><br>{
><br>    "id": 1,
><br>    "name": "Toothpaste",
><br>    "description": "Toothpaste",
><br>    "itemType": "Personal Care",
><br>    "count": 1,
><br>    "expiryDate": "2024-10-05",
><br>    "isExpired": false,
><br>    "buy": true
><br>},
><br>{
><br>    "id": 2,
><br>    "name": "Carrot",
><br>    "description": "Carrot",
><br>    "itemType": "Produce",
><br>    "count": 1,
><br>    "expiryDate": "2024-10-05",
><br>    "isExpired": false,
><br>    "buy": true
><br>},
><br>]

---
*and more....*