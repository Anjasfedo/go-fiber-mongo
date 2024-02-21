package main

import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const DB_NAME = "fiber-mongo"
const MONGO_URI = "mongodb://localhost:27017" + DB_NAME

type Employee struct {
	Id     string  `json:"id,omitempty" bson:"_id, omitempty"`
	Name   string  `json:"name"`
	Age    int64   `json:"age"`
	Salary float64 `json:"salary"`
}

func ConnectDatabase() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(MONGO_URI))

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = client.Connect(ctx)

	db := client.Database(DB_NAME)

	if err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		Db:     db,
	}

	return nil
}

func main() {

	if err := ConnectDatabase(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Get("/employee", func(c *fiber.Ctx) error {
		query := bson.D{{}}

		cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		var employees []Employee = make([]Employee, 0)

		if err := cursor.All(c.Context(), &employees); err != nil {
			return c.Status(500).SendString(err.Error())
		}

		return c.JSON(employees)
	})

	app.Post("/employee", func(c *fiber.Ctx) error {
		collection := mg.Db.Collection("employees")

		employee := new(Employee)

		if err := c.BodyParser(employee); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		employee.Id = ""

		insertResult, err := collection.InsertOne(c.Context(), employee)
		if err != nil {
			return c.Status(500).SendString(err.Error())
		}

		filter := bson.D{{Key: "_id", Value: insertResult.InsertedID}}

		createdRecord := collection.FindOne(c.Context(), filter)

		createdEmployee := &Employee{}

		createdRecord.Decode(createdEmployee)

		return c.Status(201).JSON(createdEmployee)
	})

	app.Put("/employee/:id")

	app.Delete("/employee/:id")
}
