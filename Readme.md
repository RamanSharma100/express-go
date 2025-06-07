# express-go [Work in Progress]

[![GitHub](https://img.shields.io/github/license/ramansharma100/express-go)](LICENSE)
[![GitHub issues](https://img.shields.io/github/issues/ramansharma100/express-go)](https://github.com/ramansharma100/express-go/issues)
[![Go Report Card](https://goreportcard.com/badge/github.com/ramansharma100/express-go)](https://goreportcard.com/report/github.com/ramansharma100/express-go)
[![GoDoc](https://pkg.go.dev/badge/github.com/ramansharma100/express-go)](https://pkg.go.dev/github.com/ramansharma100/express-go)

A simple and lightweight MVC Web framework for Go, inspired by Node.js Express and PHP Laravel. It provides a minimalistic API for building web applications and APIs with ease.

> - This is for educational purposes and is not intended for production use. It is a work in progress and may not be fully functional or secure.
> - The API is inspired by Express.js and php laravel, but it is not a direct port. Some features may be different or missing.

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
- Middleware support (global and route-specific)
- Grouping of routes with middleware
- Support for URL parameters
- Middleware Chaining
- Request validation
- Embedding Middleware in Route groups
- Custom Error handling [Panic handling]
- CORS support
- Request logging
- Route naming
- Support for query parameters
- Route chaining [For now only `Name` is supported, more work needed to support other features like `Middleware`, etc.]
- Static file serving [need improvements]
- Rate limiting [in memory not redis implementation - will be added while caching support]
- URL encoding/decoding [Available in Context]

## Upcoming Features

This is lot of work in progress and will be updated frequently. Some of the upcoming features include:

- Custom Template Engine Support
- Support for cookies
- Session management
- File uploads [multer like]
- WebSocket support
- Support for mixins
- Support for plugins
- Layout support
- Support iterative routing parameters

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

func middlewareTest1(ctx *http.Context, next func()) {
	fmt.Println("Middleware Group Test 1")
	next()
}

func middlewareTest2(ctx *http.Context, next func()) {
	fmt.Println("Middleware Group Test 2")
	next()
}

func main() {
	app := http.New()

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 1")
		next()
	})

	// rate limiter
	app.Use(http.RateLimit(&http.RateLimitOptions{
		Limit:     10,
		Window:    60, // in seconds
		Remaining: 10,
	}))

	// cors middleware
	app.Use(http.CORS(&http.CorsOptions{
		AllowOrigin: "*,"
	}))

	// custom error handler
	app.SetErrorHandler(func(ctx *http.Context, err error) {
		fmt.Println("Error Handler:", err)
		ctx.Response.Status(500).Json(map[string]any{
			"error":   err.Error(),
			"message": "An error occurred from custom error handler",
		})
	})


	app.Get("/", func(ctx *http.Context) {
		ctx.Request.AddField("user", "Raman Sharma")
		http.Logger().Info("Request received for / endpoint")
		// ctx.Response.Send("Hello, World!" + ctx.Request.AdditionalFields["user"].(string))
		ctx.Render("index.html", map[string]any{
			"user": ctx.Request.AdditionalFields["user"],
		})
	}).Name("home")

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 2")
		next()
	})

	app.Get("/error", func(ctx *http.Context) {
		panic("This is a test error")
	})

	// use groups to add
	app.Group("/test", []http.Middleware{middlewareTest1, middlewareTest2}, func(router *http.Router) {
		router.Get("/", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"message": "Hello from test group!"})
		})

		router.Get("/info", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"info": "Test group information."})
		})
	})

	app.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware Global 3")
		next()
	})

	app.Get("/:id", func(ctx *http.Context) {
		params := ctx.GetParams()
		ctx.Response.Json(
			map[string]any{
				"params":     params,
				"queryPrams": ctx.GetSearchParams(),
			},
		)
	}).Name("getById")

	app.UseRouter("/company", routes.CompanyRouter())

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

	app.Patch("/user", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		body := ctx.Request.GetJsonBody()
		// validate the request body
		errors := ctx.Request.Validate(map[string]string{
			"name":  "required|string",
			"email": "required|email",
		}, body)

		if len(errors) > 0 {
			ctx.Response.Status(400).Json(map[string]any{"errors": errors})
			return
		}

		ctx.Response.Status(200).Json(map[string]any{"body": body})
	})

	app.Listen(8000, func(port int, err error) {
		if err != nil {
			panic(err)
		}
		fmt.Println("Server is running on http://localhost:" + strconv.Itoa(port))
	})
}


func middlewareTest1(ctx *http.Context, next func()) {
	fmt.Println("Middleware Router Group Test 1")
	next()
}

func middlewareTest2(ctx *http.Context, next func()) {
	fmt.Println("Middleware Router Group Test 2")
	next()
}

