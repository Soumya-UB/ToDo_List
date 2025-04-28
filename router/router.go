package router

import (
	"ToDo_List/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/task/{id}", middleware.GetFile).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newTask", middleware.CreateFile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/task/{id}", middleware.UpdateFile).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/tasks", middleware.GetAllFiles).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteFile).Methods("DELETE", "OPTIONS")
	return router
}
