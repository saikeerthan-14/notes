package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ToDo struct {
	ID        primitive.ObjectID `json:"id" bson:"_id"`
	BODY      string             `json:"body"`
	COMPLETED bool               `json:"completed"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("entered")

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading env file", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Atlas")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	// app.Patch("/api/todos/:id", updateTodo)
	// app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")

	if port == "" {
		port = "4000"
	}

	log.Fatal(app.Listen("0.0.0.0:" + port))

}

func getTodos(c fiber.Ctx) error {
	var todos []ToDo

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var todo ToDo
		if err := cursor.Decode(&todo); err != nil {
			return err
		}

		todos = append(todos, todo)
	}

	return c.JSON(todos)

}

func createTodo(c fiber.Ctx) error {
	todo := new(ToDo)

	if err := c.Bind().Body(todo); err != nil {
		return err
	}

	if todo.BODY == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Body field is required"})
	}

	insertResult, err := collection.InsertOne(context.Background(), todo)
	if err != nil {
		return err
	}

	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

}

// func updateTodo(c fiber.Ctx) error {}
// func deleteTodos(c fiber.Ctx) error {}
