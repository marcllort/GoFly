package Utils

import (
	"GoFly/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strings"
	"time"
)

var dp Model.DialogflowProcessor
var resultNumber int
var place Model.PlacesResponse

func InitDialogFlow() {
	// Initialization of dialogFlow processor, with the basic info
	_ = dp.Init("flightbot-9a1fc", "credentials.json", "en", "Europe/Madrid")
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

		// Logging the request message
		l := log.New(os.Stdout, "", 0)
		LogString(l, m.Message)

		// Logging the response
		Log(l, response)

		// Call to API if searching for places
		if strings.Contains(response.ResponseMessage, "+") {
			place = RequestAPI(response.ResponseMessage)
			rand.Seed(time.Now().UnixNano())
			resultNumber = rand.Intn(5)
			apiResponse := place.Results[resultNumber].Name + " -- Do you want more information?"
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.ResponseMessage, "yes") {
			rating := fmt.Sprintf("%.1f", place.Results[resultNumber].Rating)
			response.ResponseMessage = "Direction: " + place.Results[resultNumber].FormattedAddress + " -- Rating: " + rating
		} else if strings.Contains(response.ResponseMessage, "no") {
			if len(place.Results) == resultNumber {
				resultNumber = 0
			} else {
				resultNumber++
			}
			apiResponse := "Maybe you like this one more! " + place.Results[resultNumber].Name + " -- Do you want more information?"
			response.ResponseMessage = apiResponse
		}

		// Prepare the JSON to return
		writter.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(writter).Encode(response)
		// When code reaches here, when "writter" content is sent to the user
	}
}
