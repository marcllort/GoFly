package Utils

import (
	"GoFly/Model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var dp Model.DialogflowProcessor

func InitDialogFlow() {
	// Initialization of dialogFlow processor, with the basic info
	dp.Init("flightbot-9a1fc", "credentials.json", "en", "Europe/Madrid")
}

func RequestHandler(writter http.ResponseWriter, request *http.Request) {
	if request.Method == "POST" {
		//POST will receive a JSON, and return the response (as JSON)
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			http.Error(writter, "Error reading request body",
				http.StatusInternalServerError)
		}

		// Saving the JSON in the corresponding fields
		var m Model.InboundMessage
		err = json.Unmarshal(body, &m)
		if err != nil {
			panic(err)
		}

		// Pass message to DialogFlow to process the input
		response := dp.ProcessNLP(m.Message, m.User)

		// Logging the response
		l := log.New(os.Stdout, "", 0)
		Log(l, response)

		// Prepare the JSON to return
		writter.Header().Set("Content-Type", "application/json")
		json.NewEncoder(writter).Encode(response)
		// When code reaches here, when "writter" content is sent to the user
	}
}
