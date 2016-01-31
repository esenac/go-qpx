package goqpx

//GoogleResponse maps the whole Google response
type GoogleResponse struct {
	Kind  string `json:"kind"`
	Trips `json:"trips"`
}

//Trips maps the 'trips' part of the Google response
type Trips struct {
	Kind        string `json:"kind"`
	RequestID   string `json:"requestId"`
	Data        `json:"data"`
	TripOptions []TripOption `json:"tripOption"`
}

//TripOption maps a 'tripOption' object
type TripOption struct {
	Kind           string          `json:"kind"`
	SaleTotal      string          `json:"saleTotal"`
	ID             string          `json:"id"`
	ResponseSlices []ResponseSlice `json:"slice"`
	Pricings       []Pricing       `json:"pricing"`
}

//ResponseSlice maps a 'slice' in the response
type ResponseSlice struct {
	Kind     string    `json:"kind"`
	Duration int       `json:"duration"`
	Segments []Segment `json:"segment"`
}

//Segment maps a slice segment
type Segment struct {
	Kind                string `json:"kind"`
	Duration            int    `json:"duration"`
	Flight              `json:"flight"`
	ID                  string `json:"id"`
	Cabin               string `json:"cabin"`
	BookingCode         string `json:"bookingCode"`
	BookingCodeCount    int    `json:"bookingCodeCount"`
	MarriedSegmentGroup string `json:"marriedSegmentGroup"`
	Legs                []Leg  `json:"leg"`
	ConnectionDuration  int    `json:"connectionDuration, omitempty"`
}

//Flight maps a flight
type Flight struct {
	Carrier string `json:"carrier"`
	Number  string `json:"number"`
}

//Leg maps a leg
type Leg struct {
	Kind                string `json:"kind"`
	ID                  string `json:"id"`
	Aircraft            string `json:"aircraft"`
	ArrivalTime         string `json:"arrivalTime"`
	DepartureTime       string `json:"departureTime"`
	Origin              string `json:"origin"`
	Destination         string `json:"destination"`
	OriginTerminal      string `json:"originTerminal"`
	DestinationTerminal string `json:"destinationTerminal"`
	Duration            int    `json:"duration"`
	Mileage             int    `json:"mileage"`
	Meal                string `json:"meal"`
	Secure              bool   `json:"secure"`
}

//Pricing maps the pricing part of the response
type Pricing struct {
	Kind               string           `json:"kind"`
	Fares              []Fare           `json:"fare"`
	SegmentPricings    []SegmentPricing `json:"segmentPricing"`
	BaseFareTotal      string           `json:"baseFareTotal"`
	SaleFareTotal      string           `json:"saleFareTotal"`
	SaleTaxTotal       string           `json:"saleTaxTotal"`
	SaleTotal          string           `json:"saleTotal"`
	ResponsePassengers `json:"passengers"`
	PricingTaxes       []PricingTax `json:"tax"`
	FareCalculation    string       `json:"fareCalculation"`
	LastTicketingTime  string       `json:"latestTicketingTime"`
	Ptc                string       `json:"ptc"`
	Refundable         bool         `json:"refundable"`
}

//Fare maps a fare
type Fare struct {
	Kind        string `json:"kind"`
	ID          string `json:"id"`
	Carrier     string `json:"carrier"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	BasisCode   string `json:"basisCode"`
}

//SegmentPricing maps segmentPricing
type SegmentPricing struct {
	Kind      string `json:"kind"`
	FareID    string `json:"fareId"`
	SegmentID string `json:"segmentId"`
}

//ResponsePassengers maps passengers in the response
type ResponsePassengers struct {
	Kind       string `json:"kind"`
	AdultCount int    `json:"adultCount"`
}

//PricingTax maps pricingTax
type PricingTax struct {
	Kind       string `json:"kind"`
	ID         string `json:"id"`
	ChargeType string `json:"chargeType"`
	Code       string `json:"code"`
	Country    string `json:"country"`
	SalePrice  string `json:"salePrice"`
}
