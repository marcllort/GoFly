package Utils

import (
	"GoFly/Model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const APIKEY = "AIzaSyDk6ATVIpo6S_VlgXh1subtfBXVrRmK7jU"

func RequestAPI(place string) Model.PlacesResponse {

	placeModified := strings.ReplaceAll(place, " ", "+")
	url := "https://maps.googleapis.com/maps/api/place/textsearch/json?query=" + placeModified + "&key=" + APIKEY
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

func RequestDetails(id string) Model.DetailedPlaceResponse {
	url2 := "https://maps.googleapis.com/maps/api/place/details/json?place_id=" + id + "&key=" + APIKEY
	method := "GET"
	req, err := http.NewRequest(method, url2, nil)
	client := &http.Client{}
	if err != nil {
		fmt.Println(err)
	}
	res, err := client.Do(req)
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	var result2 Model.DetailedPlaceResponse
	json.Unmarshal(body, &result2)

	return result2
}
