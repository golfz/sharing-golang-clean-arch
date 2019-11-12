package request

type AddLocation struct {
	Datetime string  `json:"datetime"`
	Lat      float64 `json:"lat"`
	Long     float64 `json:"long"`
}
