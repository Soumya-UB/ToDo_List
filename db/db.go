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
		log.Panicf("Couldn't read db password ", err3.Error())
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

func GetFile(id string) (model.File, error) {
	query := "select * from File where Id=?"
	dbConn := createConnection()
	defer dbConn.Close()
	var file model.File
	row := dbConn.QueryRow(query, id)
	err := row.Scan(&file.Id, &file.Name, &file.Size, &file.CreatedTime, &file.LastUpdatedTime, &file.IsDir)
	if errors.Is(err, sql.ErrNoRows) {
		err1 := errorTypes.NoRowsFoundError{"No row for id: " + id}
		return file, &err1
	}
	return file, nil
}

func GetAllFiles() ([]model.File, error) {
	query := "select * from File"
	db := createConnection()
	defer db.Close()
	var files []model.File
	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var file model.File
		err := rows.Scan(&file.Id, &file.Name, &file.Size, &file.CreatedTime, &file.LastUpdatedTime, &file.IsDir)
		if err != nil {
			return nil, err
		}
		files = append(files, file)
	}
	if err1 := rows.Err(); err1 != nil {
		return nil, err1
	}
	return files, nil
}

func UpdateFile(id string, file model.File) error {
	db := createConnection()
	defer db.Close()
	var paramList []any
	query := "update File set"
	if file.Name != "" {
		query = query + " name=?,"
		paramList = append(paramList, file.Name)
	}
	if &file.Size != nil {
		query = query + " size=?,"
		paramList = append(paramList, file.Size)
	}
	if !file.CreatedTime.IsZero() {
		query = query + " CreatedTime=?,"
		paramList = append(paramList, file.CreatedTime)
	}
	if !file.LastUpdatedTime.IsZero() {
		query = query + " LastUpdatedTime=?,"
		paramList = append(paramList, file.LastUpdatedTime)
	}
	if &file.IsDir != nil {
		query = query + " IsDir=?,"
		paramList = append(paramList, file.IsDir)
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

func DeleteFile(id string) error {
	db := createConnection()
	defer db.Close()
	query := "delete from File where Id=?"
	records, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	if num, _ := records.RowsAffected(); num < 1 {
		return &errorTypes.NoRowsFoundError{"No row for id: " + id}
	}
	return nil
}

func CreateFile(file models.File) error {
	log.Println("fileName: ", file.Name)
	log.Println("fileSize: ", file.Size)
	log.Println("CreatedTime: ", file.CreatedTime)
	log.Println("LastUpdatedTime: ", file.LastUpdatedTime)
	log.Println("IsDir: ", file.IsDir)
	query := "insert into File(Name, Size, CreatedTime, LastUpdatedTime, IsDir) values(?, ?, ?, ?, ?)"
	log.Println(query)
	dbConn := createConnection()
	defer dbConn.Close()
	rows, err := dbConn.Query(query, file.Name, file.Size, file.CreatedTime, file.LastUpdatedTime, file.IsDir)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
