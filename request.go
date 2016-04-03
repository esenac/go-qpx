package goqpx

import (
	"strconv"
)

//GoogleRequest represents the request params sent to GoogleAPI
type GoogleRequest struct {
	Request `json:"request,omitempty"`
}

//Request is the effective body
type Request struct {
	Passengers  `json:"passengers,omitempty"`
	Slices      []Slice `json:"slice,omitempty"`
	MaxPrice    string  `json:"maxPrice,omitempty"`
	SaleCountry string  `json:"saleCountry,omitempty"`
	Refundable  bool    `json:"refundable,omitempty"`
	Solutions   int     `json:"solutions,omitempty"`
}

//Passengers represents the Passangers data sent in the request
type Passengers struct {
	Kind              string `json:"kind,omitempty"`
	AdultCount        int    `json:"adultCount,omitempty"`
	ChildCount        int    `json:"childCount,omitempty"`
	InfantInLapCount  int    `json:"infantInLapCount,omitempty"`
	InfantInSeatCount int    `json:"infantInSeatCount,omitempty"`
	SeniorCount       int    `json:"seniorCount,omitempty"`
}

//PermittedDepartureTime represents the data relative to the temporal range allowed for departure
type PermittedDepartureTime struct {
	Kind         string `json:"kind,omitempty"`
	EarliestTime string `json:"earliestTime,omitempty"`
	LatestTime   string `json:"latestTime,omitempty"`
}

//Slice represents the 'slice' component of the Request
type Slice struct {
	Kind                    string `json:"kind,omitempty"`
	Origin                  string `json:"origin,omitempty"`
	Destination             string `json:"destination,omitempty"`
	Date                    string `json:"date,omitempty"`
	MaxStops                *int   `json:"maxStops,omitempty"`
	MaxConnectionDuration   *int   `json:"maxConnectionDuration,omitempty"`
	PreferredCabin          string `json:"preferredCabin,omitempty"`
	*PermittedDepartureTime `json:"permittedDepartureTime,omitempty"`
	PermittedCarrier        []string `json:"permittedCarrier,omitempty"`
	Alliance                string   `json:"alliance,omitempty"`
	ProhibitedCarrier       []string `json:"prohibitedCarrier,omitempty"`
}

func (r *GoogleRequest) prepare(params RequestParams) error {
	if err := params.validate(); err != nil {
		return err
	}
	var err error

	saleCountry := saleCountryDefault
	if params[SaleCountry] != "" {
		r.SaleCountry = params[SaleCountry]
	}
	r.SaleCountry = saleCountry

	refundable := true
	if params[Refundable] != "" {
		refundable, err = strconv.ParseBool(params[Refundable])
		if err != nil {
			return err
		}
	}
	r.Refundable = refundable

	adultPassengers := 1
	if params[AdultPassengers] != "" {
		adultPassengers, err = strconv.Atoi(params[AdultPassengers])
		if err != nil {
			return err
		}
	}
	r.AdultCount = adultPassengers

	childPassengers := 0
	if params[ChildPassengers] != "" {
		childPassengers, err = strconv.Atoi(params[ChildPassengers])
		if err != nil {
			return err
		}
	}
	r.ChildCount = childPassengers

	var tripTo Slice
	tripTo.Date = params[DepartureDate]
	tripTo.Origin = params[Origin]
	tripTo.Destination = params[Destination]

	var requestSlices []Slice
	requestSlices = append(requestSlices, tripTo)

	if params[ReturnDate] != "" {
		var tripBack Slice
		tripBack.Date = params[ReturnDate]
		tripBack.Origin = params[Destination]
		tripBack.Destination = params[Origin]
		requestSlices = append(requestSlices, tripBack)
	}

	r.Slices = requestSlices

	solNumber := 10
	if params[SolutionsNumber] != "" {
		solNumber, err = strconv.Atoi(params[SolutionsNumber])
		if err != nil {
			return err
		}
	}
	r.Solutions = solNumber

	return nil
}
