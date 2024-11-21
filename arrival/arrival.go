package arrival

import "github.com/ownerofglory/govvs/vvscommon"

type Arrival struct {
	StopID       string                `json:"stopID"`
	X            string                `json:"x"`
	Y            string                `json:"y"`
	MapName      string                `json:"mapName"`
	Area         string                `json:"area"`
	Platform     string                `json:"platform"`
	PlatformName string                `json:"platformName,omitempty"`
	StopName     string                `json:"stopName"`
	NameWO       string                `json:"nameWO"`
	Countdown    string                `json:"countdown"`
	DateTime     vvscommon.DateTime    `json:"dateTime"`
	RealDateTime vvscommon.DateTime    `json:"realDateTime"`
	ServingLine  vvscommon.ServingLine `json:"servingLine"`
	Operator     vvscommon.Operator    `json:"operator"`
	Attrs        []vvscommon.Attr      `json:"attrs"`
}
