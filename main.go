package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"

	"be/controller"
	"be/middleware"
	"be/model"
)

var (
	db  *gorm.DB
	err error
)

func main() {
	// Initialize Gin router
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour

	router.Use(cors.New(config))

	router.Use(func(c *gin.Context) {
		if c.Request.URL.Path != "/api/login" {
			middleware.AuthMiddleware()(c)
		}
	})

	// Connect to PostgreSQL database
	err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file")
    }
	dbConnectionString := os.Getenv("DB_CONNECTION_STRING")
	if dbConnectionString == "" {
		log.Fatal("DB_CONNECTION_STRING environment variable is not set")
	}

	db, err = gorm.Open("postgres", dbConnectionString)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}
	defer db.Close()

	db.LogMode(true)

	// Automigrate the Weather model and relation
	db.AutoMigrate(&model.Weather{}, &model.WeatherDetail{}, &model.User{})
	db.Model(&model.Weather{}).Related(&model.WeatherDetail{})

	// Define route handlers
	router.GET("/api/weather", func(c *gin.Context) {
		controller.GetWeatherData(c, db)
	})
	router.POST("/api/weather", func(c *gin.Context) {
		controller.CreateWeatherData(c, db)
	})
	// router.POST("/api/user", func(c* gin.Context) {
	// 	controller.CreateUserHandler(c, db)
	// })
	router.POST("/api/login", func(c *gin.Context) {
		controller.LoginHandler(c, db)
	})

	// Run the server
	router.Run(":8080")
}
