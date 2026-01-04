package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	"todo-backend/internal/server"
	"todo-backend/internal/todo"
)

func main() {
	// Database connection
	db, err := sql.Open("postgres", "postgres://todo:todo@db:5432/todo?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Dependency injection
	repo := todo.NewRepository(db)
	service := todo.NewService(repo)
	handler := todo.NewHandler(service)

	// HTTP server
	router := server.NewRouter(handler)

	log.Println("Server running on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
