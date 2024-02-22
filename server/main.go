package main

import (
    "fmt"
    "log"

    "github.com/rahmatadlin/Todo-Golang-React/pkg/server"
)

func main() {
    app := server.AppWithRoutes()

    port := ":4000" // Change port to 4000
    fmt.Printf("Listen on port http://0.0.0.0%s\n", port)
    log.Fatal(app.Listen(port))
}
