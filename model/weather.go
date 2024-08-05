package model

import (
	"time"
)

type Weather struct {
	Id         uint `gorm:"primary_key;auto_increment"`
	Lat        float64
	Lon        float64
	Timezone   int
	Pressure   int
	Humidity   int
	WindSpeed  float64
	Weather    []WeatherDetail `gorm:"foreignkey:WeatherID"`
	RecordTime time.Time
}