package govvs

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
