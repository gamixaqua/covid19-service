package dtos

type CovidCase struct {
	Success          bool   `json:"success"`
	LastRefreshed    string `json:"lastRefreshed"`
	LastOriginUpdate string `json:"lastOriginUpdate"`
	Data             Data   `json:"data"`
}
type Data struct {
	Summary  Summary    `json:"summary"`
	Regional []Regional `json:"regional"`
}

type Summary struct {
	Total     int64 `json:"total"`
	Recovered int64 `json:"discharged"`
	Deaths    int64 `json:"deaths"`
}
type Regional struct {
	StateName      string `json:"loc"`
	Deaths         int64  `json:"deaths"`
	Recovered      int64  `json:"discharged"`
	TotalConfirmed int64  `json:"totalConfirmed"`
}

type ReverseGeoCodingResponse struct {
	ReverseGeoCodingData []ReverseGeoCodingData `json:"data"`
}

type ReverseGeoCodingData struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Region    string  `json:"region"`
}
