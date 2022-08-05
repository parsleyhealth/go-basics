package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Todo struct {
	Content string `json:"content"`
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Home!")
}

func todosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		w.WriteHeader(http.StatusAccepted)
		json.NewEncoder(w).Encode(todos)
	case "POST":
		bytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			fmt.Println("Failed to parse request body")
		}

		var newTodo Todo
		if err := json.Unmarshal(bytes, &newTodo); err != nil {
			fmt.Println("Error parsing json", err)
		}

		todos = append(todos, newTodo)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(todos)

	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}

var todos = []Todo{
	{Content: "Buy stuff"},
	{Content: "Take out trash"},
	{Content: "Make dinner"},
}

func main() {

	// Route handler
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/todos", todosHandler)

	fmt.Println("Server is running on port 8888")
	// Spin up our server on port 8888
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		fmt.Errorf("Error firing up server: %v", err)
	}
}
