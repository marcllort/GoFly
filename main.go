package main

import (
	"GoFly/Utils"
	"fmt"
	"net/http"
)

const port = ":8081"

func main() {
	Utils.InitDialogFlow()

	// Creation of Server listener for requests
	http.HandleFunc("/", Utils.RequestHandler)
	fmt.Println("Server started. Listening to port " + port)
	http.ListenAndServe(port, nil)

}
