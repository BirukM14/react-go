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
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// POST route to add a todo
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{}

		// Parse request body into Todo struct
		if err := c.BodyParser(todo); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
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

	// PATCH route to update a todo
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")

		var updateData struct {
			Body *string `json:"body"`
		}

		if err := c.BodyParser(&updateData); err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
		}

		for i, todo := range todos {
			if fmt.Sprint(todo.ID) == id {
				if updateData.Body != nil {
					todos[i].Body = *updateData.Body
				}
				todos[i].Completed = true
				return c.Status(200).JSON(todos[i])
			}
		}
		return c.Status(404).JSON(fiber.Map{"error": "todo not found"})
	})


	app.Delete("api/todos/:id", func (c *fiber.Ctx) error {
		id := c.Params("id")

		for i, todo := range todos{
			if fmt.Sprint(todo.ID) == id {

				todos= append(todos[:i],todos[i+1:]...)
				return c.Status(200).JSON(fiber.Map{"sucess":"true"})

			}
		} 
		return c.Status(404).JSON(fiber.Map{"error":"todo not found"})
		
		
	})

	// Start the Fiber server
	log.Fatal(app.Listen(":4000"))
}
