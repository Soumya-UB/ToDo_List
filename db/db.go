package db

import (
	"ToDo_List/config"
	"ToDo_List/errorTypes"
	"ToDo_List/models"
	model "ToDo_List/models"
	"database/sql"
	"errors"
	"log"
	"os"
	"strconv"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

var (
	conf config.Config
	errN error
)

func newDbConfig() (config.Config, error) {
	var once sync.Once
	once.Do(func() {
		conf, errN = config.GetConfig()
		if errN != nil {
			log.Panicf("Fatal error reading config file: %w", errN)
		}
	})
	return conf, nil
}

func createConnection() *sql.DB { // Create connection just once
	conf, err2 := newDbConfig()
	if err2 != nil {
		log.Panicf(err2.Error())
	}
	uname := conf.Username
	passwordPath := conf.PasswordPath
	instance := conf.Instance
	port := conf.Port
	dbName := conf.DbName
	if _, err := os.Stat(passwordPath); errors.Is(err, os.ErrNotExist) {
		log.Panicf("Couldn't find db password %v", err.Error())
	}
	password, err3 := os.ReadFile(passwordPath)
	if err3 != nil {
		log.Panicf("Couldn't read db password %v", err3.Error())
	}
	connString := uname + ":" + string(password[:]) + "@tcp(" + instance + ":" + strconv.Itoa(port) + ")/" + dbName + "?parseTime=true"
	log.Println("Connection string: %v", connString)
	db, err := sql.Open("mysql", connString)

	if err != nil {
		log.Panicf(err.Error())
	}

	err1 := db.Ping()
	if err1 != nil {
		log.Panicf(err1.Error())
	}
	log.Println("db connection successful")
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
