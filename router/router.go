package router

import (
    "go-united/middleware"

    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

    router := mux.NewRouter()

    router.HandleFunc("/api/team", middleware.GetAllPlayers).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/team", middleware.CreatePlayer).Methods("POST", "OPTIONS")
    router.HandleFunc("/api/team/{id}", middleware.GetPlayer).Methods("GET", "OPTIONS")
    router.HandleFunc("/api/team/{id}", middleware.UpdatePlayer).Methods("PUT", "OPTIONS")
    router.HandleFunc("/api/deleteplayer/{id}", middleware.DeletePlayer).Methods("DELETE", "OPTIONS")


    return router
}