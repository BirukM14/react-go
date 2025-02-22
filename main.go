package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Todo struct {
	ID        int    `json:"id" bson:"_id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

var collection *mongo.Collection

func main() {
	fmt.Println("Hello")

	// Load environment variables
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	// Get MongoDB URI
	MONGOURI := os.Getenv("MONGOURI")
	if MONGOURI == "" {
		log.Fatal("MONGOURI is not set in .env file")
	}

	// Set MongoDB client options
	clientOptions := options.Client().ApplyURI(MONGOURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to MongoDB:", err)    
	}

	defer client.Disconnect(context.Background())

	// Ping MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Could not connect to MongoDB:", err)
	}

	fmt.Println("Connected to MongoDB Atlas")

	app := fiber.New()

	app.Get("/api/todos".getTodos)
	app.Post("/api/todos".createTodo)

	app.Patch("/api/todos/:id".updateTodo)

	app.Delete("/api/todos/:id",deleteTodo)

	port := os.Getenv("PORT")
	if port == ""{
		port="5000"
	}
	log.Fatal(app.listen("0.0.0.0:"+ port))

}

func getTodos(c *fiber.ctx) error{
	var todos []Todo
	cursor, err := collection.Find(context.Background(),bson.M{})

	if err != nil{
		return err
	}

	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()){
		var todo Todo
		if err := cursor.Decode(&todo); err !=nil{
			return err
		}
		todos = append(todos, todo)
	}
	return c.JSON(todos)

}
