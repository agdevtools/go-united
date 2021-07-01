package router

import (
    "go-united/middleware"
    "github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router(r *mux.Router)  {

    //router := mux.NewRouter()

    r.HandleFunc("/api/team", middleware.GetAllPlayers).Methods("GET", "OPTIONS")
    r.HandleFunc("/api/team", middleware.CreatePlayer).Methods("POST", "OPTIONS")
    r.HandleFunc("/api/team/{id}", middleware.GetPlayer).Methods("GET", "OPTIONS")
    r.HandleFunc("/api/team/{id}", middleware.UpdatePlayer).Methods("PUT", "OPTIONS")
    r.HandleFunc("/api/deleteplayer/{id}", middleware.DeletePlayer).Methods("DELETE", "OPTIONS")

   // return router
}