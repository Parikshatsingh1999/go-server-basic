package router

import (
	"encoding/json"
	"net/http"

	"github.com/go-test-app/controllers"
	"github.com/go-test-app/logs"
	"github.com/go-test-app/models"
	"github.com/gorilla/mux"
)

type SingleRoutes struct {
	route   string
	handler func(w http.ResponseWriter, r *http.Request)
	method  string
}

// var todos []models.Todo = []models.Todo{
// 	{IsDone: false, Name: "Task 1", Description: "Description for task 1"},
// 	{IsDone: true, Name: "Task 2", Description: "Description for task 2"},
// }

func handleRootRoute(w http.ResponseWriter, r *http.Request) {

	data := map[string]string{
		"message": "Welcome to Go server",
	}

	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.Write(jsonData)
}

// func handleGetTodos(w http.ResponseWriter, r *http.Request) {

// 	w.Header().Set("Content-Type", "application/json")
// 	jsonData, err := json.Marshal(todos)
// 	if err != nil {
// 		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
// 		return
// 	}
// 	//need more control over the encoding process or want to manipulate the JSON data before sending it
// 	w.Write(jsonData)
// }

// func addTodo(w http.ResponseWriter, r *http.Request) {
// 	var newTodo models.Todo
// 	err := json.NewDecoder(r.Body).Decode(&newTodo)
// 	if err != nil {
// 		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 		return
// 	}
// 	todo, err := controllers.AddTodo(newTodo)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}
// 	if todo.ID != 0 {
// 		allTodos, _ := controllers.GetAllTodos()
// 		json.NewEncoder(w).Encode(allTodos)
// 	}
// }

// // func addTodo(w http.ResponseWriter, r *http.Request) {
// // 	// Parse the JSON request body into a Todo struct
// // 	var newTodo Todo
// // 	err := json.NewDecoder(r.Body).Decode(&newTodo)
// // 	if err != nil {
// // 		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// // 		return
// // 	}
// // 	if newTodo.Name != "" && newTodo.Description != "" {
// // 		id := 1
// // 		if len(todos) != 0 {
// // 			id = todos[len(todos)-1].Id + 1
// // 		}
// // 		newTodo.Id = id
// // 		todos = append(todos, newTodo)
// // 	}

// // 	w.Header().Set("Content-Type", "application/json")

// // 	//If you prefer a concise and idiomatic approach, especially for basic use cases
// // 	json.NewEncoder(w).Encode(todos)
// // }

// func deleteTodo(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	todoID, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 	}
// 	for ind, value := range todos {
// 		if int(value.ID) == todoID {
// 			todos = append(todos[:ind], todos[ind+1:]...)
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(todos)
// }

// func updateTodo(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
// 	todoId, err := strconv.Atoi(params["id"])
// 	if err != nil {
// 		http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 	}
// 	for index, value := range todos {
// 		if int(value.ID) == todoId {
// 			var update models.Todo
// 			err := json.NewDecoder(r.Body).Decode(&update)
// 			if err != nil {
// 				http.Error(w, "Error decoding JSON", http.StatusBadRequest)
// 			}
// 			todos[index].IsDone = update.IsDone
// 			break
// 		}
// 	}
// 	json.NewEncoder(w).Encode(todos)
// }

// func testRoute(w http.ResponseWriter, r *http.Request) {
// 	json.NewEncoder(w).Encode(Test)
// }

func middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		message := "Got a request from " + r.Host + " of type " + r.Method + " to route " + r.URL.Path
		go logs.AddLog(message)
		next.ServeHTTP(w, r)
	})
}

var Test []models.Todo

func InitRouter(router *mux.Router) {
	router.Use(middleware)

	routes := []SingleRoutes{
		// {
		// 	route:   "/test",
		// 	handler: testRoute,
		// 	method:  "GET",
		// },
		{
			route:   "/",
			handler: handleRootRoute,
			method:  "GET",
		},
		{
			route:   "/todos",
			handler: controllers.Get,
			method:  "GET",
		},
		{
			route:   "/todos/{id}",
			handler: controllers.Get,
			method:  "GET",
		},
		{
			route:   "/todo",
			handler: controllers.Add,
			method:  "POST",
		},
		{
			route:   "/todos/{id}",
			handler: controllers.Delete,
			method:  "DELETE",
		},
		{
			route:   "/todos/{id}",
			handler: controllers.Update,
			method:  "PATCH",
		},
	}

	for _, value := range routes {
		router.HandleFunc(value.route, value.handler).Methods(value.method)
	}

}
