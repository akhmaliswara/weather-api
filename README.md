# Backend API for Weather Application

This is the backend API for a weather application, built using Go and the Gin framework. It provides endpoints for retrieving and creating weather data, as well as user authentication.

## Getting Started

To run the application, follow these steps:

1. Install the required dependencies by running `go get` in the terminal.
2. Create a PostgreSQL database and update the connection string in `main.go`.
3. Run the application using `go run main.go`.
4. The API will be available at `http://localhost:8080`.

## Endpoints

The following endpoints are available:

* `GET /api/weather`: Retrieves weather data.
* `POST /api/weather`: Creates new weather data.
* `POST /api/login`: Authenticates a user.

## Models

The application uses the following models:

* `Weather`: Represents weather data.
* `WeatherDetail`: Represents detailed weather data.
* `User`: Represents a user.

## Middleware

The application uses the following middleware:

* `AuthMiddleware`: Authenticates requests.

## Dependencies

The application uses the following dependencies:

* `github.com/gin-contrib/cors`
* `github.com/gin-gonic/gin`
* `github.com/jinzhu/gorm`
* `github.com/jinzhu/gorm/dialects/postgres`

## Database

The application uses a PostgreSQL database. The connection string is defined in `main.go`.

## Logging

The application logs requests and errors using the `log` package.

## Security

The application uses CORS to enable cross-origin requests. The `AuthMiddleware` authenticates requests.