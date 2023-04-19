package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

func main() {
	fmt.Print("Running Server")

	// := means variable is going to be whatever fiber.new returns. if only = used would need to declare type
	app := fiber.New()

	//Stores todos in memory not in DB
	todos := []Todo{}

	// func variable c of type fiber context.
	app.Get("/healthcheck", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	//Post function

	app.Post("/api/todos", func(c *fiber.Ctx) error {
		fmt.Print("IN POST FUNCITON")

		todo := &Todo{}
		fmt.Print("Created Todo Var")

		if err := c.BodyParser(todo); err != nil {
			fmt.Print("IN RETURN ERROR")
			return err
		}
		fmt.Print("no err")
		todo.ID = len(todos) + 1

		todos = append(todos, *todo)

		fmt.Print(c.JSON(todos))
		fmt.Print("returned Json")

		return c.JSON(todos)
	})

	log.Fatal(app.Listen((":4000")))
}
