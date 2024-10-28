package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

// Define a structure for the "To-Do" items
type ToDo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Status string `json:"status"`
}

var (
	todoList = []ToDo{}
	mutex = &sync.Mutex{}
	todoCount = 1
)

// Handler for /todos endpoint (GET, POST)
func todosHandler(w http.ResponseWriter, r *http.Request){
	switch r.Method {
	case "GET":
		getToDos(w, r)
	case "POST":
		createToDo(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}


// GET: Retrieve all To-Do items
func getToDos(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	w.Header().Set("Contetn-Type", "application/json")
	json.NewEncoder(w).Encode(todoList)
}

// POST: Create a new To-Do item
func createToDo(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	var newToDo ToDo
	err := json.NewDecoder(r.Body).Decode(&newToDo)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	newToDo.ID = todoCount
	todoCount++
	newToDo.Status = "pending"
	todoList = append(todoList, newToDo)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newToDo)
}


// Handler for /todos/{id} endpoint (GET, PUT, DELETE)
func todoHandler(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/todos/"):])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
	}

	switch r.Method {
	case "GET":
		getToDoByID(w, id)
	case "PUT":
		updateToDoByID(w, r, id)
	case "DELETE":
		deleteToDoByID(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}



// GET: Retrieve a specific To-Do item by ID
func getToDoByID(w http.ResponseWriter, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	for _, todo := range todoList {
		if todo.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	http.Error(w, "To-Do item not found", http.StatusNotFound)
}


// PUT: Update a specific To-Do item by ID
func updateToDoByID(w http.ResponseWriter, r *http.Request, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, todo := range todoList {
		if todo.ID == id {
			var updatedToDo ToDo
			err := json.NewDecoder(r.Body).Decode(&updatedToDo)
			if err != nil {
				http.Error(w, "Invalid request payload", http.StatusBadRequest)
				return
			}

			todoList[i].Title = updatedToDo.Title
			todoList[i].Status = updatedToDo.Status

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(todoList[i])
			return
		}
	}

	http.Error(w, "To-Do item not found", http.StatusNoContent)
}

// DELETE: Delete a specific To-Do item by ID
func deleteToDoByID(w http.ResponseWriter, id int) {
	mutex.Lock()
	defer mutex.Unlock()

	for i, todo := range todoList {
		if todo.ID == id {
			todoList = append(todoList[:i], todoList[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "To-Do item not found", http.StatusNotFound)
}




func main() {
	http.HandleFunc("/todos", todosHandler)
	http.HandleFunc("/todos/", todoHandler)

	fmt.Println("Server is running on porn 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
