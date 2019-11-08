package viewmodel

type LocationData struct {
	Id       int64   `json:"id"`
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
	Speed    string  `json:"speed"`
}
