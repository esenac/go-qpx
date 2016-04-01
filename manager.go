package goqpx

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

//RequestParams maps params for GoogleAPI request
type RequestParams map[string]string

func (r RequestParams) validate() error {
	// Airport codes validation could be enforced
	// by ensuring that the code refers to an
	// existing airport
	if r[Origin] == "" {
		return errors.New("Origin not set")
	}
	if r[Destination] == "" {
		return errors.New("Destination not set")
	}
	if r[DepartureDate] == "" {
		return errors.New("Departure date not set")
	}
	return nil
}

//PerformRequest performs the request to GoogleAPI
func PerformRequest(params RequestParams, apiKey string) (*GoogleResponse, error) {
	if err := params.validate(); err != nil {
		return nil, err
	}

	var requestSlices []Slice
	var googleRequest GoogleRequest
	var err error

	googleRequest.SaleCountry = saleCountryDefault
	if params[SaleCountry] != "" {
		googleRequest.SaleCountry = params[SaleCountry]
	}

	refundable := true
	if params[Refundable] != "" {
		refundable, err = strconv.ParseBool(params[Refundable])
		if err != nil {
			return nil, err
		}
	}

	googleRequest.Refundable = refundable

	adultPassengers := 1

	if params[PassengersNumber] != "" {
		adultPassengers, _ := strconv.Atoi(params[PassengersNumber])

	}
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
