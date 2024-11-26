package geoobj

import "github.com/ownerofglory/govvs/vvscommon"

// Transfer represents a transfer or transportation object.
type Transfer struct {
	ID               string                              `json:"id"`
	Name             string                              `json:"name"`
	DisassembledName string                              `json:"disassembledName"`
	Number           string                              `json:"number"`
	Description      string                              `json:"description"`
	Product          vvscommon.TransportationProduct     `json:"product"`
	Operator         vvscommon.Operator                  `json:"operator"`
	Destination      vvscommon.TransportationDestination `json:"destination"`
	Properties       vvscommon.TransportationProperties  `json:"properties"`
	AssignedStop     string                              `json:"assignedStop"`
	AssignedStopID   string                              `json:"assignedStopID"`
}
