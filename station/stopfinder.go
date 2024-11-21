package station

// Location represents a specific address or place result.
type Location struct {
	ID               string         `json:"id"`
	Name             string         `json:"name"`
	DisassembledName string         `json:"disassembledName"`
	Coord            []float64      `json:"coord"`
	StreetName       string         `json:"streetName"`
	BuildingNumber   string         `json:"buildingNumber"`
	Type             string         `json:"type"`
	MatchQuality     int            `json:"matchQuality"`
	IsBest           bool           `json:"isBest"`
	Parent           Parent         `json:"parent"`
	AssignedStops    []AssignedStop `json:"assignedStops"`
}

// Parent represents a higher-level location such as a locality (e.g., Stuttgart).
type Parent struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// AssignedStop represents nearby public transport stops related to the location.
type AssignedStop struct {
	ID             string         `json:"id"`
	IsGlobalId     bool           `json:"isGlobalId"`
	Name           string         `json:"name"`
	Type           string         `json:"type"`
	Coord          []float64      `json:"coord"`
	Parent         Parent         `json:"parent"`
	Distance       int            `json:"distance"` // Distance in meters
	Duration       int            `json:"duration"` // Duration in minutes
	ProductClasses []int          `json:"productClasses"`
	ConnectingMode int            `json:"connectingMode"`
	Properties     StopProperties `json:"properties"`
}

// StopProperties contains additional properties about a stop.
type StopProperties struct {
	StopID string `json:"stopId"`
}