func CompanyRouter() *http.Router {
	router := http.NewRouter()

	router.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware  Router Global 1")
		next()
	})

	// // use groups to add
	router.Group("/test", []http.Middleware{middlewareTest1, middlewareTest2}, func(r *http.Router) {
		r.Get("/", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"message": "Hello from test group!"})
		})

		r.Get("/info", func(ctx *http.Context) {
			ctx.Response.AddHeader("Content-Type", "application/json")
			ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
			ctx.Response.Status(200).Json(map[string]any{"info": "Test group information."})
		})
	})

	router.Get("/", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"message": "Welcome to the company!"})
	}).Name("companyHome")

	router.Get("/info", func(ctx *http.Context) {
		ctx.Response.AddHeader("Content-Type", "application/json")
		ctx.Response.AddHeader("X-Custom-Header", "CustomValue")
		ctx.Response.Status(200).Json(map[string]any{"info": "Company information goes here."})
	})

	router.Use(func(ctx *http.Context, next func()) {
		fmt.Println("Middleware  Router Global 2")
		next()
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
- `app.Use(middleware Middleware)` - Add global middleware
- `app.Group(path string, middlewares []Middleware, handler func(*Router))` - Group routes with middleware
- `app.SetErrorHandler(handler func(*Context, error))` - Set a custom error handler

### HTTP

- `http.New()` - Create a new application instance
- `http.NewRouter()` - Create a new router instance
- `http.Context` - Context for request and response handling
- `http.Handler` - Type for request handlers
- `http.Router` - Router for handling routes
- `http.CORS(options *CorsOptions)` - Middleware for handling CORS
- `http.Logger()` - Get the global logger instance
- `http.RateLimit(options *RateLimitOptions)` - Middleware for rate limiting

### Context

- `ctx.GetParams()` - Get URL parameters as a map
- `ctx.GetBody()` - Get the request body as a string
- `ctx.GetJsonBody()` - Get the request body parsed as JSON
- `ctx.GetSearchParams()` - Get query parameters as a map
- `ctx.GetSearchParam(key string)` - Get a specific query parameter by key
- `ctx.Redirect(url string)` - Redirect to a different URL
- `ctx.Render(template string, data map[string]any)` - Render an HTML template with data
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
- `ctx.EncodeURL(urls ...string) string` - Encode URLs for safe transmission
- `ctx.DecodeURL(url string) string` - Decode URLs from their encoded form

### Router

- `router.Get(path string, handler Handler)`
- `router.Post(path string, handler Handler)`
- `router.Put(path string, handler Handler)`
- `router.Patch(path string, handler Handler)`
- `router.Delete(path string, handler Handler)`
- `router.Options(path string, handler Handler)`
- `router.Use(middleware Middleware)` - Add middleware to the router
- `router.Group(path string, middlewares []Middleware, handler func(*Router))` - Group routes with middleware

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
- `ctx.Request.GetJsonBody()` - Get the request body parsed as JSON
- `ctx.Request.GetParams()` - Get URL parameters as a map
- `ctx.Request.GetSearchParams()` - Get query parameters as a map
- `ctx.Request.GetSearchParam(key string)` - Get a specific query parameter by key
- `ctx.Request.AdditionalFields` - Map for additional fields added by middleware or handlers

### Response

- `ctx.Response.Send(text string)` - Send plain text response
- `ctx.Response.Json(data any)` - Send JSON response
- `ctx.Response.Status(code int)` - Set status code
- `ctx.Response.AddHeader(key, value string)` - Add custom header
- `ctx.Response.Writer` - Get the underlying `http.ResponseWriter`

### Route Chaining

You can chain multiple handlers for a route:
note:- For now only `Name` is supported for chaining, more work needed to support other features like `Middleware`, `Group`, etc.

```go
app.Get("/example", func(ctx *http.Context) {
	ctx.Response.Send("First handler")
}).Name("example")
```

### Logging

You can use the global logger instance to log messages:

```go
http.Logger().Info("This is an info message")
http.Logger().Error("This is an error message")
http.Logger().Debug("This is a debug message")
```

### CORS

CORS middleware is included to handle cross-origin requests. You can configure it with options:

````go
app.Use(http.CORS(&http.CorsOptions{
	AllowOrigin: "*",
	AllowMethods: "GET, POST, PUT, DELETE, OPTIONS",
	AllowHeaders: "Content-Type, Authorization",
	ContentType: "application/json",
}))

### Rate Limiting

You can use the rate limiting middleware to limit the number of requests from a client:

```go
app.Use(http.RateLimit(&http.RateLimitOptions{
	Limit:     100, // Maximum requests allowed
	Window:    60,  // Time window in seconds
	Remaining: 100, // Remaining requests allowed
}))
````

### Error Handling

You can set a custom error handler to handle errors globally:

```go
app.SetErrorHandler(func(ctx *http.Context, err error) {
	fmt.Println("Custom Error Handler:", err)
	ctx.Response.Status(500).Json(map[string]any{
		"error":   err.Error(),
		"message": "An error occurred from custom error handler",
	})
})
```

### Static File Serving

You can serve static files using the `Static` middleware:

```go
app.Static("/static", "./public") // there are improvements needed
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines.

Feel free to open issues for suggestions, bugs, or questions.

## License

This project is licensed under the terms of the [MIT License](LICENSE).

```

```
