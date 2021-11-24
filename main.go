package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sprite5641/lab-golang/controllers"
)

func main() {
	app := fiber.New()

	app.Get("/api/todos", controllers.GetAllTodos)
	app.Get("/api/todos/:id", controllers.GetTodoByID)
	app.Post("/api/todos", controllers.CreateTodo)
	app.Patch("/api/todos/:id", controllers.ToggleTodoStatus)
	app.Delete("/api/todos/:id", controllers.DeleteTodo)

	app.Listen("127.0.0.1:3000")
}
