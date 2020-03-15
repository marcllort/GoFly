package main

import (
	"GoFly/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

var dp Model.DialogflowProcessor

func main() {
	dp.Init("flightbot-9a1fc", "credentials.json", "en", "Europe/Madrid")
	http.HandleFunc("/", requestHandler)
	fmt.Println("Started listening...")
	http.ListenAndServe(":5000", nil)
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		//POST method, receives a json to parse
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body",
				http.StatusInternalServerError)
		}

		var m Model.InboundMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			panic(err)
		}

		// Use NLP
		response := dp.ProcessNLP(m.Message, "testUser")
		fmt.Printf("%#v", response)
		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(response)
	}
}
