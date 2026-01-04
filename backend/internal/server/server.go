package server

import "net/http"

func NewRouter(todoHandler http.Handler) http.Handler {
	mux := http.NewServeMux()

	mux.Handle("/tasks", todoHandler)

	return mux
}
