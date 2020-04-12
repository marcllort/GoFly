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
var placeDetailed Model.DetailedPlaceResponse
var historic []string

const EMPTY = "Which place yo you want to know about?"

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
		if strings.Contains(m.Message, "historic") {
			apiResponse := "Here are your last searches!"
			for i := range historic {
				content := fmt.Sprintf("Historic %d: %s", i, historic[i])
				apiResponse = apiResponse + "\n" + content
			}
			fmt.Println(apiResponse)
			response.ResponseMessage = apiResponse

		} else if strings.Contains(response.Intent, "Open") {

			if response.ResponseMessage != EMPTY {
				place = RequestAPI(response.ResponseMessage)
				placeDetailed = RequestDetails(place.Results[0].PlaceID)
			}
			var apiResponse string
			if len(place.Results) != 0 {
				if !placeDetailed.Result.PermanentlyClosed {
					apiResponse = place.Results[0].Name + " is open!"
				} else {
					apiResponse = place.Results[0].Name + " is closed!"
				}
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.Intent, "Web") {
			if response.ResponseMessage != EMPTY {
				place = RequestAPI(response.ResponseMessage)
				placeDetailed = RequestDetails(place.Results[0].PlaceID)
			}
			var apiResponse string
			if len(place.Results) != 0 {
				apiResponse = "The website is: " + placeDetailed.Result.Website
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.Intent, "Review") {
			if response.ResponseMessage != EMPTY {
				place = RequestAPI(response.ResponseMessage)
				placeDetailed = RequestDetails(place.Results[0].PlaceID)
			}
			var apiResponse string
			rand.Seed(time.Now().UnixNano())
			resultNumber = rand.Intn(len(placeDetailed.Result.Reviews))
			if len(place.Results) != 0 {
				apiResponse = "Here you have a review: " + placeDetailed.Result.Reviews[resultNumber].Text
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.Intent, "Info") {
			place = RequestAPI(response.ResponseMessage)
			placeDetailed = RequestDetails(place.Results[0].PlaceID)
			var apiResponse string
			if len(place.Results) != 0 {
				rating := fmt.Sprintf("%.1f", place.Results[0].Rating)
				apiResponse = place.Results[0].Name + " info -- Direction: " + place.Results[0].FormattedAddress + " -- Rating: " + rating
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.Intent, "address") || strings.Contains(response.Intent, "arribe") {
			if response.ResponseMessage != EMPTY {
				place = RequestAPI(response.ResponseMessage)
				placeDetailed = RequestDetails(place.Results[0].PlaceID)
			}
			var apiResponse string

			if len(place.Results) != 0 {
				apiResponse = placeDetailed.Result.FormattedAddress
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.Intent, "Phone") {
			if response.ResponseMessage != EMPTY {
				place = RequestAPI(response.ResponseMessage)
				placeDetailed = RequestDetails(place.Results[0].PlaceID)
			}
			var apiResponse string
			rand.Seed(time.Now().UnixNano())
			resultNumber = rand.Intn(len(placeDetailed.Result.Reviews))
			if len(place.Results) != 0 {
				apiResponse = "Here is the phone number: " + placeDetailed.Result.FormattedPhoneNumber
				historic = append(historic, response.ResponseMessage)
			} else {
				apiResponse = "Which place is the one you are looking for?"
			}
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.ResponseMessage, "cities") {
			place = RequestAPI(response.ResponseMessage)
			rand.Seed(time.Now().UnixNano())
			resultNumber = rand.Intn(5)
			apiResponse := "You could like " + place.Results[resultNumber].Name + ". Ask me for places to visit there!"
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.ResponseMessage, "+") {
			place = RequestAPI(response.ResponseMessage)
			rand.Seed(time.Now().UnixNano())
			resultNumber = rand.Intn(5)
			apiResponse := place.Results[resultNumber].Name + " -- Do you want more information?"
			response.ResponseMessage = apiResponse
		} else if strings.Contains(response.ResponseMessage, "yes") {
			rating := fmt.Sprintf("%.1f", place.Results[resultNumber].Rating)
			response.ResponseMessage = "Direction: " + place.Results[resultNumber].FormattedAddress + " -- Rating: " + rating
			placeDetailed = RequestDetails(place.Results[resultNumber].PlaceID)
			// Save the search on historic table
			historic = append(historic, response.ResponseMessage)
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
