package govvs

type JourneyResponse struct {
	ServerInfo     ServerInfo      `json:"serverInfo"`
	Version        string          `json:"version"`
	SystemMessages []SystemMessage `json:"systemMessages"`
	Journeys       []Journey       `json:"journeys"`
}

type Leg struct {
	Duration             int                `json:"duration"`
	IsRealtimeControlled bool               `json:"isRealtimeControlled"`
	RealtimeStatus       []string           `json:"realtimeStatus"`
	Origin               Location           `json:"origin"`
	Destination          Location           `json:"destination"`
	Transportation       Transportation     `json:"transportation"`
	StopSequence         []StopSequenceItem `json:"stopSequence"`
	Infos                []Info             `json:"infos"`
	Coords               [][]float64        `json:"coords"`
	Properties           LegProperties      `json:"properties"`
}

type Location struct {
	IsGlobalId                 bool               `json:"isGlobalId"`
	ID                         string             `json:"id"`
	Name                       string             `json:"name"`
	Type                       string             `json:"type"`
	Coord                      []float64          `json:"coord"`
	Niveau                     int                `json:"niveau"`
	Parent                     *ParentLocation    `json:"parent,omitempty"`
	ProductClasses             []int              `json:"productClasses"`
	DepartureTimeBaseTimetable string             `json:"departureTimeBaseTimetable,omitempty"`
	DepartureTimePlanned       string             `json:"departureTimePlanned,omitempty"`
	DepartureTimeEstimated     string             `json:"departureTimeEstimated,omitempty"`
	ArrivalTimeBaseTimetable   string             `json:"arrivalTimeBaseTimetable,omitempty"`
	ArrivalTimePlanned         string             `json:"arrivalTimePlanned,omitempty"`
	ArrivalTimeEstimated       string             `json:"arrivalTimeEstimated,omitempty"`
	Properties                 LocationProperties `json:"properties"`
}

type ParentLocation struct {
	IsGlobalId bool             `json:"isGlobalId"`
	ID         string           `json:"id"`
	Name       string           `json:"name"`
	Type       string           `json:"type"`
	Parent     *ParentLocation  `json:"parent,omitempty"`
	Properties ParentProperties `json:"properties,omitempty"`
	Coord      []float64        `json:"coord,omitempty"`
	Niveau     int              `json:"niveau,omitempty"`
}

type ParentProperties struct {
	StopId string `json:"stopId"`
}

type LocationProperties struct {
	Downloads      []Download    `json:"downloads"`
	AreaNiveauDiva string        `json:"AREA_NIVEAU_DIVA"`
	AreaGid        string        `json:"areaGid"`
	Area           string        `json:"area"`
	Platform       string        `json:"platform"`
	PlatformName   string        `json:"platformName,omitempty"`
	AccessArray    []AccessArray `json:"accessArray"`
}

type Download struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

type AccessArray struct {
	StoppointAccess StoppointAccess `json:"stoppointAccess"`
}

type StoppointAccess struct {
	Height        int `json:"height"`
	Equipment     int `json:"equipment"`
	MaxWeight     int `json:"maxWeight"`
	DistanceTrack int `json:"distanceTrack"`
	LengthAvail   int `json:"lengthAvail"`
	WidthAvail    int `json:"widthAvail"`
	Attributes    int `json:"attributes"`
}

type Transportation struct {
	ID               string                    `json:"id"`
	Name             string                    `json:"name"`
	DisassembledName string                    `json:"disassembledName"`
	Number           string                    `json:"number"`
	Description      string                    `json:"description"`
	Product          TransportationProduct     `json:"product"`
	Operator         Operator                  `json:"operator"`
	Destination      TransportationDestination `json:"destination"`
	Properties       TransportationProperties  `json:"properties"`
}

type TransportationProduct struct {
	ID     int    `json:"id"`
	Class  int    `json:"class"`
	Name   string `json:"name"`
	IconId int    `json:"iconId"`
}

type TransportationDestination struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type TransportationProperties struct {
	IsTTB           bool   `json:"isTTB"`
	IsSTT           bool   `json:"isSTT"`
	IsROP           bool   `json:"isROP"`
	TripCode        int    `json:"tripCode"`
	TimetablePeriod string `json:"timetablePeriod"`
	LineDisplay     string `json:"lineDisplay"`
	GlobalId        string `json:"globalId"`
	AVMSTripID      string `json:"AVMSTripID"`
	OperatorURL     string `json:"OperatorURL"`
}

type StopSequenceItem struct {
	IsGlobalId           bool               `json:"isGlobalId"`
	ID                   string             `json:"id"`
	Name                 string             `json:"name"`
	DisassembledName     string             `json:"disassembledName,omitempty"`
	Type                 string             `json:"type"`
	PointType            string             `json:"pointType,omitempty"`
	Coord                []float64          `json:"coord"`
	Niveau               int                `json:"niveau"`
	Parent               ParentLocation     `json:"parent"`
	ProductClasses       []int              `json:"productClasses"`
	Properties           SequenceProperties `json:"properties"`
	ArrivalTimePlanned   string             `json:"arrivalTimePlanned,omitempty"`
	ArrivalTimeEstimated string             `json:"arrivalTimeEstimated,omitempty"`
}

type SequenceProperties struct {
	AreaNiveauDiva       string        `json:"AREA_NIVEAU_DIVA"`
	StoppingPointPlanned string        `json:"stoppingPointPlanned,omitempty"`
	AreaGid              string        `json:"areaGid"`
	Area                 string        `json:"area"`
	Platform             string        `json:"platform"`
	PlatformName         string        `json:"platformName,omitempty"`
	AccessArray          []AccessArray `json:"accessArray,omitempty"`
	Zone                 string        `json:"zone,omitempty"`
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

type ServerInfo struct {
	ControllerVersion string  `json:"controllerVersion"`
	ServerID          string  `json:"serverID"`
	VirtDir           string  `json:"virtDir"`
	ServerTime        string  `json:"serverTime"`
	CalcTime          float64 `json:"calcTime"`
	LogRequestId      string  `json:"logRequestId"`
}

type SystemMessage struct {
	Type    string `json:"type"`
	Module  string `json:"module"`
	Code    int    `json:"code"`
	Text    string `json:"text"`
	SubType string `json:"subType"`
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
