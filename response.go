package govvs

import (
	"github.com/ownerofglory/govvs/arrival"
	"github.com/ownerofglory/govvs/common"
	"github.com/ownerofglory/govvs/departure"
	"github.com/ownerofglory/govvs/journey"
)

// ArrivalResponse encapsulates the response data for an API request querying arrival information.
//
// Fields:
// - Parameters ([]Parameter): A list of parameters that were used in the request, providing context for the response.
// - DM (DM): Detailed information about the requested destination, including metadata.
// - Arr (Arr): Contains specific data related to the arrival query.
// - DateTime (DateTime): The date and time information relevant to the arrival request.
// - DateRange ([]DateRange): A list of date ranges applicable to the arrival information provided.
// - Option (Option): Options that were applied to the arrival request, such as filters or search criteria.
// - ServingLines (ServingLines): Information about the lines serving the queried station or location.
// - ArrivalList ([]Arrival): A comprehensive list of arrivals matching the request criteria, including details about each arrival.
//
// ArrivalResponse provides a structured format for interpreting the data returned from an arrival information request to the API.
type ArrivalResponse struct {
	Parameters   []common.Parameter  `json:"parameters"`
	DM           common.DM           `json:"dm"`
	Arr          common.Arr          `json:"arr"`
	DateTime     common.DateTime     `json:"dateTime"`
	DateRange    []common.DateRange  `json:"dateRange"`
	Option       common.Option       `json:"option"`
	ServingLines common.ServingLines `json:"servingLines"`
	ArrivalList  []arrival.Arrival   `json:"arrivalList"`
}

// DepartureResponse encapsulates the response data for an API request querying departure information.
//
// Fields:
// - Parameters ([]Parameter): A list of parameters that were used in the request, providing insight into the response context.
// - DM (DM): Provides detailed information about the destination or station from which departures are being queried.
// - Arr (Arr): Contains specific data related to the departure query.
// - DateTime (DateTime): The date and time information pertinent to the departure request.
// - DateRange ([]DateRange): A series of date ranges that are relevant to the departure information provided.
// - Option (Option): Options applied to the departure request, such as filters or search criteria.
// - ServingLines (ServingLines): Details about the lines that serve the queried station or location.
// - DepartureList ([]Departure): A detailed list of departures that match the request criteria, including information about each departure.
//
// DepartureResponse offers a structured approach to understanding the data returned from a departure information request to the API.
type DepartureResponse struct {
	Parameters    []common.Parameter    `json:"parameters"`
	DM            common.DM             `json:"dm"`
	Arr           common.Arr            `json:"arr"`
	DateTime      common.DateTime       `json:"dateTime"`
	DateRange     []common.DateRange    `json:"dateRange"`
	Option        common.Option         `json:"option"`
	ServingLines  common.ServingLines   `json:"servingLines"`
	DepartureList []departure.Departure `json:"departureList"`
}

// JourneyResponse represents the data structure for a response to a journey planning request.
//
// Fields:
// - ServerInfo (ServerInfo): Contains metadata about the server that processed the request, including version and performance metrics.
// - Version (string): The version of the API or data format used to generate the response.
// - SystemMessages ([]SystemMessage): A list of system messages related to the request, which may include warnings or errors encountered during processing.
// - Journeys ([]Journey): A list of journey options that match the request criteria, providing comprehensive details about each available journey.
//
// JourneyResponse provides a detailed overview of journey options between specified points, including routes, timings, and other relevant information as per the request parameters.
type JourneyResponse struct {
	ServerInfo     journey.ServerInfo      `json:"serverInfo"`
	Version        string                  `json:"version"`
	SystemMessages []journey.SystemMessage `json:"systemMessages"`
	Journeys       []journey.Journey       `json:"journeys"`
}
