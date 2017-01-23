# goqpx
[![Build Status](https://travis-ci.org/esenac/goqpx.svg?branch=master)](https://travis-ci.org/esenac/goqpx)

[Google QPX Express API](https://developers.google.com/qpx-express/) client written in Go.

## Usage example

```go

...
params := make(goqpx.RequestParams)
params[goqpx.Origin] = "JFK"
params[goqpx.Destination] = "ORD"
params[goqpx.AdultPassengers] = "1"
params[goqpx.DepartureDate] = "2016-08-01"
params[goqpx.SolutionsNumber] = "10"

response, err := goqpx.PerformRequest(params, _yourGoogleAPIKey_)
...

```
