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
	// Your code here
}

func ToDoListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Your code here
}
