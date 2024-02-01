package govvs

type ArrivalResponse struct {
	Parameters   []Parameter  `json:"parameters"`
	DM           DM           `json:"dm"`
	Arr          Arr          `json:"arr"`
	DateTime     DateTime     `json:"dateTime"`
	DateRange    []DateRange  `json:"dateRange"`
	Option       Option       `json:"option"`
	ServingLines ServingLines `json:"servingLines"`
	ArrivalList  []Arrival    `json:"arrivalList"`
}

type ServingLines struct {
	Lines []Line `json:"lines"`
}

type Parameter struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

type DM struct {
	Input               Input               `json:"input"`
	Points              DMPoints            `json:"points"`
	ItdOdvAssignedStops ItdOdvAssignedStops `json:"itdOdvAssignedStops"`
}

type Arr struct {
	Input  ArrInput    `json:"input"`
	Points interface{} `json:"points"`
}

type ArrInput struct {
	Input string `json:"input"`
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

type Arrival struct {
	StopID       string      `json:"stopID"`
	X            string      `json:"x"`
	Y            string      `json:"y"`
	MapName      string      `json:"mapName"`
	Area         string      `json:"area"`
	Platform     string      `json:"platform"`
	PlatformName string      `json:"platformName,omitempty"`
	StopName     string      `json:"stopName"`
	NameWO       string      `json:"nameWO"`
	Countdown    string      `json:"countdown"`
	DateTime     DateTime    `json:"dateTime"`
	RealDateTime DateTime    `json:"realDateTime"`
	ServingLine  ServingLine `json:"servingLine"`
	Operator     Operator    `json:"operator"`
	Attrs        []Attr      `json:"attrs"`
}

type Attr struct {
	Name  string `json:"name"`
	Value string `json:"value"`
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

type DateRange struct {
	Day     string `json:"day"`
	Month   string `json:"month"`
	Year    string `json:"year"`
	Weekday string `json:"weekday"`
}
