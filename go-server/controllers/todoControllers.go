package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-test-app/models"
	"github.com/gorilla/mux"
)

func checkError(w http.ResponseWriter, err error, status int) {
	if err != nil {
		if status == 0 {
			status = http.StatusBadRequest
		}
		http.Error(w, err.Error(), status)
	}
}

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var (
		result []models.Todo
		err    error
	)
	if params["id"] != "" {
		todoID, idError := strconv.Atoi(params["id"])
		checkError(w, idError, http.StatusBadRequest)
		result, err = models.GetTodoById(todoID)
	} else {
		result, err = models.GetAllTodos()
	}
	checkError(w, err, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(result)
}

func Add(w http.ResponseWriter, r *http.Request) {
	var newTodo models.Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	checkError(w, err, http.StatusBadRequest)
	_, err = models.InsertTodo(newTodo)
	checkError(w, err, http.StatusInternalServerError)
	allTodos, err := models.GetAllTodos()
	checkError(w, err, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(allTodos)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	todoID, err := strconv.Atoi(params["id"])
	checkError(w, err, http.StatusBadRequest)
	_, err = models.DeleteTodo(todoID)
	checkError(w, err, http.StatusInternalServerError)
	allTodos, err := models.GetAllTodos()
	checkError(w, err, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(allTodos)
}

func Update(w http.ResponseWriter, r *http.Request) {
	var updates map[string]interface{}
	params := mux.Vars(r)
	todoID, err := strconv.Atoi(params["id"])
	checkError(w, err, http.StatusBadRequest)
	err = json.NewDecoder(r.Body).Decode(&updates)
	checkError(w, err, http.StatusBadRequest)
	_, err = models.UpdateTodo(updates, todoID)
	checkError(w, err, http.StatusInternalServerError)
	allTodos, err := models.GetAllTodos()
	checkError(w, err, http.StatusInternalServerError)
	json.NewEncoder(w).Encode(allTodos)
}
