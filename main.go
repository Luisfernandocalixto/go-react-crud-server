package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/luis/go-react-crud/models"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	app := fiber.New()

	uri := os.Getenv("MONGODB_URI")
	docs := "www.mongodb.com/docs/drivers/go/current/"

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable. " +
			"See: " + docs +
			"usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(options.Client().
		ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	app.Use(cors.New())

	app.Static("/", "./public")

	app.Post("/users", func(c *fiber.Ctx) error {
		var user models.User

		c.BodyParser(&user)

		coll := client.Database("gomongodb").Collection("users")
		result, err := coll.InsertOne(context.TODO(), user)

		fmt.Print(err)

		return c.JSON(&fiber.Map{
			"data": result,
		})
	})

	app.Get("/users", func(c *fiber.Ctx) error {
		var users []models.User

		coll := client.Database("gomongodb").Collection("users")
		results, error := coll.Find(context.TODO(), bson.M{})

		if error != nil {
			panic(error)
		}

		for results.Next(context.TODO()) {
			var user models.User
			results.Decode(&user)
			users = append(users, user)
		}

		return c.JSON(&fiber.Map{
			"users": users,
		})

	})

	app.Listen(":" + port)
	fmt.Println("Server on port " + port)
}

/* TODO
show id primitive
*/
