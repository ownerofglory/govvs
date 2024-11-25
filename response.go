package govvs

import (
	"github.com/ownerofglory/govvs/arrival"
	"github.com/ownerofglory/govvs/departure"
	"github.com/ownerofglory/govvs/geoobj"
	"github.com/ownerofglory/govvs/journey"
	"github.com/ownerofglory/govvs/station"
	"github.com/ownerofglory/govvs/vvscommon"
)

// ArrivalResponse encapsulates the response data for an API request querying arrival information.
//
// Fields:
// - Parameters ([]vvscommon.Parameter): A list of parameters that were used in the request, providing context for the response.
// - DM (vvscommon.DM): Detailed information about the requested destination, including metadata.
// - Arr (vvscommon.Arr): Contains specific data related to the arrival query.
// - DateTime (vvscommon.DateTime): The date and time information relevant to the arrival request.
// - DateRange ([]vvscommon.DateRange): A list of date ranges applicable to the arrival information provided.
// - Option (vvscommon.Option): Options that were applied to the arrival request, such as filters or search criteria.
// - ServingLines (vvscommon.ServingLines): Information about the lines serving the queried station or location.
// - ArrivalList ([]arrival.Arrival): A comprehensive list of arrivals matching the request criteria, including details about each arrival.
//
// ArrivalResponse provides a structured format for interpreting the data returned from an arrival information request to the API.
type ArrivalResponse struct {
	Parameters   []vvscommon.Parameter  `json:"parameters"`
	DM           vvscommon.DM           `json:"dm"`
	Arr          vvscommon.Arr          `json:"arr"`
	DateTime     vvscommon.DateTime     `json:"dateTime"`
	DateRange    []vvscommon.DateRange  `json:"dateRange"`
	Option       vvscommon.Option       `json:"option"`
	ServingLines vvscommon.ServingLines `json:"servingLines"`
	ArrivalList  []arrival.Arrival      `json:"arrivalList"`
}

// DepartureResponse encapsulates the response data for an API request querying departure information.
//
// Fields:
// - Parameters ([]vvscommon.Parameter): A list of parameters that were used in the request, providing insight into the response context.
// - DM (vvscommon.DM): Provides detailed information about the destination or station from which departures are being queried.
// - Arr (vvscommon.Arr): Contains specific data related to the departure query.
// - DateTime (vvscommon.DateTime): The date and time information pertinent to the departure request.
// - DateRange ([]vvscommon.DateRange): A series of date ranges that are relevant to the departure information provided.
// - Option (vvscommon.Option): Options applied to the departure request, such as filters or search criteria.
// - ServingLines (vvscommon.ServingLines): Details about the lines that serve the queried station or location.
// - DepartureList ([]departure.Departure): A detailed list of departures that match the request criteria, including information about each departure.
//
// DepartureResponse offers a structured approach to understanding the data returned from a departure information request to the API.
type DepartureResponse struct {
	Parameters    []vvscommon.Parameter  `json:"parameters"`
	DM            vvscommon.DM           `json:"dm"`
	Arr           vvscommon.Arr          `json:"arr"`
	DateTime      vvscommon.DateTime     `json:"dateTime"`
	DateRange     []vvscommon.DateRange  `json:"dateRange"`
	Option        vvscommon.Option       `json:"option"`
	ServingLines  vvscommon.ServingLines `json:"servingLines"`
	DepartureList []departure.Departure  `json:"departureList"`
}

// JourneyResponse represents the data structure for a response to a journey planning request.
//
// Fields:
// - ServerInfo (vvscommon.ServerInfo): Contains metadata about the server that processed the request, including version and performance metrics.
// - Version (string): The version of the API or data format used to generate the response.
// - SystemMessages ([]vvscommon.SystemMessage): A list of system messages related to the request, which may include warnings or errors encountered during processing.
// - Journeys ([]journey.Journey): A list of journey options that match the request criteria, providing comprehensive details about each available journey.
//
// JourneyResponse provides a detailed overview of journey options between specified points, including routes, timings, and other relevant information as per the request parameters.
type JourneyResponse struct {
	ServerInfo     vvscommon.ServerInfo      `json:"serverInfo"`
	Version        string                    `json:"version"`
	SystemMessages []vvscommon.SystemMessage `json:"systemMessages"`
	Journeys       []journey.Journey         `json:"journeys"`
}

// StopFinderResponse represents the data structure for a response to a stop finder request.
//
// Fields:
// - Version (string): The version of the API or data format used to generate the response.
// - ServerInfo (vvscommon.ServerInfo): Contains metadata about the server that processed the request, including version, server ID, and processing time metrics.
// - SystemMessages ([]vvscommon.SystemMessage): A list of system messages related to the request, such as status messages or warnings.
// - Locations ([]station.Location): A list of locations that match the stop finder request, each containing details about the location, including nearby public transport stops.
//
// StopFinderResponse provides a detailed overview of locations and associated public transport stops that match the requested input, along with relevant server information and system messages.
type StopFinderResponse struct {
	Version        string                    `json:"version"`
	ServerInfo     vvscommon.ServerInfo      `json:"serverInfo"`
	SystemMessages []vvscommon.SystemMessage `json:"systemMessages"`
	Locations      []station.Location        `json:"locations"`
}

// GeoObjectResponse represents the data structure for a response from the geo-object request.
//
// Fields:
// - ServerInfo (ServerInfo): Contains metadata about the server that processed the request, including version and performance metrics.
// - Version (string): The version of the API used to generate the response.
// - Transportations ([]Transportation): A list of transport objects, including lines, stops, and their associated data.
// - Transfers ([]Transfer): A list of transfer options associated with the requested line or object.
//
// GeoObjectResponse provides comprehensive details about a specific line or transportation object, including its stops, coordinates, and transfer options.
type GeoObjectResponse struct {
	ServerInfo      vvscommon.ServerInfo       `json:"serverInfo"`
	Version         string                     `json:"version"`
	Transportations []vvscommon.Transportation `json:"transportations"`
	Transfers       []geoobj.Transfer          `json:"transfers"`
}
