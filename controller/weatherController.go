package controller

import (
	"net/http"
	"time"
	
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"be/model"
)

func GetWeatherData(c *gin.Context, db *gorm.DB) {
	var weatherData []model.Weather
	if err := db.Preload("Weather").Find(&weatherData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch weather data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   weatherData,
	})
}

func CreateWeatherData(c *gin.Context, db *gorm.DB) {
	var reqBody struct {
		Coord struct {
			Lat float64
			Lon float64
		}
		Main struct {
			Pressure int
			Humidity int
		}
		Timezone int
		Wind     struct {
			Speed float64
		}
		Weather []struct {
			Id          uint
			Main        string
			Description string
			Icon        string
		}
	}

	if err := c.ShouldBindJSON(&reqBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Insert Weather record
	weather := model.Weather{
		Lat:        reqBody.Coord.Lat,
		Lon:        reqBody.Coord.Lon,
		Timezone:   reqBody.Timezone,
		Pressure:   reqBody.Main.Pressure,
		Humidity:   reqBody.Main.Humidity,
		WindSpeed:  reqBody.Wind.Speed,
		RecordTime: time.Now(),
	}

	tx := db.Begin()

	if err := tx.Create(&weather).Error; err != nil {
		tx.Rollback()
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create weather record"})
		return
	}

	for _, body := range reqBody.Weather {
		mappedDetail := model.WeatherDetail{
			WeatherId:   weather.Id,
			Id:          body.Id,
			Main:        body.Main,
			Description: body.Description,
			Icon:        body.Icon,
			RecordTime:  time.Now(),
		}

		if err := tx.Create(&mappedDetail).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	tx.Commit()

	c.JSON(http.StatusOK, gin.H{"message": "Weather data created successfully"})
}