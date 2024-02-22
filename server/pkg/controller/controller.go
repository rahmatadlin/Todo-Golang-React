package controller

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

var Todos = []Todo{}

func HealthCheck(c *fiber.Ctx) error {
	return c.SendString(fmt.Sprintf("OK. Allow origin from - %s.", os.Getenv("ALLOW_ORIGIN_FROM")))
}

func Root(c *fiber.Ctx) error {
	return c.SendString("Todo backend server")
}

func AddTodo(c *fiber.Ctx) error {
	newTodo := &Todo{}
	if err := c.BodyParser(newTodo); err != nil {
		return err
	}

	newTodo.ID = len(Todos) + 1
	Todos = append(Todos, *newTodo)

	return c.Status(201).JSON(newTodo)
}

func GetTodos(c *fiber.Ctx) error {
	return c.JSON(Todos)
}

func findTodo(c *fiber.Ctx) (int, int, error) {
	id, err := c.ParamsInt("id")
	if err != nil {
		return -1, 401, fmt.Errorf("invalid id")
	}
	for k, v := range Todos {
		if v.ID == id {
			return k, 200, nil
		}
	}
	return -1, 404, fmt.Errorf("Todo not found")
}

func GetTodo(c *fiber.Ctx) error {
	foundIndex, status, err := findTodo(c)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	return c.JSON(Todos[foundIndex])
}

func SetTodoDone(c *fiber.Ctx) error {
	foundIndex, status, err := findTodo(c)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	Todos[foundIndex].Done = !Todos[foundIndex].Done

	return c.JSON(Todos[foundIndex])
}

func DelTodo(c *fiber.Ctx) error {
	foundIndex, status, err := findTodo(c)
	if err != nil {
		return c.Status(status).SendString(err.Error())
	}
	Todos = append(Todos[:foundIndex], Todos[foundIndex+1:]...)
	return c.Status(204).JSON("")
}