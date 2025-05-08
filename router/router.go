package router

import (
	"ToDo_List/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/file/{id}", middleware.GetFile).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/newFile", middleware.CreateFile).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/file/{id}", middleware.UpdateFile).Methods("PUT", "OPTIONS")
	router.HandleFunc("/api/files", middleware.GetAllFiles).Methods("GET", "OPTIONS")
	router.HandleFunc("/api/deleteFile/{id}", middleware.DeleteFile).Methods("DELETE", "OPTIONS")
	return router
}
