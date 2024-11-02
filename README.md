# Go To-Do REST API

A simple REST API built using Go's `net/http` package for managing a "To-Do" list, with endpoints to create, read, update, and delete To-Do items.

## Prerequisites
- Install Go (if you haven't already): [Go installation guide](https://golang.org/doc/install)
- Basic knowledge of Go syntax and programming concepts

## Project Setup

### Step 2: Set up and Run the Server
Navigate to the project directory, then run the Go server:

```
go run main.go
```

The server will start on localhost:8080.

API Endpoints and curl Commands
Below are the available endpoints and examples of how to test them using curl.

## 1. Create a New To-Do Item
To create a new To-Do item, send a POST request to /todos with JSON data.

```
curl -X POST -H "Content-Type: application/json" -d '{"title":"Buy groceries"}' http://localhost:8080/todos
```

Expected Response:

```
{
  "id": 1,
  "title": "Buy groceries",
  "status": "pending"
}
```

## 2. Get All To-Do Items
To retrieve a list of all To-Do items, send a GET request to /todos.

```
curl http://localhost:8080/todos
```

Expected Response:

```
[
  {
    "id": 1,
    "title": "Buy groceries",
    "status": "pending"
  }
]
```

## 3. Get a Specific To-Do Item by ID
To retrieve a specific To-Do item by ID, send a GET request to /todos/{id}.

```
curl http://localhost:8080/todos/1
```

Expected Response:
```
{
  "id": 1,
  "title": "Buy groceries",
  "status": "pending"
}
```

## 4. Update a To-Do Item by ID
To update an existing To-Do item, send a PUT request to /todos/{id} with the updated data in JSON format.

```
curl -X PUT -H "Content-Type: application/json" -d '{"title":"Buy groceries and fruits", "status":"completed"}' http://localhost:8080/todos/1
```

Expected Response:
```
{
  "id": 1,
  "title": "Buy groceries and fruits",
  "status": "completed"
}
```

## 5. Delete a To-Do Item by ID
To delete a specific To-Do item, send a DELETE request to /todos/{id}.

```
curl -X DELETE http://localhost:8080/todos/1
Expected Response: The server responds with status 204 No Content if successful.
```