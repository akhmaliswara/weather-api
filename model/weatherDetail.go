package model

import (
	"time"
)

type WeatherDetail struct {
	Id          uint `gorm:"-"`
	WeatherId   uint
	Main        string
	Description string
	Icon        string
	RecordTime  time.Time
}