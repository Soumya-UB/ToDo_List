package main

import (
	"ToDo_List/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("App started")
	setLogFile() // Use goroutine to close the log file
	r := router.Router()
	http.ListenAndServe(":8080", r)
}

func setLogFile() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Panicf("Failed to open file")
	}
	log.SetOutput(file)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.Println("Log file created!")
}
