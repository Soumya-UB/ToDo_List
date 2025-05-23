package middleware

import (
	"ToDo_List/db"
	"ToDo_List/errorTypes"
	"ToDo_List/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	Message string `json:"message,omitempty"`
}

func GetFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)

	id := params["id"]
	task, err := db.GetFile(id)

	if err != nil {
		_, ok := err.(*errorTypes.NoRowsFoundError)
		if ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func DeleteFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	id := params["id"]
	err := db.DeleteFile(id)
	if err != nil {
		if _, ok := err.(*errorTypes.NoRowsFoundError); ok {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	res2 := response{
		Message: "Record successfully updated",
	}
	json.NewEncoder(w).Encode(res2)
}

func GetAllFiles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	tasks, err := db.GetAllFiles()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(tasks)
}

func UpdateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	id := params["id"]
	var file models.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := db.UpdateFile(id, file)
	if err1 != nil {
		if _, ok := err1.(*errorTypes.NoRowsFoundError); ok {
			http.Error(w, err1.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	res2 := response{
		Message: "Record successfully updated",
	}
	json.NewEncoder(w).Encode(res2)
}

func CreateFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var file models.File
	err := json.NewDecoder(r.Body).Decode(&file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	err1 := db.CreateFile(file)
	if err1 != nil {
		http.Error(w, err1.Error(), http.StatusInternalServerError)
		return
	}
	res2 := response{
		Message: "Record successfully inserted",
	}
	json.NewEncoder(w).Encode(res2)
}
