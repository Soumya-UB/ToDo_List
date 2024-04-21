package middleware

import (
	"ToDo_List/db"
	"ToDo_List/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type response struct {
	Message string `json:"message,omitempty"`
}

// type errResponse struct {
// 	ErrorMessage string `json:"message,omitempty"`
// }

func GetTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	id := params["id"]
	task, err := db.GetTask(id)
	if err != nil {
		log.Printf("Error getting task: %v", err)
		res := response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	json.NewEncoder(w).Encode(task)
}

func UpdateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	id := params["id"]
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		res := response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	err1 := db.UpdateTask(id, task)
	if err1 != nil {
		log.Printf("Error updating task: %v", err1.Error())
		res := response{
			Message: err1.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	res2 := response{
		Message: "Record successfully updated",
	}
	json.NewEncoder(w).Encode(res2)
}

func CreateTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	var task models.Task
	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		res := response{
			Message: err.Error(),
		}
		json.NewEncoder(w).Encode(res)
		return
	}
	err1 := db.CreateTask(task)
	if err1 != nil {
		res1 := response{
			Message: err1.Error(),
		}
		json.NewEncoder(w).Encode(res1)
		return
	}
	res2 := response{
		Message: "Record successfully inserted",
	}
	json.NewEncoder(w).Encode(res2)
}
