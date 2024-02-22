package main

import (
	"fmt"
	"log"
	"os"

	"github.com/rahmatadlin/Todo-Golang-React/pkg/server"
)

func main() {
	app := server.AppWithRoutes()

	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	fmt.Printf("Listen on port http://0.0.0.0%s", port)
	log.Fatal(app.Listen(port))
}