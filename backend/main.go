package main

import "net/http"


// API data has to be in json format as per openAPI spec
// used this for reference: https://go.dev/doc/tutorial/web-service-gin#design_endpoints
type Todo struct {
    Title  string  `json:"title"`
    Description string  `json:"description"`
}

var todolist []Todo // basically a python list(global var)

func main() {
	// this sends all requests to the handler function (reference: https://go.dev/doc/articles/wiki/#tmp_3)
	http.HandleFunc("/", ToDoListHandler)
	http.ListenAndServe(":8080", nil)
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	
	// main reference for json encoding/decoding: https://medium.com/@amanlalwani0807/implement-rest-api-in-golang-using-net-http-deadfd527452
	if r.Method == "GET" {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	} 
	else if r.Method == "POST" {
		var newTodo Todo
		json.NewDecoder(r.Body).Decode(&newTodo)
		todos = append(todos, newTodo)	

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(newTodo)
	}
	else {
		// add error handling for invalid methods
	}
}
