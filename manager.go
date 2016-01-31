package goqpx

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//RequestParams maps params for GoogleAPI request
type RequestParams map[string]string

//PerformRequest performs the request to GoogleAPI
func PerformRequest(params RequestParams, apiKey string) (*GoogleResponse, error) {
	var requestSlices []Slice
	var googleRequest GoogleRequest
	googleRequest.SaleCountry = SaleCountryDefault
	googleRequest.Refundable = false
	adultPassengers, _ := strconv.Atoi(params[PassengersNumber])
	googleRequest.AdultCount = adultPassengers
	solNumber, _ := strconv.Atoi(params[SolutionsNumber])
	googleRequest.Solutions = solNumber

	var tripTo Slice
	tripTo.Date = params[DepartureDate]
	tripTo.Origin = params[Origin]
	tripTo.Destination = params[Destination]

	requestSlices = append(requestSlices, tripTo)

	if params[ReturnDate] != "" {
		var tripBack Slice
		tripBack.Date = params[ReturnDate]
		tripBack.Origin = params[Destination]
		tripBack.Destination = params[Origin]
		requestSlices = append(requestSlices, tripBack)
	}

	googleRequest.Slices = requestSlices

	jsonPost, err := json.Marshal(googleRequest)
	if err != nil {
		return nil, err
	}
	request, err := http.NewRequest("POST", getGoogleURL(apiKey), bytes.NewBuffer(jsonPost))
	if err != nil {
		return nil, err
	}
	request.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	googleResponse, err := parseResponse(body)
	if err != nil {
		return nil, err
	}
	return googleResponse, nil
}

func parseResponse(response []byte) (*GoogleResponse, error) {
	r := &GoogleResponse{}
	err := json.Unmarshal(response, r)
	return r, err
}

func getGoogleURL(apiKey string) string {
	firstParam := fmt.Sprintf("?%s=%s", "key", apiKey)
	return fmt.Sprintf("%s%s", googleBaseURL, firstParam)
}
