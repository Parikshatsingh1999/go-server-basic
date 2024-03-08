package main

import (
	"net/http"

	"github.com/go-test-app/models"
	"github.com/go-test-app/router"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func init() {
	models.ConnectDB()
}

func main() {
	newRouter := mux.NewRouter()
	router.InitRouter(newRouter)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "PATCH"})
	origins := handlers.AllowedOrigins([]string{"*"})

	http.Handle("/", newRouter)
	http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(newRouter))
}
