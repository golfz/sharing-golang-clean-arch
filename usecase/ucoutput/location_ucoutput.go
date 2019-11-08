package ucoutput

import "time"

type Location struct {
	Id    int64
	Time  time.Time
	Lat   float64
	Long  float64
	Speed int
}
