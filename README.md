# Todo-Golang-React

## Initialize Go app
go mod init github.com/rahmatadlin/Todo-Golang-React

## Install Fiber v2
go get -u github.com/gofiber/fiber/v2

# Install Testify Assert
go get github.com/stretchr/testify/assert

# Test Server
go test -v

## Create client app with Vite
npm create vite@latest

## Install dependencies
npm i @mantine/hooks @mantine/core swr @primer/octicons-react

## Run Server
ALLOW_ORIGIN_FROM=http://localhost:5173 PORT=4000 go run main.go

## Run Client
VITE_BACKEND_URL=http://localhost:4000/api npm run dev