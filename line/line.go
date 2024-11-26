package line

import "github.com/ownerofglory/govvs/vvscommon"

// Line represents a single line in the serving lines response.
//
// Fields:
// - ID (string): The unique identifier for the line.
// - Name (string): The full name of the line.
// - DisassembledName (string): A shortened version of the line's name.
// - Number (string): The line number (e.g., "U7").
// - Description (string): A textual description of the line's route.
// - Product (Product): Information about the line's product type.
// - Operator (Operator): The operator of the line.
// - Destination (Destination): The line's destination details.
// - Properties (LineProperties): Additional properties related to the line.
type Line struct {
	ID               string                              `json:"id"`
	Name             string                              `json:"name"`
	DisassembledName string                              `json:"disassembledName"`
	Number           string                              `json:"number"`
	Description      string                              `json:"description"`
	Product          vvscommon.TransportationProduct     `json:"product"`
	Operator         vvscommon.Operator                  `json:"operator"`
	Destination      vvscommon.TransportationDestination `json:"destination"`
	Properties       LineProperties                      `json:"properties"`
}

type LineProperties struct {
	IsTTB           bool     `json:"isTTB"`
	IsSTT           bool     `json:"isSTT"`
	IsROP           bool     `json:"isROP"`
	TripCode        int      `json:"tripCode"`
	TimetablePeriod string   `json:"timetablePeriod"`
	Validity        Validity `json:"validity"`
	LineDisplay     string   `json:"lineDisplay"`
}

type Validity struct {
	From string `json:"from"`
	To   string `json:"to"`
}
