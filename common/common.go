package common

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
