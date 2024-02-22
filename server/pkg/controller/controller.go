package controller

import (
	"fmt"
	"os"
	"sync"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
	Body  string `json:"body"`
}

var (
	Todos      = make(map[int]*Todo)
	LastTodoID int
	mu         sync.Mutex
)

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

	mu.Lock()
	defer mu.Unlock()

	LastTodoID++
	newTodo.ID = LastTodoID
	Todos[newTodo.ID] = newTodo

	return c.Status(201).JSON(newTodo)
}

func GetTodos(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	return c.JSON(Todos)
}

func GetTodo(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	todo, ok := Todos[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}
	return c.JSON(todo)
}

func SetTodoDone(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	todo, ok := Todos[id]
	if !ok {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}
	todo.Done = !todo.Done
	return c.JSON(todo)
}

func DelTodo(c *fiber.Ctx) error {
	mu.Lock()
	defer mu.Unlock()

	id, err := c.ParamsInt("id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid ID")
	}

	if _, ok := Todos[id]; !ok {
		return c.Status(fiber.StatusNotFound).SendString("Todo not found")
	}

	delete(Todos, id)
	return c.SendStatus(fiber.StatusNoContent)
}
