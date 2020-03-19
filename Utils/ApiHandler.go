package Utils

import (
	"GoFly/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const APIKEY = "AIzaSyDk6ATVIpo6S_VlgXh1subtfBXVrRmK7jU"

func RequestAPI(place string) Model.PlacesResponse {

	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + place + "&key=" + APIKEY
	method := "POST"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	var result Model.PlacesResponse
	json.Unmarshal(body, &result)

	return result
}
