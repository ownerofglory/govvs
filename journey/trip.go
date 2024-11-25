package journey

import "github.com/ownerofglory/govvs/vvscommon"

type Leg struct {
	Duration             int                      `json:"duration"`
	IsRealtimeControlled bool                     `json:"isRealtimeControlled"`
	RealtimeStatus       []string                 `json:"realtimeStatus"`
	Origin               vvscommon.Location       `json:"origin"`
	Destination          vvscommon.Location       `json:"destination"`
	Transportation       vvscommon.Transportation `json:"transportation"`
	StopSequence         []StopSequenceItem       `json:"stopSequence"`
	Infos                []Info                   `json:"infos"`
	Coords               [][]float64              `json:"coords"`
	Properties           LegProperties            `json:"properties"`
}

type StopSequenceItem struct {
	IsGlobalId           bool                     `json:"isGlobalId"`
	ID                   string                   `json:"id"`
	Name                 string                   `json:"name"`
	DisassembledName     string                   `json:"disassembledName,omitempty"`
	Type                 string                   `json:"type"`
	PointType            string                   `json:"pointType,omitempty"`
	Coord                []float64                `json:"coord"`
	Niveau               int                      `json:"niveau"`
	Parent               vvscommon.ParentLocation `json:"parent"`
	ProductClasses       []int                    `json:"productClasses"`
	Properties           SequenceProperties       `json:"properties"`
	ArrivalTimePlanned   string                   `json:"arrivalTimePlanned,omitempty"`
	ArrivalTimeEstimated string                   `json:"arrivalTimeEstimated,omitempty"`
}

type SequenceProperties struct {
	AreaNiveauDiva       string                  `json:"AREA_NIVEAU_DIVA"`
	StoppingPointPlanned string                  `json:"stoppingPointPlanned,omitempty"`
	AreaGid              string                  `json:"areaGid"`
	Area                 string                  `json:"area"`
	Platform             string                  `json:"platform"`
	PlatformName         string                  `json:"platformName,omitempty"`
	AccessArray          []vvscommon.AccessArray `json:"accessArray,omitempty"`
	Zone                 string                  `json:"zone,omitempty"`
}

type Info struct {
	Priority   string         `json:"priority"`
	ID         string         `json:"id"`
	Version    string         `json:"version"`
	Type       string         `json:"type"`
	UrlText    string         `json:"urlText"`
	Url        string         `json:"url"`
	Content    string         `json:"content"`
	Subtitle   string         `json:"subtitle"`
	Title      string         `json:"title"`
	Properties InfoProperties `json:"properties"`
}

type InfoProperties struct {
	Publisher string `json:"publisher"`
	InfoType  string `json:"infoType"`
	HtmlText  string `json:"htmlText"`
	SmsText   string `json:"smsText"`
}

type LegProperties struct {
	VehicleAccess       []VehicleAccess `json:"vehicleAccess"`
	PlanLowFloorVehicle string          `json:"PlanLowFloorVehicle"`
}

type VehicleAccess struct {
	Height     int `json:"height"`
	HalfWidth  int `json:"halfWidth"`
	DoorWidth  int `json:"doorWidth"`
	Attributes int `json:"attributes"`
	AreaToStay int `json:"areaToStay"`
	AreaToMove int `json:"areaToMove"`
}

type Ticket struct {
	ID                      string           `json:"id"`
	Name                    string           `json:"name"`
	Comment                 string           `json:"comment"`
	URL                     string           `json:"URL"`
	Currency                string           `json:"currency"`
	PriceLevel              string           `json:"priceLevel"`
	PriceBrutto             float64          `json:"priceBrutto"`
	PriceNetto              float64          `json:"priceNetto"`
	TaxPercent              float64          `json:"taxPercent"`
	FromLeg                 int              `json:"fromLeg"`
	ToLeg                   int              `json:"toLeg"`
	Net                     string           `json:"net"`
	Person                  string           `json:"person"`
	TravellerClass          string           `json:"travellerClass"`
	TimeValidity            string           `json:"timeValidity"`
	ValidMinutes            int              `json:"validMinutes"`
	ValidityExtent          string           `json:"validityExtent"`
	IsShortHaul             string           `json:"isShortHaul"`
	ReturnsAllowed          string           `json:"returnsAllowed"`
	ValidForOneJourneyOnly  string           `json:"validForOneJourneyOnly"`
	ValidForOneOperatorOnly string           `json:"validForOneOperatorOnly"`
	NumberOfChanges         int              `json:"numberOfChanges"`
	NameValidityArea        string           `json:"nameValidityArea"`
	Properties              TicketProperties `json:"properties"`
}

type TicketProperties struct {
	TariffProductDefault []interface{} `json:"tariffProductDefault"`
	TariffProductOption  []interface{} `json:"tariffProductOption"`
}

type Journey struct {
	Rating        int           `json:"rating"`
	IsAdditional  bool          `json:"isAdditional"`
	Interchanges  int           `json:"interchanges"`
	Legs          []Leg         `json:"legs"`
	Fare          Fare          `json:"fare"`
	DaysOfService DaysOfService `json:"daysOfService"`
}

type Fare struct {
	Tickets []Ticket `json:"tickets"`
	Zones   []Zone   `json:"zones"`
}

type Zone struct {
	Net         string   `json:"net"`
	ToLeg       int      `json:"toLeg"`
	FromLeg     int      `json:"fromLeg"`
	NeutralZone string   `json:"neutralZone"`
	Zones       []string `json:"zones"`
}

type DaysOfService struct {
	Rvb string `json:"rvb"`
}
