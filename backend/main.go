package main

import (
	"net/http"
	"log"
	"encoding/json"
)


// API data has to be in json format as per openAPI spec
// used this for reference: https://go.dev/doc/tutorial/web-service-gin#design_endpoints
type Todo struct {
	Title string `json:"title"`
	Description string `json:"description"`
}

var todos []Todo // basically a python list(global var)

func main() {
	// this sends all requests to the handler function (reference: https://go.dev/doc/articles/wiki/#tmp_3)
	http.HandleFunc("/", ToDoListHandler)
	log.Println("Running the server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	// main reference for json encoding/decoding: https://medium.com/@amanlalwani0807/implement-rest-api-in-golang-using-net-http-deadfd527452
	// this is my router for handling different methods
	if r.Method == "GET" {
		listTodos(w, r)
	} else if r.Method == "POST" {
		addTodo(w, r)
	} else {
		// add error handling for invalid methods
	}
}

// the get "endpoint"(technically ToDoListHandler is the endpoint and this is just a helper function to handle method requests)
func listTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// the post "endpoint"
func addTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	json.NewDecoder(r.Body).Decode(&newTodo)
	todos = append(todos, newTodo)	

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTodo)
}
