package main

import (
	"ToDo_List/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("App started")
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
