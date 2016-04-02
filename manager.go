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
	var googleRequest GoogleRequest

	err := googleRequest.prepare(params)
	if err != nil {
		return nil, err
	}

	solNumber, _ := strconv.Atoi(params[SolutionsNumber])
	googleRequest.Solutions = solNumber

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
