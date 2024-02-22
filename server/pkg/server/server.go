package server

import (

	"github.com/rahmatadlin/Todo-Golang-React/pkg/controller"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func AppWithRoutes() *fiber.App {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", controller.Root)
	app.Get("/healthcheck", controller.HealthCheck)

	app.Get("/api/todos", controller.GetTodos)
	app.Post("/api/todos", controller.AddTodo)
	app.Get("/api/todos/:id", controller.GetTodo)
	app.Delete("/api/todos/:id", controller.DelTodo)
	app.Patch("/api/todos/:id/done", controller.SetTodoDone)

	return app
}