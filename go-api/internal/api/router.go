package api

import (
    "my-go-api/internal/api/handlers"
    "github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/getNames", handlers.GetNames).Methods("GET")
    return router
}
