package main

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/rahmatadlin/Todo-Golang-React/pkg/controller"
	"github.com/rahmatadlin/Todo-Golang-React/pkg/server"
	"github.com/stretchr/testify/assert" // add Testify package
)

func TestRootRoute(t *testing.T) {
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		method       string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		// First test case
		{
			description:  "get HTTP status 200",
			route:        "/",
			method:       "GET",
			expectedCode: 200,
		},
		// Second test case
		{
			description:  "run not allowed method POST",
			route:        "/",
			method:       "POST",
			expectedCode: 405,
		},
	}

	// Define Fiber app.
	app := server.AppWithRoutes()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		defer resp.Body.Close()
		bodyData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		if test.expectedCode == 200 {
			assert.Equal(t, "Todo backend server", string(bodyData))
		}
	}
}

func TestListTodosRoute(t *testing.T) {
	route := "/api/todos"
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string // description of the test case
		route        string // route path to test
		method       string // route path to test
		expectedCode int    // expected HTTP status code
	}{
		{
			description:  "get HTTP status 200",
			route:        route,
			method:       "GET",
			expectedCode: 200,
		},
	}
	controller.Todos[1] = &controller.Todo{
		ID:    1,
		Title: "Todo Title",
		Body:  "Todo Body",
		Done:  false,
	}
	// Define Fiber app.
	app := server.AppWithRoutes()

	// Iterate through test single test cases
	for _, test := range tests {
		// r := strings.NewReader(`{"title": 123, "body": "body-123"}`)
		// Create a new http request with the route from the test case
		req := httptest.NewRequest(test.method, test.route, nil)

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)
		// fmt.Printf("%#v", resp)
		defer resp.Body.Close()
		bodyData, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("expected error to be nil got %v", err)
		}
		if test.expectedCode == 200 {
			expStr, err := json.Marshal(controller.Todos)
			if err != nil {
				t.Errorf("expected error to be nil got %v", err)
			}
			assert.Equal(t, string(expStr), string(bodyData))
		}
	}
}

func TestAddTodoRoute(t *testing.T) {
	route := "/api/todos"
	// Define a structure for specifying input and output data
	// of a single test case
	tests := []struct {
		description  string    // description of the test case
		route        string    // route path to test
		method       string    // route path to test
		expectedCode int       // expected HTTP status code
		body         io.Reader // expected HTTP status code
		isJson       bool
	}{
		// Second test case
		{
			description:  "validation error with empty body",
			route:        route,
			expectedCode: 422,
			body:         nil,
		},
		{
			description:  "validation error with empty body",
			route:        route,
			expectedCode: 500,
			body:         nil,
			isJson:       true,
		},
		{
			description:  "validation error with empty body",
			route:        route,
			expectedCode: 201,
			body:         strings.NewReader(`{"title": "123", "body": "body-123"}`),
			isJson:       true,
		},
		{
			description:  "validation error with empty body",
			route:        route,
			expectedCode: 201,
			body:         strings.NewReader(`{"title": "234", "body": "body-234"}`),
			isJson:       true,
		},
	}

	// Define Fiber app.
	app := server.AppWithRoutes()

	// Iterate through test single test cases
	for _, test := range tests {
		// Create a new http request with the route from the test case
		req := httptest.NewRequest("POST", test.route, test.body)
		if test.isJson {
			req.Header.Add("Content-Type", "application/json")
		}

		// Perform the request plain with the app,
		// the second argument is a request latency
		// (set to -1 for no latency)
		resp, _ := app.Test(req, 1)

		// Verify, if the status code is as expected
		assert.Equalf(t, test.expectedCode, resp.StatusCode, test.description)

		// fmt.Printf("%#v", resp)
	}
}