package data

type Meter struct {
	Root struct {
		ID string `json:"_id"`
	} `json:"root"`
	EAN         string `json:"ean"`
	Consumption string
	Box         string
	Revoked     bool
}

type Response struct {
	Meters []Meter `json:"data"`
}
