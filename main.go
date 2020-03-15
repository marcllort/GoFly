package main

import (
	"GoFly/Utils"
	"fmt"
	"net/http"
)

func main() {
	Utils.InitDialogFlow()

	// Creation of Server listener for requests
	http.HandleFunc("/", Utils.RequestHandler)
	fmt.Println("Started listening...")
	http.ListenAndServe(":5000", nil)
}
