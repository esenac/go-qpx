package goqpx

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequestPrepareMissingParams(t *testing.T) {
	var request GoogleRequest
	params := make(RequestParams, 0)
	err := request.prepare(params)
	assert.NotNil(t, err, "Expected error")
	params[Origin] = "JFK"
	params[Destination] = "ORD"
	params[DepartureDate] = "2016-01-01"
	err = request.prepare(params)
	assert.Nil(t, err, fmt.Sprintf("Unexpected error: %s\n", err))
	assert.Equal(t, "JFK", request.Slices[0].Origin, "Unexpected origin")
	assert.Equal(t, "ORD", request.Slices[0].Destination, "Unexpected destination")
	assert.Equal(t, "2016-01-01", request.Slices[0].Date, "Unexpected date")
}

func TestRequestPrepareDefaultValues(t *testing.T) {
	var request GoogleRequest
	params := make(RequestParams, 0)
	params[Origin] = "JFK"
	params[Destination] = "ORD"
	params[DepartureDate] = "2016-01-01"
	err := request.prepare(params)
	assert.Nil(t, err, fmt.Sprintf("Unexpected error: %s\n", err))
	assert.Equal(t, saleCountryDefault, request.SaleCountry, fmt.Sprintf("Unexpected sale country: %s\n", request.SaleCountry))
	assert.True(t, request.Refundable, "Expected refundable request")
	assert.Equal(t, 1, request.AdultCount, fmt.Sprintf("Unexpected adult count: %d\n", request.AdultCount))
	assert.Equal(t, 0, request.ChildCount, fmt.Sprintf("Unexpected child count: %d\n", request.ChildCount))
	assert.Equal(t, 10, request.Solutions, fmt.Sprintf("Unexpected solutions count: %d\n", request.Solutions))
}
