package data

type Meter struct {
	Root struct {
		ID string `json:"_id"`
	} `json:"root"`
	EAN         string `json:"ean"`
	Consumption string `json:"consumption"`
	Box         string `json:"box"`
	Revoked     bool   `json:"revoked"`
}

type Response struct {
	Meters []Meter `json:"data"`
}
