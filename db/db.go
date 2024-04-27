package db

import (
	"ToDo_List/errorTypes"
	"ToDo_List/models"
	model "ToDo_List/models"
	"database/sql"
	"errors"
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
	err := row.Scan(&task.Id, &task.Title, &task.Description, &task.StartTime)
	if errors.Is(err, sql.ErrNoRows) {
		err1 := errorTypes.NoRowsFoundError{"No row for id: " + id}
		return task, &err1
	}
	return task, nil
}

func GetAllTasks() ([]model.Task, error) {
	query := "select * from Task"
	db := createConnection()
	defer db.Close()
	var tasks []model.Task
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var task model.Task
		err := rows.Scan(&task.Id, &task.Title, &task.Description, &task.StartTime)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err1 := rows.Err(); err1 != nil {
		return nil, err1
	}
	return tasks, nil
}

func UpdateTask(id string, task model.Task) error {
	db := createConnection()
	defer db.Close()
	var paramList []any
	query := "update task set"
	if task.Description != "" {
		query = query + " Description=?,"
		paramList = append(paramList, task.Description)
	}
	if task.Title != "" {
		query = query + " Title=?,"
		paramList = append(paramList, task.Title)
	}
	if !task.StartTime.IsZero() {
		query = query + " StartTime=?,"
		paramList = append(paramList, task.StartTime)
	}
	query = query[:len(query)-1]
	query = query + " where ID=?"
	paramList = append(paramList, id)
	records, err := db.Exec(query, paramList...)
	if err != nil {
		return err
	}
	if num, _ := records.RowsAffected(); num < 1 {
		return &errorTypes.NoRowsFoundError{"No row for id: " + id}
	}
	return nil
}

func DeleteTask(id string) error {
	db := createConnection()
	defer db.Close()
	query := "delete from Task where Id=?"
	records, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	if num, _ := records.RowsAffected(); num < 1 {
		return &errorTypes.NoRowsFoundError{"No row for id: " + id}
	}
	return nil
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
