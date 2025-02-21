package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("hello bura")

	app := fiber.New()

	todos := []Todo{}

	// Basic route
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "hello world"})
	})

	// POST route to add a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// Parse request body into Todo struct
		if err := c.BodyParser(todo); err != nil {
			return err
		}

		// Validate that Body is not empty
		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error": "todo body is required"})
		}

		// Assign an ID based on the current length of the todos slice
		todo.ID = len(todos) + 1

		// Append the new todo to the slice
		todos = append(todos, *todo)

		// Return the newly created todo
		return c.Status(201).JSON(todo)
	})

	// Start the Fiber server
	log.Fatal(app.Listen(":4000"))
}
