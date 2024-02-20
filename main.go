package main

import "github.com/gofiber/fiber/v2"

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const DB_NAME = "fiber-mongo"
const MONGO_URL = "mongodb://localhost:27017" + DB_NAME

type Employee struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Salary float64 `json:"salary"`
}

func ConnectDatabase() error {

}

func main() {
	app := fiber.New()

	app.Get("/employee")

	app.Post("/employee")

	app.Put("/employee/:id")

	app.Delete("/employee/:id")
}
