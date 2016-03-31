package goqpx

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateParams(t *testing.T) {
	params := make(RequestParams)
	err := params.validate()
	assert.NotNil(t, err, "Expected error")
	assert.EqualError(t, err, "Origin not set", "Error different from expected")
	params[Origin] = "JFK"
	err = params.validate()
	assert.NotNil(t, err, "Expected error")
	assert.EqualError(t, err, "Destination not set", "Error different from expected")
	params[Destination] = "ORD"
	err = params.validate()
	assert.NotNil(t, err, "Expected error")
	assert.EqualError(t, err, "Departure date not set", "Error different from expected")
	params[DepartureDate] = "2016-01-01"
	err = params.validate()
	assert.Nil(t, err, fmt.Sprintf("Unexpected error: %s", err))
}
