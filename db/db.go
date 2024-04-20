package db

import (
	"ToDo_List/models"
	model "ToDo_List/models"
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func createConnection() *sql.DB {
	db, err := sql.Open("mysql", "root:Jan@2019@tcp(127.0.0.1:3306)/test?parseTime=true")

	if err != nil {
		panic(err.Error())
	}

	db.Ping()
	return db
}

func GetTask(id string) (model.Task, error) {
	query := "select * from Task where Id=?"
	dbConn := createConnection()
	defer dbConn.Close()
	var task model.Task
	row := dbConn.QueryRow(query, id)
	// t := time.Time{}
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.StartTime)
	if err != nil {
		return task, err
	}
	return task, nil
}

func UpdateTask(id string, task model.Task) error {
	// desc := task.Description
	// title := task.Title
	// startTime := task.StartTime
	// query := "update task set Description=?, Title=?, StartTime=? where ID=?"

}

func CreateTask(task models.Task) error {
	query := "insert into Task values(?, ?, ?, ?)"
	dbConn := createConnection()
	defer dbConn.Close()
	rows, err := dbConn.Query(query, task.Id, task.Title, task.Description, task.StartTime)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
