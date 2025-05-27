# express-go [Work in Progress]

[![GitHub](https://img.shields.io/github/license/ramansharma100/express-go)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/ramansharma100/express-go)](https://github.com/ramansharma100/express-go/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramansharma100/express-go)](https://goreportcard.com/report/github.com/ramansharma100/express-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/ramansharma100/express-go)](https://pkg.go.dev/github.com/ramansharma100/express-go)

A simple and lightweight HTTP server framework for Go, inspired by Node.js Express. It provides a minimalistic API for building web applications and APIs with ease.

> - This is for educational purposes and is not intended for production use. It is a work in progress and may not be fully functional or secure.
> - The API is inspired by Express.js, but it is not a direct port. Some features may be different or missing.

## Features

- Simple routing (GET, POST, PUT, DELETE, PATCH, OPTIONS, HEAD)
- Context-based request and response handling [More work needed]
- Easy to use API inspired by Node.js Express
- Add custom headers to requests and responses
- Add additional fields to requests
- Easily extendable and lightweight
- JSON response support
- Plain text response support
- JSON parsing support
- Parameterized routes
- Nested routing
- Create Router instances [For modular routing]
- UseRouter function to use a router in the main application
- HTML template rendering

## Upcoming Features

This is lot of work in progress and will be updated frequently. Some of the upcoming features include:

- Custom Template Engine Support
- Custom Error Pages
- Custom Error Handlers
- Request logging
- Route grouping
- Route naming
- Middleware chaining
- Embedding Middleware in Route groups
- Rendering HTML templates
- Support for query parameters
- Support for URL parameters
- Support for cookies
- Middleware support
- Custom Error handling
- Static file serving
- Session management
- Global error handling
- CORS support
- Rate limiting
- Request validation
- File uploads [multer like]
- URL encoding
- WebSocket support
- Support for mixins
- Support for plugins
- Layout support

## Upcoming Extended Features

- CLI Support
- Migrations Support
- AI Studio Support
- Will Add more...

## Installation

```bash
go get github.com/ramansharma100/express-go # may not work yet
```

or

```bash
# or
# clone the repo and run
go build
go install
```

## Usage

```go
package main

import (
	"fmt"
	"strconv"

	"github.com/ramansharma100/express-go/http"
)

func main() {
	app := http.New()

	app.Get("/", func(ctx *http.Context) {
		ctx.Request.AddField("user", "Raman Sharma")
		// ctx.Response.Send("Hello, World!" + ctx.Request.AdditionalFields["user"].(string))
		ctx.Render("index.html", map[string]any{
			"user": ctx.Request.AdditionalFields["user"],
		})
	})

	app.UseRouter("/company", CompanyRouter())

	app.Get("/user/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(
			map[string]any{
				"params": ctx.GetParams(),
			},
		)
	})

	app.Get("/user/:id/:name", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"context": ctx.GetParams()})
	})

	app.Get("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Send("Hello from user")
	})

	app.Post("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.GetBody()})
	})

	app.Put("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.Request.GetJsonBody()})
	})

	app.Patch("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		fmt.Println("Body:", ctx.GetBody())
		ctx.Response.Status(200).Json(map[string]any{"body": ctx.Request.GetJsonBody()})
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}

func CompanyRouter() *http.Router {
	router := http.NewRouter()

	router.Get("/", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"message": "Welcome to the company!"})
	})

	router.Get("/info", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"info": "Company information goes here."})
	})

	router.Post("/create", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		body := ctx.GetBody()
		ctx.Response.Status(201).Json(map[string]any{"message": "Company created successfully", "data": body})
	})

	router.Put("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		body := ctx.GetBody()
		ctx.Response.Status(200).Json(map[string]any{"message": "Company updated successfully", "id": id, "data": body})
	})

	router.Patch("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		body := ctx.GetBody()
		ctx.Response.Status(200).Json(map[string]any{"message": "Company patched successfully", "id": id, "data": body})
	})

	router.Delete("/:id", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		params := ctx.GetParams()
		id := params["id"]
		ctx.Response.Status(200).Json(map[string]any{"message": "Company deleted successfully", "id": id})
	})

	return router
}
```

```bash
# templates/index.html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Welcome to Express GO</title>
</head>

<body>
    <h1>Welcome to Express GO</h1>
    <p>This is a simple Express GO application for go developers.</p>
    <p>
        User Logged In: <strong>{{.user}}</strong>
    </p>
</body>

</html>
```

## API

### Application

- `app.Get(path string, handler Handler)`
- `app.Post(path string, handler Handler)`
- `app.Put(path string, handler Handler)`
- `app.Patch(path string, handler Handler)`
- `app.Delete(path string, handler Handler)`
- `app.Options(path string, handler Handler)`
- `app.Listen(port int, callback func(int, error))`
- `app.UseRouter(path string, router *Router)` - Use a router for a specific path

### HTTP

- `http.New()` - Create a new application instance
- `http.NewRouter()` - Create a new router instance
- `http.Context` - Context for request and response handling
- `http.Handler` - Type for request handlers
- `http.Router` - Router for handling routes

### Context

- `ctx.GetParams()` - Get URL parameters as a map
- `ctx.GetBody()` - Get the request body as a string
- `ctx.GetJsonBody()` - Get the request body parsed as JSON
- `ctx.Request` - Access the request object
- `ctx.Response` - Access the response object
- `ctx.Response.Status(code int)` - Set the response status code
- `ctx.Response.Send(text string)` - Send a plain text response
- `ctx.Response.Json(data any)` - Send a JSON response
- `ctx.Response.AddHeader(key, value string)` - Add a custom header to the response
- `ctx.Response.Writer` - Get the underlying `http.ResponseWriter`
- `ctx.Request.AddHeader(key, value string)` - Add a custom header to the request
- `ctx.Request.AddField(key, value string)` - Add a custom field to the request
- `ctx.Request.ParseBody()` - Parse the request body (for POST requests)
- `ctx.Request.Method` - Get the HTTP method of the request
- `ctx.Request.Url` - Get the URL string of the request

### Router

- `router.Get(path string, handler Handler)`
- `router.Post(path string, handler Handler)`
- `router.Put(path string, handler Handler)`
- `router.Patch(path string, handler Handler)`
- `router.Delete(path string, handler Handler)`
- `router.Options(path string, handler Handler)`

### Handler

A handler receives a `*http.Context`:

```go
func(ctx *http.Context) {
    // ctx.Request, ctx.Response
}
```

### Request

- `ctx.Request.Method` - HTTP method
- `ctx.Request.Url` - URL string
- `ctx.Request.Body` - Request body (for POST)
- `ctx.Request.Headers` - Map of headers
- `ctx.Request.AddHeader(key, value string)` - Add custom header
- `ctx.Request.AddField(key, value string)` - Add custom field to `request.AdditionalFields`
- `ctx.Request.ParseBody()` - Parse request body (for POST)
- `ctx.Request.r` - Get the underlying `http.Request`

### Response

- `ctx.Response.Send(text string)` - Send plain text response
- `ctx.Response.Json(data any)` - Send JSON response
- `ctx.Response.Status(code int)` - Set status code
- `ctx.Response.AddHeader(key, value string)` - Add custom header
- `ctx.Response.Writer` - Get the underlying `http.ResponseWriter`

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

Feel free to open issues for suggestions, bugs, or questions.

## License

This project is licensed under the terms of the [MIT License](LICENSE).
