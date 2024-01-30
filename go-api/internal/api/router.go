package api

import (
    "go-api/internal/api/handlers"
    "go-api/internal/api/middleware"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.Use(middleware.LoggingMiddleware)
    router.HandleFunc("/health", handlers.GetHeatlh).Methods("GET")
    router.HandleFunc("/", handlers.GetHeatlh).Methods("GET")
    router.HandleFunc("/getNames", handlers.GetNames).Methods("GET")
    return router
}
