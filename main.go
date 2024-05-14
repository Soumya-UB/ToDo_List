package main

import (
	"ToDo_List/router"
	"fmt"
	"net/http"
)

func main() {
	// go func() {
	// 	fmt.Println("App started")
	// 	r := router.Router()
	// 	http.ListenAndServe(":8080", r)
	// }()
	fmt.Println("App started")
	r := router.Router()
	http.ListenAndServe(":8080", r)
}
