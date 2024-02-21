package main

import (
	"context"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client
	Db
}

var mg MongoInstance

const DB_NAME = "fiber-mongo"
const MONGO_URI = "mongodb://localhost:27017" + DB_NAME

type Employee struct {
	Id     int64   `json:"id"`
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Salary float64 `json:"salary"`
}

func ConnectDatabase() error {
	client, err := mongo.newClient(options.Client().ApplyURI(MONGO_URI))

	context.withTimeOut(context.Background())
	if err != nil {

	}
}

func main() {

	if err := ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {

	})

	app.Post("/employee")

	app.Put("/employee/:id")

	app.Delete("/employee/:id")
}
