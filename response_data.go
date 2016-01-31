package goqpx

//Data maps the data part of the response
type Data struct {
	Kind      string     `json:"kind"`
	Airports  []Airport  `json:"airport"`
	Cities    []City     `json:"city"`
	Aircrafts []Aircraft `json:"aircraft"`
	Taxes     []Tax      `json:"tax"`
	Carriers  []Carrier  `json:"carrier"`
}

//Airport maps an airport
type Airport struct {
	Code string `json:"code"`
	City string `json:"city"`
	Name string `json:"name"`
}

//BaseStruct is a struct composed by a Kind and a Name
type BaseStruct struct {
	Kind string `json:"kind"`
	Name string `json:"name"`
}

//CodeStruct is a BaseStruct extended with a Code
type CodeStruct struct {
	BaseStruct
	Code string `json:"code"`
}

//NameStruct s a BaseStruct extended with a Name
type NameStruct struct {
	BaseStruct
	Name string `json:"name"`
}

//City maps a city
type City CodeStruct

//Aircraft maps an aircraft
type Aircraft CodeStruct

//Tax maps a tax
type Tax NameStruct

//Carrier --
type Carrier CodeStruct

//GetCarrierName utlity method for getting a carrier name given its code
func (d Data) GetCarrierName(code string) string {
	for _, c := range d.Carriers {
		if c.Code == code {
			return c.BaseStruct.Name
		}
	}
	return "N.A."
}

//GetAirport utlity method for getting an airport name given its code
func (d Data) GetAirport(code string) Airport {
	for _, a := range d.Airports {
		if a.Code == code {
			return a
		}
	}
	return Airport{}
}

//GetCity utlity method for getting a city name given its code
func (d Data) GetCity(code string) string {
	for _, c := range d.Cities {
		if c.Code == code {
			return c.Name
		}
	}
	return "N.A."
}
