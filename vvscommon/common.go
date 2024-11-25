package vvscommon

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

type ServingLine struct {
	Key           string      `json:"key"`
	Code          string      `json:"code"`
	Number        string      `json:"number"`
	Symbol        string      `json:"symbol"`
	MotType       string      `json:"motType"`
	MtSubcode     string      `json:"mtSubcode"`
	Realtime      string      `json:"realtime"`
	Direction     string      `json:"direction"`
	DirectionFrom string      `json:"directionFrom"`
	Name          string      `json:"name"`
	LiErgRiProj   LiErgRiProj `json:"liErgRiProj"`
	DestID        string      `json:"destID"`
	Stateless     string      `json:"stateless"`
	LineDisplay   string      `json:"lineDisplay"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type LiErgRiProj struct {
	Line       string `json:"line"`
	Project    string `json:"project"`
	Direction  string `json:"direction"`
	Supplement string `json:"supplement"`
	Network    string `json:"network"`
	Gid        string `json:"gid"`
}

type Operator struct {
	Code       string `json:"code"`
	Name       string `json:"name"`
	PublicCode string `json:"publicCode,omitempty"`
}

type DateTime struct {
	Year    string `json:"year"`
	Month   string `json:"month"`
	Day     string `json:"day"`
	Weekday string `json:"weekday"`
	Hour    string `json:"hour"`
	Minute  string `json:"minute"`
}

type DM struct {
	Input               Input               `json:"input"`
	Points              DMPoints            `json:"points"`
	ItdOdvAssignedStops ItdOdvAssignedStops `json:"itdOdvAssignedStops"`
}

type Input struct {
	Input string `json:"input"`
}

type DMPoints struct {
	Point DMPoint `json:"point"`
}

type DMPoint struct {
	Usage     string      `json:"usage"`
	Type      string      `json:"type"`
	Name      string      `json:"name"`
	Stateless string      `json:"stateless"`
	AnyType   string      `json:"anyType"`
	Sort      string      `json:"sort"`
	Quality   string      `json:"quality"`
	Best      string      `json:"best"`
	Object    string      `json:"object"`
	Ref       Ref         `json:"ref"`
	Infos     interface{} `json:"infos"`
}

type Ref struct {
	ID      string `json:"id"`
	GID     string `json:"gid"`
	OMC     string `json:"omc"`
	PlaceID string `json:"placeID"`
	Place   string `json:"place"`
	Coords  string `json:"coords"`
}

type ItdOdvAssignedStops struct {
	StopID         string `json:"stopID"`
	Name           string `json:"name"`
	X              string `json:"x"`
	Y              string `json:"y"`
	MapName        string `json:"mapName"`
	Value          string `json:"value"`
	Place          string `json:"place"`
	NameWithPlace  string `json:"nameWithPlace"`
	DistanceTime   string `json:"distanceTime"`
	IsTransferStop string `json:"isTransferStop"`
	VM             string `json:"vm"`
	GID            string `json:"gid"`
}

type ServingLines struct {
	Lines []Line `json:"lines"`
}

type Line struct {
	Mode  Mode   `json:"mode"`
	Index string `json:"index"`
}

type Mode struct {
	Name            string `json:"name"`
	Number          string `json:"number"`
	Product         string `json:"product"`
	ProductId       string `json:"productId"`
	Type            string `json:"type"`
	Code            string `json:"code"`
	Destination     string `json:"destination"`
	DestID          string `json:"destID"`
	Desc            string `json:"desc"`
	TimetablePeriod string `json:"timetablePeriod"`
	Diva            Diva   `json:"diva"`
}

type Diva struct {
	Branch      string `json:"branch"`
	Line        string `json:"line"`
	Supplement  string `json:"supplement"`
	Dir         string `json:"dir"`
	Project     string `json:"project"`
	Network     string `json:"network"`
	Stateless   string `json:"stateless"`
	GlobalId    string `json:"globalId"`
	TripCode    string `json:"tripCode"`
	Operator    string `json:"operator"`
	OpCode      string `json:"opCode"`
	VF          string `json:"vF"`
	VTo         string `json:"vTo"`
	LineDisplay string `json:"lineDisplay"`
	IsTTB       string `json:"isTTB"`
	IsSTT       string `json:"isSTT"`
	IsROP       string `json:"isROP"`
	Attrs       []Attr `json:"attrs"`
}

type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Arr struct {
	Input  ArrInput    `json:"input"`
	Points interface{} `json:"points"`
}

type ArrInput struct {
	Input string `json:"input"`
}

type DateRange struct {
	Day     string `json:"day"`
	Month   string `json:"month"`
	Year    string `json:"year"`
	Weekday string `json:"weekday"`
}

type Option struct {
	PtOption PtOption `json:"ptOption"`
}

type PtOption struct {
	Active                       string          `json:"active"`
	MaxChanges                   string          `json:"maxChanges"`
	MaxTime                      string          `json:"maxTime"`
	MaxWait                      string          `json:"maxWait"`
	RouteType                    string          `json:"routeType"`
	ChangeSpeed                  string          `json:"changeSpeed"`
	LineRestriction              string          `json:"lineRestriction"`
	UseProxFootSearch            string          `json:"useProxFootSearch"`
	UseProxFootSearchOrigin      string          `json:"useProxFootSearchOrigin"`
	UseProxFootSearchDestination string          `json:"useProxFootSearchDestination"`
	Bike                         string          `json:"bike"`
	Plane                        string          `json:"plane"`
	NoCrowded                    string          `json:"noCrowded"`
	NoSolidStairs                string          `json:"noSolidStairs"`
	NoEscalators                 string          `json:"noEscalators"`
	NoElevators                  string          `json:"noElevators"`
	LowPlatformVhcl              string          `json:"lowPlatformVhcl"`
	Wheelchair                   string          `json:"wheelchair"`
	NeedElevatedPlt              string          `json:"needElevatedPlt"`
	Assistance                   string          `json:"assistance"`
	SOSAvail                     string          `json:"SOSAvail"`
	NoLonelyTransfer             string          `json:"noLonelyTransfer"`
	IllumTransfer                string          `json:"illumTransfer"`
	OvergroundTransfer           string          `json:"overgroundTransfer"`
	NoInsecurePlaces             string          `json:"noInsecurePlaces"`
	PrivateTransport             string          `json:"privateTransport"`
	ExcludedMeans                []ExcludedMeans `json:"excludedMeans"`
	ActiveImp                    string          `json:"activeImp"`
	ActiveCom                    string          `json:"activeCom"`
	ActiveSec                    string          `json:"activeSec"`
}

type ExcludedMeans struct {
	Means    string `json:"means"`
	Value    string `json:"value"`
	Selected string `json:"selected"`
}

// Transportation represents a line or transportation object.
//
// Fields:
// - ID (string): The unique identifier for the transportation object.
// - Name (string): The name of the transportation object, e.g., "Stadtbahn U7".
// - Description (string): A description of the transportation line, including its route.
// - Product (Product): Details about the transportation product (e.g., type, icon).
// - Operator (Operator): Details about the operator managing the transportation line.
// - Destination (Destination): Information about the destination of the line.
// - Properties (TransportationProperties): Additional properties related to the line.
// - Coords ([][][]float64): A list of coordinates representing the route.
// - LocationSequence ([]Location): A sequence of locations or stops on the line.
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

type LocationProperties struct {
	Downloads      []Download    `json:"downloads"`
	AreaNiveauDiva string        `json:"AREA_NIVEAU_DIVA"`
	AreaGid        string        `json:"areaGid"`
	Area           string        `json:"area"`
	Platform       string        `json:"platform"`
	PlatformName   string        `json:"platformName,omitempty"`
	AccessArray    []AccessArray `json:"accessArray"`
}

type ParentProperties struct {
	StopId string `json:"stopId"`
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
