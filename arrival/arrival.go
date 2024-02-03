package arrival

import "github.com/ownerofglory/govvs/common"

type Arrival struct {
	StopID       string             `json:"stopID"`
	X            string             `json:"x"`
	Y            string             `json:"y"`
	MapName      string             `json:"mapName"`
	Area         string             `json:"area"`
	Platform     string             `json:"platform"`
	PlatformName string             `json:"platformName,omitempty"`
	StopName     string             `json:"stopName"`
	NameWO       string             `json:"nameWO"`
	Countdown    string             `json:"countdown"`
	DateTime     common.DateTime    `json:"dateTime"`
	RealDateTime common.DateTime    `json:"realDateTime"`
	ServingLine  common.ServingLine `json:"servingLine"`
	Operator     common.Operator    `json:"operator"`
	Attrs        []common.Attr      `json:"attrs"`
}
