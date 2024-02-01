package govvs

type DepartureResponse struct {
	Parameters    []Parameter  `json:"parameters"`
	DM            DM           `json:"dm"`
	Arr           Arr          `json:"arr"`
	DateTime      DateTime     `json:"dateTime"`
	DateRange     []DateRange  `json:"dateRange"`
	Option        Option       `json:"option"`
	ServingLines  ServingLines `json:"servingLines"`
	DepartureList []Departure  `json:"departureList"`
}

type Departure struct {
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
	ServingLine  ServingLine `json:"servingLine"`
	Operator     Operator    `json:"operator"`
	StopInfos    *StopInfos  `json:"stopInfos,omitempty"`
	LineInfos    interface{} `json:"lineInfos,omitempty"`
	TripInfos    interface{} `json:"tripInfos,omitempty"`
}

type DateTime struct {
	Year    string `json:"year"`
	Month   string `json:"month"`
	Day     string `json:"day"`
	Weekday string `json:"weekday"`
	Hour    string `json:"hour"`
	Minute  string `json:"minute"`
}

type StopInfos struct {
	StopInfo StopInfo `json:"stopInfo"`
}

type StopInfo struct {
	InfoLinkText string   `json:"infoLinkText"`
	InfoLinkURL  string   `json:"infoLinkURL"`
	InfoText     InfoText `json:"infoText"`
	ParamList    []Param  `json:"paramList"`
}

type InfoText struct {
	Content        string `json:"content"`
	Subtitle       string `json:"subtitle"`
	Subject        string `json:"subject"`
	AdditionalText string `json:"additionalText,omitempty"`
	HtmlText       string `json:"htmlText"`
	WmlText        string `json:"wmlText,omitempty"`
	SmsText        string `json:"smsText"`
	SpeechText     string `json:"speechText,omitempty"`
}

type Param struct {
	Type  string `json:"type"`
	Name  string `json:"name"`
	Value string `json:"value"`
	Edit  string `json:"edit"`
}
