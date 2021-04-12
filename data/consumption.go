package data

type Consumption struct {
	MeterID string `json:"meter_id"`
	Periods []struct {
		Energy    []int  `json:"energy"`
		StartDate string `json:"start_date"`
	} `json:"periods"`
}

type ConsumptionResponse struct {
	Data []Consumption `json:"data"`
}
