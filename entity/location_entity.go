package entity

import (
	"math/rand"
	"time"
)

type Location struct {
	Id       *int64
	Time     time.Time
	Lat      float64
	Long     float64
	DeviceId int64
}

func (d *Location) GetSpeedMPH() int64 {
	return int64(rand.Intn(40) + 60) // 60 - 100
}
