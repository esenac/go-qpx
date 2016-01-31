package goqpx

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
